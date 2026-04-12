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
		// {"beehive", 4, 3, "b2ob$o2bo$b2o!"},
		// {"glider gun", 36, 9, "24bo11b$22bobo11b$12b2o6b2o12b2o$11bo3bo4b2o12b2o$2o8bo5bo3b2o14b$2o8bo3bob2o4bobo11b$10bo5bo7bo11b$11bo3bo20b$12b2o!"},
	}
	for _, testCase := range testCases {
		t.Run("test pattern parsing", func(t *testing.T) {

			pattern, err := parser.ParsePattern(testCase.x, testCase.y, testCase.pattern)
			if err != nil {
				t.Fatal(err.Error())
			}

			got := pattern.ToString()
			if got != testCase.want {
				t.Fatalf("Parser failed to parse pattern, want %q, got %q", testCase.want, got)
			}

		})
	}
}
