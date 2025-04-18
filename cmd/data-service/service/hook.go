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

	"gorm.io/gorm"

	"github.com/TencentBlueKing/bk-bscp/internal/criteria/constant"
	"github.com/TencentBlueKing/bk-bscp/pkg/criteria/enumor"
	"github.com/TencentBlueKing/bk-bscp/pkg/criteria/errf"
	"github.com/TencentBlueKing/bk-bscp/pkg/dal/table"
	"github.com/TencentBlueKing/bk-bscp/pkg/i18n"
	"github.com/TencentBlueKing/bk-bscp/pkg/kit"
	"github.com/TencentBlueKing/bk-bscp/pkg/logs"
	pbbase "github.com/TencentBlueKing/bk-bscp/pkg/protocol/core/base"
	pbhook "github.com/TencentBlueKing/bk-bscp/pkg/protocol/core/hook"
	pbds "github.com/TencentBlueKing/bk-bscp/pkg/protocol/data-service"
	"github.com/TencentBlueKing/bk-bscp/pkg/tools"
	"github.com/TencentBlueKing/bk-bscp/pkg/types"
)

// CreateHook create hook.
func (s *Service) CreateHook(ctx context.Context, req *pbds.CreateHookReq) (*pbds.CreateResp, error) {
	kt := kit.FromGrpcContext(ctx)

	// GetByName get hook by name
	if _, err := s.dao.Hook().GetByName(kt, req.Attachment.BizId, req.Spec.Name); err == nil {
		return nil, errf.Errorf(errf.InvalidArgument, i18n.T(kt, "hook name %s already exists", req.Spec.Name))
	}

	res := &table.Revision{
		Creator: kt.User,
		Reviser: kt.User,
	}

	tx := s.dao.GenQuery().Begin()

	// 1. create hook
	hook := &table.Hook{
		Spec:       req.Spec.HookSpec(),
		Attachment: req.Attachment.HookAttachment(),
		Revision:   res,
	}

	// CreateWithTx create one hook instance with transaction.
	id, err := s.dao.Hook().CreateWithTx(kt, tx, hook, false)
	if err != nil {
		logs.Errorf("create hook failed, err: %v, rid: %s", err, kt.Rid)
		if rErr := tx.Rollback(); rErr != nil {
			logs.Errorf("transaction rollback failed, err: %v, rid: %s", rErr, kt.Rid)
		}
		return nil, err
	}

	// 2. create hook revision
	// it must be the first hook revision, so no need to check the revision name uniqueness
	if req.Spec.RevisionName == "" {
		req.Spec.RevisionName = tools.GenerateRevisionName()
	}
	revision := &table.HookRevision{
		Spec: &table.HookRevisionSpec{
			Name:    req.Spec.RevisionName,
			Content: req.Spec.Content,
			Memo:    req.Spec.Memo,
			State:   table.HookRevisionStatusDeployed,
		},
		Attachment: &table.HookRevisionAttachment{
			BizID:  req.Attachment.BizId,
			HookID: id,
		},
		Revision: res,
	}
	// CreateWithTx create hook revision instance with transaction.
	_, err = s.dao.HookRevision().CreateWithTx(kt, tx, revision)
	if err != nil {
		logs.Errorf("create hook revision failed, err: %v, rid: %s", err, kt.Rid)
		if rErr := tx.Rollback(); rErr != nil {
			logs.Errorf("transaction rollback failed, err: %v, rid: %s", rErr, kt.Rid)
		}
		return nil, err
	}

	err = s.dao.AuditDao().Decorator(kt, req.Attachment.BizId, &table.AuditField{
		ResourceInstance: fmt.Sprintf(constant.HookName+constant.ResSeparator+constant.HookRevisionName,
			hook.Spec.Name, revision.Spec.Name),
		Status: enumor.Success,
		Detail: revision.Spec.Memo,
	}).PrepareCreate(revision).Do(tx.Query)
	if err != nil {
		logs.Errorf("PrepareCreate HookRevision failed, err: %v, rid: %s", err, kt.Rid)
		if rErr := tx.Rollback(); rErr != nil {
			logs.Errorf("transaction rollback failed, err: %v, rid: %s", rErr, kt.Rid)
		}
		return nil, err
	}

	if e := tx.Commit(); e != nil {
		logs.Errorf("commit transaction failed, err: %v, rid: %s", e, kt.Rid)
		return nil, e
	}

	resp := &pbds.CreateResp{Id: id}
	return resp, nil
}

// ListHooks list hooks.
func (s *Service) ListHooks(ctx context.Context, req *pbds.ListHooksReq) (*pbds.ListHooksResp, error) {

	// FromGrpcContext used only to obtain Kit through grpc context.
	kt := kit.FromGrpcContext(ctx)

	page := &types.BasePage{Start: req.Start, Limit: uint(req.Limit)}
	opt := &types.ListHooksWithReferOption{
		BizID:     req.BizId,
		Name:      req.Name,
		Tag:       req.Tag,
		All:       req.All,
		NotTag:    req.NotTag,
		Page:      page,
		SearchKey: req.SearchKey,
	}

	po := &types.PageOption{
		EnableUnlimitedLimit: true,
	}
	if err := opt.Validate(po); err != nil {
		return nil, err
	}
	// StrToUint32Slice the comma separated string goes to uint32 slice
	topIds, _ := tools.StrToUint32Slice(req.TopIds)
	// ListWithRefer hooks with refer info.
	details, count, err := s.dao.Hook().ListWithRefer(kt, opt, topIds)
	if err != nil {
		logs.Errorf("list hook failed, err: %v, rid: %s", err, kt.Rid)
		return nil, err
	}

	uncitedCount, err := s.dao.Hook().CountNumberUnReferences(kt, req.BizId, opt)
	if err != nil {
		return nil, err
	}

	result := []*pbds.ListHooksResp_Detail{}
	for _, detail := range details {
		result = append(result, &pbds.ListHooksResp_Detail{
			Hook:                pbhook.PbHook(detail.Hook),
			BoundNum:            uint32(detail.ReferCount),
			ConfirmDelete:       detail.BoundEditingRelease,
			PublishedRevisionId: detail.PublishedRevisionID,
		})
	}

	resp := &pbds.ListHooksResp{
		Count:          uint32(count),
		Details:        result,
		ExclusionCount: uint32(uncitedCount),
	}

	return resp, nil
}

// DeleteHook delete hook.
func (s *Service) DeleteHook(ctx context.Context, req *pbds.DeleteHookReq) (*pbbase.EmptyResp, error) {
	kt := kit.FromGrpcContext(ctx)

	tx := s.dao.GenQuery().Begin()

	// 1. check if hook was bound to an editing release
	count, err := s.dao.ReleasedHook().CountByHookIDAndReleaseID(kt, req.BizId, req.HookId, 0)
	if err != nil {
		logs.Errorf("count hook bound editing releases failed, err: %v, rid: %s", err, kt.Rid)
		if rErr := tx.Rollback(); rErr != nil {
			logs.Errorf("transaction rollback failed, err: %v, rid: %s", rErr, kt.Rid)
		}
		return nil, err
	}
	if count > 0 && !req.Force {
		if rErr := tx.Rollback(); rErr != nil {
			logs.Errorf("transaction rollback failed, err: %v, rid: %s", rErr, kt.Rid)
		}
		return nil, fmt.Errorf("hook was bound to %d editing releases, "+
			"set force=true to delete hook with references, rid: %s", count, kt.Rid)
	}
	// 2. delete released hook that release_id = 0
	if e := s.dao.ReleasedHook().DeleteByHookIDAndReleaseIDWithTx(kt, tx, req.BizId, req.HookId, 0); e != nil {
		logs.Errorf("delete released hook failed, err: %v, rid: %s", e, kt.Rid)
		if rErr := tx.Rollback(); rErr != nil {
			logs.Errorf("transaction rollback failed, err: %v, rid: %s", rErr, kt.Rid)
		}
		return nil, e
	}
	// 3. delete all hook revisions by hook id
	if e := s.dao.HookRevision().DeleteByHookIDWithTx(kt, tx, req.HookId, req.BizId); e != nil {
		logs.Errorf("delete hook revision failed, err: %v, rid: %s", e, kt.Rid)
		if rErr := tx.Rollback(); rErr != nil {
			logs.Errorf("transaction rollback failed, err: %v, rid: %s", rErr, kt.Rid)
		}
		return nil, e
	}
	// 4. delete hook
	hook := &table.Hook{
		ID: req.HookId,
		Attachment: &table.HookAttachment{
			BizID: req.BizId,
		},
	}
	// DeleteWithTx delete hook instance with transaction.
	if e := s.dao.Hook().DeleteWithTx(kt, tx, hook); e != nil {
		logs.Errorf("delete hook failed, err: %v, rid: %s", e, kt.Rid)
		if rErr := tx.Rollback(); rErr != nil {
			logs.Errorf("transaction rollback failed, err: %v, rid: %s", rErr, kt.Rid)
		}
		return nil, e
	}

	if e := tx.Commit(); e != nil {
		logs.Errorf("commit transaction failed, err: %v, rid: %s", e, kt.Rid)
		return nil, e
	}
	return new(pbbase.EmptyResp), nil
}

// UpdateHook update hook.
func (s *Service) UpdateHook(ctx context.Context, req *pbds.UpdateHookReq) (*pbbase.EmptyResp, error) {
	kt := kit.FromGrpcContext(ctx)

	hook := &table.Hook{
		ID:         req.Id,
		Spec:       req.Spec.HookSpec(),
		Attachment: req.Attachment.HookAttachment(),
		Revision: &table.Revision{
			Reviser: kt.User,
		},
	}
	// Update one hook's info.
	if err := s.dao.Hook().Update(kt, hook); err != nil {
		logs.Errorf("update hook failed, err: %v, rid: %s", err, kt.Rid)
		return nil, err
	}

	return new(pbbase.EmptyResp), nil
}

// ListHookTags list tag
func (s *Service) ListHookTags(ctx context.Context, req *pbds.ListHookTagReq) (*pbds.ListHookTagResp, error) {

	// FromGrpcContext used only to obtain Kit through grpc context.
	kt := kit.FromGrpcContext(ctx)

	// CountHookTag count hook tag
	ht, err := s.dao.Hook().CountHookTag(kt, req.BizId)
	if err != nil {
		logs.Errorf("list hook failed, err: %v, rid: %s", err, kt.Rid)
		return nil, err
	}

	resp := &pbds.ListHookTagResp{}

	for _, count := range ht {
		resp.Details = append(resp.Details, &pbhook.CountHookTags{
			Tag:    count.Tag,
			Counts: count.Counts,
		})
	}

	return resp, nil
}

// GetHook get a hook
func (s *Service) GetHook(ctx context.Context, req *pbds.GetHookReq) (*pbds.GetHookResp, error) {

	// FromGrpcContext used only to obtain Kit through grpc context.
	kt := kit.FromGrpcContext(ctx)

	h, err := s.dao.Hook().GetByID(kt, req.BizId, req.HookId)
	if err != nil {
		logs.Errorf("list hook failed, err: %v, rid: %s", err, kt.Rid)
		return nil, err
	}

	opt := &types.GetByPubStateOption{
		BizID:  req.BizId,
		HookID: req.HookId,
		State:  table.HookRevisionStatusNotDeployed,
	}
	var revisionID uint32
	revision, err := s.dao.HookRevision().GetByPubState(kt, opt)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		logs.Errorf("get hook failed, err: %v, rid: %s", err, kt.Rid)
		return nil, err
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		revisionID = 0
	} else {
		revisionID = revision.ID
	}

	resp := &pbds.GetHookResp{
		Id: h.ID,
		Spec: &pbds.GetHookInfoSpec{
			Name:     h.Spec.Name,
			Type:     string(h.Spec.Type),
			Tags:     h.Spec.Tags,
			Memo:     h.Spec.Memo,
			Releases: &pbds.GetHookInfoSpec_Releases{NotReleaseId: revisionID},
		},
		Attachment: pbhook.PbHookAttachment(h.Attachment),
		Revision:   pbbase.PbRevision(h.Revision),
	}

	return resp, nil
}

// ListHookReferences ..
func (s *Service) ListHookReferences(ctx context.Context,
	req *pbds.ListHookReferencesReq) (*pbds.ListHookReferencesResp, error) {

	kt := kit.FromGrpcContext(ctx)

	page := &types.BasePage{Start: req.Start, Limit: uint(req.Limit)}
	opt := &types.ListHookReferencesOption{
		BizID:     req.BizId,
		HookID:    req.HookId,
		Page:      page,
		SearchKey: req.SearchKey,
	}
	if err := opt.Validate(types.DefaultPageOption); err != nil {
		return nil, err
	}

	results, count, err := s.dao.Hook().ListHookReferences(kt, opt)
	if err != nil {
		logs.Errorf("list hook references failed, err: %v, rid: %s", err, kt.Rid)
		return nil, err
	}

	details := make([]*pbds.ListHookReferencesResp_Detail, 0, len(results))
	for _, result := range results {
		details = append(details, &pbds.ListHookReferencesResp_Detail{
			HookRevisionId:   result.HookRevisionID,
			HookRevisionName: result.HookRevisionName,
			AppId:            result.AppID,
			AppName:          result.AppName,
			ReleaseId:        result.ReleaseID,
			ReleaseName:      result.ReleaseName,
			Type:             result.HookType,
			Deprecated:       result.Deprecated,
		})
	}

	resp := &pbds.ListHookReferencesResp{
		Count:   uint32(count),
		Details: details,
	}

	return resp, nil
}

// GetReleaseHook get release's pre hook and post hook
func (s *Service) GetReleaseHook(ctx context.Context, req *pbds.GetReleaseHookReq) (*pbds.GetReleaseHookResp, error) {

	grpcKit := kit.FromGrpcContext(ctx)

	preHook, err := s.dao.ReleasedHook().Get(grpcKit, req.BizId, req.AppId, req.ReleaseId, table.PreHook)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		logs.Errorf("get pre hook failed, err: %v, rid: %s", err, grpcKit.Rid)
		return nil, err
	}
	postHook, err := s.dao.ReleasedHook().Get(grpcKit, req.BizId, req.AppId, req.ReleaseId, table.PostHook)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		logs.Errorf("get post hook failed, err: %v, rid: %s", err, grpcKit.Rid)
		return nil, err
	}
	var pre, post *pbds.GetReleaseHookResp_Hook
	if preHook != nil {
		pre = &pbds.GetReleaseHookResp_Hook{
			HookId:           preHook.HookID,
			HookName:         preHook.HookName,
			HookRevisionId:   preHook.HookRevisionID,
			HookRevisionName: preHook.HookRevisionName,
			Type:             preHook.ScriptType.String(),
			Content:          preHook.Content,
		}
	}
	if postHook != nil {
		post = &pbds.GetReleaseHookResp_Hook{
			HookId:           postHook.HookID,
			HookName:         postHook.HookName,
			HookRevisionId:   postHook.HookRevisionID,
			HookRevisionName: postHook.HookRevisionName,
			Type:             postHook.ScriptType.String(),
			Content:          postHook.Content,
		}
	}
	resp := &pbds.GetReleaseHookResp{
		PreHook:  pre,
		PostHook: post,
	}

	return resp, nil
}

// HookFetchIDsExcluding 获取指定ID后排除的ID
func (s *Service) HookFetchIDsExcluding(ctx context.Context, req *pbds.HookFetchIDsExcludingReq) (
	*pbds.HookFetchIDsExcludingResp, error) {
	grpcKit := kit.FromGrpcContext(ctx)

	ids, err := s.dao.Hook().FetchIDsExcluding(grpcKit, req.BizId, req.GetIds())
	if err != nil {
		return nil, errf.Errorf(errf.DBOpFailed, i18n.T(grpcKit, "get excluded hook failed, err: %s", err))
	}

	return &pbds.HookFetchIDsExcludingResp{
		Ids: ids,
	}, nil
}

// GetHookReferencedIDs 获取脚本被引用的IDs
func (s *Service) GetHookReferencedIDs(ctx context.Context, req *pbds.GetHookReferencedIDsReq) (
	*pbds.GetHookReferencedIDsResp, error) {
	grpcKit := kit.FromGrpcContext(ctx)

	ids, err := s.dao.Hook().GetReferencedIDs(grpcKit, req.BizId)
	if err != nil {
		return nil, errf.Errorf(errf.DBOpFailed, i18n.T(grpcKit, "retrieve the referenced script failed, err: %s", err))
	}

	return &pbds.GetHookReferencedIDsResp{
		Ids: ids,
	}, nil
}
