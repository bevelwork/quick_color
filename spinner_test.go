package quick_color

import (
	"bytes"
	"testing"
	"time"
)

func TestSpinnerWritesFramesAndStops(t *testing.T) {
	var buf bytes.Buffer
	sp := NewSpinner(&buf, "Testing...", 50*time.Millisecond)
	sp.Start()
	time.Sleep(160 * time.Millisecond)
	sp.Stop()
	if buf.Len() == 0 {
		t.Fatal("expected spinner to write output")
	}
}

func TestWithProgress(t *testing.T) {
	var buf bytes.Buffer
	res, err := WithProgress(&buf, "Work", 20*time.Millisecond, func() (int, error) {
		time.Sleep(70 * time.Millisecond)
		return 42, nil
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if res != 42 {
		t.Fatalf("expected 42, got %d", res)
	}
	if buf.Len() == 0 {
		t.Fatal("expected spinner output during progress")
	}
}

func TestSpinnerVisualFormat(t *testing.T) {
	var buf bytes.Buffer
	sp := NewSpinner(&buf, "X", 10*time.Millisecond)
	// Use a single deterministic frame to simplify assertions
	sp.frames = []string{"*"}
	sp.Start()
	time.Sleep(30 * time.Millisecond)
	sp.Stop()
	out := buf.String()
	if !bytes.Contains([]byte(out), []byte("\r* X")) {
		t.Fatalf("expected visual frame + message, got: %q", out)
	}
	if !bytes.Contains([]byte(out), []byte("\r\033[K")) {
		t.Fatalf("expected clear-line sequence on stop, got: %q", out)
	}
}
