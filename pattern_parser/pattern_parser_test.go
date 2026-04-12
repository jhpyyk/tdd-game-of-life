package pattern_parser_test

import (
	"testing"

	parser "github.com/jhpyyk/tdd-game-of-life/pattern_parser"
)

func TestPatternParser(t *testing.T) {
	type TestCase struct {
		name    string
		x       int
		y       int
		pattern string
		want    string
	}
	testCases := []TestCase{
		{"block", 2, 2, "2o$2o!", "##\n##\n"},
		{"short", 2, 1, "bo!", ".#\n"},
		{"preceeding", 13, 1, "12bo!", "............#\n"},
		{"trailing", 13, 1, "o!", "#............\n"},
		{"beehive", 4, 3, "b2ob$o2bo$b2o!", ".##.\n#..#\n.##.\n"},
	}
	for _, testCase := range testCases {
		t.Run("test pattern parsing", func(t *testing.T) {

			pattern, err := parser.ParsePattern(testCase.x, testCase.y, testCase.pattern)
			if err != nil {
				t.Fatal(err.Error())
			}

			got := pattern.ToString()
			if got != testCase.want {
				t.Fatalf("Parser failed to parse pattern %q, want %q, got %q", testCase.name, testCase.want, got)
			}

		})
	}
}
