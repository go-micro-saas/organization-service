package po

const (
	KeyPrefix = "org_"
)

func Key(k string) string {
	return KeyPrefix + k
}
