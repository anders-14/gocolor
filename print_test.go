package gocolor

import (
	"testing"
)

var f = (rgb{255, 0, 0}).toANSIfg()
var b = (rgb{255, 0, 0}).toANSIbg()

func TestPrint(t *testing.T) {
	Print(f, "Hi this is a test", "\n")
	Print(b, "Hi this is a test", "\n")
}

func TestPrintf(t *testing.T) {
	name := "Anders"
	Printf(f, "%s is my name\n", name)
	Printf(b, "%s is my name\n", name)
}

func TestPrintln(t *testing.T) {
	Println(f, "Hi this is a test")
	Println(b, "Hi this is a test")
}
