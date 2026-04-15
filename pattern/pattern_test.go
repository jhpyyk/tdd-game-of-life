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
	blocksSeparated = `
		##...##
		##...##
		`

	blockNoCorner = `
		#.
		##
		`

	glider = `
		.#.
		..#
		###
		`

	gliderAfter1 = `
			#.#
			.##
			.#.
		`
	gliderAfter2 = `
			..#
			#.#
			.##
		`
	gliderAfter3 = `
			#..
			.##
			##.
		`

	// dieHard = `
	// ......#.
	// ##......
	// .#...###
	// `

	// dieHard129 = `
	// #
	// #
	// `
	// dieHard130 = ""

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
		name        string
		x           int
		y           int
		generations int
		pattern     string
		want        string
	}

	testCases := []TestCase{
		// {"block", 2, 2, 1, block, block},
		// {"block stays still", 2, 2, 50, block, block},
		// {"block without corner", 2, 2, 1, blockNoCorner, block},
		// {"block without corner stays still", 2, 2, 50, blockNoCorner, block},
		// {"glider after 1", 3, 3, 1, glider, gliderAfter1},
		// {"glider after 2", 3, 3, 2, glider, gliderAfter2},
		// {"glider after 3", 3, 3, 3, glider, gliderAfter3},
		// {"glider after 4", 3, 3, 4, glider, glider},
		// {"glider after 8", 3, 3, 4, glider, glider},
		// {"glider after 16", 3, 3, 4, glider, glider},
		// {"separated blocks", 7, 2, 1, blocksSeparated, blocksSeparated},
		// {"die hard generation before vanishing", 8, 3, 129, dieHard, dieHard129},
		// {"die hard empty", 8, 3, 130, dieHard, dieHard130},
	}
	for _, testCase := range testCases {
		t.Run("test pattern generation", func(t *testing.T) {

			pattern, err := pattern.FromString(testCase.x, testCase.y, testCase.pattern)
			if err != nil {
				t.Fatal(err.Error())
			}

			for range testCase.generations {
				pattern = pattern.GetNextGeneration()
			}
			got := utils.StripPattern(pattern.ToString())
			want := utils.StripPattern(testCase.want)
			if got != want {
				t.Fatalf("failed to get next generation for %q, want %q, got %q", testCase.name, want, got)
			}

		})
	}
}

func TestPatternStringConversion(t *testing.T) {
	t.Run("converts block", func(t *testing.T) {
		pattern, err := pattern.FromString(2, 2, block)
		if err != nil {
			t.Fatal(err.Error())
		}

		got := utils.StripPattern(pattern.ToString())
		want := utils.StripPattern(block)
		if got != want {
			t.Fatalf("Parser failed to parse pattern, want %q, got %q", want, got)
		}
	})

	t.Run("converts two blocks", func(t *testing.T) {
		pattern, err := pattern.FromString(7, 2, blocksSeparated)
		if err != nil {
			t.Fatal(err.Error())
		}

		got := utils.StripPattern(pattern.ToString())
		want := utils.StripPattern(blocksSeparated)
		if got != want {
			t.Fatalf("Parser failed to parse pattern, want %q, got %q", want, got)
		}
	})
}
