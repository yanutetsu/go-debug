package debug

import (
	"fmt"
	"path"
	"runtime"
)

// Print prints a message with a file name and a line number.
func Print(msg string) {
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		printWithColor(msg)
	}

	printWithCallaer(file, line, msg)
}

// Printf prints a formatted message with a file name and a line number.
func Printf(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)

	_, file, line, ok := runtime.Caller(1)
	if !ok {
		printWithColor(msg)
	}

	printWithCallaer(file, line, msg)
}

func printWithColor(msg string) {
	fmt.Printf("\x1b[33m"+"%s"+"\n\x1b[0m", msg)
}

func printWithCallaer(file string, line int, msg string) {
	printWithColor(fmt.Sprintf("%s:%d: %s", path.Base(file), line, msg))
}
