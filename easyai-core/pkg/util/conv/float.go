package conv

import (
	"fmt"
	"strconv"
	"strings"
)

// FormatFloat2 format float
func FormatFloat2(f float32) string {
	return strconv.FormatFloat(float64(f), 'f', -1, 64)
}

// FormatFloat format float, with zero trim
func FormatFloat(v float32, p int) string {
	f := fmt.Sprintf("%%.%df", p)
	return strings.TrimRight(strings.TrimRight(fmt.Sprintf(f, v), "0"), ".")
}
