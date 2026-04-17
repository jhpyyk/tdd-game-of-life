package visualizer

import (
	"fmt"
	"io"
)

const (
	CursorToHome = "\x1b[H"
)

func DrawFrame(writer io.Writer, frame string) {
	fmt.Fprint(writer, CursorToHome)
	fmt.Fprint(writer, frame)
}
