package gocolor

import "fmt"

// ANSIEscape is an ANSI color escape code
type ANSIEscape string

var (
	Reset        ANSIEscape = "\033[0m"
	Bold         ANSIEscape = "\033[1m"
	Underline    ANSIEscape = "\033[4m"
	Blinking     ANSIEscape = "\033[5m"
	BlackFg      ANSIEscape = "\033[30m"
	RedFg        ANSIEscape = "\033[31m"
	GreenFg      ANSIEscape = "\033[32m"
	BrownFg      ANSIEscape = "\033[33m"
	BlueFg       ANSIEscape = "\033[34m"
	PurpleFg     ANSIEscape = "\033[35m"
	CyanFg       ANSIEscape = "\033[36m"
	GrayFg       ANSIEscape = "\033[37m"
	BlackBg      ANSIEscape = "\033[40m"
	RedBg        ANSIEscape = "\033[41m"
	GreenBg      ANSIEscape = "\033[42m"
	BrownBg      ANSIEscape = "\033[43m"
	BlueBg       ANSIEscape = "\033[44m"
	PurpleBg     ANSIEscape = "\033[45m"
	CyanBg       ANSIEscape = "\033[46m"
	GrayBg       ANSIEscape = "\033[47m"
)

// Print just like `fmt.Print()`, but with color
func Print(color ANSIEscape, a ...interface{}) {
	fmt.Print(string(color))
	fmt.Print(a...)
	fmt.Print(string(Reset))
}

// Printf just like `fmt.Printf()`, but with color
func Printf(color ANSIEscape, format string, a ...interface{}) {
	fmt.Print(string(color))
	fmt.Printf(format, a...)
	fmt.Print(string(Reset))
}

// Println just like `fmt.Println()`, but with color
func Println(color ANSIEscape, a ...interface{}) {
	fmt.Print(string(color))
	fmt.Println(a...)
	fmt.Print(string(Reset))
}
