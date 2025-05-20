package bizrepos

import (
	"context"
	"github.com/go-micro-saas/organization-service/app/org-service/internal/biz/bo"
)

type OrgBizRepo interface {
	CreateOrg(ctx context.Context, param *bo.CreateOrgParam) (*bo.CreateOrgReply, error)
	AddEmployee(ctx context.Context, param *bo.AddEmployeeParam) (*bo.AddEmployeeReply, error)
}
