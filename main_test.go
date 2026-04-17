package main

import (
	"os/exec"
	"testing"
)

func TestMain(t *testing.T) {
	t.Run("test block for 1 generation", func(t *testing.T) {
		cmd := exec.Command("go", "run", "main.go", "patterns/block.rle", "1")

		out, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatalf("error while running main %q", err.Error())
		}

		got := string(out)
		want := "x = 2, y = 2\n2o$2o!\n"

		if got != want {
			t.Fatalf("wrong output, wanted %q, got %q", want, got)
		}
	})
}
