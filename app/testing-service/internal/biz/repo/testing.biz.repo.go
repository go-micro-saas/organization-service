package bizrepos

import (
	"context"
	"github.com/go-micro-saas/organization-service/app/testing-service/internal/biz/bo"
)

type TestingBizRepo interface {
	HelloWorld(ctx context.Context, param *bo.HelloWorldParam) (*bo.HelloWorldReply, error)
}
