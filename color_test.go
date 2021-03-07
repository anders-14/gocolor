package gocolor

import (
	"testing"
)

func Test(t *testing.T) {
	Println("#3355ff", true, "hei")
	Println("#3355ff", false, "hei")
	Print("#3355ff", true, "hei")
	Print("#3355ff", true, "hei")
	Printf("#3355ff", false, "hei")
	Printf("#3355ff", false, "hei")
}
