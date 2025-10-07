package quick_color

import (
	"fmt"
	"io"
	"strings"
)

// Subdisplay maintains a scrolling buffer and renders it with optional header/footer.
// Rendering is line-oriented to keep it simple and portable.
type Subdisplay struct {
	buffer     []string
	maxLines   int
	separator  string
	linePrefix string
}

// NewSubdisplay creates a new subdisplay with a maximum number of lines to retain.
func NewSubdisplay(maxLines int) *Subdisplay {
	if maxLines <= 0 {
		maxLines = 100
	}
	return &Subdisplay{maxLines: maxLines, separator: strings.Repeat("-", 60)}

}

// Append adds a line to the buffer, trimming to maxLines.
func (s *Subdisplay) Append(line string) {
	s.buffer = append(s.buffer, line)
	if len(s.buffer) > s.maxLines {
		start := len(s.buffer) - s.maxLines
		s.buffer = s.buffer[start:]
	}
}

// Lines returns a copy of the buffered lines.
func (s *Subdisplay) Lines() []string {
	dup := make([]string, len(s.buffer))
	copy(dup, s.buffer)
	return dup
}

// SetLinePrefix sets a prefix for each rendered line (e.g., indentation or bullet).
func (s *Subdisplay) SetLinePrefix(prefix string) { s.linePrefix = prefix }

// SetSeparator overrides the default horizontal separator.
func (s *Subdisplay) SetSeparator(sep string) { s.separator = sep }

// Render writes the subdisplay content to w with an optional title.
func (s *Subdisplay) Render(w io.Writer, title string) {
	if title != "" {
		fmt.Fprintf(w, "%s\n", Color(title, ColorBlue))
	}
	fmt.Fprintf(w, "%s\n", s.separator)
	for _, l := range s.buffer {
		if s.linePrefix != "" {
			fmt.Fprintf(w, "%s%s\n", s.linePrefix, l)
		} else {
			fmt.Fprintf(w, "%s\n", l)
		}
	}
	fmt.Fprintf(w, "%s\n", s.separator)
}
