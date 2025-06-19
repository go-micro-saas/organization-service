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

func (s *orgV1Service) CreateInviteRecordForLink(ctx context.Context, req *resourcev1.CreateInviteRecordForLinkReq) (*resourcev1.CreateInviteRecordForLinkResp, error) {
	param := dto.OrgDto.ToBoCreateInviteRecordForLinkParam(req)
	reply, err := s.orgBiz.CreateInviteRecordForLink(ctx, param)
	if err != nil {
		return nil, err
	}
	return &resourcev1.CreateInviteRecordForLinkResp{
		Data: dto.OrgDto.ToPbCreateInviteRecordForLinkRespData(reply),
	}, nil
}

func (s *orgV1Service) CreateInviteRecordForEmployee(ctx context.Context, req *resourcev1.CreateInviteRecordForEmployeeReq) (*resourcev1.CreateInviteRecordForEmployeeResp, error) {
	param := dto.OrgDto.ToBoCreateInviteRecordForEmployeeParam(req)
	reply, err := s.orgBiz.CreateInviteRecordForEmployee(ctx, param)
	if err != nil {
		return nil, err
	}
	return &resourcev1.CreateInviteRecordForEmployeeResp{
		Data: dto.OrgDto.ToPbCreateInviteRecordForEmployeeRespData(reply),
	}, nil
}

func (s *orgV1Service) JoinByInviteLink(ctx context.Context, req *resourcev1.JoinByInviteLinkReq) (*resourcev1.JoinByInviteLinkResp, error) {
	param := dto.OrgDto.ToBoJoinByInviteLinkParam(req)
	reply, err := s.orgBiz.JoinByInviteLink(ctx, param)
	if err != nil {
		return nil, err
	}
	return &resourcev1.JoinByInviteLinkResp{
		Data: dto.OrgDto.ToPbJoinByInviteLinkRespData(reply),
	}, nil
}

func (s *orgV1Service) AgreeOrRefuseInvite(ctx context.Context, req *resourcev1.AgreeOrRefuseInviteReq) (*resourcev1.AgreeOrRefuseInviteResp, error) {
	param := dto.OrgDto.ToBoAgreeOrRefuseInviteParam(req)
	reply, err := s.orgBiz.AgreeOrRefuseInvite(ctx, param)
	if err != nil {
		return nil, err
	}
	return &resourcev1.AgreeOrRefuseInviteResp{
		Data: dto.OrgDto.ToPbAgreeOrRefuseInviteRespData(reply),
	}, nil
}

func (s *orgV1Service) GetOrgInfo(ctx context.Context, req *resourcev1.GetOrgInfoReq) (*resourcev1.GetOrgInfoResp, error) {
	return s.UnimplementedSrvOrgV1Server.GetOrgInfo(ctx, req)
}

func (s *orgV1Service) GetOrgInfoList(ctx context.Context, req *resourcev1.GetOrgInfoListReq) (*resourcev1.GetOrgInfoListResp, error) {
	return s.UnimplementedSrvOrgV1Server.GetOrgInfoList(ctx, req)
}

func (s *orgV1Service) GetOrgEmployeeInfo(ctx context.Context, req *resourcev1.GetOrgEmployeeInfoReq) (*resourcev1.GetOrgEmployeeInfoResp, error) {
	return s.UnimplementedSrvOrgV1Server.GetOrgEmployeeInfo(ctx, req)
}

func (s *orgV1Service) GetOrgEmployeeInfoList(ctx context.Context, req *resourcev1.GetOrgEmployeeInfoListReq) (*resourcev1.GetOrgEmployeeInfoListResp, error) {
	return s.UnimplementedSrvOrgV1Server.GetOrgEmployeeInfoList(ctx, req)
}

func (s *orgV1Service) GetOrgList(ctx context.Context, req *resourcev1.GetOrgListReq) (*resourcev1.GetOrgListResp, error) {
	return s.UnimplementedSrvOrgV1Server.GetOrgList(ctx, req)
}

func (s *orgV1Service) GetOrgEmployeeList(ctx context.Context, req *resourcev1.GetOrgEmployeeListReq) (*resourcev1.GetOrgEmployeeListResp, error) {
	return s.UnimplementedSrvOrgV1Server.GetOrgEmployeeList(ctx, req)
}

func (s *orgV1Service) GetOrgInviteRecordList(ctx context.Context, req *resourcev1.GetOrgInviteRecordListReq) (*resourcev1.GetOrgInviteRecordListResp, error) {
	return s.UnimplementedSrvOrgV1Server.GetOrgInviteRecordList(ctx, req)
}
