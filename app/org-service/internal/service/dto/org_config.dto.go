package dto

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-micro-saas/organization-service/app/org-service/internal/conf"
	"github.com/go-micro-saas/organization-service/app/org-service/internal/data/po"
	nodeidresourcev1 "github.com/go-micro-saas/service-api/api/nodeid-service/v1/resources"
	accountapi "github.com/go-micro-saas/service-api/app/account-service"
	snowflakeapi "github.com/go-micro-saas/service-api/app/snowflake-service"
	errorpkg "github.com/ikaiguang/go-srv-kit/kratos/error"
	clientutil "github.com/ikaiguang/go-srv-kit/service/cluster_service_api"
)

func GetNodeIDOptions(logger log.Logger, cfg *conf.ServiceConfig) []snowflakeapi.Option {
	opts := snowflakeapi.DefaultOptions(logger)
	if cfg.GetOrgService().GetSnowflake().GetEnable() {
		opts = append(opts, snowflakeapi.WithMustGetNodeIdFromAPI(true))
	}
	return opts
}

func ToPbGetNodeIdReq(cfg *conf.ServiceConfig) (*nodeidresourcev1.GetNodeIdReq, error) {
	snowflakeConf := cfg.GetOrgService().GetSnowflake()
	if snowflakeConf == nil {
		e := errorpkg.ErrorInvalidParameter("snowflake config is nil")
		return nil, errorpkg.WithStack(e)
	}
	if err := snowflakeConf.Validate(); err != nil {
		e := errorpkg.ErrorInvalidParameter(err.Error())
		return nil, errorpkg.WithStack(e)
	}
	res := &nodeidresourcev1.GetNodeIdReq{
		InstanceId:   snowflakeConf.InstanceId,
		InstanceName: snowflakeConf.InstanceName,
		Metadata:     snowflakeConf.Metadata,
	}
	return res, nil
}

func GetAccountV1ServiceNameForGRPC() []clientutil.ServiceName {
	return []clientutil.ServiceName{
		accountapi.AccountServiceGRPC,
	}
}

func GetAccountV1ServiceNameForHTTP() []clientutil.ServiceName {
	return []clientutil.ServiceName{
		accountapi.AccountServiceHTTP,
	}
}

func GetBusinessSetting(cfg *conf.ServiceConfig) *po.BusinessSetting {
	setting := po.DefaultBusinessSetting()
	config := cfg.GetOrgService().GetBusinessSetting()
	if config == nil {
		return setting
	}
	if config.GetUserCreateOrgMaxCount() > 0 {
		setting.UserCreateOrgMaxCount = config.GetUserCreateOrgMaxCount()
	}
	if config.GetUserBelongOrgMaxCount() > 0 {
		setting.UserBelongOrgMaxCount = config.GetUserBelongOrgMaxCount()
	}
	return setting
}
