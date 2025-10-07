package quick_color

import (
	"bytes"
	"strings"
	"testing"
)

func TestSubdisplayAppendTrim(t *testing.T) {
	s := NewSubdisplay(3)
	s.Append("a")
	s.Append("b")
	s.Append("c")
	if len(s.Lines()) != 3 {
		t.Fatal("expected 3 lines")
	}
	s.Append("d")
	lines := s.Lines()
	if len(lines) != 3 {
		t.Fatal("expected trim to 3 lines")
	}
	if strings.Join(lines, ",") != "b,c,d" {
		t.Fatalf("unexpected lines: %v", lines)
	}

}

func TestSubdisplayRender(t *testing.T) {
	s := NewSubdisplay(5)
	s.SetLinePrefix("  ")
	s.Append("x")
	var buf bytes.Buffer
	s.Render(&buf, "Title")
	out := buf.String()
	if !strings.Contains(out, "Title") {
		t.Fatal("expected title")
	}
	if !strings.Contains(out, "  x\n") {
		t.Fatal("expected prefixed line")
	}
}
