package quick_color

import (
	"fmt"
	"io"
	"time"
)

// Spinner provides a throbber/spinner suitable for CLI feedback.
// It writes frames to the provided writer until Stop is called.
type Spinner struct {
	frames   []string
	message  string
	writer   io.Writer
	interval time.Duration
	stopCh   chan struct{}
	doneCh   chan struct{}
}

// DefaultSpinnerFrames mirrors the braille spinner used in quick_ecs.
var DefaultSpinnerFrames = []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"}

// NewSpinner constructs a Spinner.
func NewSpinner(w io.Writer, message string, interval time.Duration) *Spinner {
	return &Spinner{
		frames:   DefaultSpinnerFrames,
		message:  message,
		writer:   w,
		interval: interval,
		stopCh:   make(chan struct{}),
		doneCh:   make(chan struct{}),
	}
}

// Start begins the spinner loop in a goroutine. Safe to call once.
func (s *Spinner) Start() {
	go func() {
		i := 0
		t := time.NewTicker(s.interval)
		defer t.Stop()
		defer close(s.doneCh)
		for {
			select {
			case <-s.stopCh:
				// Clear line
				fmt.Fprint(s.writer, "\r\033[K")
				return
			case <-t.C:
				fmt.Fprintf(s.writer, "\r%s %s", s.frames[i%len(s.frames)], s.message)
				i++
			}
		}
	}()
}

// Stop signals the spinner to terminate and clears the line.
func (s *Spinner) Stop() {
	select {
	case <-s.stopCh:
		// already stopped
	default:
		close(s.stopCh)
	}
	<-s.doneCh
}

// WithProgress runs fn while showing a spinner with message.
// It clears the line before returning.
func WithProgress[T any](w io.Writer, message string, interval time.Duration, fn func() (T, error)) (T, error) {
	sp := NewSpinner(w, message, interval)
	sp.Start()
	res, err := fn()
	sp.Stop()
	return res, err
}
