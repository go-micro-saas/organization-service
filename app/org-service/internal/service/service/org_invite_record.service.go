package service

import (
	"context"
	resourcev1 "github.com/go-micro-saas/organization-service/api/org-service/v1/resources"
	"github.com/go-micro-saas/organization-service/app/org-service/internal/service/dto"
	gormpkg "github.com/ikaiguang/go-srv-kit/data/gorm"
	pagepkg "github.com/ikaiguang/go-srv-kit/kit/page"
)

func (s *orgV1Service) CreateInviteRecordForLink(ctx context.Context, req *resourcev1.CreateInviteRecordForLinkReq) (*resourcev1.CreateInviteRecordForLinkResp, error) {
	param := dto.OrgDto.ToBoCreateInviteRecordForLinkParam(req)
	reply, err := s.orgBiz.CreateInviteRecordForLink(ctx, param)
	if err != nil {
		return nil, err
	}
	return &resourcev1.CreateInviteRecordForLinkResp{
		Data: dto.OrgDto.ToPbCreateInviteRecordForLinkRespData(reply),
	}, nil
}

func (s *orgV1Service) CreateInviteRecordForEmployee(ctx context.Context, req *resourcev1.CreateInviteRecordForEmployeeReq) (*resourcev1.CreateInviteRecordForEmployeeResp, error) {
	param := dto.OrgDto.ToBoCreateInviteRecordForEmployeeParam(req)
	reply, err := s.orgBiz.CreateInviteRecordForEmployee(ctx, param)
	if err != nil {
		return nil, err
	}
	return &resourcev1.CreateInviteRecordForEmployeeResp{
		Data: dto.OrgDto.ToPbCreateInviteRecordForEmployeeRespData(reply),
	}, nil
}

func (s *orgV1Service) JoinByInviteLink(ctx context.Context, req *resourcev1.JoinByInviteLinkReq) (*resourcev1.JoinByInviteLinkResp, error) {
	param := dto.OrgDto.ToBoJoinByInviteLinkParam(req)
	reply, err := s.orgBiz.JoinByInviteLink(ctx, param)
	if err != nil {
		return nil, err
	}
	return &resourcev1.JoinByInviteLinkResp{
		Data: dto.OrgDto.ToPbJoinByInviteLinkRespData(reply),
	}, nil
}

func (s *orgV1Service) AgreeOrRefuseInvite(ctx context.Context, req *resourcev1.AgreeOrRefuseInviteReq) (*resourcev1.AgreeOrRefuseInviteResp, error) {
	param := dto.OrgDto.ToBoAgreeOrRefuseInviteParam(req)
	reply, err := s.orgBiz.AgreeOrRefuseInvite(ctx, param)
	if err != nil {
		return nil, err
	}
	return &resourcev1.AgreeOrRefuseInviteResp{
		Data: dto.OrgDto.ToPbAgreeOrRefuseInviteRespData(reply),
	}, nil
}

func (s *orgV1Service) GetOrgInviteRecordInfo(ctx context.Context, req *resourcev1.GetOrgInviteRecordInfoReq) (*resourcev1.GetOrgInviteRecordInfoResp, error) {
	dataModel, err := s.orgBiz.GetInviteRecordInfo(ctx, req.InviteId)
	if err != nil {
		return nil, err
	}
	return &resourcev1.GetOrgInviteRecordInfoResp{
		Data: dto.OrgDto.ToPbOrgInviteRecord(dataModel),
	}, nil
}

func (s *orgV1Service) GetOrgInviteRecordInfoList(ctx context.Context, req *resourcev1.GetOrgInviteRecordInfoListReq) (*resourcev1.GetOrgInviteRecordInfoListResp, error) {
	dataModels, err := s.orgBiz.GetInviteRecordInfoList(ctx, req.InviteIds)
	if err != nil {
		return nil, err
	}
	return &resourcev1.GetOrgInviteRecordInfoListResp{
		Data: dto.OrgDto.ToPbOrgInviteRecordList(dataModels),
	}, nil
}

func (s *orgV1Service) GetOrgInviteRecordList(ctx context.Context, req *resourcev1.GetOrgInviteRecordListReq) (*resourcev1.GetOrgInviteRecordListResp, error) {
	// paging args
	pageRequest, pageOption := pagepkg.ParsePageRequest(req.PageRequest)

	// list
	param := dto.OrgDto.ToBoOrgInviteRecordListParam(req)
	param.PaginatorArgs = &gormpkg.PaginatorArgs{
		PageRequest: pageRequest,
		PageOption:  pageOption,
	}
	dataModels, dataCount, err := s.orgBiz.ListOrgInviteRecord(ctx, param)
	if err != nil {
		return nil, err
	}

	// paging result
	return &resourcev1.GetOrgInviteRecordListResp{
		Data: &resourcev1.GetOrgInviteRecordListRespData{
			List:     dto.OrgDto.ToPbOrgInviteRecordList(dataModels),
			PageInfo: pagepkg.CalcPageResponse(pageRequest, uint32(dataCount)),
		},
	}, nil
}
