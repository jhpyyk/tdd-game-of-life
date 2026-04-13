package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestMain(t *testing.T) {
	t.Run("test block", func(t *testing.T) {
		cmd := exec.Command("go", "run", "main.go", "patterns/block.rle", "1")

		out, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatalf("error while running main %q", err.Error())
		}

		got := stripPattern(t, string(out))
		block := `
		##
		##
		`

		want := stripPattern(t, block)
		if got != want {
			t.Fatalf("wrong output, wanted %q, got %q", want, got)
		}
	})
}

func stripPattern(t testing.TB, pattern string) string {
	t.Helper()
	var sb strings.Builder
	for _, c := range pattern {
		if c == '.' || c == '#' {
			sb.WriteRune(c)
		}
	}
	return sb.String()
}
