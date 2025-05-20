package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	resourcev1 "github.com/go-micro-saas/organization-service/api/org-service/v1/resources"
	servicev1 "github.com/go-micro-saas/organization-service/api/org-service/v1/services"
	bizrepos "github.com/go-micro-saas/organization-service/app/org-service/internal/biz/repo"
	"github.com/go-micro-saas/organization-service/app/org-service/internal/service/dto"
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

func (s *orgV1Service) AddEmployee(ctx context.Context, req *resourcev1.AddEmployeeReq) (*resourcev1.AddEmployeeResp, error) {
	param := dto.OrgDto.ToBoAddEmployeeParam(req)
	reply, err := s.orgBiz.AddEmployee(ctx, param)
	if err != nil {
		return nil, err
	}
	return &resourcev1.AddEmployeeResp{
		Data: dto.OrgDto.ToPbAddEmployeeRespData(reply),
	}, nil
}

func (s *orgV1Service) GenerateInviteLink(ctx context.Context, req *resourcev1.GenerateInviteLinkReq) (*resourcev1.GenerateInviteLinkResp, error) {
	return s.UnimplementedSrvOrgV1Server.GenerateInviteLink(ctx, req)
}

func (s *orgV1Service) CreateInviteEmployee(ctx context.Context, req *resourcev1.CreateInviteEmployeeReq) (*resourcev1.CreateInviteEmployeeResp, error) {
	return s.UnimplementedSrvOrgV1Server.CreateInviteEmployee(ctx, req)
}
