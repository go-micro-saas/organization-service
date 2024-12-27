package datarepos

import (
	"context"
	"github.com/go-micro-saas/organization-service/app/org-service/internal/data/po"
)

type OrgDataRepo interface {
	HelloWorld(ctx context.Context, dataModel *po.HelloWorld) error
	QueryByID(ctx context.Context, id uint64) (dataModel *po.HelloWorld, isNotFound bool, err error)
	Create(ctx context.Context, dataModel *po.HelloWorld) error
}
