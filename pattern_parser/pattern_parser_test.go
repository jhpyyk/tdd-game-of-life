package pattern_parser_test

import (
	"strings"
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
	block := `
			##
			##
			`
	short := `
	.#
	`
	preceeding := `
	............#
	............#
	`
	trailing := `
	#............
	#............
	`
	beehive := `
	.##.
	#..#
	.##.
	`
	testCases := []TestCase{
		{"block", 2, 2, "2o$2o!", block},
		{"short", 2, 1, "bo!", short},
		{"preceeding", 13, 2, "12bo$12bo!", preceeding},
		{"trailing", 13, 2, "o$o!", trailing},
		{"beehive", 4, 3, "b2ob$o2bo$b2o!", beehive},
	}
	for _, testCase := range testCases {
		t.Run("test pattern parsing", func(t *testing.T) {

			pattern, err := parser.ParsePattern(testCase.x, testCase.y, testCase.pattern)
			if err != nil {
				t.Fatal(err.Error())
			}

			got := stripPattern(t, pattern.ToString())
			want := stripPattern(t, testCase.want)
			if got != want {
				t.Fatalf("Parser failed to parse pattern %q, want %q, got %q", testCase.name, want, got)
			}

		})
	}
}

func stripPattern(t testing.TB, pattern string) string {
	t.Helper()
	var sb strings.Builder
	for _, c := range pattern {
		if c == '.' || c == '#' {
			sb.WriteRune(c)
		}
	}
	return sb.String()
}
