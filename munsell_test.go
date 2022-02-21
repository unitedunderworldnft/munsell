package munsell_test

import (
	"testing"

	"github.com/unitedunderworldnft/munsell"
)

func TestMatchFromHexWhite(t *testing.T) {
	color, err := munsell.GetColorFromHex("#f0f4ef")
	if err != nil {
		t.Errorf("Expected color, received %v", err)
	}
	if color != munsell.White {
		t.Errorf("Expected White, received %v", color)
	}
}

func TestMatchFromHexRed(t *testing.T) {
	color, err := munsell.GetColorFromHex("#d33d32")
	if err != nil {
		t.Errorf("Expected color, received %v", err)
	}
	if color != munsell.Red {
		t.Errorf("Expected Red, received %v", color)
	}
}

func TestMatchFromHexOrange(t *testing.T) {
	color, err := munsell.GetColorFromHex("#f28c32")
	if err != nil {
		t.Errorf("Expected color, received %v", err)
	}
	if color != munsell.Orange {
		t.Errorf("Expected Orange, received %v", color)
	}
}

func TestMatchFromHexYellow(t *testing.T) {
	color, err := munsell.GetColorFromHex("#f2de00")
	if err != nil {
		t.Errorf("Expected color, received %v", err)
	}
	if color != munsell.Yellow {
		t.Errorf("Expected Yellow, received %v", color)
	}
}

func TestMatchFromHexGreen(t *testing.T) {
	color, err := munsell.GetColorFromHex("#6cce21")
	if err != nil {
		t.Errorf("Expected color, received %v", err)
	}
	if color != munsell.Green {
		t.Errorf("Expected Green, received %v", color)
	}
}

func TestMatchFromHexBlue(t *testing.T) {
	color, err := munsell.GetColorFromHex("#4fa1c6")
	if err != nil {
		t.Errorf("Expected color, received %v", err)
	}
	if color != munsell.Blue {
		t.Errorf("Expected Blue, received %v", color)
	}
}

func TestMatchFromHexPurple(t *testing.T) {
	color, err := munsell.GetColorFromHex("#9328ba")
	if err != nil {
		t.Errorf("Expected color, received %v", err)
	}
	if color != munsell.Purple {
		t.Errorf("Expected Purple, received %v", color)
	}
}

func TestMatchFromHexPink(t *testing.T) {
	color, err := munsell.GetColorFromHex("#d84eac")
	if err != nil {
		t.Errorf("Expected color, received %v", err)
	}
	if color != munsell.Pink {
		t.Errorf("Expected Pink, received %v", color)
	}
}
