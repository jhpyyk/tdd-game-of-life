package rle_parser_test

import (
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

func TestParseRleFiles(t *testing.T) {
	type TestCase struct {
		name    string
		path    string
		x       int
		y       int
		pattern string
	}

	testCases := []TestCase{
		{"block", filepath.Join("../", "patterns", "block.rle"), 2, 2, "2o$2o!"},
		{"beehive", filepath.Join("../", "patterns", "beehive.rle"), 4, 3, "b2ob$o2bo$b2o!"},
		{"glider gun", filepath.Join("../", "patterns", "gosper_glider_gun.rle"), 36, 9, "24bo11b$22bobo11b$12b2o6b2o12b2o$11bo3bo4b2o12b2o$2o8bo5bo3b2o14b$2o8bo3bob2o4bobo11b$10bo5bo7bo11b$11bo3bo20b$12b2o!"},
	}

	for _, testCase := range testCases {
		t.Run("test parsing pattern string", func(t *testing.T) {
			patternData, err := parser.ParseRleFile(testCase.path)
			if err != nil {
				t.Fatalf("error parsing file")
			}
			assertStringsEqual(t, testCase.pattern, patternData.PatternString)
		})

		t.Run("test parsing x dimension", func(t *testing.T) {
			patternData, err := parser.ParseRleFile(testCase.path)
			if err != nil {
				t.Fatalf("error parsing file")
			}

			want := testCase.x
			if patternData.X != want {
				t.Fatalf("incorrect pattern x dimension, wanted %v, got %v", want, patternData.X)
			}
		})

		t.Run("test parsing y dimension", func(t *testing.T) {
			patternData, err := parser.ParseRleFile(testCase.path)
			if err != nil {
				t.Fatalf("error parsing file")
			}

			want := testCase.y
			if patternData.Y != want {
				t.Fatalf("incorrect pattern y dimension, wanted %v, got %v", want, patternData.Y)
			}
		})
	}
}
