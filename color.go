package gocolor

import (
	"errors"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

var hexError = errors.New("invalid hex color")

var hexPattern = regexp.MustCompile(`^#[a-fA-F0-9]{6}$`)
var fmtPattern = regexp.MustCompile(`\$\{([fb]#[a-fA-F0-9]{6}|reset)\}`)

var reset string = "\033[0m"

func hexToANSI(hex string, fg bool) (string, error) {
	if !hexPattern.MatchString(hex) {
		return "", hexError
	}
	var where int
	if fg {
		where = 38
	} else {
		where = 48
	}
	r, _ := strconv.ParseInt(hex[1:3], 16, 0)
	g, _ := strconv.ParseInt(hex[3:5], 16, 0)
	b, _ := strconv.ParseInt(hex[5:], 16, 0)

	return fmt.Sprintf("\033[%d;2;%d;%d;%dm", where, r, g, b), nil
}

func findColorBlocks(str string) [][]string {
	return fmtPattern.FindAllStringSubmatch(str, -1)
}

func replaceColorBlocks(str string, formatStrings [][]string) string {
	for _, v := range formatStrings {
		switch v[1] {
		case "reset":
			str = strings.Replace(str, v[0], reset, 1)
		default:
			fg := strings.HasPrefix(v[1], "f")
			esc, err := hexToANSI(v[1][1:], fg)
			if err != nil {
				log.Fatalf("err: %+v", err)
			}
			str = strings.Replace(str, v[0], esc, 1)
		}
	}
	return str
}

// Print replaces the colorblocks with ANSI color escape codes
// and prints the colored string.
func Print(str string) {
	fmtStrs := findColorBlocks(str)
	cStr := replaceColorBlocks(str, fmtStrs)
	fmt.Print(cStr)
}

// Printf does the same as Print, but with string formatting.
func Printf(str string, a ...interface{}) {
	fmtStrs := findColorBlocks(str)
	cStr := replaceColorBlocks(str, fmtStrs)
	fmt.Printf(cStr, a...)
}

// Println does the same as Print, but with a newline at the end.
func Println(str string) {
	fmtStrs := findColorBlocks(str)
	cStr := replaceColorBlocks(str, fmtStrs)
	fmt.Println(cStr)
}
