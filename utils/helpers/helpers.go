package helpers

func IsInvalid(value string) bool {
	if value == "" {
		return true
	}
	return false
}
