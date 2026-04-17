package main

import (
	"os"
	"os/exec"
	"path/filepath"
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

	t.Run("test output can be used as input again", func(t *testing.T) {
		cmd1 := exec.Command("go", "run", "main.go", "patterns/block.rle", "1")
		out1, err := cmd1.CombinedOutput()
		if err != nil {
			t.Fatalf("error while running main %q", err.Error())
		}

		tmpDir := t.TempDir()
		tmpPath := filepath.Join(tmpDir, "out.rle")
		if err := os.WriteFile(tmpPath, out1, 0o600); err != nil {
			t.Fatalf("failed to write temp file: %v", err)
		}

		cmd2 := exec.Command("go", "run", "main.go", tmpPath, "1")

		out, err := cmd2.CombinedOutput()
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
