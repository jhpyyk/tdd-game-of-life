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
	t.Run("test parsing cells", func(t *testing.T) {
		blockPath := filepath.Join("patterns", "block.rle")
		parsed := ParseRleFile(blockPath)

		want := "2o$2o"
		assertPatternCellsStringsEqual(t, want, parsed)
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
