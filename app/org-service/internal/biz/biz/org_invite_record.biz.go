package biz

import (
	"context"
	enumv1 "github.com/go-micro-saas/organization-service/api/org-service/v1/enums"
	"github.com/go-micro-saas/organization-service/app/org-service/internal/biz/bo"
	"github.com/go-micro-saas/organization-service/app/org-service/internal/data/po"
	errorpkg "github.com/ikaiguang/go-srv-kit/kratos/error"
)

func (s *orgBiz) CreateInviteRecordForLink(ctx context.Context, param *bo.CreateInviteRecordForLinkParam) (*po.OrgInviteRecord, error) {
	_, err := s.canInviteEmployee(ctx, param.OperatorUid, param.OrgId)
	if err != nil {
		return nil, err
	}

	recordModel, err := po.DefaultInviteRecordWithID(s.idGenerator)
	if err != nil {
		return nil, err
	}
	recordModel.OrgId = param.OrgId
	recordModel.InvitedType = enumv1.OrgInviteTypeEnum_LINK
	recordModel.InvitedEmployeeRole = param.InviteEmployeeRole
	recordModel.CheckAndSetExpireTime(param.ExpireTime)
	err = s.inviteRecordData.Create(ctx, recordModel)
	if err != nil {
		return nil, err
	}
	return recordModel, nil
}

func (s *orgBiz) CreateInviteRecordForEmployee(ctx context.Context, param *bo.CreateInviteRecordForEmployeeParam) (*po.OrgInviteRecord, error) {
	inviterModel, err := s.canInviteEmployee(ctx, param.OperatorUid, param.OrgId)
	if err != nil {
		return nil, err
	}
	employeeUUID := po.GenEmployeeUUID(param.OrgId, param.InviteUserId)
	if err = s.isEmployeeExistsByEmployeeUUID(ctx, employeeUUID); err != nil {
		return nil, err
	}
	if err = s.cannotBeOwner(param.InviteEmployeeRole); err != nil {
		return nil, err
	}

	recordModel, err := po.DefaultInviteRecordWithID(s.idGenerator)
	if err != nil {
		return nil, err
	}
	recordModel.OrgId = param.OrgId
	recordModel.InviterUserId = param.OperatorUid
	recordModel.InviterEmployeeId = inviterModel.EmployeeId
	recordModel.InvitedType = enumv1.OrgInviteTypeEnum_ACCOUNT
	recordModel.InvitedUserId = param.InviteUserId
	recordModel.InvitedAccount = param.InviteAccount
	recordModel.InvitedAccountType = param.InviteAccountType
	recordModel.InvitedEmployeeRole = param.InviteEmployeeRole
	recordModel.CheckAndSetExpireTime(param.ExpireTime)
	recordModel.AssignEmployeeId, err = s.idGenerator.NextID()
	if err != nil {
		e := errorpkg.ErrorInternalServer(err.Error())
		return nil, errorpkg.WithStack(e)
	}

	err = s.inviteRecordData.Create(ctx, recordModel)
	if err != nil {
		return nil, err
	}
	return recordModel, nil
}

func (s *orgBiz) JoinByInviteLink(ctx context.Context, param *bo.JoinByInviteLinkParam) (*po.OrgEmployee, error) {
	inviteModel, isNotFound, err := s.inviteRecordData.QueryOneByInviteID(ctx, param.InviteId)
	if err != nil {
		return nil, err
	}
	if isNotFound {
		e := errorpkg.ErrorRecordNotFound("invite not found")
		return nil, errorpkg.WithStack(e)
	}
	if err = inviteModel.IsValidInviteRecord(param.InviteCode); err != nil {
		return nil, err
	}

	employeeModel, err := po.DefaultOrgEmployeeWithID(s.idGenerator)
	if err != nil {
		return nil, err
	}
	employeeModel.UserId = param.UserId
	employeeModel.OrgId = inviteModel.OrgId
	employeeModel.InviterRecordId = inviteModel.InviteId
	employeeModel.InviterUserId = inviteModel.InviterUserId
	employeeModel.EmployeeName = param.UserName
	employeeModel.EmployeeAvatar = param.UserAvatar
	employeeModel.EmployeePhone = param.UserPhone
	employeeModel.EmployeeEmail = param.UserEmail
	employeeModel.EmployeeUuid = employeeModel.GenUUID()
	employeeModel.EmployeeRole = enumv1.OrgEmployeeRoleEnum_NORMAL
	if inviteModel.InvitedEmployeeRole != enumv1.OrgEmployeeRoleEnum_UNSPECIFIED {
		employeeModel.EmployeeRole = inviteModel.InvitedEmployeeRole
	}
	if err = s.isEmployeeExists(ctx, employeeModel); err != nil {
		return nil, err
	}
	// create
	err = s.employeeData.Create(ctx, employeeModel)
	if err != nil {
		return nil, err
	}
	return employeeModel, nil
}
