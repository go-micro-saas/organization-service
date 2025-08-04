package po

const (
	DefaultUserCreateOrgMaxCount = 10
	DefaultUserBelongOrgMaxCount = 20
)

type BusinessSetting struct {
	UserCreateOrgMaxCount int32
	UserBelongOrgMaxCount int32
}

func DefaultBusinessSetting() *BusinessSetting {
	return &BusinessSetting{
		UserCreateOrgMaxCount: DefaultUserCreateOrgMaxCount,
		UserBelongOrgMaxCount: DefaultUserBelongOrgMaxCount,
	}
}
