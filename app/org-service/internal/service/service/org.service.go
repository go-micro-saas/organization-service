package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	resourcev1 "github.com/go-micro-saas/organization-service/api/org-service/v1/resources"
	servicev1 "github.com/go-micro-saas/organization-service/api/org-service/v1/services"
	bizrepos "github.com/go-micro-saas/organization-service/app/org-service/internal/biz/repo"
	"github.com/go-micro-saas/organization-service/app/org-service/internal/service/dto"
	gormpkg "github.com/ikaiguang/go-srv-kit/data/gorm"
	pagepkg "github.com/ikaiguang/go-srv-kit/kit/page"
)

type orgV1Service struct {
	servicev1.UnimplementedSrvOrgV1Server

	log    *log.Helper
	orgBiz bizrepos.OrgBizRepo
}

func NewOrgV1Service(logger log.Logger, orgBiz bizrepos.OrgBizRepo) servicev1.SrvOrgV1Server {
	logHelper := log.NewHelper(log.With(logger, "module", "test-service/service/service"))
	return &orgV1Service{
		log:    logHelper,
		orgBiz: orgBiz,
	}
}

func (s *orgV1Service) Ping(ctx context.Context, req *resourcev1.PingReq) (*resourcev1.PingResp, error) {
	return &resourcev1.PingResp{
		Data: dto.OrgDto.ToPbTestRespData(req.Message),
	}, nil
}

func (s *orgV1Service) CreateOrg(ctx context.Context, req *resourcev1.CreateOrgReq) (*resourcev1.CreateOrgResp, error) {
	param := dto.OrgDto.ToBoCreateOrgParam(req)
	reply, err := s.orgBiz.CreateOrg(ctx, param)
	if err != nil {
		return nil, err
	}
	return &resourcev1.CreateOrgResp{
		Data: dto.OrgDto.ToPbCreateOrgRespData(reply),
	}, nil
}

func (s *orgV1Service) OnlyCreateOrg(ctx context.Context, req *resourcev1.OnlyCreateOrgReq) (*resourcev1.CreateOrgResp, error) {
	param := dto.OrgDto.ToBoCreateOrgParam2(req)
	reply, err := s.orgBiz.CreateOrg(ctx, param)
	if err != nil {
		return nil, err
	}
	return &resourcev1.CreateOrgResp{
		Data: dto.OrgDto.ToPbCreateOrgRespData(reply),
	}, nil
}

func (s *orgV1Service) SetOrgStatus(ctx context.Context, req *resourcev1.SetOrgStatusReq) (*resourcev1.SetOrgStatusResp, error) {
	param := dto.OrgDto.ToBoSetOrgStatusParam(req)
	reply, err := s.orgBiz.SetOrgStatus(ctx, param)
	if err != nil {
		return nil, err
	}
	return &resourcev1.SetOrgStatusResp{
		Data: dto.OrgDto.ToPbSetOrgStatusRespData(reply),
	}, nil
}

func (s *orgV1Service) GetOrgInfo(ctx context.Context, req *resourcev1.GetOrgInfoReq) (*resourcev1.GetOrgInfoResp, error) {
	dataModel, err := s.orgBiz.GetOrgInfo(ctx, req.OrgId)
	if err != nil {
		return nil, err
	}
	return &resourcev1.GetOrgInfoResp{
		Data: dto.OrgDto.ToPbOrg(dataModel),
	}, nil
}

func (s *orgV1Service) GetOrgInfoList(ctx context.Context, req *resourcev1.GetOrgInfoListReq) (*resourcev1.GetOrgInfoListResp, error) {
	dataModels, err := s.orgBiz.GetOrgInfoList(ctx, req.OrgIds)
	if err != nil {
		return nil, err
	}
	return &resourcev1.GetOrgInfoListResp{
		Data: dto.OrgDto.ToPbOrgList(dataModels),
	}, nil
}

func (s *orgV1Service) GetOrgList(ctx context.Context, req *resourcev1.GetOrgListReq) (*resourcev1.GetOrgListResp, error) {
	// paging args
	pageRequest, pageOption := pagepkg.ParsePageRequest(req.PageRequest)

	// list
	param := dto.OrgDto.ToBoOrgListParam(req)
	param.PaginatorArgs = &gormpkg.PaginatorArgs{
		PageRequest: pageRequest,
		PageOption:  pageOption,
	}
	dataModels, dataCount, err := s.orgBiz.ListOrg(ctx, param)
	if err != nil {
		return nil, err
	}

	// paging result
	return &resourcev1.GetOrgListResp{
		Data: &resourcev1.GetOrgListRespData{
			List:     dto.OrgDto.ToPbOrgList(dataModels),
			PageInfo: pagepkg.CalcPageResponse(pageRequest, uint32(dataCount)),
		},
	}, nil
}

func (s *orgV1Service) GetUserLastOrg(ctx context.Context, req *resourcev1.GetUserLastOrgReq) (*resourcev1.GetUserLastOrgResp, error) {
	param := dto.OrgDto.ToBoGetUserLastOrgParam(req)
	reply, err := s.orgBiz.GetUserLastOrg(ctx, param)
	if err != nil {
		return nil, err
	}
	return &resourcev1.GetUserLastOrgResp{
		Data: dto.OrgDto.ToPbCreateOrgRespData(reply),
	}, nil
}
