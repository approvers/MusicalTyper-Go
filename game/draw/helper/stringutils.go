package helper

// Substring slices string from a to b
func Substring(s string, a, b int) string {
	return string([]rune(s)[a:b])
}
