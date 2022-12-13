package stdlogger

import (
	"fmt"
	"log"

	"github.com/fatih/color"
)

const (
	// DebugLevel debug
	DebugLevel = iota
	// InfoLevel info
	InfoLevel
	// WarnLevel warn
	WarnLevel
	// ErrorLevel error
	ErrorLevel
)

var (
	// LogLevel global level
	LogLevel = DebugLevel
)

// Debug logs a message with severity DEBUG.
func Debug(format string, v ...interface{}) {
	if LogLevel > DebugLevel {
		return
	}
	c := color.New(color.BgBlue, color.FgWhite)
	output(c.Sprintf(fmt.Sprintf("[DEBUG] %s", format), v...))
}

// Info logs a message with severity INFO.
func Info(format string, v ...interface{}) {
	if LogLevel > InfoLevel {
		return
	}
	c := color.New(color.FgHiBlue)
	output(c.Sprintf(fmt.Sprintf("[INFO] %s", format), v...))
}

// Warn logs a message with severity WARN.
func Warn(format string, v ...interface{}) {
	if LogLevel > WarnLevel {
		return
	}
	c := color.New(color.FgHiYellow)
	output(c.Sprintf(fmt.Sprintf("[WARN] %s", format), v...))
}

// Error logs a message with severity ERROR.
func Error(format string, v ...interface{}) {
	if LogLevel > ErrorLevel {
		return
	}
	c := color.New(color.FgHiRed)
	output(c.Sprintf(fmt.Sprintf("[ERROR] %s", format), v...))
}

// Fatal logs a message with severity ERROR followed by a call to os.Exit().
func Fatal(format string, v ...interface{}) {
	c := color.New(color.FgHiRed)
	panic(c.Sprintf(fmt.Sprintf("[ERROR] %s", format), v...))
}

func output(msg string) {
	log.Println(msg)
}
