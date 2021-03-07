# Gocolor

Simple color printing in golang.

## Usage

Gocolor uses a color block syntax like this:

```go
redFg := "${f#ff0000}This whole string has red foreground color${reset}"
blueBg := "${b#0000ff}This${reset}, only this, has blue background color"
```

The color blocks starts with either `f` or `b` for foreground and
background. Then there must be a valid 6 digit hex color starting with a
`#`. The `${reset}` block is used to reset the text to normal.

The color string can be printed via these functions:

```go
gocolor.Print()
gocolor.Printf()
gocolor.Println()
```

