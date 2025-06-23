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
	accountservicev1 "github.com/go-micro-saas/service-api/api/account-service/v1/services"
	idpkg "github.com/ikaiguang/go-srv-kit/kit/id"
	errorpkg "github.com/ikaiguang/go-srv-kit/kratos/error"
)

type orgBiz struct {
	log         *log.Helper
	idGenerator idpkg.Snowflake

	orgData          datarepos.OrgRepo
	employeeData     datarepos.OrgEmployeeRepo
	inviteRecordData datarepos.OrgInviteRecordRepo

	accountV1Client accountservicev1.SrvAccountV1Client
}

func NewOrgBiz(
	logger log.Logger,
	idGenerator idpkg.Snowflake,

	orgData datarepos.OrgRepo,
	employeeData datarepos.OrgEmployeeRepo,
	inviteRecordData datarepos.OrgInviteRecordRepo,

	accountV1Client accountservicev1.SrvAccountV1Client,
) bizrepos.OrgBizRepo {
	logHelper := log.NewHelper(log.With(logger, "module", "test-service/biz/biz"))

	return &orgBiz{
		log:              logHelper,
		idGenerator:      idGenerator,
		orgData:          orgData,
		employeeData:     employeeData,
		inviteRecordData: inviteRecordData,
		accountV1Client:  accountV1Client,
	}
}

func (s *orgBiz) CreateOrg(ctx context.Context, param *bo.CreateOrgParam) (*bo.CreateOrgReply, error) {
	// check account
	accountInfo, err := s.GetAccountInfo(ctx, param.CreatorID)
	if err != nil {
		return nil, err
	}

	// model
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
	// set employee
	s.SetOrgEmployeeByAccountInfo(ctx, employeeModel, accountInfo)

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

func (s *orgBiz) GetOrgInfo(ctx context.Context, orgID uint64) (*po.Org, error) {
	dataModel, isNotFound, err := s.orgData.QueryOneByOrgID(ctx, orgID)
	if err != nil {
		return nil, err
	}
	if isNotFound {
		e := errorv1.DefaultErrorS105OrgNotFound()
		return nil, errorpkg.WithStack(e)
	}
	return dataModel, nil
}

func (s *orgBiz) GetOrgInfoList(ctx context.Context, orgIDList []uint64) ([]*po.Org, error) {
	dataModels, err := s.orgData.QueryByOrgIDList(ctx, orgIDList)
	if err != nil {
		return nil, err
	}
	return dataModels, nil
}

func (s *orgBiz) ListOrg(ctx context.Context, param *bo.OrgListParam) ([]*po.Org, int64, error) {
	queryParam := &po.OrgListParam{
		OrgIDList:     param.OrgIDList,
		OrgName:       param.OrgName,
		PaginatorArgs: param.PaginatorArgs,
	}
	dataModels, counter, err := s.orgData.ListOrg(ctx, queryParam, param.PaginatorArgs)
	if err != nil {
		return dataModels, counter, err
	}
	return dataModels, counter, err
}
