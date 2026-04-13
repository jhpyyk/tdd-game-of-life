package main

import (
	"os/exec"
	"testing"

	helpers "github.com/jhpyyk/tdd-game-of-life/test_helpers"
)

func TestMain(t *testing.T) {
	t.Run("test block for 1 generation", func(t *testing.T) {
		cmd := exec.Command("go", "run", "main.go", "patterns/block.rle", "1")

		out, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatalf("error while running main %q", err.Error())
		}

		got := helpers.StripPattern(t, string(out))
		block := `
		##
		##
		`

		want := helpers.StripPattern(t, block)
		if got != want {
			t.Fatalf("wrong output, wanted %q, got %q", want, got)
		}
	})
}
