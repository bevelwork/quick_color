package quick_color

import (
	"bytes"
	"testing"
)

func TestPinToTopSequences(t *testing.T) {
	var buf bytes.Buffer
	PinToTop(&buf, "TOP")
	out := buf.String()
	if want := "\033[s"; !bytes.Contains([]byte(out), []byte(want)) {
		t.Fatalf("missing save cursor: %q", out)
	}
	if want := "\033[1;1H"; !bytes.Contains([]byte(out), []byte(want)) {
		t.Fatalf("missing move to top: %q", out)
	}
	if want := "\033[K"; !bytes.Contains([]byte(out), []byte(want)) {
		t.Fatalf("missing clear line: %q", out)
	}
	if want := "TOP"; !bytes.Contains([]byte(out), []byte(want)) {
		t.Fatalf("missing text: %q", out)
	}
	if want := "\033[u"; !bytes.Contains([]byte(out), []byte(want)) {
		t.Fatalf("missing restore cursor: %q", out)
	}
}

func TestPinToBottomSequences(t *testing.T) {
	var buf bytes.Buffer
	PinToBottom(&buf, "BOTTOM")
	out := buf.String()
	if want := "\033[s"; !bytes.Contains([]byte(out), []byte(want)) {
		t.Fatalf("missing save cursor: %q", out)
	}
	if want := "\033[999;1H"; !bytes.Contains([]byte(out), []byte(want)) {
		t.Fatalf("missing move to bottom-ish: %q", out)
	}
	if want := "\033[K"; !bytes.Contains([]byte(out), []byte(want)) {
		t.Fatalf("missing clear line: %q", out)
	}
	if want := "BOTTOM"; !bytes.Contains([]byte(out), []byte(want)) {
		t.Fatalf("missing text: %q", out)
	}
	if want := "\033[u"; !bytes.Contains([]byte(out), []byte(want)) {
		t.Fatalf("missing restore cursor: %q", out)
	}
}
