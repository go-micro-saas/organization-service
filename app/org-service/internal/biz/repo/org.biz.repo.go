package bizrepos

import (
	"context"
	"github.com/go-micro-saas/organization-service/app/org-service/internal/biz/bo"
	"github.com/go-micro-saas/organization-service/app/org-service/internal/data/po"
)

type OrgBizRepo interface {
	CreateOrg(ctx context.Context, param *bo.CreateOrgParam) (*bo.CreateOrgReply, error)
	AddEmployee(ctx context.Context, param *bo.AddEmployeeParam) (*bo.AddEmployeeReply, error)
	CreateInviteRecordForLink(ctx context.Context, param *bo.CreateInviteRecordForLinkParam) (*po.OrgInviteRecord, error)
	CreateInviteRecordForEmployee(ctx context.Context, param *bo.CreateInviteRecordForEmployeeParam) (*po.OrgInviteRecord, error)
}
