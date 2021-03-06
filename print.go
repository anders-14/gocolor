package gocolor

import "fmt"

const reset string = "\033[0m"

// Print just like `fmt.Print()`, but with color
func Print(color string, a ...interface{}) {
	fmt.Print(color)
	fmt.Print(a...)
	fmt.Print(reset)
}

// Printf just like `fmt.Printf()`, but with color
func Printf(color string, format string, a ...interface{}) {
	fmt.Print(color)
	fmt.Printf(format, a...)
	fmt.Print(reset)
}

// Println just like `fmt.Println()`, but with color
func Println(color string, a ...interface{}) {
	fmt.Print(color)
	fmt.Println(a...)
	fmt.Print(reset)
}
