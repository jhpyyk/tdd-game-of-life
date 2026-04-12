package pattern_parser_test

import (
	"testing"

	parser "github.com/jhpyyk/tdd-game-of-life/pattern_parser"
)

func TestPatternParser(t *testing.T) {
	t.Run("test pattern parsing", func(t *testing.T) {

		block := "2o$2o!"
		want := "##\n##\n"

		pattern, err := parser.ParsePattern(2, 2, block)
		if err != nil {
			t.Fatal(err.Error())
		}

		got := pattern.ToString()
		if got != want {
			t.Fatalf("Parser failed to parse pattern, want %q, got %q", want, got)
		}

	})
}
