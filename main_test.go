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
	t.Run("test block for 5 generations", func(t *testing.T) {
		cmd := exec.Command("go", "run", "main.go", "patterns/block.rle", "5")

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
	t.Run("test die hard should be empty after 130 generations", func(t *testing.T) {
		cmd := exec.Command("go", "run", "main.go", "patterns/die_hard.rle", "130")

		out, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatalf("error while running main %q", err.Error())
		}

		got := string(out)
		want := "x = 0, y = 0\n\n"

		if got != want {
			t.Fatalf("wrong output, wanted %q, got %q", want, got)
		}
	})
	t.Run("test die hard should be empty after 150 generations", func(t *testing.T) {
		cmd := exec.Command("go", "run", "main.go", "patterns/die_hard.rle", "130")

		out, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatalf("error while running main %q", err.Error())
		}

		got := string(out)
		want := "x = 0, y = 0\n\n"

		if got != want {
			t.Fatalf("wrong output, wanted %q, got %q", want, got)
		}
	})
}
