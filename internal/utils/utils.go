package utils

// InArray to check if in array.
func InArray(arr []string, str string) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
}

// Ellipsis cut long string.
func Ellipsis(str string, length int) string {
	if len(str) > length {
		return str[:length] + "..."
	}
	return str
}
