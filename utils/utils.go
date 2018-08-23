package utils

import "strings"

func Trim(input string) string {
	return strings.TrimRightFunc(input, func(c rune) bool {
		return c == '\r' || c == '\n'
	})
}
