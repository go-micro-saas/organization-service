package dto

import (
	resourcev1 "github.com/go-micro-saas/organization-service/api/org-service/v1/resources"
	"github.com/go-micro-saas/organization-service/app/org-service/internal/biz/bo"
)

var (
	OrgDto orgDto
)

type orgDto struct{}

func (s *orgDto) ToBoHelloWorldParam(req *resourcev1.PingReq) *bo.HelloWorldParam {
	res := &bo.HelloWorldParam{
		Message: req.Message,
	}
	return res
}

func (s *orgDto) ToPbTestRespData(msg string) *resourcev1.PingRespData {
	res := &resourcev1.PingRespData{
		Message: msg,
	}

	return res
}
