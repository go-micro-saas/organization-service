package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	enumv1 "github.com/go-micro-saas/organization-service/api/org-service/v1/enums"
	"github.com/go-micro-saas/organization-service/app/org-service/internal/biz/bo"
	bizrepos "github.com/go-micro-saas/organization-service/app/org-service/internal/biz/repo"
	"github.com/go-micro-saas/organization-service/app/org-service/internal/data/po"
	datarepos "github.com/go-micro-saas/organization-service/app/org-service/internal/data/repo"
	idpkg "github.com/ikaiguang/go-srv-kit/kit/id"
)

type orgBiz struct {
	log         *log.Helper
	idGenerator idpkg.Snowflake

	orgData      datarepos.OrgRepo
	employeeData datarepos.OrgEmployeeRepo
}

func NewOrgBiz(
	logger log.Logger,
	idGenerator idpkg.Snowflake,
	orgData datarepos.OrgRepo,
	employeeData datarepos.OrgEmployeeRepo,
) bizrepos.OrgBizRepo {
	logHelper := log.NewHelper(log.With(logger, "module", "test-service/biz/biz"))

	return &orgBiz{
		log:          logHelper,
		idGenerator:  idGenerator,
		orgData:      orgData,
		employeeData: employeeData,
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
