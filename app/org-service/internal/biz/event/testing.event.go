package events

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	bizrepos "github.com/go-micro-saas/organization-service/app/org-service/internal/biz/repo"
	errorpkg "github.com/ikaiguang/go-srv-kit/kratos/error"
	"google.golang.org/protobuf/types/known/emptypb"
)

type testingEvent struct {
	log *log.Helper
}

func NewTestingEvent(
	logger log.Logger,
) bizrepos.TestingEventRepo {
	logHelper := log.NewHelper(log.With(logger, "module", "test-service/biz/event"))

	return &testingEvent{
		log: logHelper,
	}
}

func (t *testingEvent) Publish(ctx context.Context, msg *emptypb.Empty) error {
	return errorpkg.WithStack(errorpkg.ErrorUnimplemented("implement me"))
}

func (t *testingEvent) Consume(ctx context.Context, handler bizrepos.Handler) error {
	return errorpkg.WithStack(errorpkg.ErrorUnimplemented("implement me"))
}

func (t *testingEvent) Close(ctx context.Context) error {
	return nil
}
