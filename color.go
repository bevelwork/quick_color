package quick_color

// ANSI color codes and helpers centralizing patterns used across quick_* tools.

const (
	ColorReset  = "\033[0m"
	ColorRed    = "\033[31m"
	ColorGreen  = "\033[32m"
	ColorYellow = "\033[33m"
	ColorBlue   = "\033[34m"
	ColorPurple = "\033[35m"
	ColorCyan   = "\033[36m"
	ColorWhite  = "\033[37m"
	ColorBold   = "\033[1m"
)

// Color wraps text with the specified ANSI color code and resets at the end.
func Color(text, colorCode string) string {
	return colorCode + text + ColorReset
}

// ColorizeBold wraps text with the specified ANSI color code, adds bold, and resets.
func ColorizeBold(text, colorCode string) string {
	return colorCode + ColorBold + text + ColorReset
}

// AlternatingColor returns two colors alternating by index (even/odd).
// Defaults mirror existing quick_* tools: white for even, cyan for odd.
func AlternatingColor(index int, evenColor, oddColor string) string {
	if index%2 == 0 {
		return evenColor
	}
	return oddColor
}
