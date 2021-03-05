package gocolor_test

import (
	"testing"

	"github.com/anders-14/gocolor"
)

func TestPrint(t *testing.T) {
	gocolor.Print(gocolor.GreenFg, "Hi this is a test", "\n")
	gocolor.Print(gocolor.BlueBg, "Hi this is a test", "\n")
}

func TestPrintf(t *testing.T) {
	name := "Anders"
	gocolor.Printf(gocolor.GreenFg, "%s is my name\n", name)
	gocolor.Printf(gocolor.BlueBg, "%s is my name\n", name)
}

func TestPrintln(t *testing.T) {
	gocolor.Println(gocolor.GreenFg, "Hi this is a test")
	gocolor.Println(gocolor.Underline, "Hi this is a test")
}
