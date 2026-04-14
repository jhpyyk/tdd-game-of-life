package utils

import (
	"strings"
)

func StripPattern(pattern string) string {
	var sb strings.Builder
	for _, c := range pattern {
		if c == '.' || c == '#' {
			sb.WriteRune(c)
		}
	}
	return sb.String()
}
