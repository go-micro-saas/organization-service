package biz

import (
	"context"
	enumv1 "github.com/go-micro-saas/organization-service/api/org-service/v1/enums"
	errorv1 "github.com/go-micro-saas/organization-service/api/org-service/v1/errors"
	"github.com/go-micro-saas/organization-service/app/org-service/internal/biz/bo"
	"github.com/go-micro-saas/organization-service/app/org-service/internal/data/po"
	errorpkg "github.com/ikaiguang/go-srv-kit/kratos/error"
	"time"
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
	// check account
	_, err := s.GetAccountInfo(ctx, param.InviteUserId)
	if err != nil {
		return nil, err
	}

	// invite
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
	// check account
	accountInfo, err := s.GetAccountInfo(ctx, param.UserId)
	if err != nil {
		return nil, err
	}

	// invite
	inviteModel, err := s.GetInviteRecordInfo(ctx, param.InviteId)
	if err != nil {
		return nil, err
	}
	if err = inviteModel.IsValidInviteRecord(param.InviteCode); err != nil {
		return nil, err
	}
	if inviteModel.InvitedType != enumv1.OrgInviteTypeEnum_LINK {
		e := errorv1.DefaultErrorS105IncorrectInviteType()
		return nil, errorpkg.WithStack(e)
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
	s.SetOrgEmployeeByAccountInfo(employeeModel, accountInfo)
	if err = s.isEmployeeExists(ctx, employeeModel); err != nil {
		return nil, err
	}
	if err = s.CheckCreateOrJoinOrgLimit(ctx, employeeModel); err != nil {
		return nil, err
	}

	// create
	err = s.employeeData.Create(ctx, employeeModel)
	if err != nil {
		return nil, err
	}
	return employeeModel, nil
}

func (s *orgBiz) AgreeOrRefuseInvite(ctx context.Context, param *bo.AgreeOrRefuseInviteParam) (employeeModel *po.OrgEmployee, err error) {
	// check account
	accountInfo, err := s.GetAccountInfo(ctx, param.UserId)
	if err != nil {
		return nil, err
	}

	// invite
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
	if inviteModel.InvitedUserId != param.UserId {
		e := errorv1.DefaultErrorS105NotInvitedUser()
		return nil, errorpkg.WithStack(e)
	}
	if inviteModel.InvitedType != enumv1.OrgInviteTypeEnum_ACCOUNT {
		e := errorv1.DefaultErrorS105IncorrectInviteType()
		return nil, errorpkg.WithStack(e)
	}

	inviteModel.UpdatedTime = time.Now()
	inviteModel.InviteStatus = enumv1.OrgInviteStatusEnum_AGREED
	if !param.IsAgree {
		inviteModel.InviteStatus = enumv1.OrgInviteStatusEnum_REJECTED
	}

	// employee
	employeeModel, err = po.DefaultOrgEmployeeWithID(s.idGenerator)
	if err != nil {
		return employeeModel, err
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
	s.SetOrgEmployeeByAccountInfo(employeeModel, accountInfo)
	// check
	if param.IsAgree {
		if err = s.isEmployeeExists(ctx, employeeModel); err != nil {
			return employeeModel, err
		}
	}
	if param.IsAgree {
		if err = s.CheckCreateOrJoinOrgLimit(ctx, employeeModel); err != nil {
			return nil, err
		}
	}

	// 事务
	tx := s.inviteRecordData.NewTransaction(ctx)
	defer func() {
		commitErr := tx.CommitAndErrRollback(ctx, err)
		if commitErr != nil {
			s.log.WithContext(ctx).Errorw(
				"mgs", "AgreeOrRefuseInvite tx.CommitAndErrRollback failed!",
				"err", commitErr,
			)
		}
	}()
	err = s.inviteRecordData.UpdateInviteStatusWithTransaction(ctx, tx, inviteModel)
	if err != nil {
		return employeeModel, err
	}
	// refuse
	if !param.IsAgree {
		return employeeModel, err
	}
	// create
	err = s.employeeData.CreateWithTransaction(ctx, tx, employeeModel)
	if err != nil {
		return employeeModel, err
	}
	return employeeModel, err
}

func (s *orgBiz) GetInviteRecordInfo(ctx context.Context, inviteID uint64) (*po.OrgInviteRecord, error) {
	dataModel, isNotFound, err := s.inviteRecordData.QueryOneByInviteID(ctx, inviteID)
	if err != nil {
		return nil, err
	}
	if isNotFound {
		e := errorpkg.ErrorRecordNotFound("invite not found")
		return nil, errorpkg.WithStack(e)
	}
	return dataModel, nil
}

func (s *orgBiz) GetInviteRecordInfoList(ctx context.Context, inviteIDList []uint64) ([]*po.OrgInviteRecord, error) {
	dataModels, err := s.inviteRecordData.QueryByInviteIDList(ctx, inviteIDList)
	if err != nil {
		return nil, err
	}
	return dataModels, nil
}

func (s *orgBiz) ListOrgInviteRecord(ctx context.Context, param *bo.OrgInviteRecordListParam) ([]*po.OrgInviteRecord, int64, error) {
	queryParam := &po.OrgInviteRecordListParam{
		OrgIDList:         param.OrgIDList,
		InviterUserIDList: param.InviterUserIDList,
		InviteIDList:      param.InviteIDList,
		InviteCode:        param.InviteCode,
		InviteAccount:     param.InviteAccount,
		PaginatorArgs:     param.PaginatorArgs,
	}
	dataModels, counter, err := s.inviteRecordData.ListOrgInviteRecord(ctx, queryParam, param.PaginatorArgs)
	if err != nil {
		return dataModels, counter, err
	}
	return dataModels, counter, err
}
