package gocolor

import (
	"testing"
)

func TestPrint(t *testing.T) {
	str := "${f#0000ff}Hello${reset} is blue, ${b#ff0000}I${reset} am red.\n"
	Print(str)
}

func TestPrintf(t *testing.T) {
	name := "Bob"
  color := "#99f910"
	str := "${b%s}%s${reset} is a name.\n"
	Printf(str, color, name)
}

func TestPrintln(t *testing.T) {
	str := "${f#923874}Hello this is all colored${reset}"
	Println(str)
}
