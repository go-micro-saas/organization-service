package bo

import (
	enumv1 "github.com/go-micro-saas/organization-service/api/org-service/v1/enums"
	"github.com/go-micro-saas/organization-service/app/org-service/internal/data/po"
	gormpkg "github.com/ikaiguang/go-srv-kit/data/gorm"
	"time"
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

type AddEmployeeParam struct {
	OperatorUid  uint64                                     // 操作人ID
	OrgId        uint64                                     // 组织ID
	UserId       uint64                                     // 成员ID
	UserName     string                                     // 成员名称
	EmployeeRole enumv1.OrgEmployeeRoleEnum_OrgEmployeeRole // 成员角色
	UserAvatar   string                                     // 成员头像
	UserPhone    string                                     // 成员电话
	UserEmail    string                                     // 成员邮箱
}

func (s *AddEmployeeParam) NewEmployeeModel() *po.OrgEmployee {
	employeeModel := po.DefaultOrgEmployee()
	employeeModel.EmployeeId = 0
	employeeModel.OrgId = s.OrgId
	employeeModel.UserId = s.UserId
	employeeModel.EmployeeUuid = employeeModel.GenUUID()
	employeeModel.EmployeeName = s.UserName
	employeeModel.EmployeeAvatar = s.UserAvatar
	employeeModel.EmployeePhone = s.UserPhone
	employeeModel.EmployeeEmail = s.UserEmail
	employeeModel.EmployeeRole = s.EmployeeRole
	employeeModel.EmployeeStatus = enumv1.OrgEmployeeStatusEnum_ENABLE

	return employeeModel
}

type AddEmployeeReply struct {
	OrgId          uint64
	UserId         uint64
	EmployeeId     uint64
	EmployeeName   string
	EmployeeAvatar string
	EmployeeStatus enumv1.OrgEmployeeStatusEnum_OrgEmployeeStatus // 成员状态
	EmployeeRole   enumv1.OrgEmployeeRoleEnum_OrgEmployeeRole     // 成员角色
}

func (s *AddEmployeeReply) SetByEmployee(employeeModel *po.OrgEmployee) {
	s.OrgId = employeeModel.OrgId
	s.UserId = employeeModel.UserId
	s.EmployeeId = employeeModel.EmployeeId
	s.EmployeeName = employeeModel.EmployeeName
	s.EmployeeAvatar = employeeModel.EmployeeAvatar
	s.EmployeeStatus = employeeModel.EmployeeStatus
	s.EmployeeRole = employeeModel.EmployeeRole
}

type InviteEmployeeParam struct {
	InviteID uint64
}

type CreateInviteRecordForLinkParam struct {
	OperatorUid        uint64                                     // 操作人ID
	OrgId              uint64                                     // 组织ID
	ExpireTime         time.Time                                  //
	InviteEmployeeRole enumv1.OrgEmployeeRoleEnum_OrgEmployeeRole //
}

type CreateInviteRecordForEmployeeParam struct {
	OperatorUid        uint64    // 操作人ID
	OrgId              uint64    // 组织ID
	ExpireTime         time.Time //
	InviteUserId       uint64
	InviteAccount      string
	InviteAccountType  enumv1.OrgInviteAccountTypeEnum_OrgInviteAccountType
	InviteEmployeeRole enumv1.OrgEmployeeRoleEnum_OrgEmployeeRole
}

type JoinByInviteLinkParam struct {
	InviteId   uint64
	InviteCode string
	UserId     uint64 // 成员ID
	UserName   string // 成员名称
	UserAvatar string // 成员头像
	UserPhone  string // 成员电话
	UserEmail  string // 成员邮箱
}

type AgreeOrRefuseInviteParam struct {
	InviteId   uint64
	InviteCode string
	IsAgree    bool
	UserId     uint64 // 成员ID
	UserName   string // 成员名称
	UserAvatar string // 成员头像
	UserPhone  string // 成员电话
	UserEmail  string // 成员邮箱
}

type OrgListParam struct {
	OrgIDList []uint64 // 用户ID列表
	OrgName   string   // 组织名称

	PaginatorArgs *gormpkg.PaginatorArgs
}

type OrgEmployeeListParam struct {
	OrgIDList      []uint64 // 用户ID列表
	EmployeeIDList []uint64 // 成员ID列表
	UserIDList     []uint64 // 用户ID列表
	EmployeeName   string   // 成员名称

	PaginatorArgs *gormpkg.PaginatorArgs
}

type OrgInviteRecordListParam struct {
	OrgIDList         []uint64 // 用户ID列表
	InviterUserIDList []uint64 // 邀请者用户ID列表
	InviteIDList      []uint64 // 邀请ID列表
	InviteCode        string   // 邀请码
	InviteAccount     string   // 邀请账号

	PaginatorArgs *gormpkg.PaginatorArgs
}
