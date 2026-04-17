package visualizer_test

import (
	"bytes"
	"testing"

	"github.com/jhpyyk/tdd-game-of-life/visualizer"
)

func TestVisualizer(t *testing.T) {
	t.Run("test visualizer draws frame", func(t *testing.T) {
		buffer := bytes.Buffer{}
		visualizer.DrawFrame(&buffer, "#")
		got := buffer.String()

		want := visualizer.CursorToHome + "#"

		if got != want {
			t.Errorf("wanted %q, got %q", want, got)
		}
	})
	t.Run("test visualizer clears frame before drawing new one", func(t *testing.T) {
		buffer := bytes.Buffer{}
		visualizer.DrawFrame(&buffer, "#")
		visualizer.DrawFrame(&buffer, "#")

		got := buffer.String()
		want := visualizer.CursorToHome + "#" + visualizer.CursorToHome + "#"

		if got != want {
			t.Errorf("wanted %q, got %q", want, got)
		}
	})
}
