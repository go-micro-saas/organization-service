package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	enumv1 "github.com/go-micro-saas/organization-service/api/org-service/v1/enums"
	errorv1 "github.com/go-micro-saas/organization-service/api/org-service/v1/errors"
	"github.com/go-micro-saas/organization-service/app/org-service/internal/biz/bo"
	bizrepos "github.com/go-micro-saas/organization-service/app/org-service/internal/biz/repo"
	"github.com/go-micro-saas/organization-service/app/org-service/internal/data/po"
	datarepos "github.com/go-micro-saas/organization-service/app/org-service/internal/data/repo"
	idpkg "github.com/ikaiguang/go-srv-kit/kit/id"
	errorpkg "github.com/ikaiguang/go-srv-kit/kratos/error"
)

type orgBiz struct {
	log         *log.Helper
	idGenerator idpkg.Snowflake

	orgData          datarepos.OrgRepo
	employeeData     datarepos.OrgEmployeeRepo
	inviteRecordData datarepos.OrgInviteRecordRepo
}

func NewOrgBiz(
	logger log.Logger,
	idGenerator idpkg.Snowflake,
	orgData datarepos.OrgRepo,
	employeeData datarepos.OrgEmployeeRepo,
	inviteRecordData datarepos.OrgInviteRecordRepo,
) bizrepos.OrgBizRepo {
	logHelper := log.NewHelper(log.With(logger, "module", "test-service/biz/biz"))

	return &orgBiz{
		log:              logHelper,
		idGenerator:      idGenerator,
		orgData:          orgData,
		employeeData:     employeeData,
		inviteRecordData: inviteRecordData,
	}
}

func (s *orgBiz) CreateOrg(ctx context.Context, param *bo.CreateOrgParam) (*bo.CreateOrgReply, error) {
	orgModel, err := po.DefaultOrgWithID(s.idGenerator)
	if err != nil {
		return nil, err
	}
	orgModel.OrgName = param.OrgName
	orgModel.OrgAvatar = param.OrgAvatar
	orgModel.OrgType = param.OrgType
	orgModel.OrgCreatorId = param.CreatorID

	employeeModel, err := po.DefaultOrgEmployeeWithID(s.idGenerator)
	if err != nil {
		return nil, err
	}
	employeeModel.UserId = orgModel.OrgCreatorId
	employeeModel.OrgId = orgModel.OrgId
	employeeModel.EmployeeUuid = employeeModel.GenUUID()
	employeeModel.EmployeeName = param.CreatorName
	employeeModel.EmployeeAvatar = param.CreatorAvatar
	employeeModel.EmployeeRole = enumv1.OrgEmployeeRoleEnum_CREATOR

	// 事务
	tx := s.orgData.NewTransaction(ctx)
	defer func() {
		commitErr := tx.CommitAndErrRollback(ctx, err)
		if commitErr != nil {
			s.log.WithContext(ctx).Errorw(
				"mgs", "CreateOrg tx.CommitAndErrRollback failed!",
				"err", commitErr,
			)
		}
	}()
	err = s.orgData.CreateWithTransaction(ctx, tx, orgModel)
	if err != nil {
		return nil, err
	}
	if !param.SkipCreateEmployee {
		err = s.employeeData.CreateWithTransaction(ctx, tx, employeeModel)
		if err != nil {
			return nil, err
		}
	}

	// result
	res := &bo.CreateOrgReply{}
	res.SetByOrg(orgModel)
	return res, nil
}

func (s *orgBiz) AddEmployee(ctx context.Context, param *bo.AddEmployeeParam) (*bo.AddEmployeeReply, error) {
	var (
		employee = param.NewEmployeeModel()
		err      error
	)
	employee.EmployeeId, err = s.idGenerator.NextID()
	if err != nil {
		e := errorpkg.ErrorInternalServer(err.Error())
		return nil, errorpkg.WithStack(e)
	}

	if err = s.cannotBeOwner(employee.EmployeeRole); err != nil {
		return nil, err
	}
	if _, err = s.canAddEmployee(ctx, param); err != nil {
		return nil, err
	}
	if err = s.isEmployeeExists(ctx, employee); err != nil {
		return nil, err
	}

	// create
	err = s.employeeData.Create(ctx, employee)
	if err != nil {
		return nil, err
	}

	res := &bo.AddEmployeeReply{}
	res.SetByEmployee(employee)
	return res, nil
}

func (s *orgBiz) isEmployeeExists(ctx context.Context, dataModel *po.OrgEmployee) error {
	return s.isEmployeeExistsByEmployeeUUID(ctx, dataModel.EmployeeUuid)
}

func (s *orgBiz) isEmployeeExistsByEmployeeUUID(ctx context.Context, employeeUUID string) error {
	_, isNotFound, err := s.employeeData.ExistCreateByEmployeeUUID(ctx, employeeUUID)
	if err != nil {
		return err
	}
	if !isNotFound {
		e := errorv1.DefaultErrorS105EmployeeExists()
		return errorpkg.WithStack(e)
	}
	return nil
}

func (s *orgBiz) cannotBeOwner(employeeRole enumv1.OrgEmployeeRoleEnum_OrgEmployeeRole) error {
	if po.IsOrgOwner(employeeRole) {
		e := errorv1.DefaultErrorS105CannotBeOwner()
		return errorpkg.WithStack(e)
	}
	return nil
}

func (s *orgBiz) canAddEmployee(ctx context.Context, param *bo.AddEmployeeParam) (*po.OrgEmployee, error) {
	employee, err := s.canInviteEmployee(ctx, param.OperatorUid, param.OrgId)
	if err != nil {
		return nil, err
	}
	if !employee.CanAddEmployee() {
		e := errorv1.DefaultErrorS105EmployeeNoPermission()
		return nil, errorpkg.WithStack(e)
	}
	return employee, nil
}

func (s *orgBiz) canInviteEmployee(ctx context.Context, operatorUid, orgID uint64) (*po.OrgEmployee, error) {
	queryParam := &po.QueryEmployeeParam{
		OrgID:      orgID,
		UserID:     operatorUid,
		EmployeeId: 0,
	}
	employee, isNotFound, err := s.employeeData.QueryOneByUserID(ctx, queryParam)
	if err != nil {
		return nil, err
	}
	if isNotFound {
		e := errorv1.DefaultErrorS105InvalidOperator()
		return nil, errorpkg.WithStack(e)
	}
	if !employee.IsValid() {
		e := errorv1.DefaultErrorS105InvalidOperator()
		return nil, errorpkg.WithStack(e)
	}
	return employee, nil
}
