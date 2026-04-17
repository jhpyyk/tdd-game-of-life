package visualizer

import (
	"fmt"
	"io"
)

const (
	CursorToHomeString   = "\x1b[H"
	EnterAltScreenString = "\033[?1049h"
	ExitAltScreenString  = "\033[?1049l"
	ShowCursorString     = "\033[?25l"
	HideCursorString     = "\033[?25h"
)

type Frame struct {
	Header  string
	Content string
}

func EnterAltScreen() {
	fmt.Print(EnterAltScreenString)
}

func ExitAltScreen() {
	fmt.Print(ExitAltScreenString)
}

func HideCursor() {
	fmt.Print(HideCursorString)
}
func ShowCursor() {
	fmt.Print(ShowCursorString)
}

func DrawFrame(writer io.Writer, frame Frame) {
	fmt.Fprint(writer, CursorToHomeString)
	fmt.Fprintln(writer, frame.Header)
	fmt.Fprint(writer, frame.Content)
}
