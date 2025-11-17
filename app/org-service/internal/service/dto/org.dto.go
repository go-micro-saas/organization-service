package dto

import (
	resourcev1 "github.com/go-micro-saas/organization-service/api/org-service/v1/resources"
	"github.com/go-micro-saas/organization-service/app/org-service/internal/biz/bo"
	"github.com/go-micro-saas/organization-service/app/org-service/internal/data/po"
	timepkg "github.com/ikaiguang/go-srv-kit/kit/time"
	"google.golang.org/protobuf/types/known/timestamppb"
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
		OrgStatus: dataModel.OrgStatus,
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

func (s *orgDto) ToBoCreateInviteRecordForLinkParam(req *resourcev1.CreateInviteRecordForLinkReq) *bo.CreateInviteRecordForLinkParam {
	res := &bo.CreateInviteRecordForLinkParam{
		OperatorUid:        req.OperatorUid,
		OrgId:              req.OrgId,
		ExpireTime:         req.ExpireTime.AsTime(),
		InviteEmployeeRole: req.InviteEmployeeRole,
	}
	return res
}

func (s *orgDto) ToPbCreateInviteRecordForLinkRespData(dataModel *po.OrgInviteRecord) *resourcev1.CreateInviteRecordForLinkRespData {
	res := &resourcev1.CreateInviteRecordForLinkRespData{
		InviteId:   dataModel.InviteId,
		InviteCode: dataModel.InviteCode,
		ExpireTime: timestamppb.New(dataModel.ExpireTime),
	}
	return res
}

func (s *orgDto) ToBoCreateInviteRecordForEmployeeParam(req *resourcev1.CreateInviteRecordForEmployeeReq) *bo.CreateInviteRecordForEmployeeParam {
	res := &bo.CreateInviteRecordForEmployeeParam{
		OperatorUid:        req.OperatorUid,
		OrgId:              req.OrgId,
		ExpireTime:         req.ExpireTime.AsTime(),
		InviteUserId:       req.InviteUserId,
		InviteAccount:      req.InviteAccount,
		InviteAccountType:  req.InviteAccountType,
		InviteEmployeeRole: req.InviteEmployeeRole,
	}
	return res
}

func (s *orgDto) ToPbCreateInviteRecordForEmployeeRespData(dataModel *po.OrgInviteRecord) *resourcev1.CreateInviteRecordForEmployeeRespData {
	res := &resourcev1.CreateInviteRecordForEmployeeRespData{
		InviteId:         dataModel.InviteId,
		InviteCode:       dataModel.InviteCode,
		ExpireTime:       timestamppb.New(dataModel.ExpireTime),
		AssignEmployeeId: dataModel.AssignEmployeeId,
	}
	return res
}

func (s *orgDto) ToBoJoinByInviteLinkParam(req *resourcev1.JoinByInviteLinkReq) *bo.JoinByInviteLinkParam {
	res := &bo.JoinByInviteLinkParam{
		InviteId:   req.InviteId,
		InviteCode: req.GetInviteCode(),
		UserId:     req.UserId,
		UserName:   req.UserName,
		UserAvatar: req.UserAvatar,
		UserPhone:  req.UserPhone,
		UserEmail:  req.UserEmail,
	}
	return res
}

func (s *orgDto) ToPbJoinByInviteLinkRespData(dataModel *po.OrgEmployee) *resourcev1.JoinByInviteLinkRespData {
	res := &resourcev1.JoinByInviteLinkRespData{
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

func (s *orgDto) ToBoAgreeOrRefuseInviteParam(req *resourcev1.AgreeOrRefuseInviteReq) *bo.AgreeOrRefuseInviteParam {
	res := &bo.AgreeOrRefuseInviteParam{
		InviteId:   req.InviteId,
		InviteCode: req.GetInviteCode(),
		IsAgree:    req.IsAgree,
		UserId:     req.UserId,
		UserName:   req.UserName,
		UserAvatar: req.UserAvatar,
		UserPhone:  req.UserPhone,
		UserEmail:  req.UserEmail,
	}
	return res
}

func (s *orgDto) ToPbAgreeOrRefuseInviteRespData(dataModel *po.OrgEmployee) *resourcev1.AgreeOrRefuseInviteRespData {
	res := &resourcev1.AgreeOrRefuseInviteRespData{
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

func (s *orgDto) ToPbOrgList(dataModels []*po.Org) []*resourcev1.Org {
	res := make([]*resourcev1.Org, 0, len(dataModels))
	for i := range dataModels {
		res = append(res, s.ToPbOrg(dataModels[i]))
	}
	return res
}

func (s *orgDto) ToPbOrg(dataModel *po.Org) *resourcev1.Org {
	res := &resourcev1.Org{
		Id:              dataModel.Id,
		CreatedTime:     dataModel.CreatedTime.Format(timepkg.YmdHms),
		UpdatedTime:     dataModel.UpdatedTime.Format(timepkg.YmdHms),
		DeletedTime:     dataModel.DeletedTime,
		OrgId:           dataModel.OrgId,
		OrgName:         dataModel.OrgName,
		OrgAvatar:       dataModel.OrgAvatar,
		OrgContactName:  dataModel.OrgContactName,
		OrgContactPhone: dataModel.OrgContactPhone,
		OrgContactEmail: dataModel.OrgContactEmail,
		OrgType:         dataModel.OrgType,
		OrgStatus:       dataModel.OrgStatus,
		OrgIndustryId:   dataModel.OrgIndustryId,
		OrgScaleId:      dataModel.OrgScaleId,
		OrgAddress:      dataModel.OrgAddress,
		OrgZipCode:      dataModel.OrgZipCode,
		OrgCreatorId:    dataModel.OrgCreatorId,
	}
	return res
}

func (s *orgDto) ToPbOrgEmployeeList(dataModels []*po.OrgEmployee) []*resourcev1.OrgEmployee {
	res := make([]*resourcev1.OrgEmployee, 0, len(dataModels))
	for i := range dataModels {
		res = append(res, s.ToPbOrgEmployee(dataModels[i]))
	}
	return res
}

func (s *orgDto) ToPbOrgEmployee(dataModel *po.OrgEmployee) *resourcev1.OrgEmployee {
	res := &resourcev1.OrgEmployee{
		Id:              dataModel.Id,
		CreatedTime:     "" + dataModel.CreatedTime.Format(timepkg.YmdHms),
		UpdatedTime:     dataModel.UpdatedTime.Format(timepkg.YmdHms),
		DeletedTime:     dataModel.DeletedTime,
		EmployeeId:      dataModel.EmployeeId,
		UserId:          dataModel.UserId,
		OrgId:           dataModel.OrgId,
		EmployeeName:    dataModel.EmployeeName,
		EmployeeAvatar:  dataModel.EmployeeAvatar,
		EmployeePhone:   dataModel.EmployeePhone,
		EmployeeEmail:   dataModel.EmployeeEmail,
		EmployeeRole:    dataModel.EmployeeRole,
		EmployeeStatus:  dataModel.EmployeeStatus,
		InviterRecordId: dataModel.InviterRecordId,
		InviterUserId:   dataModel.InviterUserId,
	}
	return res
}

func (s *orgDto) ToPbOrgInviteRecordList(dataModels []*po.OrgInviteRecord) []*resourcev1.OrgInviteRecord {
	res := make([]*resourcev1.OrgInviteRecord, 0, len(dataModels))
	for i := range dataModels {
		res = append(res, s.ToPbOrgInviteRecord(dataModels[i]))
	}
	return res
}

func (s *orgDto) ToPbOrgInviteRecord(dataModel *po.OrgInviteRecord) *resourcev1.OrgInviteRecord {
	res := &resourcev1.OrgInviteRecord{
		Id:                  dataModel.Id,
		CreatedTime:         dataModel.CreatedTime.Format(timepkg.YmdHms),
		UpdatedTime:         dataModel.UpdatedTime.Format(timepkg.YmdHms),
		InviteId:            dataModel.InviteId,
		OrgId:               dataModel.OrgId,
		InviterEmployeeId:   dataModel.InviterEmployeeId,
		InvitedType:         dataModel.InvitedType,
		InvitedUserId:       dataModel.InviterUserId,
		InvitedAccount:      dataModel.InvitedAccount,
		InvitedAccountType:  dataModel.InvitedAccountType,
		InvitedEmployeeRole: dataModel.InvitedEmployeeRole,
		InviteStatus:        dataModel.InviteStatus,
		AssignEmployeeId:    dataModel.AssignEmployeeId,
		InviteCode:          dataModel.InviteCode,
	}
	return res
}

func (s *orgDto) ToBoOrgListParam(req *resourcev1.GetOrgListReq) *bo.OrgListParam {
	res := &bo.OrgListParam{
		OrgIDList:     req.OrgIds,
		OrgName:       req.OrgName,
		PaginatorArgs: nil,
	}
	return res
}

func (s *orgDto) ToBoOrgEmployeeListParam(req *resourcev1.GetOrgEmployeeListReq) *bo.OrgEmployeeListParam {
	res := &bo.OrgEmployeeListParam{
		OrgIDList:      req.OrgIds,
		EmployeeIDList: req.EmployeeIds,
		UserIDList:     req.UserIds,
		EmployeeName:   req.EmployeeName,
		PaginatorArgs:  nil,
	}
	return res
}

func (s *orgDto) ToBoOrgInviteRecordListParam(req *resourcev1.GetOrgInviteRecordListReq) *bo.OrgInviteRecordListParam {
	res := &bo.OrgInviteRecordListParam{
		OrgIDList:         req.OrgIds,
		InviterUserIDList: req.InviterUserIds,
		InviteIDList:      req.InviteIds,
		InviteCode:        req.InviteCode,
		InviteAccount:     req.InviteAccount,
		PaginatorArgs:     nil,
	}
	return res
}

func (s *orgDto) ToBoSetOrgStatusParam(req *resourcev1.SetOrgStatusReq) *bo.SetOrgStatusParam {
	res := &bo.SetOrgStatusParam{
		OperatorEid: req.OperatorEid,
		OrgId:       req.OrgId,
		OrgStatus:   req.OrgStatus,
	}
	return res
}

func (s *orgDto) ToBoRemoveEmployeeParam(req *resourcev1.RemoveEmployeeReq) *bo.RemoveEmployeeParam {
	res := &bo.RemoveEmployeeParam{
		OperatorEid: req.OperatorEid,
		EmployeeId:  req.EmployeeId,
	}
	return res
}

func (s *orgDto) ToBoSetEmployeeRoleParam(req *resourcev1.SetEmployeeRoleReq) *bo.SetEmployeeRoleParam {
	res := &bo.SetEmployeeRoleParam{
		OperatorEid:  req.OperatorEid,
		EmployeeId:   req.EmployeeId,
		EmployeeRole: req.EmployeeRole,
	}
	return res
}

func (s *orgDto) ToBoSetEmployeeStatusParam(req *resourcev1.SetEmployeeStatusReq) *bo.SetEmployeeStatusParam {
	res := &bo.SetEmployeeStatusParam{
		OperatorEid:    req.OperatorEid,
		EmployeeId:     req.EmployeeId,
		EmployeeStatus: req.EmployeeStatus,
	}
	return res
}

func (s *orgDto) ToPbSetOrgStatusRespData(dataModel *po.Org) *resourcev1.SetOrgStatusRespData {
	res := &resourcev1.SetOrgStatusRespData{
		Success:   true,
		OrgId:     dataModel.OrgId,
		OrgStatus: dataModel.OrgStatus,
	}
	return res
}

func (s *orgDto) ToPbRemoveEmployeeRespData(dataModel *po.OrgEmployee) *resourcev1.RemoveEmployeeRespData {
	res := &resourcev1.RemoveEmployeeRespData{
		Success:    true,
		EmployeeId: dataModel.EmployeeId,
	}
	return res
}

func (s *orgDto) ToPbSetEmployeeRoleRespData(dataModel *po.OrgEmployee) *resourcev1.SetEmployeeRoleRespData {
	res := &resourcev1.SetEmployeeRoleRespData{
		Success:      true,
		EmployeeId:   dataModel.EmployeeId,
		EmployeeRole: dataModel.EmployeeRole,
	}
	return res
}

func (s *orgDto) ToPbSetEmployeeStatusRespData(dataModel *po.OrgEmployee) *resourcev1.SetEmployeeStatusRespData {
	res := &resourcev1.SetEmployeeStatusRespData{
		Success:        true,
		EmployeeId:     dataModel.EmployeeId,
		EmployeeStatus: dataModel.EmployeeStatus,
	}
	return res
}

func (s *orgDto) ToBoGetUserLastOrgParam(req *resourcev1.GetUserLastOrgReq) *bo.GetUserLastOrgParam {
	res := &bo.GetUserLastOrgParam{
		UserID:     req.UserId,
		UserName:   req.UserName,
		UserAvatar: req.UserAvatar,

		CreatePersonalOrgIfNotExist: req.CreatePersonalOrgIfNotExist,
	}
	return res
}
