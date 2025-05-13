package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-micro-saas/organization-service/app/org-service/internal/biz/bo"
	bizrepos "github.com/go-micro-saas/organization-service/app/org-service/internal/biz/repo"
	"github.com/go-micro-saas/organization-service/app/org-service/internal/data/po"
	datarepos "github.com/go-micro-saas/organization-service/app/org-service/internal/data/repo"
	idpkg "github.com/ikaiguang/go-srv-kit/kit/id"
	"time"
)

type orgBiz struct {
	log         *log.Helper
	idGenerator idpkg.Snowflake

	orgData datarepos.OrgDataRepo
}

func NewOrgBiz(
	logger log.Logger,
	idGenerator idpkg.Snowflake,
	orgData datarepos.OrgDataRepo,
) bizrepos.OrgBizRepo {
	logHelper := log.NewHelper(log.With(logger, "module", "test-service/biz/biz"))

	return &orgBiz{
		log:         logHelper,
		idGenerator: idGenerator,
		orgData:     orgData,
	}
}

func (s *orgBiz) HelloWorld(ctx context.Context, param *bo.HelloWorldParam) (*bo.HelloWorldReply, error) {
	dataModel := s.toHelloWorkModel(param)
	err := s.orgData.HelloWorld(ctx, dataModel)
	if err != nil {
		return nil, err
	}
	return &bo.HelloWorldReply{Message: dataModel.RequestMessage}, nil
}

func (s *orgBiz) toHelloWorkModel(param *bo.HelloWorldParam) *po.HelloWorld {
	res := &po.HelloWorld{
		Id:             0,
		CreatedTime:    time.Now(),
		UpdatedTime:    time.Now(),
		RequestMessage: param.Message,
	}
	return res
}

func (s *orgBiz) CreateOrg(ctx context.Context, param *bo.HelloWorldParam) (*bo.HelloWorldReply, error) {
	dataModel := po.DefaultOrg()

	_ = dataModel
	return nil, nil
}
