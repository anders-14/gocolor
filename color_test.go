package gocolor

import (
	"fmt"
	"testing"
)

func TestHexToRGB(t *testing.T) {
	fmt.Printf("%+v\n", hexToRGB("#33ee33"))
	fmt.Printf("%+v\n", hexToRGB("#ee3333"))
	fmt.Printf("%+v\n", hexToRGB("#3333ee"))
	fmt.Printf("%+v\n", hexToRGB("#FFFFFF"))
	fmt.Printf("%+v\n", hexToRGB("#000000"))
}

func TestRGBToANSI(t *testing.T) {
	a := (rgb{r: 255, g: 0, b: 0}).toANSIfg()
	Println(a, "hei")
	b := (rgb{r: 255, g: 0, b: 0}).toANSIbg()
	Println(b, "hei")
}
