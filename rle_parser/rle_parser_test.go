package rle_parser_test

import (
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	parser "github.com/jhpyyk/tdd-game-of-life/rle_parser"
)

func assertStringsEqual(t testing.TB, want string, got string) {
	t.Helper()
	if want != got {
		t.Fatalf("Patterns are not equal, wanted %q, got %q", want, got)
	}
}

func TestParseBlockRleFile(t *testing.T) {
	blockPath := filepath.Join("../", "patterns", "block.rle")
	t.Run("test parsing pattern string", func(t *testing.T) {
		patternData := parser.ParseRleFile(blockPath)

		want := "2o$2o"
		assertStringsEqual(t, want, patternData.PatternString)
	})

	t.Run("test parsing x dimension", func(t *testing.T) {
		patternData := parser.ParseRleFile(blockPath)

		want := 2
		if patternData.X != want {
			t.Fatalf("incorrect pattern x dimension, wanted %v, got %v", want, patternData.X)
		}
	})

	t.Run("test parsing y dimension", func(t *testing.T) {
		patternData := parser.ParseRleFile(blockPath)

		want := 2
		if patternData.Y != want {
			t.Fatalf("incorrect pattern y dimension, wanted %v, got %v", want, patternData.Y)
		}
	})
}

func TestParseBeehiveRleFile(t *testing.T) {
	beehivePath := filepath.Join("../", "patterns", "beehive.rle")
	t.Run("test parsing pattern string", func(t *testing.T) {
		patternData := parser.ParseRleFile(beehivePath)

		want := "b2ob$o2bo$b2o"
		assertStringsEqual(t, want, patternData.PatternString)
	})

	t.Run("test parsing x dimension", func(t *testing.T) {
		patternData := parser.ParseRleFile(beehivePath)

		want := 4
		if patternData.X != want {
			t.Fatalf("incorrect pattern x dimension, wanted %v, got %v", want, patternData.X)
		}
	})

	t.Run("test parsing y dimension", func(t *testing.T) {
		patternData := parser.ParseRleFile(beehivePath)

		want := 3
		if patternData.Y != want {
			t.Fatalf("incorrect pattern y dimension, wanted %v, got %v", want, patternData.Y)
		}
	})
}

func TestParseRleFileCrash(t *testing.T) {
	if os.Getenv("BE_CRASHER") == "1" {
		parser.ParseRleFile("invalidpath")
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
