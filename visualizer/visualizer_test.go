package visualizer_test

import (
	"bytes"
	"strings"
	"testing"

	"github.com/jhpyyk/tdd-game-of-life/visualizer"
)

func TestVisualizer(t *testing.T) {
	t.Run("should draw frame", func(t *testing.T) {
		frame := visualizer.Frame{
			Header:  "header",
			Content: "#",
		}
		buffer := bytes.Buffer{}
		visualizer.DrawFrame(&buffer, frame)

		got := buffer.String()

		want := "#"

		if !strings.Contains(got, want) {
			t.Errorf("got %q did not contain wanted %q", got, want)
		}
	})
	t.Run("should start drawing next frame from the top left", func(t *testing.T) {
		frame := visualizer.Frame{
			Header:  "header",
			Content: "#",
		}
		buffer := bytes.Buffer{}
		visualizer.DrawFrame(&buffer, frame)
		visualizer.DrawFrame(&buffer, frame)

		got := buffer.String()
		want := visualizer.CursorToHomeString

		if !strings.Contains(got, want) {
			t.Errorf("got %q did not contain wanted %q", got, want)
		}
	})

	t.Run("should contain header", func(t *testing.T) {
		frame := visualizer.Frame{
			Header:  "header",
			Content: "#",
		}
		buffer := bytes.Buffer{}
		visualizer.DrawFrame(&buffer, frame)

		got := buffer.String()

		want := "header"

		if !strings.Contains(got, want) {
			t.Errorf("got %q did not contain wanted %q", got, want)
		}
	})
}
