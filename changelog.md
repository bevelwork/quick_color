## 1.0.20251007

Initial public release with core utilities for simple, readable CLI outputs.

- **ANSI color constants**: Basic colors and bold reset helpers.
- **Style helpers**: Dim, italic, underline, inverse, strike.
- **Color utilities**:
  - `Color(text, color)` and `Colorize(text, color, ...styles)`.
  - `ApplyStyle(text, ...styles)` and convenience wrappers (`Bold`, `Italic`, etc.).
- **Alternating color**: `AlternatingColor(index, evenColor, oddColor)` for list UIs.
- **Spinner utility**: Lightweight terminal spinner for progress indication.
- **Pinned line display**: Render and update a single pinned status line.
- **Subdisplay**: Sectioned output with consistent headers and indentation.
- **Alternate list renderer**: Simple alternating-color list printer.

Version string: see `version.go` (`Version = "1.0.20251007"`).


