package visualizer_test

import (
	"bytes"
	"testing"

	"github.com/jhpyyk/tdd-game-of-life/visualizer"
)

func TestVisualizer(t *testing.T) {
	t.Run("should draw frame", func(t *testing.T) {
		buffer := bytes.Buffer{}
		visualizer.DrawFrame(&buffer, "#")
		got := buffer.String()

		want := visualizer.CursorToHome + "#"

		if got != want {
			t.Errorf("wanted %q, got %q", want, got)
		}
	})
	t.Run("should start drawing next frame from the start", func(t *testing.T) {
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
