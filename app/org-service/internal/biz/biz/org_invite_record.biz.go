package biz

import (
	"context"
	enumv1 "github.com/go-micro-saas/organization-service/api/org-service/v1/enums"
	"github.com/go-micro-saas/organization-service/app/org-service/internal/biz/bo"
	"github.com/go-micro-saas/organization-service/app/org-service/internal/data/po"
	errorpkg "github.com/ikaiguang/go-srv-kit/kratos/error"
	"time"
)

func (s *orgBiz) GenerateInviteLink(ctx context.Context, param *bo.GenerateInviteLinkParam) (*po.OrgInviteRecord, error) {
	recordModel, err := po.DefaultInviteRecordWithID(s.idGenerator)
	if err != nil {
		return nil, err
	}
	recordModel.InvitedType = enumv1.OrgInviteTypeEnum_LINK
	if param.ExpireTime.After(time.Now()) {
		recordModel.ExpireTime = param.ExpireTime
	}
	err = s.inviteRecordData.Create(ctx, recordModel)
	if err != nil {
		return nil, err
	}
	return recordModel, nil
}

func (s *orgBiz) GenerateInviteEmployee(ctx context.Context, param *bo.GenerateInviteEmployeeParam) (*po.OrgInviteRecord, error) {
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
	if param.ExpireTime.After(time.Now()) {
		recordModel.ExpireTime = param.ExpireTime
	}
	recordModel.OrgId = param.OrgId
	recordModel.InviterUserId = param.OperatorUid
	recordModel.InviterEmployeeId = inviterModel.EmployeeId
	recordModel.InvitedType = enumv1.OrgInviteTypeEnum_ACCOUNT
	recordModel.InvitedUserId = param.InviteUserId
	recordModel.InvitedAccount = param.InviteAccount
	recordModel.InvitedAccountType = param.InviteAccountType
	recordModel.InvitedEmployeeRole = param.InviteEmployeeRole
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
