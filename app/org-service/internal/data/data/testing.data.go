package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-micro-saas/organization-service/app/org-service/internal/data/po"
	datarepos "github.com/go-micro-saas/organization-service/app/org-service/internal/data/repo"
	gormpkg "github.com/ikaiguang/go-srv-kit/data/gorm"
	idpkg "github.com/ikaiguang/go-srv-kit/kit/id"
	errorpkg "github.com/ikaiguang/go-srv-kit/kratos/error"
)

type orgData struct {
	log *log.Helper
}

func NewOrgData(
	logger log.Logger,
) datarepos.OrgDataRepo {
	logHelper := log.NewHelper(log.With(logger, "module", "test-service/data/data"))

	return &orgData{
		log: logHelper,
	}
}

func (s *orgData) HelloWorld(ctx context.Context, dataModel *po.HelloWorld) error {
	dataModel.Id, _ = idpkg.NextID()
	var err error
	if err != nil {
		e := errorpkg.ErrorInternalServer("")
		return errorpkg.Wrap(e, err)
	}
	return nil
}

func (s *orgData) QueryByID(ctx context.Context, id uint64) (dataModel *po.HelloWorld, isNotFound bool, err error) {
	if err != nil {
		if gormpkg.IsErrRecordNotFound(err) {
			err = nil
			isNotFound = true
		} else {
			e := errorpkg.ErrorInternalServer("")
			err = errorpkg.Wrap(e, err)
		}
		return
	}
	return dataModel, isNotFound, err
}

func (s *orgData) Create(ctx context.Context, dataModel *po.HelloWorld) error {
	dataModel.Id, _ = idpkg.NextID()
	var err error
	if err != nil {
		if gormpkg.IsErrDuplicatedKey(err) {
			e := errorpkg.ErrorDuplicateKey("")
			err = errorpkg.Wrap(e, err)
		} else {
			e := errorpkg.ErrorInternalServer("")
			err = errorpkg.Wrap(e, err)
		}
		return err
	}
	return nil
}
