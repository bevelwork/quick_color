package quick_color

import (
	"bytes"
	"strings"
	"testing"
)

func TestRenderAlternatingList(t *testing.T) {
	items := []string{"one", "two", "three"}
	var buf bytes.Buffer
	RenderAlternatingList(&buf, items, "", "")
	out := buf.String()
	lines := strings.Split(strings.TrimSpace(out), "\n")
	if len(lines) != 3 {
		t.Fatalf("expected 3 lines, got %d", len(lines))
	}
	if !strings.Contains(lines[0], ColorWhite) {
		t.Fatal("line 0 should be white")
	}
	if !strings.Contains(lines[1], ColorCyan) {
		t.Fatal("line 1 should be cyan")
	}
}
