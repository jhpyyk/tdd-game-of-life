package main

import (
	"os/exec"
	"testing"

	"github.com/jhpyyk/tdd-game-of-life/utils"
)

func TestMain(t *testing.T) {
	t.Run("test block for 1 generation", func(t *testing.T) {
		cmd := exec.Command("go", "run", "main.go", "patterns/block.rle", "1")

		out, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatalf("error while running main %q", err.Error())
		}

		got := utils.StripPattern(string(out))
		want := "x = 2, y = 2\n2o$2o!"

		if got != want {
			t.Fatalf("wrong output, wanted %q, got %q", want, got)
		}
	})
}
