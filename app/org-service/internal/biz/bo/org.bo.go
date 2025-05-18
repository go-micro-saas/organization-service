package bo

import (
	enumv1 "github.com/go-micro-saas/organization-service/api/org-service/v1/enums"
	"github.com/go-micro-saas/organization-service/app/org-service/internal/data/po"
)

type Testdata struct {
}

type HelloWorldParam struct {
	Message string
}

type HelloWorldReply struct {
	Message string
}

type CreateOrgParam struct {
	CreatorID     uint64
	CreatorName   string
	CreatorAvatar string

	OrgName   string
	OrgAvatar string
	OrgType   enumv1.OrgTypeEnum_OrgType

	SkipCreateEmployee bool
}

type CreateOrgReply struct {
	OrgId     uint64
	OrgName   string
	OrgAvatar string
	OrgType   enumv1.OrgTypeEnum_OrgType
}

func (s *CreateOrgReply) SetByOrg(orgModel *po.Org) {
	s.OrgId = orgModel.OrgId
	s.OrgName = orgModel.OrgName
	s.OrgAvatar = orgModel.OrgAvatar
	s.OrgType = orgModel.OrgType
}
