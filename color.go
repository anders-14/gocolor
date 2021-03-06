package gocolor

import (
	"fmt"
	"strconv"
	"strings"
)

type rgb struct {
	r int64
	g int64
	b int64
}

func hexToRGB(hex string) rgb {
	if strings.HasPrefix(hex, "#") && len(hex) == 7 {
		hex = hex[1:]
	}
	r, _ := strconv.ParseInt(hex[:2], 16, 64)
	g, _ := strconv.ParseInt(hex[2:4], 16, 64)
	b, _ := strconv.ParseInt(hex[4:6], 16, 64)
	return rgb{r, g, b}
}

func (color rgb) toANSIfg() string {
	return fmt.Sprintf("\033[38;2;%+v;%+v;%+vm", color.r, color.g, color.b)
}

func (color rgb) toANSIbg() string {
	return fmt.Sprintf("\033[48;2;%+v;%+v;%+vm", color.r, color.g, color.b)
}
