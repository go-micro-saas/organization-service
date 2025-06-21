package service

import (
	"context"
	resourcev1 "github.com/go-micro-saas/organization-service/api/org-service/v1/resources"
	"github.com/go-micro-saas/organization-service/app/org-service/internal/service/dto"
	gormpkg "github.com/ikaiguang/go-srv-kit/data/gorm"
	pagepkg "github.com/ikaiguang/go-srv-kit/kit/page"
)

func (s *orgV1Service) AddEmployee(ctx context.Context, req *resourcev1.AddEmployeeReq) (*resourcev1.AddEmployeeResp, error) {
	param := dto.OrgDto.ToBoAddEmployeeParam(req)
	reply, err := s.orgBiz.AddEmployee(ctx, param)
	if err != nil {
		return nil, err
	}
	return &resourcev1.AddEmployeeResp{
		Data: dto.OrgDto.ToPbAddEmployeeRespData(reply),
	}, nil
}

func (s *orgV1Service) GetOrgEmployeeInfo(ctx context.Context, req *resourcev1.GetOrgEmployeeInfoReq) (*resourcev1.GetOrgEmployeeInfoResp, error) {
	dataModel, err := s.orgBiz.GetEmployeeInfo(ctx, req.EmployeeId)
	if err != nil {
		return nil, err
	}
	return &resourcev1.GetOrgEmployeeInfoResp{
		Data: dto.OrgDto.ToPbOrgEmployee(dataModel),
	}, nil
}

func (s *orgV1Service) GetOrgEmployeeInfoList(ctx context.Context, req *resourcev1.GetOrgEmployeeInfoListReq) (*resourcev1.GetOrgEmployeeInfoListResp, error) {
	dataModels, err := s.orgBiz.GetEmployeeInfoList(ctx, req.EmployeeIds)
	if err != nil {
		return nil, err
	}
	return &resourcev1.GetOrgEmployeeInfoListResp{
		Data: dto.OrgDto.ToPbOrgEmployeeList(dataModels),
	}, nil
}

func (s *orgV1Service) GetOrgEmployeeList(ctx context.Context, req *resourcev1.GetOrgEmployeeListReq) (*resourcev1.GetOrgEmployeeListResp, error) {
	// paging args
	pageRequest, pageOption := pagepkg.ParsePageRequest(req.PageRequest)

	// list
	param := dto.OrgDto.ToBoOrgEmployeeListParam(req)
	param.PaginatorArgs = &gormpkg.PaginatorArgs{
		PageRequest: pageRequest,
		PageOption:  pageOption,
	}
	dataModels, dataCount, err := s.orgBiz.ListOrgEmployee(ctx, param)
	if err != nil {
		return nil, err
	}

	// paging result
	return &resourcev1.GetOrgEmployeeListResp{
		Data: &resourcev1.GetOrgEmployeeListRespData{
			List:     dto.OrgDto.ToPbOrgEmployeeList(dataModels),
			PageInfo: pagepkg.CalcPageResponse(pageRequest, uint32(dataCount)),
		},
	}, nil
}
