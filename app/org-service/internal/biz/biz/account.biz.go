package biz

import (
	"context"
	"github.com/go-micro-saas/organization-service/app/org-service/internal/data/po"
	accountresourcev1 "github.com/go-micro-saas/service-api/api/account-service/v1/resources"
	apiutil "github.com/go-micro-saas/service-api/util"
	errorpkg "github.com/ikaiguang/go-srv-kit/kratos/error"
)

func (s *orgBiz) GetAccountInfo(ctx context.Context, accountID uint64) (*accountresourcev1.AccountInfo, error) {
	req := &accountresourcev1.GetUserInfoReq{
		UserId: accountID,
	}
	resp, err := s.accountV1Client.GetUserInfo(ctx, req)
	if e := apiutil.CheckAPIResponse(resp, err); e != nil {
		return nil, errorpkg.WithStack(e)
	}
	return resp.GetData(), nil
}

func (s *orgBiz) SetOrgEmployeeByAccountInfo(dataModel *po.OrgEmployee, info *accountresourcev1.AccountInfo) {
	if dataModel.EmployeeName == "" {
		dataModel.EmployeeName = info.GetUserNickname()
	}
	if dataModel.EmployeeAvatar == "" {
		dataModel.EmployeeAvatar = info.GetUserAvatar()
	}
	if dataModel.EmployeePhone == "" {
		dataModel.EmployeePhone = info.GetUserPhone()
	}
	if dataModel.EmployeeEmail == "" {
		dataModel.EmployeeEmail = info.GetUserEmail()
	}
}
