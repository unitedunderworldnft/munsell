package munsell_test

import (
	"testing"

	"github.com/worldies/munsell"
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

func TestMatchFromHexBlack(t *testing.T) {
	color, err := munsell.GetColorFromHex("#0c1921")
	if err != nil {
		t.Errorf("Expected color, received %v", err)
	}
	if color != munsell.Black {
		t.Errorf("Expected Black, received %v", color)
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

func TestMatchFromHexLightBlue(t *testing.T) {
	color, err := munsell.GetColorFromHex("#4fa1c6")
	if err != nil {
		t.Errorf("Expected color, received %v", err)
	}
	if color != munsell.LightBlue {
		t.Errorf("Expected LightBlue, received %v", color)
	}
}

func TestMatchFromHexBlue(t *testing.T) {
	color, err := munsell.GetColorFromHex("#2a57c1")
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

func TestMatchFromHexShorthand(t *testing.T) {
	color, err := munsell.GetColorFromHex("#f00")
	if err != nil {
		t.Errorf("Expected color, received %v", err)
	}
	if color != munsell.Red {
		t.Errorf("Expected Red, received %v", color)
	}
}

/* Test Premier League teams colors */
func testPremierLeague(t *testing.T, inputColor string, expectedColor munsell.Color) {
	color, err := munsell.GetColorFromHex(inputColor)
	if err != nil {
		t.Errorf("Expected color, received %v", err)
	}
	if color != expectedColor {
		t.Errorf("Expected %s, received %v", expectedColor.String(), color)
	}
}
func TestPremierLeague_Arsenal(t *testing.T) {
	testPremierLeague(t, "#EF0107", munsell.Red)
}
func TestPremierLeague_ArsenalAlt(t *testing.T) {
	testPremierLeague(t, "#fbf2b0", munsell.Yellow)
}

func TestPremierLeague_Aston(t *testing.T) {
	testPremierLeague(t, "#7A003C", munsell.Purple)
}

func TestPremierLeague_AstonAlt(t *testing.T) {
	testPremierLeague(t, "#ffffff", munsell.White)
}

func TestPremierLeague_Brentford(t *testing.T) {
	testPremierLeague(t, "#f42727", munsell.Red)
}

func TestPremierLeague_BrentfordAlt(t *testing.T) {
	testPremierLeague(t, "#ffff00", munsell.Yellow)
}

func TestPremierLeague_Brighton(t *testing.T) {
	testPremierLeague(t, "#0000ff", munsell.Blue)
}

func TestPremierLeague_BrightonAlt(t *testing.T) {
	testPremierLeague(t, "#39c3cd", munsell.LightBlue)
}

func TestPremierLeague_Burnley(t *testing.T) {
	testPremierLeague(t, "#53162F", munsell.Purple)
}

func TestPremierLeague_BurnleyAlt(t *testing.T) {
	testPremierLeague(t, "#ffffff", munsell.White)
}

func TestPremierLeague_Chelsea(t *testing.T) {
	testPremierLeague(t, "#0054cd", munsell.Blue) // FIXME:
}

func TestPremierLeague_ChelseaAlt(t *testing.T) {
	testPremierLeague(t, "#ffff00", munsell.Yellow)
}

func TestPremierLeague_CrystalPalace(t *testing.T) {
	testPremierLeague(t, "#0202fb", munsell.Blue)
}

func TestPremierLeague_CrystalPalaceAlt(t *testing.T) {
	testPremierLeague(t, "#ffde00", munsell.Yellow)
}

func TestPremierLeague_Everton(t *testing.T) {
	testPremierLeague(t, "#0000ff", munsell.Blue)
}

func TestPremierLeague_EvertonAlt(t *testing.T) {
	testPremierLeague(t, "#231F20", munsell.Black)
}

func TestPremierLeague_Leeds(t *testing.T) {
	testPremierLeague(t, "#ffffff", munsell.White)
}

func TestPremierLeague_LeedsAlt(t *testing.T) {
	testPremierLeague(t, "#081d54", munsell.Blue)
}

func TestPremierLeague_Leicester(t *testing.T) {
	testPremierLeague(t, "#0202fb", munsell.Blue)
}

func TestPremierLeague_LeicesterAlt(t *testing.T) {
	testPremierLeague(t, "#bfe5e6", munsell.LightBlue)
}

func TestPremierLeague_Liverpool(t *testing.T) {
	testPremierLeague(t, "#D00027", munsell.Red)
}

func TestPremierLeague_LiverpoolAlt(t *testing.T) {
	testPremierLeague(t, "#ece2d9", munsell.White)
}

func TestPremierLeague_ManchesterCity(t *testing.T) {
	testPremierLeague(t, "#61acdf", munsell.LightBlue)
}

func TestPremierLeague_ManchesterCityAlt(t *testing.T) {
	testPremierLeague(t, "#ffffff", munsell.White)
}

func TestPremierLeague_ManchesterUnited(t *testing.T) {
	testPremierLeague(t, "#DA020E", munsell.Red)
}

func TestPremierLeague_ManchesterUnitedAlt(t *testing.T) {
	testPremierLeague(t, "#b7ddff", munsell.LightBlue)
}

func TestPremierLeague_Newcastle(t *testing.T) {
	testPremierLeague(t, "#231F20", munsell.Black)
}

func TestPremierLeague_NewcastleAlt(t *testing.T) {
	testPremierLeague(t, "#2e4b60", munsell.Blue)
}

func TestPremierLeague_Norwich(t *testing.T) {
	testPremierLeague(t, "#ffff00", munsell.Yellow)
}

func TestPremierLeague_NorwichAlt(t *testing.T) {
	testPremierLeague(t, "#081d54", munsell.Blue)
}

func TestPremierLeague_Southampton(t *testing.T) {
	testPremierLeague(t, "#ED1A3B", munsell.Red)
}

func TestPremierLeague_SouthamptonAlt(t *testing.T) {
	testPremierLeague(t, "#ffff00", munsell.Yellow)
}

func TestPremierLeague_Tottenham(t *testing.T) {
	testPremierLeague(t, "#ffffff", munsell.White)
}

func TestPremierLeague_TottenhamAlt(t *testing.T) {
	testPremierLeague(t, "#231F20", munsell.Black)
}

func TestPremierLeague_Watford(t *testing.T) {
	testPremierLeague(t, "#ffff00", munsell.Yellow)
}

func TestPremierLeague_WatfordAlt(t *testing.T) {
	testPremierLeague(t, "#ffffff", munsell.White)
}

func TestPremierLeague_WestHam(t *testing.T) {
	testPremierLeague(t, "#7A003C", munsell.Purple)
}

func TestPremierLeague_WestHamAlt(t *testing.T) {
	testPremierLeague(t, "#b7ddff", munsell.LightBlue)
}

func TestPremierLeague_Wolverhampton(t *testing.T) {
	testPremierLeague(t, "#F7AA25", munsell.Yellow)
}

func TestPremierLeague_WolverhamptonAlt(t *testing.T) {
	testPremierLeague(t, "#2e4b60", munsell.Blue)
}
