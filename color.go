package quick_color

import "strings"

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

// Style (non-color) codes
const (
	StyleDim       = "\033[2m"
	StyleItalic    = "\033[3m"
	StyleUnderline = "\033[4m"
	StyleInverse   = "\033[7m"
	StyleStrike    = "\033[9m"
)

// Color wraps text with the specified ANSI color code and resets at the end.
func Color(text, colorCode string) string {
	return colorCode + text + ColorReset
}

// ColorizeBold wraps text with the specified ANSI color code, adds bold, and resets.
func ColorizeBold(text, colorCode string) string {
	return colorCode + ColorBold + text + ColorReset
}

// Colorize applies a color and any number of style codes, then resets.
func Colorize(text, colorCode string, styles ...string) string {
	prefix := colorCode + strings.Join(styles, "")
	return prefix + text + ColorReset
}

// ApplyStyle applies one or more non-color styles and resets.
func ApplyStyle(text string, styles ...string) string {
	return strings.Join(styles, "") + text + ColorReset
}

// Convenience wrappers for common styles.
func Bold(text string) string      { return ApplyStyle(text, ColorBold) }
func Italic(text string) string    { return ApplyStyle(text, StyleItalic) }
func Underline(text string) string { return ApplyStyle(text, StyleUnderline) }
func Dim(text string) string       { return ApplyStyle(text, StyleDim) }
func Inverse(text string) string   { return ApplyStyle(text, StyleInverse) }
func Strike(text string) string    { return ApplyStyle(text, StyleStrike) }

// AlternatingColor returns two colors alternating by index (even/odd).
// Defaults mirror existing quick_* tools: white for even, cyan for odd.
func AlternatingColor(index int, evenColor, oddColor string) string {
	if index%2 == 0 {
		return evenColor
	}
	return oddColor
}
