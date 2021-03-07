package gocolor

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
)

var hexError = errors.New("invalid hex color")
var hexPattern = regexp.MustCompile(`^#[a-fA-F0-9]{6}$`)

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

var reset string = "\033[0m"

// Print just like `fmt.Print()`, but with color
func Print(hex string, fg bool, a ...interface{}) error {
  esc, err := hexToANSI(hex, fg) 
  if err != nil {
    return err
  }
	fmt.Print(esc)
	fmt.Print(a...)
	fmt.Print(reset)
  return nil
}

// Printf just like `fmt.Printf()`, but with color
func Printf(hex string, fg bool, format string, a ...interface{}) error {
  esc, err := hexToANSI(hex, fg) 
  if err != nil {
    return err
  }
	fmt.Print(esc)
	fmt.Printf(format, a...)
	fmt.Print(reset)
  return nil
}

// Println just like `fmt.Println()`, but with color
func Println(hex string, fg bool, a ...interface{}) error {
  esc, err := hexToANSI(hex, fg) 
  if err != nil {
    return err
  }
	fmt.Print(esc)
	fmt.Println(a...)
	fmt.Print(reset)
  return nil
}
