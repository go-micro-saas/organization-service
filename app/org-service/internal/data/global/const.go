package global

const (
	KeyPrefix = "org_service_"
)

func Key(k string) string {
	return KeyPrefix + k
}
