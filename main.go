package main

import "fmt"

type Color int

const (
	Unknown Color = iota
	Red
	Orange
	Yellow
	Green
	Blue
	Purple
)

func (c Color) String() string {
	colors := [...]string{"red", "orange", "yellow", "green", "blue", "purple"}
	if c < Red || c > Purple {
		return fmt.Sprintf("Color(%d)", int(c))
	}
	return colors[c-Red]
}

// Returns a color from Color when passed a hex color code, defaults to Unknown
func GetColorFromHex(hexColor string) Color {
	return Color(0)
}
