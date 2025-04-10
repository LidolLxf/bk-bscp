/*
 * Tencent is pleased to support the open source community by making Blueking Container Service available.
 * Copyright (C) 2019 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 * http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 */

package service

import (
	"context"
	"errors"
	"fmt"
	"path"
	"time"

	pbstruct "github.com/golang/protobuf/ptypes/struct"
	"gorm.io/gorm"

	"github.com/TencentBlueKing/bk-bscp/internal/dal/gen"
	"github.com/TencentBlueKing/bk-bscp/pkg/criteria/errf"
	"github.com/TencentBlueKing/bk-bscp/pkg/dal/table"
	"github.com/TencentBlueKing/bk-bscp/pkg/i18n"
	"github.com/TencentBlueKing/bk-bscp/pkg/kit"
	"github.com/TencentBlueKing/bk-bscp/pkg/logs"
	pbbase "github.com/TencentBlueKing/bk-bscp/pkg/protocol/core/base"
	pbci "github.com/TencentBlueKing/bk-bscp/pkg/protocol/core/config-item"
	pbcontent "github.com/TencentBlueKing/bk-bscp/pkg/protocol/core/content"
	pbkv "github.com/TencentBlueKing/bk-bscp/pkg/protocol/core/kv"
	pbrelease "github.com/TencentBlueKing/bk-bscp/pkg/protocol/core/release"
	pbrkv "github.com/TencentBlueKing/bk-bscp/pkg/protocol/core/released-kv"
	pbtv "github.com/TencentBlueKing/bk-bscp/pkg/protocol/core/template-variable"
	pbds "github.com/TencentBlueKing/bk-bscp/pkg/protocol/data-service"
	"github.com/TencentBlueKing/bk-bscp/pkg/tools"
	"github.com/TencentBlueKing/bk-bscp/pkg/types"
)

// CreateRelease create release.
// nolint:funlen,gocyclo
func (s *Service) CreateRelease(ctx context.Context, req *pbds.CreateReleaseReq) (*pbds.CreateResp, error) {
	grpcKit := kit.FromGrpcContext(ctx)

	app, err := s.dao.App().GetByID(grpcKit, req.Attachment.AppId)
	if err != nil {
		logs.Errorf("get app failed, err: %v, rid: %s", err, grpcKit.Rid)
		return nil, err
	}

	if _, e := s.dao.Release().GetByName(grpcKit, req.Attachment.BizId, req.Attachment.AppId, req.Spec.Name); e == nil {
		return nil, fmt.Errorf("release name %s already exists", req.Spec.Name)
	}
	// begin transaction to create release and released config item.
	tx := s.dao.GenQuery().Begin()
	// 1. create release, and create release and released config item need to begin tx.
	release := &table.Release{
		Spec:       req.Spec.ReleaseSpec(),
		Attachment: req.Attachment.ReleaseAttachment(),
		Revision: &table.CreatedRevision{
			Creator: grpcKit.User,
		},
	}
	id, err := s.dao.Release().CreateWithTx(grpcKit, tx, release)
	if err != nil {
		logs.Errorf("create release failed, err: %v, rid: %s", err, grpcKit.Rid)
		if rErr := tx.Rollback(); rErr != nil {
			logs.Errorf("transaction rollback failed, err: %v, rid: %s", rErr, grpcKit.Rid)
		}
		return nil, err
	}

	// 2. create released hook.
	pre, err := s.dao.ReleasedHook().Get(grpcKit, req.Attachment.BizId, req.Attachment.AppId, 0, table.PreHook)
	if err == nil {
		pre.ID = 0
		pre.ReleaseID = release.ID
		if _, e := s.dao.ReleasedHook().CreateWithTx(grpcKit, tx, pre); e != nil {
			logs.Errorf("create released pre-hook failed, err: %v, rid: %s", e, grpcKit.Rid)
			if rErr := tx.Rollback(); rErr != nil {
				logs.Errorf("transaction rollback failed, err: %v, rid: %s", rErr, grpcKit.Rid)
			}
			return nil, e
		}
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		logs.Errorf("query released pre-hook failed, err: %v, rid: %s", err, grpcKit.Rid)
		if rErr := tx.Rollback(); rErr != nil {
			logs.Errorf("transaction rollback failed, err: %v, rid: %s", rErr, grpcKit.Rid)
		}
		return nil, err
	}
	post, err := s.dao.ReleasedHook().Get(grpcKit, req.Attachment.BizId, req.Attachment.AppId, 0, table.PostHook)
	if err == nil {
		post.ID = 0
		post.ReleaseID = release.ID
		if _, e := s.dao.ReleasedHook().CreateWithTx(grpcKit, tx, post); e != nil {
			logs.Errorf("create released post-hook failed, err: %v, rid: %s", e, grpcKit.Rid)
			if rErr := tx.Rollback(); rErr != nil {
				logs.Errorf("transaction rollback failed, err: %v, rid: %s", rErr, grpcKit.Rid)
			}
			return nil, e
		}
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		logs.Errorf("query released post-hook failed, err: %v, rid: %s", err, grpcKit.Rid)
		if rErr := tx.Rollback(); rErr != nil {
			logs.Errorf("transaction rollback failed, err: %v, rid: %s", rErr, grpcKit.Rid)
		}
		return nil, err
	}

	switch app.Spec.ConfigType {
	case table.File:
		conflictNums, _, errC := s.checkNonTmpAndTmpConflicts(grpcKit, req.Attachment.BizId,
			req.Attachment.AppId, []string{})
		if errC != nil {
			return nil, errC
		}
		if conflictNums > 0 {
			logs.Errorf("create release failed there is a file conflict, err: %v, rid: %s", err, grpcKit.Rid)
			return nil, errors.New(i18n.T(grpcKit, "create release failed there is a file conflict"))
		}
		// Note: need to change batch operator to query config item and its commit.
		// get app's all config items.
		cis, fErr := s.getAppConfigItems(grpcKit)
		if fErr != nil {
			if rErr := tx.Rollback(); rErr != nil {
				logs.Errorf("transaction rollback failed, err: %v, rid: %s", rErr, grpcKit.Rid)
			}
			logs.Errorf("get app's all config items failed, err: %v, rid: %s", fErr, grpcKit.Rid)
			return nil, fErr
		}

		// get app template revisions which are template config items
		tmplRevisions, fErr := s.getAppTmplRevisions(grpcKit)
		if fErr != nil {
			if rErr := tx.Rollback(); rErr != nil {
				logs.Errorf("transaction rollback failed, err: %v, rid: %s", rErr, grpcKit.Rid)
			}
			logs.Errorf("get app template revisions failed, err: %v, rid: %s", fErr, grpcKit.Rid)
			return nil, fErr
		}

		// if no config item, return directly.
		if len(cis) == 0 && len(tmplRevisions) == 0 {
			return nil, errors.New("app config items is empty")
		}

		// 3: do template and non-template config item related operations for create release.
		if err = s.doConfigItemOperations(grpcKit, req.Variables, tx, release.ID, tmplRevisions, cis); err != nil {
			if rErr := tx.Rollback(); rErr != nil {
				logs.Errorf("transaction rollback failed, err: %v, rid: %s", rErr, grpcKit.Rid)
			}
			logs.Errorf("do template action for create release failed, err: %v, rid: %s", err, grpcKit.Rid)
			return nil, err
		}
	case table.KV:
		expirationNumber, errC := s.checkForExpiredCertificates(grpcKit, req.Attachment.BizId, req.Attachment.AppId)
		if errC != nil {
			return nil, errC
		}
		if expirationNumber > 0 {
			return nil, errors.New(i18n.T(grpcKit, "create release failed there is a certificate expiration exists"))
		}

		if err = s.doKvOperations(grpcKit, tx, req.Attachment.AppId, req.Attachment.BizId, release.ID); err != nil {
			if rErr := tx.Rollback(); rErr != nil {
				logs.Errorf("transaction rollback failed, err: %v, rid: %s", rErr, grpcKit.Rid)
			}
			logs.Errorf("do kv action for create release failed, err: %v, rid: %s", err, grpcKit.Rid)
			return nil, err
		}
	}

	// commit transaction.
	if err = tx.Commit(); err != nil {
		logs.Errorf("commit transaction failed, err: %v, rid: %s", err, grpcKit.Rid)
		return nil, err
	}
	return &pbds.CreateResp{Id: id}, nil
}

// doConfigItemOperations do config item related operations for create release.
/*
1.下载服务的所有模版和非模版配置文件内容，提取服务模版变量
2.获取入参变量和业务变量，判断是否缺少变量，缺少则报错
3.使用变量渲染模版和非模版配置文件，上传渲染后的内容
4.创建已生成版本服务的模版和非模版配置项
5.创建已生成版本服务的服务模版详情
6.创建已生成版本服务的模版变量
7.将当前使用变量更新到未命名版本的服务模版变量
*/
//nolint:funlen
func (s *Service) doConfigItemOperations(kt *kit.Kit, variables []*pbtv.TemplateVariableSpec,
	tx *gen.QueryTx, releaseID uint32, tmplRevisions []*table.TemplateRevision, cis []*pbci.ConfigItem) error {
	// validate input variables and get the map
	inputVarMap := make(map[string]*table.TemplateVariableSpec)
	for _, v := range variables {
		if v == nil {
			continue
		}
		if err := v.TemplateVariableSpec().ValidateCreate(kt); err != nil {
			logs.Errorf("validate template variables failed, err: %v, rid: %s", err, kt.Rid)
			return err
		}
		inputVarMap[v.Name] = v.TemplateVariableSpec()
	}

	tmplsNeedRender := filterSizeForTmplRevisions(tmplRevisions)
	cisNeedRender := filterSizeForConfigItems(cis)

	vars, ciVars, allVars, err := s.getVariables(kt, tmplsNeedRender, cisNeedRender)
	if err != nil {
		logs.Errorf("get variables failed, err: %v, rid: %s", err, kt.Rid)
		return err
	}
	tmplsNeedRender = filterVarsForTmplRevisions(tmplsNeedRender, vars)
	cisNeedRender = filterVarsForConfigItems(cisNeedRender, ciVars)

	contents, err := s.downloadTmplContent(kt, tmplsNeedRender)
	if err != nil {
		logs.Errorf("download template content failed, err: %v, rid: %s", err, kt.Rid)
		return err
	}
	ciContents, err := s.downloadCIContent(kt, cisNeedRender)
	if err != nil {
		logs.Errorf("download config item content failed, err: %v, rid: %s", err, kt.Rid)
		return err
	}

	usedVars, renderKV, err := s.getRenderedVars(kt, allVars, inputVarMap)
	if err != nil {
		logs.Errorf("get rendered variables failed, err: %v, rid: %s", err, kt.Rid)
		return err
	}

	// get rendered content map which is template revision id => rendered content
	renderedContentMap := make(map[uint32][]byte, len(tmplRevisions))
	signatureMap := make(map[uint32]string, len(tmplRevisions))
	md5Map := make(map[uint32]string, len(tmplRevisions))
	byteSizeMap := make(map[uint32]uint64, len(tmplRevisions))
	revisionMap := make(map[uint32]*table.TemplateRevision, len(tmplRevisions))
	// data which need render
	for idx, r := range tmplsNeedRender {
		revisionMap[r.ID] = r
		renderedContentMap[r.ID] = s.tmplProc.Render(contents[idx], renderKV)
		signatureMap[r.ID] = tools.ByteSHA256(renderedContentMap[r.ID])
		md5Map[r.ID] = tools.ByteMD5(renderedContentMap[r.ID])
		byteSizeMap[r.ID] = uint64(len(renderedContentMap[r.ID]))
	}
	// data which doesn't need render
	for _, r := range tmplRevisions {
		if _, ok := revisionMap[r.ID]; ok {
			continue
		}
		revisionMap[r.ID] = r
		signatureMap[r.ID] = r.Spec.ContentSpec.Signature
		md5Map[r.ID] = r.Spec.ContentSpec.Md5
		byteSizeMap[r.ID] = r.Spec.ContentSpec.ByteSize
	}

	// get rendered content map which is config item id => rendered content
	ciRenderedContentMap := make(map[uint32][]byte, len(cis))
	ciSignatureMap := make(map[uint32]string, len(cis))
	ciMd5Map := make(map[uint32]string, len(cis))
	ciByteSizeMap := make(map[uint32]uint64, len(cis))
	ciMap := make(map[uint32]*pbci.ConfigItem, len(cis))
	// data which need render
	for idx, ci := range cisNeedRender {
		ciMap[ci.Id] = ci
		ciRenderedContentMap[ci.Id] = s.tmplProc.Render(ciContents[idx], renderKV)
		ciSignatureMap[ci.Id] = tools.ByteSHA256(ciRenderedContentMap[ci.Id])
		ciMd5Map[ci.Id] = tools.ByteMD5(ciRenderedContentMap[ci.Id])
		ciByteSizeMap[ci.Id] = uint64(len(ciRenderedContentMap[ci.Id]))
	}
	// data which doesn't need render
	for _, ci := range cis {
		if _, ok := ciMap[ci.Id]; ok {
			continue
		}
		ciMap[ci.Id] = ci
		ciMd5Map[ci.Id] = ci.CommitSpec.Content.Md5
		ciSignatureMap[ci.Id] = ci.CommitSpec.Content.Signature
		ciByteSizeMap[ci.Id] = ci.CommitSpec.Content.ByteSize
	}

	// upload rendered template content
	if e := s.uploadRenderedTmplContent(kt, renderedContentMap, signatureMap, revisionMap); e != nil {
		logs.Errorf("upload rendered template failed, err: %v, rid: %s", e, kt.Rid)
		return e
	}
	// upload rendered config item content
	if e := s.uploadRenderedCIContent(kt, ciRenderedContentMap, ciSignatureMap, ciMap); e != nil {
		logs.Errorf("upload rendered config item failed, err: %v, rid: %s", e, kt.Rid)
		return e
	}

	if e := s.createReleasedRenderedTemplateCIs(kt, tx, releaseID, tmplRevisions, renderedContentMap, byteSizeMap,
		signatureMap, md5Map); e != nil {
		logs.Errorf("create released rendered template config items failed, err: %v, rid: %s", e, kt.Rid)
		return e
	}

	if err = s.createReleasedRenderedCIs(kt, tx, releaseID, cis, ciRenderedContentMap, ciByteSizeMap,
		ciSignatureMap, ciMd5Map); err != nil {
		if rErr := tx.Rollback(); rErr != nil {
			logs.Errorf("transaction rollback failed, err: %v, rid: %s", rErr, kt.Rid)
		}
		logs.Errorf("create released config items failed, err: %v, rid: %s", err, kt.Rid)
		return err
	}

	if e := s.createReleasedAppTemplates(kt, tx, releaseID, renderedContentMap, byteSizeMap, signatureMap,
		md5Map); e != nil {
		logs.Errorf("create released rendered template config items failed, err: %v, rid: %s", e, kt.Rid)
		return e
	}

	if e := s.createReleasedAppTemplateVariable(kt, tx, releaseID, usedVars); e != nil {
		logs.Errorf("create released app template variable failed, err: %v, rid: %s", e, kt.Rid)
		return e
	}

	if e := s.updateAppTemplateVariable(kt, tx, usedVars); e != nil {
		logs.Errorf("update app template variable failed, err: %v, rid: %s", e, kt.Rid)
		return e
	}

	return nil
}

func (s *Service) getRenderedVars(kt *kit.Kit, allVars []string, inputVarMap map[string]*table.TemplateVariableSpec) (
	[]*table.TemplateVariableSpec, map[string]interface{}, error) {
	// get biz template variables
	bizVars, _, err := s.dao.TemplateVariable().List(kt, kt.BizID, nil, &types.BasePage{All: true})
	if err != nil {
		logs.Errorf("list template variables failed, err: %v, rid: %s", err, kt.Rid)
		return nil, nil, err
	}
	bizVarMap := make(map[string]*table.TemplateVariableSpec, len(bizVars))
	for _, v := range bizVars {
		bizVarMap[v.Spec.Name] = v.Spec
	}

	// get variables which are used to render the template
	usedVars := make([]*table.TemplateVariableSpec, 0)
	renderKV := make(map[string]interface{})
	var missingVars []string
	for _, name := range allVars {
		if _, ok := inputVarMap[name]; ok {
			usedVars = append(usedVars, inputVarMap[name])
			renderKV[name] = inputVarMap[name].DefaultVal
			continue
		}
		if _, ok := bizVarMap[name]; ok {
			usedVars = append(usedVars, bizVarMap[name])
			renderKV[name] = bizVarMap[name].DefaultVal
			continue
		}
		missingVars = append(missingVars, name)
	}
	if len(missingVars) > 0 {
		return nil, nil, fmt.Errorf("variable name in %v is missing for render the app's template config", missingVars)
	}

	return usedVars, renderKV, nil
}

// createReleasedRenderedTemplateCIs create released rendered templates config items.
func (s *Service) createReleasedRenderedTemplateCIs(kt *kit.Kit, tx *gen.QueryTx, releaseID uint32,
	tmplRevisions []*table.TemplateRevision, renderedContentMap map[uint32][]byte, byteSizeMap map[uint32]uint64,
	signatureMap map[uint32]string, md5Map map[uint32]string) error {
	releasedCIs := make([]*table.ReleasedConfigItem, len(tmplRevisions))
	for idx, r := range tmplRevisions {
		creator := r.Revision.Creator
		reviser := creator
		createdAt := r.Revision.CreatedAt
		updatedAt := createdAt
		// if rendered with variables, which means the template config item is new generated, update the user and time
		if _, ok := renderedContentMap[r.ID]; ok {
			reviser = kt.User
			updatedAt = time.Now().UTC()
		}

		releasedCIs[idx] = &table.ReleasedConfigItem{
			ReleaseID: releaseID,
			CommitSpec: &table.ReleasedCommitSpec{
				ContentID: 0,
				Content: &table.ReleasedContentSpec{
					Signature:       signatureMap[r.ID],
					Md5:             md5Map[r.ID],
					ByteSize:        byteSizeMap[r.ID],
					OriginSignature: r.Spec.ContentSpec.Signature,
					OriginByteSize:  r.Spec.ContentSpec.ByteSize,
				},
				Memo: "",
			},
			ConfigItemSpec: &table.ConfigItemSpec{
				Name:       r.Spec.Name,
				Path:       r.Spec.Path,
				FileType:   r.Spec.FileType,
				FileMode:   r.Spec.FileMode,
				Memo:       r.Spec.RevisionMemo,
				Permission: r.Spec.Permission,
				Charset:    r.Spec.Charset,
			},
			Attachment: &table.ConfigItemAttachment{
				BizID: kt.BizID,
				AppID: kt.AppID,
			},
			Revision: &table.Revision{
				Creator:   creator,
				Reviser:   reviser,
				CreatedAt: createdAt,
				UpdatedAt: updatedAt,
			},
		}
	}
	if err := s.dao.ReleasedCI().BulkCreateWithTx(kt, tx, releasedCIs); err != nil {
		logs.Errorf("bulk create released rendered template config item failed, err: %v, rid: %s", err, kt.Rid)
		return err
	}

	return nil
}

// createReleasedRenderedCIs create released rendered config items
func (s *Service) createReleasedRenderedCIs(kt *kit.Kit, tx *gen.QueryTx, releaseID uint32, cis []*pbci.ConfigItem,
	ciRenderedContentMap map[uint32][]byte, byteSizeMap map[uint32]uint64, signatureMap map[uint32]string,
	md5Map map[uint32]string) error {
	releasedCIs := make([]*table.ReleasedConfigItem, 0)
	if len(cis) == 0 {
		return nil
	}

	for _, ci := range cis {
		// query config item newest commit
		commit, e := s.dao.Commit().GetLatestCommit(kt, kt.BizID, kt.AppID, ci.Id)
		if e != nil {
			logs.Errorf("query config item latest commit failed, err: %v, rid: %s", e, kt.Rid)
			return e
		}

		creator := ci.Revision.Creator
		reviser := ci.Revision.Reviser
		var createdAt, updatedAt time.Time
		var err error
		createdAt, err = time.Parse(time.RFC3339, ci.Revision.CreateAt)
		if err != nil {
			return fmt.Errorf("parse time from createAt string failed, err: %v", err)
		}
		updatedAt, err = time.Parse(time.RFC3339, ci.Revision.UpdateAt)
		if err != nil {
			return fmt.Errorf("parse time from UpdateAt string failed, err: %v", err)
		}
		// if rendered with variables, which means the config item is new generated, update the user and time
		if _, ok := ciRenderedContentMap[ci.Id]; ok {
			reviser = kt.User
			updatedAt = time.Now().UTC()
		}

		releasedCIs = append(releasedCIs, &table.ReleasedConfigItem{
			CommitID: commit.ID,
			CommitSpec: &table.ReleasedCommitSpec{
				ContentID: commit.Spec.ContentID,
				Content: &table.ReleasedContentSpec{
					Signature:       signatureMap[ci.Id],
					ByteSize:        byteSizeMap[ci.Id],
					OriginSignature: ci.CommitSpec.Content.Signature,
					OriginByteSize:  ci.CommitSpec.Content.ByteSize,
					Md5:             md5Map[ci.Id],
				},
				Memo: commit.Spec.Memo,
			},
			ConfigItemID: ci.Id,
			ConfigItemSpec: &table.ConfigItemSpec{
				Name:     ci.Spec.Name,
				Path:     ci.Spec.Path,
				FileType: table.FileFormat(ci.Spec.FileType),
				FileMode: table.FileMode(ci.Spec.FileMode),
				Memo:     ci.Spec.Memo,
				Permission: &table.FilePermission{
					User:      ci.Spec.Permission.User,
					UserGroup: ci.Spec.Permission.UserGroup,
					Privilege: ci.Spec.Permission.Privilege,
				},
				Charset: ci.Spec.ConfigItemSpec().Charset,
			},
			Attachment: &table.ConfigItemAttachment{
				BizID: ci.Attachment.BizId,
				AppID: ci.Attachment.AppId,
			},
			Revision: &table.Revision{
				Creator:   creator,
				Reviser:   reviser,
				CreatedAt: createdAt,
				UpdatedAt: updatedAt,
			},
		})
	}

	// create released config item
	for _, rci := range releasedCIs {
		rci.ReleaseID = releaseID
	}
	if err := s.dao.ReleasedCI().BulkCreateWithTx(kt, tx, releasedCIs); err != nil {
		logs.Errorf("bulk create released config item failed, err: %v, rid: %s", err, kt.Rid)
		return err
	}

	return nil
}

// createReleasedAppTemplates create released app templates.
func (s *Service) createReleasedAppTemplates(kt *kit.Kit, tx *gen.QueryTx, releaseID uint32,
	renderedContentMap map[uint32][]byte, byteSizeMap map[uint32]uint64, signatureMap map[uint32]string,
	md5Map map[uint32]string) error {
	revisionsResp, err := s.ListAppBoundTmplRevisions(kt.Ctx, &pbds.ListAppBoundTmplRevisionsReq{
		BizId: kt.BizID,
		AppId: kt.AppID,
		All:   true,
	})
	if err != nil {
		logs.Errorf("list app bound template revisions failed, err: %v, rid: %s", err, kt.Rid)
		return err
	}
	revisions := revisionsResp.Details

	releasedATs := make([]*table.ReleasedAppTemplate, len(revisions))
	for idx, r := range revisions {
		creator := r.Creator
		reviser := creator
		var createdAt time.Time
		createdAt, err = time.Parse(time.RFC3339, r.CreateAt)
		if err != nil {
			return fmt.Errorf("parse time from createAt string failed, err: %v", err)
		}
		updatedAt := createdAt
		// if rendered with variables, which means the template config item is new generated, update the user and time
		if _, ok := renderedContentMap[r.TemplateRevisionId]; ok {
			reviser = kt.User
			updatedAt = time.Now().UTC()
		}

		releasedATs[idx] = &table.ReleasedAppTemplate{
			Spec: &table.ReleasedAppTemplateSpec{
				ReleaseID:            releaseID,
				TemplateSpaceID:      r.TemplateSpaceId,
				TemplateSpaceName:    r.TemplateSpaceName,
				TemplateSetID:        r.TemplateSetId,
				TemplateSetName:      r.TemplateSetName,
				TemplateID:           r.TemplateId,
				Name:                 r.Name,
				Path:                 r.Path,
				TemplateRevisionID:   r.TemplateRevisionId,
				IsLatest:             r.IsLatest,
				TemplateRevisionName: r.TemplateRevisionName,
				TemplateRevisionMemo: r.TemplateRevisionMemo,
				FileType:             r.FileType,
				FileMode:             r.FileMode,
				User:                 r.User,
				UserGroup:            r.UserGroup,
				Privilege:            r.Privilege,
				Signature:            signatureMap[r.TemplateRevisionId],
				ByteSize:             byteSizeMap[r.TemplateRevisionId],
				OriginSignature:      r.Signature,
				OriginByteSize:       r.ByteSize,
				Md5:                  md5Map[r.TemplateRevisionId],
				Charset:              r.Charset,
			},
			Attachment: &table.ReleasedAppTemplateAttachment{
				BizID: kt.BizID,
				AppID: kt.AppID,
			},
			Revision: &table.Revision{
				Creator:   creator,
				Reviser:   reviser,
				CreatedAt: createdAt,
				UpdatedAt: updatedAt,
			},
		}
	}
	if err = s.dao.ReleasedAppTemplate().BulkCreateWithTx(kt, tx, releasedATs); err != nil {
		logs.Errorf("bulk create released app template config item failed, err: %v, rid: %s", err, kt.Rid)
		return err
	}

	return nil
}

// createReleasedAppTemplates create released app template variable.
func (s *Service) createReleasedAppTemplateVariable(kt *kit.Kit, tx *gen.QueryTx, releaseID uint32,
	usedVars []*table.TemplateVariableSpec) error {
	releasedAppVar := &table.ReleasedAppTemplateVariable{
		Spec: &table.ReleasedAppTemplateVariableSpec{
			ReleaseID: releaseID,
			Variables: usedVars,
		},
		Attachment: &table.ReleasedAppTemplateVariableAttachment{
			BizID: kt.BizID,
			AppID: kt.AppID,
		},
		Revision: &table.CreatedRevision{
			Creator: kt.User,
		},
	}
	if _, err := s.dao.ReleasedAppTemplateVariable().CreateWithTx(kt, tx, releasedAppVar); err != nil {
		logs.Errorf("create released app template variable failed, err: %v, rid: %s", err, kt.Rid)
		return err
	}

	return nil
}

// updateAppTemplateVariable update app template variable.
func (s *Service) updateAppTemplateVariable(kt *kit.Kit, tx *gen.QueryTx,
	usedVars []*table.TemplateVariableSpec) error {

	appVar := &table.AppTemplateVariable{
		Spec: &table.AppTemplateVariableSpec{
			Variables: usedVars,
		},
		Attachment: &table.AppTemplateVariableAttachment{
			BizID: kt.BizID,
			AppID: kt.AppID,
		},
		Revision: &table.Revision{
			Creator: kt.User,
			Reviser: kt.User,
		},
	}
	if err := s.dao.AppTemplateVariable().UpsertWithTx(kt, tx, appVar); err != nil {
		logs.Errorf("upsert app template variables failed, err: %v, rid: %s", err, kt.Rid)
		return err
	}
	return nil
}

// ListReleases list releases.
func (s *Service) ListReleases(ctx context.Context, req *pbds.ListReleasesReq) (*pbds.ListReleasesResp, error) {
	grpcKit := kit.FromGrpcContext(ctx)

	query := &types.ListReleasesOption{
		BizID: req.BizId,
		AppID: req.AppId,
		Page: &types.BasePage{
			Start: req.Start,
			Limit: uint(req.Limit),
		},
		Deprecated: req.Deprecated,
		SearchKey:  req.SearchKey,
	}
	if req.All {
		query.Page.Start = 0
		query.Page.Limit = 0
	}

	details, err := s.dao.Release().List(grpcKit, query)
	if err != nil {
		logs.Errorf("list release failed, err: %v, rid: %s", err, grpcKit.Rid)
		return nil, err
	}

	releases := pbrelease.PbReleases(details.Details)

	gcrs, err := s.dao.ReleasedGroup().ListAllByAppID(grpcKit, req.AppId, req.BizId)
	if err != nil {
		logs.Errorf("list group current releases failed, err: %v, rid: %s", err, grpcKit.Rid)
		return nil, err
	}

	groups, err := s.dao.Group().ListAppGroups(grpcKit, req.BizId, req.AppId)
	if err != nil {
		logs.Errorf("list app groups failed, err: %v, rid: %s", err, grpcKit.Rid)
		return nil, err
	}

	// 构建map以便按 ID 快速查找群组
	groupByID := make(map[uint32]*table.Group)
	for _, group := range groups {
		groupByID[group.ID] = group
	}

	var releaseIDs []uint32
	for _, release := range releases {
		releaseIDs = append(releaseIDs, release.Id)
		status, selected := s.queryPublishStatus(gcrs, release.Id)

		releasedGroups := s.buildReleasedGroups(selected, groupByID)

		release.Status.PublishStatus = status
		release.Status.ReleasedGroups = releasedGroups
	}

	// 追加发布状态
	st, err := s.dao.Strategy().ListStrategyByReleasesIDs(grpcKit, releaseIDs)
	if err != nil {
		return nil, err
	}

	for _, r := range releases {
		for _, s := range st {
			if r.Id == s.Spec.ReleaseID {
				r.Status.StrategyStatus = string(s.Spec.PublishStatus)
			}
		}
	}

	resp := &pbds.ListReleasesResp{
		Count:   details.Count,
		Details: releases,
	}
	return resp, nil
}

// buildReleasedGroups builds the released groups for a given set of selected groups.
func (s *Service) buildReleasedGroups(selected []*table.ReleasedGroup,
	groupByID map[uint32]*table.Group) []*pbrelease.ReleaseStatus_ReleasedGroup {
	var releasedGroups []*pbrelease.ReleaseStatus_ReleasedGroup

	for _, gcr := range selected {
		// Default group handling
		if gcr.GroupID == 0 {
			releasedGroups = append(releasedGroups, &pbrelease.ReleaseStatus_ReleasedGroup{
				Id:   0,
				Name: "默认分组",
				Mode: table.GroupModeDefault.String(),
			})
		} else {
			// Process group by ID
			group, exists := groupByID[gcr.GroupID]
			if exists {
				oldSelector := new(pbstruct.Struct)
				newSelector := new(pbstruct.Struct)
				if gcr.Selector != nil {
					s, err := gcr.Selector.MarshalPB()
					if err == nil {
						oldSelector = s
					}
				}
				if group.Spec.Selector != nil {
					s, err := group.Spec.Selector.MarshalPB()
					if err == nil {
						newSelector = s
					}
				}

				releasedGroups = append(releasedGroups, &pbrelease.ReleaseStatus_ReleasedGroup{
					Id:          group.ID,
					Name:        group.Spec.Name,
					Mode:        gcr.Mode.String(),
					OldSelector: oldSelector,
					NewSelector: newSelector,
					Edited:      gcr.Edited,
				})
			}
		}
	}
	return releasedGroups
}

// GetReleaseByName get release by release name.
func (s *Service) GetReleaseByName(ctx context.Context, req *pbds.GetReleaseByNameReq) (*pbrelease.Release, error) {
	grpcKit := kit.FromGrpcContext(ctx)

	release, err := s.dao.Release().GetByName(grpcKit, req.GetBizId(), req.GetAppId(), req.GetReleaseName())
	if err != nil {
		logs.Errorf("get release by name failed, err: %v, rid: %s", err, grpcKit.Rid)
		return nil, fmt.Errorf("query release by name %s failed", req.GetReleaseName())
	}

	return pbrelease.PbRelease(release), nil
}

// GetRelease get release
func (s *Service) GetRelease(ctx context.Context, req *pbds.GetReleaseReq) (*pbrelease.Release, error) {
	grpcKit := kit.FromGrpcContext(ctx)

	release, err := s.dao.Release().Get(grpcKit, req.GetBizId(), req.GetAppId(), req.GetReleaseId())
	if err != nil {
		logs.Errorf("get release failed, err: %v, rid: %s", err, grpcKit.Rid)
		return nil, fmt.Errorf("query release %d failed", req.GetReleaseId())
	}

	return pbrelease.PbRelease(release), nil
}

// DeprecateRelease deprecate a release
func (s *Service) DeprecateRelease(ctx context.Context, req *pbds.DeprecateReleaseReq) (*pbbase.EmptyResp, error) {
	grpcKit := kit.FromGrpcContext(ctx)

	// check if release was published
	rgs, err := s.dao.ReleasedGroup().ListAllByReleaseID(grpcKit, req.ReleaseId, req.BizId)
	if err != nil {
		return nil, err
	}
	if len(rgs) > 0 {
		return nil, fmt.Errorf("release %d was published, can not deprecate", req.ReleaseId)
	}
	err = s.dao.Release().UpdateDeprecated(grpcKit, req.BizId, req.AppId, req.ReleaseId, true)
	return new(pbbase.EmptyResp), err
}

// UnDeprecateRelease undeprecate a release
func (s *Service) UnDeprecateRelease(ctx context.Context, req *pbds.UnDeprecateReleaseReq) (*pbbase.EmptyResp, error) {
	grpcKit := kit.FromGrpcContext(ctx)

	err := s.dao.Release().UpdateDeprecated(grpcKit, req.BizId, req.AppId, req.ReleaseId, false)
	return new(pbbase.EmptyResp), err
}

// DeleteRelease delete a release
func (s *Service) DeleteRelease(ctx context.Context, req *pbds.DeleteReleaseReq) (*pbbase.EmptyResp, error) {
	grpcKit := kit.FromGrpcContext(ctx)

	release, err := s.dao.Release().Get(grpcKit, req.BizId, req.AppId, req.ReleaseId)
	if err != nil {
		return nil, err
	}
	if !release.Spec.Deprecated {
		return nil, fmt.Errorf("release %d can not delete, you should deprecate it first", req.ReleaseId)
	}

	// get app type
	app, err := s.dao.App().Get(grpcKit, req.BizId, req.AppId)
	if err != nil {
		return nil, err
	}

	tx := s.dao.GenQuery().Begin()

	if err := s.deleteReleaseRelatedResources(grpcKit, tx, req.BizId, req.AppId, req.ReleaseId, app); err != nil {
		logs.Errorf("delete release related resources failed, err: %v, rid: %s", err, grpcKit.Rid)
		if rErr := tx.Rollback(); rErr != nil {
			logs.Errorf("transaction rollback failed, err: %v, rid: %s", rErr, grpcKit.Rid)
		}
		return nil, err
	}

	if err := s.dao.Release().DeleteWithTx(grpcKit, tx, req.BizId, req.AppId, req.ReleaseId); err != nil {
		logs.Errorf("delete release failed, err: %v, rid: %s", err, grpcKit.Rid)
		if rErr := tx.Rollback(); rErr != nil {
			logs.Errorf("transaction rollback failed, err: %v, rid: %s", rErr, grpcKit.Rid)
		}
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		logs.Errorf("commit transaction failed, err: %v, rid: %s", err, grpcKit.Rid)
		return nil, err
	}
	return new(pbbase.EmptyResp), nil
}

func (s *Service) queryPublishStatus(gcrs []*table.ReleasedGroup, releaseID uint32) (
	string, []*table.ReleasedGroup) {
	var includeDefault = false
	var inRelease = make([]*table.ReleasedGroup, 0)
	var outRelease = make([]*table.ReleasedGroup, 0)
	for _, gcr := range gcrs {
		if gcr.ReleaseID == releaseID {
			inRelease = append(inRelease, gcr)
			if gcr.GroupID == 0 {
				includeDefault = true
			}
		} else {
			outRelease = append(outRelease, gcr)
		}
	}

	// len(inRelease) == 0: not released
	if len(inRelease) == 0 {
		return table.NotReleased.String(), inRelease
		// len(inRelease) != 0 && len(outRelease) != 0: gray released
	} else if len(outRelease) != 0 {
		return table.PartialReleased.String(), inRelease
		// len(inRelease) != 0 && len(outRelease) == 0 && includeDefault: full released
	} else if includeDefault {
		return table.FullReleased.String(), inRelease
		// len(inRelease) != 0 && len(outRelease) == 0 && !includeDefault: gray released
	}
	return table.PartialReleased.String(), inRelease
}

// doKvOperations do kv related operations for create release.
func (s *Service) doKvOperations(kt *kit.Kit, tx *gen.QueryTx, appID, bizID, releaseID uint32) error {

	kvs, err := s.genCreateKv(kt, bizID, appID)
	if err != nil {
		return err
	}

	versionMap, err := s.doBatchReleasedVault(kt, kvs, releaseID)
	if err != nil {
		return err
	}

	rkvs, err := pbrkv.RKvs(kvs, versionMap, releaseID)
	if err != nil {
		return err
	}

	if err = s.dao.ReleasedKv().BulkCreateWithTx(kt, tx, rkvs); err != nil {
		logs.Errorf("bulk create released kv failed, err: %v, rid: %s", err, kt.Rid)
		return err
	}

	if err = s.cleanUpKV(kt, tx, bizID, appID); err != nil {
		logs.Errorf("clean failed, err: %v, rid: %s", err, kt.Rid)
		return err
	}

	return nil

}

func (s *Service) genCreateKv(kt *kit.Kit, bizID, appID uint32) ([]*pbkv.Kv, error) {

	kvState := []string{
		string(table.KvStateAdd),
		string(table.KvStateRevise),
		string(table.KvStateUnchange),
	}
	details, err := s.dao.Kv().ListAllByAppID(kt, appID, bizID, kvState)
	if err != nil {
		logs.Errorf("list kv failed, err: %v, rid: %s", err, kt.Rid)
		return nil, err
	}

	var kvs []*pbkv.Kv
	for _, detail := range details {
		var value string
		_, value, err = s.getKv(kt, bizID, appID, detail.Spec.Version, detail.Spec.Key)
		if err != nil {
			logs.Errorf("get vault kv failed, err: %v, rid: %s", err, kt.Rid)
			return nil, err
		}
		kvs = append(kvs, &pbkv.Kv{
			Spec:        pbkv.PbKvSpec(detail.Spec, value),
			Attachment:  pbkv.PbKvAttachment(detail.Attachment),
			Revision:    pbbase.PbRevision(detail.Revision),
			ContentSpec: pbcontent.PbContentSpec(detail.ContentSpec),
		})
	}

	return kvs, nil
}

func (s *Service) doBatchReleasedVault(kt *kit.Kit, kvs []*pbkv.Kv, releaseId uint32) (map[string]int, error) {

	versionMap := make(map[string]int, len(kvs))
	for _, rkv := range kvs {
		opt := &types.CreateReleasedKvOption{
			BizID:     rkv.Attachment.BizId,
			AppID:     rkv.Attachment.AppId,
			ReleaseID: releaseId,
			Key:       rkv.Spec.Key,
			Value:     rkv.Spec.Value,
			KvType:    table.DataType(rkv.Spec.KvType),
		}
		version, err := s.vault.CreateRKv(kt, opt)
		if err != nil {
			logs.Errorf("create vault kv failed, err: %v, rid: %s", err, kt.Rid)
			return nil, err
		}
		versionMap[rkv.Spec.Key] = version
	}

	return versionMap, nil

}

// cleanUpKV 创建版本后更新kv 状态
func (s *Service) cleanUpKV(kt *kit.Kit, tx *gen.QueryTx, bizID, appID uint32) error {

	result, err := s.dao.Kv().ListAllByAppID(kt, appID, bizID, []string{string(table.KvStateDelete)})
	if err != nil {
		logs.Errorf("delete kv failed, err: %v, rid: %s", err, kt.Rid)
		return err
	}

	if len(result) > 0 {
		deleteKv := &table.Kv{
			KvState: table.KvStateDelete,
			Attachment: &table.KvAttachment{
				BizID: bizID,
				AppID: appID,
			},
		}
		if e := s.dao.Kv().DeleteByStateWithTx(kt, tx, deleteKv); e != nil {
			logs.Errorf("delete kv failed, err: %v, rid: %s", e, kt.Rid)
			return e
		}
	}

	targetKVStates := []string{
		string(table.KvStateAdd),
		string(table.KvStateRevise),
	}
	if e := s.dao.Kv().UpdateSelectedKVStates(kt, tx, bizID, appID, targetKVStates, table.KvStateUnchange); e != nil {
		logs.Errorf("delete kv failed, err: %v, rid: %s", e, kt.Rid)
		return e
	}

	return nil
}

func (s *Service) deleteReleaseRelatedResources(grpcKit *kit.Kit, tx *gen.QueryTx, bizID, appID, releaseID uint32,
	app *table.App) error {
	switch app.Spec.ConfigType {
	case table.File:
		if err := s.dao.ReleasedAppTemplate().BatchDeleteByReleaseIDWithTx(grpcKit, tx,
			bizID, appID, releaseID); err != nil {
			return err
		}
		if err := s.dao.ReleasedAppTemplateVariable().BatchDeleteByReleaseIDWithTx(grpcKit, tx,
			bizID, appID, releaseID); err != nil {
			return err
		}
		if err := s.dao.ReleasedHook().BatchDeleteByReleaseIDWithTx(grpcKit, tx,
			bizID, appID, releaseID); err != nil {
			return err
		}
		if err := s.dao.ReleasedCI().BatchDeleteByReleaseIDWithTx(grpcKit, tx,
			bizID, appID, releaseID); err != nil {
			return err
		}
	case table.KV:
		if err := s.dao.ReleasedKv().BatchDeleteByReleaseIDWithTx(grpcKit, tx,
			bizID, appID, releaseID); err != nil {
			return err
		}
	}
	return nil
}

// CheckReleaseName 检测某个服务下已发布的名称是否存在
func (s *Service) CheckReleaseName(ctx context.Context, req *pbds.CheckReleaseNameReq) (
	*pbds.CheckReleaseNameResp, error) {
	grpcKit := kit.FromGrpcContext(ctx)

	result, err := s.dao.Release().GetByName(grpcKit, req.GetBizId(), req.GetAppId(), req.GetName())
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errf.Errorf(errf.DBOpFailed,
			i18n.T(grpcKit, fmt.Sprintf("get release by name failed, err: %s", err.Error())))
	}

	var exist bool
	if result != nil && result.ID != 0 {
		exist = true
	}

	return &pbds.CheckReleaseNameResp{
		Exist: exist,
	}, nil
}

// 检测非模板配置和模板配置文件是否存在冲突
// 生成版本的前置检测条件
func (s *Service) checkNonTmpAndTmpConflicts(kit *kit.Kit, bizID, appID uint32,
	comparisonFiles []string) (uint32, map[string]bool, error) {
	var (
		conflictNums  uint32
		conflictPaths map[string]bool
	)

	if len(comparisonFiles) == 0 {
		// 1. 查询非模板
		details, err := s.dao.ConfigItem().ListAllByAppID(kit, appID, bizID)
		if err != nil {
			logs.Errorf("list editing config items failed, err: %v, rid: %s", err, kit.Rid)
			return conflictNums, conflictPaths, err
		}

		for _, v := range details {
			comparisonFiles = append(comparisonFiles, path.Join(v.Spec.Path, v.Spec.Name))
		}
	}

	// 2. 获取服务关联的套餐
	binding, err := s.dao.AppTemplateBinding().GetAppTemplateBindingByAppID(kit, bizID, appID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return conflictNums, conflictPaths, errf.Newf(errf.DBOpFailed,
			i18n.T(kit, "get the package associated with the app failed, err: %v"), err)
	}
	// 未绑定
	if binding == nil {
		return conflictNums, conflictPaths, nil
	}
	// 存在记录，但没有绑定模板
	if len(binding.Spec.TemplateIDs) == 0 {
		return conflictNums, conflictPaths, nil
	}

	// 套餐和套餐之间也会存在冲突
	for _, v := range binding.Spec.Bindings {
		templateIDs := []uint32{}
		for _, v := range v.TemplateRevisions {
			templateIDs = append(templateIDs, v.TemplateID)
		}
		// 如有绑定，根据绑定的模板ID查询模板文件
		templates, err := s.dao.Template().ListByIDs(kit, templateIDs)
		if err != nil {
			return conflictNums, conflictPaths, errf.Newf(errf.DBOpFailed,
				i18n.T(kit, "get template file failed, err: %v"), err)
		}
		// 把模板文件加入需要对比的切片中
		for _, v := range templates {
			comparisonFiles = append(comparisonFiles, path.Join(v.Spec.Path, v.Spec.Name))
		}
	}

	// 检测文件路径冲突
	conflictNums, conflictPaths = tools.CheckExistingPathConflict(comparisonFiles)

	return conflictNums, conflictPaths, nil
}

// 检测是否存在过期证书
// 生成版本的前置检测条件
func (s Service) checkForExpiredCertificates(kit *kit.Kit, bizID, appID uint32) (int64, error) {
	// 0 代表过期
	_, expirationNumber, err := s.dao.Kv().FindNearExpiryCertKvs(kit, bizID, appID, 0, &types.BasePage{All: true})
	if err != nil {
		return expirationNumber, errf.Errorf(errf.DBOpFailed,
			i18n.T(kit, "get a list of expired certificates failed, err: %v"), err)
	}

	return expirationNumber, nil
}

// ListAllReleasedConfigItems implements pbds.DataServer.
func (s *Service) ListAllReleasedConfigItems(ctx context.Context, req *pbds.ListAllReleasedConfigItemsReq) (
	*pbds.ListAllReleasedConfigItemsResp, error) {
	grpcKit := kit.FromGrpcContext(ctx)
	// 获取服务信息
	app, err := s.dao.App().Get(grpcKit, req.BizId, req.AppId)
	if err != nil {
		return nil, err
	}
	// 通过服务获取已经上线的版本
	releasedGroup, err := s.dao.ReleasedGroup().ListAllByAppID(grpcKit, req.AppId, req.BizId)
	if err != nil {
		return nil, err
	}

	releaseIds := []uint32{}
	for _, v := range releasedGroup {
		releaseIds = append(releaseIds, v.ReleaseID)
	}

	items := make([]*pbds.ListAllReleasedConfigItemsResp_Item, 0)

	seen := map[string]bool{}

	// 获取已上线的配置项
	if app.Spec.ConfigType == table.File {
		releasedConfigItem, err := s.dao.ReleasedCI().ListAllByReleaseIDs(grpcKit, releaseIds, req.BizId)
		if err != nil {
			return nil, err
		}
		for _, v := range releasedConfigItem {
			if seen[v.CommitSpec.Content.Signature] {
				continue
			}
			items = append(items, &pbds.ListAllReleasedConfigItemsResp_Item{
				Name: path.Join(v.ConfigItemSpec.Path, v.ConfigItemSpec.Name),
				Sign: v.CommitSpec.Content.Signature,
			})
			seen[v.CommitSpec.Content.Signature] = true
		}
	} else {
		releasedKv, err := s.dao.ReleasedKv().ListAllByReleaseIDs(grpcKit, releaseIds, req.BizId)
		if err != nil {
			return nil, err
		}
		for _, v := range releasedKv {
			if seen[v.ContentSpec.Signature] {
				continue
			}
			items = append(items, &pbds.ListAllReleasedConfigItemsResp_Item{
				Name: v.Spec.Key,
				Sign: v.ContentSpec.Signature,
			})
			seen[v.ContentSpec.Signature] = true
		}
	}

	return &pbds.ListAllReleasedConfigItemsResp{
		Items: items,
	}, nil
}
