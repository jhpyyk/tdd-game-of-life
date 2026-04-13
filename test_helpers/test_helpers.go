package test_helpers

import (
	"strings"
	"testing"
)

func StripPattern(t testing.TB, pattern string) string {
	t.Helper()
	var sb strings.Builder
	for _, c := range pattern {
		if c == '.' || c == '#' {
			sb.WriteRune(c)
		}
	}
	return sb.String()
}
