package munsell_test

import (
	"testing"

	"github.com/unitedunderworldnft/munsell"
)

func TestParseFromHex(t *testing.T) {
	color, err := munsell.GetColorFromHex("#ff0000")
	if err != nil {
		t.Errorf("Expected color, received %v", err)
	}
	if color != munsell.Red {
		t.Errorf("Expected Red, received %v", color)
	}
}
