package global

const (
	KeyPrefix = "org_s_"
)

func Key(k string) string {
	return KeyPrefix + k
}
