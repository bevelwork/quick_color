package quick_color

import (
	"fmt"
	"io"
)

// RenderAlternatingList renders items with alternating row colors.
// Colors default to white/cyan if empty strings are provided.
func RenderAlternatingList(w io.Writer, items []string, evenColor, oddColor string) {
	if evenColor == "" {
		evenColor = ColorWhite
	}
	if oddColor == "" {
		oddColor = ColorCyan
	}
	for i, item := range items {
		c := AlternatingColor(i, evenColor, oddColor)
		fmt.Fprintln(w, Color(item, c))
	}
}
