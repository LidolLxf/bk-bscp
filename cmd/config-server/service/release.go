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

	"github.com/TencentBlueKing/bk-bscp/pkg/iam/meta"
	"github.com/TencentBlueKing/bk-bscp/pkg/kit"
	"github.com/TencentBlueKing/bk-bscp/pkg/logs"
	pbcs "github.com/TencentBlueKing/bk-bscp/pkg/protocol/config-server"
	pbrelease "github.com/TencentBlueKing/bk-bscp/pkg/protocol/core/release"
	pbds "github.com/TencentBlueKing/bk-bscp/pkg/protocol/data-service"
)

// CreateRelease create a release
func (s *Service) CreateRelease(ctx context.Context, req *pbcs.CreateReleaseReq) (*pbcs.CreateReleaseResp, error) {
	grpcKit := kit.FromGrpcContext(ctx)
	// the url path doesn't include appID, set the appID to use later
	grpcKit.AppID = req.AppId

	res := []*meta.ResourceAttribute{
		{Basic: meta.Basic{Type: meta.Biz, Action: meta.FindBusinessResource}, BizID: req.BizId},
		{Basic: meta.Basic{Type: meta.App, Action: meta.GenerateRelease, ResourceID: req.AppId}, BizID: req.BizId},
	}
	err := s.authorizer.Authorize(grpcKit, res...)
	if err != nil {
		return nil, err
	}

	r := &pbds.CreateReleaseReq{
		Attachment: &pbrelease.ReleaseAttachment{
			BizId: req.BizId,
			AppId: req.AppId,
		},
		Spec: &pbrelease.ReleaseSpec{
			Name: req.Name,
			Memo: req.Memo,
		},
		Variables: req.Variables,
	}
	rp, err := s.client.DS.CreateRelease(grpcKit.RpcCtx(), r)
	if err != nil {
		logs.Errorf("create release failed, err: %v, rid: %s", err, grpcKit.Rid)
		return nil, err
	}

	resp := &pbcs.CreateReleaseResp{
		Id: rp.Id,
	}
	return resp, nil
}

// ListReleases list releases with options
func (s *Service) ListReleases(ctx context.Context, req *pbcs.ListReleasesReq) (*pbcs.ListReleasesResp, error) {
	grpcKit := kit.FromGrpcContext(ctx)

	res := []*meta.ResourceAttribute{
		{Basic: meta.Basic{Type: meta.Biz, Action: meta.FindBusinessResource}, BizID: req.BizId},
		{Basic: meta.Basic{Type: meta.App, Action: meta.View, ResourceID: req.AppId}, BizID: req.BizId},
	}
	if err := s.authorizer.Authorize(grpcKit, res...); err != nil {
		return nil, err
	}

	r := &pbds.ListReleasesReq{
		BizId:      grpcKit.BizID,
		AppId:      req.AppId,
		Deprecated: req.Deprecated,
		Start:      req.Start,
		Limit:      req.Limit,
		All:        req.All,
		SearchKey:  req.SearchKey,
	}
	rp, err := s.client.DS.ListReleases(grpcKit.RpcCtx(), r)
	if err != nil {
		logs.Errorf("list releases failed, err: %v, rid: %s", err, grpcKit.Rid)
		return nil, err
	}

	resp := &pbcs.ListReleasesResp{
		Count:   rp.Count,
		Details: rp.Details,
	}
	return resp, nil
}

// GetReleaseByName get release by name
func (s *Service) GetReleaseByName(ctx context.Context, req *pbcs.GetReleaseByNameReq) (*pbrelease.Release, error) {
	kt := kit.FromGrpcContext(ctx)

	res := []*meta.ResourceAttribute{
		{Basic: meta.Basic{Type: meta.Biz, Action: meta.FindBusinessResource}, BizID: req.BizId},
		{Basic: meta.Basic{Type: meta.App, Action: meta.View, ResourceID: req.AppId}, BizID: req.BizId},
	}
	err := s.authorizer.Authorize(kt, res...)
	if err != nil {
		return nil, err
	}

	r := &pbds.GetReleaseByNameReq{
		BizId:       req.BizId,
		AppId:       req.AppId,
		ReleaseName: req.ReleaseName,
	}
	rp, err := s.client.DS.GetReleaseByName(kt.RpcCtx(), r)
	if err != nil {
		logs.Errorf("get release by name %s failed, err: %v, rid: %s", req.ReleaseName, err, kt.Rid)
		return nil, err
	}

	return rp, nil
}

// GetRelease get release
func (s *Service) GetRelease(ctx context.Context, req *pbcs.GetReleaseReq) (*pbrelease.Release, error) {
	kt := kit.FromGrpcContext(ctx)

	res := []*meta.ResourceAttribute{
		{Basic: meta.Basic{Type: meta.Biz, Action: meta.FindBusinessResource}, BizID: req.BizId},
		{Basic: meta.Basic{Type: meta.App, Action: meta.View, ResourceID: req.AppId}, BizID: req.BizId},
	}
	err := s.authorizer.Authorize(kt, res...)
	if err != nil {
		return nil, err
	}

	r := &pbds.GetReleaseReq{
		BizId:     req.BizId,
		AppId:     req.AppId,
		ReleaseId: req.GetReleaseId(),
	}
	rp, err := s.client.DS.GetRelease(kt.RpcCtx(), r)
	if err != nil {
		logs.Errorf("get release %d failed, err: %v, rid: %s", req.GetReleaseId(), err, kt.Rid)
		return nil, err
	}

	return rp, nil
}

// DeprecateRelease deprecate a release
func (s *Service) DeprecateRelease(ctx context.Context, req *pbcs.DeprecateReleaseReq) (
	*pbcs.DeprecateReleaseResp, error) {
	kt := kit.FromGrpcContext(ctx)

	res := []*meta.ResourceAttribute{
		{Basic: meta.Basic{Type: meta.Biz, Action: meta.FindBusinessResource}, BizID: req.BizId},
		{Basic: meta.Basic{Type: meta.App, Action: meta.Update, ResourceID: req.AppId}, BizID: req.BizId},
	}
	err := s.authorizer.Authorize(kt, res...)
	if err != nil {
		return nil, err
	}

	r := &pbds.DeprecateReleaseReq{
		BizId:     req.BizId,
		AppId:     req.AppId,
		ReleaseId: req.ReleaseId,
	}
	_, err = s.client.DS.DeprecateRelease(kt.RpcCtx(), r)
	if err != nil {
		logs.Errorf("deprecate release %s failed, err: %v, rid: %s", req.ReleaseId, err, kt.Rid)
		return nil, err
	}

	resp := &pbcs.DeprecateReleaseResp{}
	return resp, nil
}

// UnDeprecateRelease undeprecate a release
func (s *Service) UnDeprecateRelease(ctx context.Context, req *pbcs.UnDeprecateReleaseReq) (
	*pbcs.UnDeprecateReleaseResp, error) {
	kt := kit.FromGrpcContext(ctx)

	res := []*meta.ResourceAttribute{
		{Basic: meta.Basic{Type: meta.Biz, Action: meta.FindBusinessResource}, BizID: req.BizId},
		{Basic: meta.Basic{Type: meta.App, Action: meta.Update, ResourceID: req.AppId}, BizID: req.BizId},
	}
	err := s.authorizer.Authorize(kt, res...)
	if err != nil {
		return nil, err
	}

	r := &pbds.UnDeprecateReleaseReq{
		BizId:     req.BizId,
		AppId:     req.AppId,
		ReleaseId: req.ReleaseId,
	}
	_, err = s.client.DS.UnDeprecateRelease(kt.RpcCtx(), r)
	if err != nil {
		logs.Errorf("undeprecate release %s failed, err: %v, rid: %s", req.ReleaseId, err, kt.Rid)
		return nil, err
	}

	resp := &pbcs.UnDeprecateReleaseResp{}
	return resp, nil
}

// DeleteRelease delete a release
func (s *Service) DeleteRelease(ctx context.Context, req *pbcs.DeleteReleaseReq) (*pbcs.DeleteReleaseResp, error) {
	kt := kit.FromGrpcContext(ctx)

	res := []*meta.ResourceAttribute{
		{Basic: meta.Basic{Type: meta.Biz, Action: meta.FindBusinessResource}, BizID: req.BizId},
		{Basic: meta.Basic{Type: meta.App, Action: meta.Update, ResourceID: req.AppId}, BizID: req.BizId},
	}
	err := s.authorizer.Authorize(kt, res...)
	if err != nil {
		return nil, err
	}

	r := &pbds.DeleteReleaseReq{
		BizId:     req.BizId,
		AppId:     req.AppId,
		ReleaseId: req.ReleaseId,
	}
	_, err = s.client.DS.DeleteRelease(kt.RpcCtx(), r)
	if err != nil {
		logs.Errorf("delete release %s failed, err: %v, rid: %s", req.ReleaseId, err, kt.Rid)
		return nil, err
	}

	resp := &pbcs.DeleteReleaseResp{}
	return resp, nil
}

// CheckReleaseName 检测某个服务下已发布的名称是否存在
func (s *Service) CheckReleaseName(ctx context.Context, req *pbcs.CheckReleaseNameReq) (
	*pbcs.CheckReleaseNameResp, error) {

	kt := kit.FromGrpcContext(ctx)

	res := []*meta.ResourceAttribute{
		{Basic: meta.Basic{Type: meta.Biz, Action: meta.FindBusinessResource}, BizID: req.BizId},
		{Basic: meta.Basic{Type: meta.App, Action: meta.Find, ResourceID: req.AppId}, BizID: req.BizId},
	}
	err := s.authorizer.Authorize(kt, res...)
	if err != nil {
		return nil, err
	}

	result, err := s.client.DS.CheckReleaseName(kt.RpcCtx(), &pbds.CheckReleaseNameReq{
		BizId: req.GetBizId(),
		AppId: req.GetAppId(),
		Name:  req.GetName(),
	})
	if err != nil {
		return nil, err
	}

	return &pbcs.CheckReleaseNameResp{Exist: result.GetExist()}, nil
}

// ListAllReleasedConfigItems implements pbcs.ConfigServer.
func (s *Service) ListAllReleasedConfigItems(ctx context.Context, req *pbcs.ListAllReleasedConfigItemsReq) (
	*pbcs.ListAllReleasedConfigItemsResp, error) {
	kt := kit.FromGrpcContext(ctx)

	res := []*meta.ResourceAttribute{
		{Basic: meta.Basic{Type: meta.Biz, Action: meta.FindBusinessResource}, BizID: req.BizId},
		{Basic: meta.Basic{Type: meta.App, Action: meta.Find, ResourceID: req.AppId}, BizID: req.BizId},
	}
	err := s.authorizer.Authorize(kt, res...)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.DS.ListAllReleasedConfigItems(kt.RpcCtx(), &pbds.ListAllReleasedConfigItemsReq{
		BizId: req.BizId,
		AppId: req.AppId,
	})
	if err != nil {
		return nil, err
	}

	items := make([]*pbcs.ListAllReleasedConfigItemsResp_Item, 0)

	for _, v := range resp.GetItems() {
		items = append(items, &pbcs.ListAllReleasedConfigItemsResp_Item{
			Name: v.Name,
			Sign: v.Sign,
		})
	}

	return &pbcs.ListAllReleasedConfigItemsResp{
		Items: items,
	}, nil
}
