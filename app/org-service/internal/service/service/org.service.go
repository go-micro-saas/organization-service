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
	param := dto.OrgDto.ToBoHelloWorldParam(req)
	reply, err := s.orgBiz.HelloWorld(ctx, param)
	if err != nil {
		return nil, err
	}
	return &resourcev1.PingResp{
		Data: dto.OrgDto.ToPbTestRespData(reply.Message),
	}, nil
}
