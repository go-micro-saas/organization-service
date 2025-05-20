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

func (s *orgDto) ToBoCreateOrgParam(req *resourcev1.CreateOrgReq) *bo.CreateOrgParam {
	res := &bo.CreateOrgParam{
		CreatorID:          req.CreatorId,
		CreatorName:        req.CreatorName,
		CreatorAvatar:      req.CreatorAvatar,
		OrgName:            req.OrgName,
		OrgAvatar:          req.OrgAvatar,
		OrgType:            req.OrgType,
		SkipCreateEmployee: false,
	}
	return res
}

func (s *orgDto) ToBoCreateOrgParam2(req *resourcev1.OnlyCreateOrgReq) *bo.CreateOrgParam {
	res := &bo.CreateOrgParam{
		CreatorID:          req.CreatorId,
		CreatorName:        "",
		CreatorAvatar:      "",
		OrgName:            req.OrgName,
		OrgAvatar:          req.OrgAvatar,
		OrgType:            req.OrgType,
		SkipCreateEmployee: true,
	}
	return res
}

func (s *orgDto) ToPbCreateOrgRespData(dataModel *bo.CreateOrgReply) *resourcev1.CreateOrgRespData {
	res := &resourcev1.CreateOrgRespData{
		OrgId:     dataModel.OrgId,
		OrgName:   dataModel.OrgName,
		OrgAvatar: dataModel.OrgAvatar,
		OrgType:   dataModel.OrgType,
	}
	return res
}

func (s *orgDto) ToBoAddEmployeeParam(req *resourcev1.AddEmployeeReq) *bo.AddEmployeeParam {
	res := &bo.AddEmployeeParam{
		OperatorUid:  req.OperatorUid,
		OrgId:        req.OrgId,
		UserId:       req.UserId,
		UserName:     req.UserName,
		EmployeeRole: req.EmployeeRole,
		UserAvatar:   req.UserAvatar,
		UserPhone:    req.UserPhone,
		UserEmail:    req.UserEmail,
	}
	return res
}

func (s *orgDto) ToPbAddEmployeeRespData(dataModel *bo.AddEmployeeReply) *resourcev1.AddEmployeeRespData {
	res := &resourcev1.AddEmployeeRespData{
		OrgId:          dataModel.OrgId,
		UserId:         dataModel.UserId,
		EmployeeId:     dataModel.EmployeeId,
		EmployeeName:   dataModel.EmployeeName,
		EmployeeAvatar: dataModel.EmployeeAvatar,
		EmployeeStatus: dataModel.EmployeeStatus,
		EmployeeRole:   dataModel.EmployeeRole,
	}
	return res
}
