package debug

import (
	"fmt"
	"path"
	"runtime"
)

// PrintColor is type string of color list you can use
type PrintColor string

// color list
const (
	PrintColorBlack   PrintColor = "\x1b[30m"
	PrintColorRed     PrintColor = "\x1b[31m"
	PrintColorGreen   PrintColor = "\x1b[32m"
	PrintColorYellow  PrintColor = "\x1b[33m"
	PrintColorBlue    PrintColor = "\x1b[34m"
	PrintColorMagenta PrintColor = "\x1b[35m"
	PrintColorCyan    PrintColor = "\x1b[36m"
	PrintColorWhite   PrintColor = "\x1b[37m"
	PrintColorReset   PrintColor = "\x1b[0m"
)

var choicedColor = PrintColorYellow

// Print prints a message with a file name and a line number.
// It works like fmt.Print.
// You can use it only by changing from `fmt` to `debug`.
//
// fmt.Print("hoge")   // => hoge
// debug.Print("hoge") // => debug_print.go:12 hoge
func Print(a ...interface{}) {
	msg := fmt.Sprint(a...)

	_, file, line, ok := runtime.Caller(1)
	if !ok {
		printf(msg)
	}

	printWithFileLine(file, line, msg)
}

// Println prints a message with a file name and a line number.
// It works like fmt.Print.
// You can use it only by changing from `fmt` to `debug`.
//
// fmt.Println("hoge")   // => hoge
// debug.Println("hoge") // => debug_print.go:12 hoge
func Println(a ...interface{}) {
	msg := fmt.Sprintln(a...)

	_, file, line, ok := runtime.Caller(1)
	if !ok {
		printf(msg)
	}

	printWithFileLine(file, line, msg)
}

// Printf prints a formatted message with a file name and a line number.
// It works like fmt.Printf.
// You can use it only by changing from `fmt` to `debug`.
//
// hoge := "piyo"
// fmt.Print("%s", hoge)   // => piyo
// debug.Print("%s", hoge) // => debug_print.go:12 piyo
func Printf(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)

	_, file, line, ok := runtime.Caller(1)
	if !ok {
		printf(msg)
	}

	printWithFileLine(file, line, msg)
}

// SetPrintColor set color on console message
func SetPrintColor(colorName PrintColor) {
	choicedColor = colorName
}

func printf(msg string) {
	fmt.Printf("%s%s%s", choicedColor, msg, PrintColorReset)
}

func printWithFileLine(file string, line int, msg string) {
	printf(fmt.Sprintf("%s:%d: %s", path.Base(file), line, msg))
}
