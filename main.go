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

type RGBComponents struct {
	Red, Green, Blue uint8
}

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

func EnsureHexIsValid(hexColor string) bool {
	return true
}

func GetRGBFromHex(hexColor string) RGBComponents {
	return RGBComponents{}
}

func MatchColorToRGB(rgb RGBComponents) Color {
	return Color(0)
}
