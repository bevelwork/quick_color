package quick_color

import (
	"fmt"
	"io"
)

const (
	ansiSaveCursor    = "\033[s"
	ansiRestoreCursor = "\033[u"
	ansiClearLine     = "\033[K"
)

// PinToTop writes text pinned to the first terminal line without moving the current cursor.
// It saves the current cursor, moves to row 1 col 1, clears the line, writes text, then restores.
func PinToTop(w io.Writer, text string) {
	// Save current cursor position
	fmt.Fprint(w, ansiSaveCursor)
	// Move to top-left (row 1, col 1)
	fmt.Fprint(w, "\033[1;1H")
	// Clear the line and write text
	fmt.Fprint(w, "\r")
	fmt.Fprint(w, ansiClearLine)
	fmt.Fprint(w, text)
	// Restore original cursor position
	fmt.Fprint(w, ansiRestoreCursor)
}

// PinToBottom writes text pinned to the last terminal line without moving the current cursor.
// It saves the current cursor, moves to a very large row (clamped by terminal), clears the line,
// writes text, then restores. This is a common trick to address the bottom-most line.
func PinToBottom(w io.Writer, text string) {
	// Save current cursor position
	fmt.Fprint(w, ansiSaveCursor)
	// Move to last visible row, column 1 (large row index clamps to bottom on most terminals)
	fmt.Fprint(w, "\033[999;1H")
	// Clear the line and write text
	fmt.Fprint(w, "\r")
	fmt.Fprint(w, ansiClearLine)
	fmt.Fprint(w, text)
	// Restore original cursor position
	fmt.Fprint(w, ansiRestoreCursor)
}
