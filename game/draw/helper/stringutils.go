package helper

import "unicode/utf8"

// Substring slices string from a to b
func Substring(s string, a, b int) string {
	return string([]rune(s)[a:b])
}

// Length calculates length of string
func Length(s string) int {
	return utf8.RuneCountInString(s)
}
