package fastmux

import (
	"fmt"
)

// debugPrint prints a formatted debug message prefixed with [Fastmux]
func debugPrint(format string, args ...any) {
	const prefix = "\033[36m[Fastmux]\033[0m"
	message := fmt.Sprintf(format, args...)
	fmt.Printf("%s %s\n", prefix, message)
}
