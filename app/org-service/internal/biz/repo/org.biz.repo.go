package bizrepos

import (
	"context"
	"github.com/go-micro-saas/organization-service/app/org-service/internal/biz/bo"
)

type OrgBizRepo interface {
	HelloWorld(ctx context.Context, param *bo.HelloWorldParam) (*bo.HelloWorldReply, error)
}
