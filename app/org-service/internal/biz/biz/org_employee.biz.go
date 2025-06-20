package biz

import (
	"context"
	enumv1 "github.com/go-micro-saas/organization-service/api/org-service/v1/enums"
	errorv1 "github.com/go-micro-saas/organization-service/api/org-service/v1/errors"
	"github.com/go-micro-saas/organization-service/app/org-service/internal/biz/bo"
	"github.com/go-micro-saas/organization-service/app/org-service/internal/data/po"
	errorpkg "github.com/ikaiguang/go-srv-kit/kratos/error"
)

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

func (s *orgBiz) GetEmployeeInfo(ctx context.Context, employeeID uint64) (*po.OrgEmployee, error) {
	dataModel, isNotFound, err := s.employeeData.QueryOneByEmployeeID(ctx, employeeID)
	if err != nil {
		return nil, err
	}
	if isNotFound {
		e := errorv1.DefaultErrorS105EmployeeNotFound()
		return nil, errorpkg.WithStack(e)
	}
	return dataModel, nil
}

func (s *orgBiz) GetEmployeeInfoList(ctx context.Context, employeeIDList []uint64) ([]*po.OrgEmployee, error) {
	dataModels, err := s.employeeData.QueryByEmployeeIDList(ctx, employeeIDList)
	if err != nil {
		return nil, err
	}
	return dataModels, nil
}

func (s *orgBiz) ListOrgEmployee(ctx context.Context, param *bo.OrgEmployeeListParam) ([]*po.OrgEmployee, int64, error) {
	queryParam := &po.OrgEmployeeListParam{
		OrgIDList:      param.OrgIDList,
		EmployeeIDList: param.EmployeeIDList,
		UserIDList:     param.UserIDList,
		EmployeeName:   param.EmployeeName,
		PaginatorArgs:  param.PaginatorArgs,
	}
	dataModels, counter, err := s.employeeData.ListOrgEmployee(ctx, queryParam, param.PaginatorArgs)
	if err != nil {
		return dataModels, counter, err
	}
	return dataModels, counter, err
}
