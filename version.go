package quick_color

import "time"

// Version components
const (
	// MAJOR, MINOR follow major.minor.YYYYMMDD
	MAJOR = "1"
	MINOR = "1"
)

// DATE and Version are computed at init time so DATE uses the current day.
var (
	DATE    = time.Now().Format("20060102")
	Version = MAJOR + "." + MINOR + "." + DATE
)
