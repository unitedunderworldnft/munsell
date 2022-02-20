package main

import (
	"testing"
)

func TestParseFromHex(t *testing.T) {
	color, err := GetColorFromHex("#ff0000")
	if err != nil {
		t.Errorf("Expected color, received %v", err)
	}
	if color != Red {
		t.Errorf("Expected Red, received %v", color)
	}
}
