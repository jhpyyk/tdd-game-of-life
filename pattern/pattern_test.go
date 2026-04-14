package pattern_test

import (
	"testing"

	"github.com/jhpyyk/tdd-game-of-life/pattern"
	"github.com/jhpyyk/tdd-game-of-life/utils"
)

const (
	block = `
			##
			##
			`
	short = `
	.#
	`
	preceeding = `
	............#
	............#
	`
	trailing = `
	#............
	#............
	`
	beehive = `
	.##.
	#..#
	.##.
	`
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
		{"block", 2, 2, "2o$2o!", block},
		{"short", 2, 1, "bo!", short},
		{"preceeding", 13, 2, "12bo$12bo!", preceeding},
		{"trailing", 13, 2, "o$o!", trailing},
		{"beehive", 4, 3, "b2ob$o2bo$b2o!", beehive},
	}
	for _, testCase := range testCases {
		t.Run("test pattern parsing", func(t *testing.T) {

			pattern, err := pattern.ParsePatternFromRLEPatternString(testCase.x, testCase.y, testCase.pattern)
			if err != nil {
				t.Fatal(err.Error())
			}

			got := utils.StripPattern(pattern.ToString())
			want := utils.StripPattern(testCase.want)
			if got != want {
				t.Fatalf("Parser failed to parse pattern %q, want %q, got %q", testCase.name, want, got)
			}

		})
	}
}

func TestPatternGeneration(t *testing.T) {
	type TestCase struct {
		name    string
		x       int
		y       int
		pattern string
		want    string
	}

	testCases := []TestCase{
		{"block", 2, 2, block, block},
	}
	for _, testCase := range testCases {
		t.Run("test pattern generation", func(t *testing.T) {

			pattern, err := pattern.FromString(testCase.x, testCase.y, testCase.pattern)
			if err != nil {
				t.Fatal(err.Error())
			}

			nextGen := pattern.GetNextGeneration()

			got := utils.StripPattern(nextGen.ToString())
			want := utils.StripPattern(testCase.want)
			if got != want {
				t.Fatalf("Parser failed to parse pattern %q, want %q, got %q", testCase.name, want, got)
			}

		})
	}
}

func TestPatternStringConversion(t *testing.T) {
	pattern, err := pattern.FromString(2, 2, block)
	if err != nil {
		t.Fatal(err.Error())
	}

	got := utils.StripPattern(pattern.ToString())
	want := utils.StripPattern(block)
	if got != want {
		t.Fatalf("Parser failed to parse pattern, want %q, got %q", want, got)
	}

}
