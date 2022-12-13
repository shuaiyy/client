package exitcode

import "fmt"

// Meaning ...
func Meaning(code int32) string {
	switch code {
	case 0:
		return "exit with success"
	case 1:
		return "general exit code, may represent miscellaneous errors"
	case 2:
		return "misuse of shell built-in command/variables"
	case 127:
		return "shell command not found; or invalid exit code out of range[0,255]"
	case 128:
		return "invalid argument to exit, linux exit code must be an integer"
	case 255:
		return "exit code out of range[0,255]"
	}
	if code > 128 && code < 255 {
		return fmt.Sprintf("exit with kill signal: %d", code-128)
	}

	return "custom defined error code"
}
