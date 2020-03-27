package Util

import "unicode/utf8"

func Substring(s string, a, b int) string {
	return string([]rune(s)[a:b])
}

func Length(s string) int {
	return utf8.RuneCountInString(s)
}
