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
	JoinByInviteLink(ctx context.Context, param *bo.JoinByInviteLinkParam) (*po.OrgEmployee, error)
	AgreeOrRefuseInvite(ctx context.Context, param *bo.AgreeOrRefuseInviteParam) (*po.OrgEmployee, error)
	GetOrgInfo(ctx context.Context, orgID uint64) (*po.Org, error)
	GetOrgInfoList(ctx context.Context, orgIDList []uint64) ([]*po.Org, error)
	GetEmployeeInfo(ctx context.Context, employeeID uint64) (*po.OrgEmployee, error)
	GetEmployeeInfoList(ctx context.Context, employeeIDList []uint64) ([]*po.OrgEmployee, error)
	GetInviteRecordInfo(ctx context.Context, inviteID uint64) (*po.OrgInviteRecord, error)
	GetInviteRecordInfoList(ctx context.Context, inviteIDList []uint64) ([]*po.OrgInviteRecord, error)
}
