package pretty

import (
	"github.com/fatih/color"
)

// Blue string
func Blue(s string) string {
	return color.New(color.BgBlue, color.FgWhite).Sprint(s)
}

// Red string
func Red(s string) string {
	return color.New(color.BgRed, color.FgWhite).Sprint(s)
}

// Green string
func Green(s string) string {
	return color.New(color.BgGreen, color.FgWhite).Sprint(s)
}
