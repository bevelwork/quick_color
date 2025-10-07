package quick_color

import "testing"

func TestColor(t *testing.T) {
	got := Color("hello", ColorGreen)
	wantPrefix := ColorGreen + "hello"
	if got[:len(wantPrefix)] != wantPrefix {
		t.Fatalf("expected prefix %q, got %q", wantPrefix, got)
	}
	if got[len(got)-len(ColorReset):] != ColorReset {
		t.Fatalf("expected suffix reset, got %q", got)
	}
}

func TestColorizeBold(t *testing.T) {
	got := ColorizeBold("world", ColorBlue)
	if got[:len(ColorBlue)] != ColorBlue {
		t.Fatalf("expected blue prefix")
	}
	if got[len(got)-len(ColorReset):] != ColorReset {
		t.Fatalf("expected reset suffix")
	}
}

func TestAlternatingColor(t *testing.T) {
	if AlternatingColor(0, ColorWhite, ColorCyan) != ColorWhite {
		t.Fatal("even index should use evenColor")
	}
	if AlternatingColor(1, ColorWhite, ColorCyan) != ColorCyan {
		t.Fatal("odd index should use oddColor")
	}
}
