package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

func assertPatternCellsStringsEqual(t testing.TB, want string, got string) {
	t.Helper()
	if want != got {
		t.Fatalf("Patterns are not equal, wanted %q, got %q", want, got)
	}
}

func TestParseRleFile(t *testing.T) {
	blockPath := filepath.Join("patterns", "block.rle")
	t.Run("test parsing pattern string", func(t *testing.T) {
		patternData := ParseRleFile(blockPath)

		want := "2o$2o"
		assertPatternCellsStringsEqual(t, want, patternData.patternString)
	})

	t.Run("test parsing x dimension", func(t *testing.T) {
		patternData := ParseRleFile(blockPath)

		want := 2
		if patternData.x != want {
			t.Fatalf("incorrect pattern x dimension, wanted %v, got %v", want, patternData.x)
		}
	})

}

func TestParseRleFileCrash(t *testing.T) {
	if os.Getenv("BE_CRASHER") == "1" {
		ParseRleFile("invalidpath")
		return
	}
	t.Run("Test program exits when invalid path", func(t *testing.T) {
		cmd := exec.Command(os.Args[0], "-test.run=TestParseRleFileCrash")
		cmd.Env = append(os.Environ(), "BE_CRASHER=1")
		err := cmd.Run()
		if e, ok := err.(*exec.ExitError); ok && !e.Success() {
			return
		}
		t.Fatalf("process ran with err %v, want exit status 1", err)
	})
}
