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
func TestPremierLeague_AFCBournemouth(t *testing.T) {
	testPremierLeague(t, "#e62333", munsell.Red)
}
func TestPremierLeague_AFCBournemouthAlt(t *testing.T) {
	testPremierLeague(t, "#ffffff", munsell.White)
}

func TestPremierLeague_Arsenal(t *testing.T) {
	testPremierLeague(t, "#fe0002", munsell.Red)
}
func TestPremierLeague_ArsenalAlt(t *testing.T) {
	testPremierLeague(t, "#003671", munsell.Blue)
}

func TestPremierLeague_AstonVilla(t *testing.T) {
	testPremierLeague(t, "#480025", munsell.Purple)
}
func TestPremierLeague_AstonVillaAlt(t *testing.T) {
	testPremierLeague(t, "#94bee5", munsell.LightBlue)
}

func TestPremierLeague_Brentford(t *testing.T) {
	testPremierLeague(t, "#c61d23", munsell.Red)
}
func TestPremierLeague_BrentfordAlt(t *testing.T) {
	testPremierLeague(t, "#b4dafd", munsell.LightBlue)
}

func TestPremierLeague_BrightonHoveAlbion(t *testing.T) {
	testPremierLeague(t, "#0054a6", munsell.Blue)
}
func TestPremierLeague_BrightonHoveAlbionAlt(t *testing.T) {
	testPremierLeague(t, "#ffffff", munsell.White)
}

func TestPremierLeague_Burnley(t *testing.T) {
	testPremierLeague(t, "#53162F", munsell.Purple)
}

func TestPremierLeague_Chelsea(t *testing.T) {
	testPremierLeague(t, "#0a4595", munsell.Blue)
}
func TestPremierLeague_ChelseaAlt(t *testing.T) {
	testPremierLeague(t, "#ffffff", munsell.White)
}

func TestPremierLeague_CrystalPalace(t *testing.T) {
	testPremierLeague(t, "#eb302e", munsell.Red)
}
func TestPremierLeague_CrystalPalaceAlt(t *testing.T) {
	testPremierLeague(t, "#0055a5", munsell.Blue)
}

func TestPremierLeague_Everton(t *testing.T) {
	testPremierLeague(t, "#00369c", munsell.Blue)
}
func TestPremierLeague_EvertonAlt(t *testing.T) {
	testPremierLeague(t, "#231F20", munsell.Black)
}

func TestPremierLeague_Fulham(t *testing.T) {
	testPremierLeague(t, "#f5f5f5", munsell.White)
}
func TestPremierLeague_FulhamAlt(t *testing.T) {
	testPremierLeague(t, "#1a1a1a", munsell.Black)
}

func TestPremierLeague_LeedsUnited(t *testing.T) {
	testPremierLeague(t, "#f5f5f5", munsell.White)
}
func TestPremierLeague_LeedsUnitedAlt(t *testing.T) {
	testPremierLeague(t, "#fddf12", munsell.Yellow)
}

func TestPremierLeague_LeicesterCity(t *testing.T) {
	testPremierLeague(t, "#273e8a", munsell.Blue)
}
func TestPremierLeague_LeicesterCityAlt(t *testing.T) {
	testPremierLeague(t, "#f7a800", munsell.Yellow)
}

func TestPremierLeague_Liverpool(t *testing.T) {
	testPremierLeague(t, "#e31b23", munsell.Red)
}
func TestPremierLeague_LiverpoolAlt(t *testing.T) {
	testPremierLeague(t, "#ffffff", munsell.White)
}

func TestPremierLeague_ManchesterCity(t *testing.T) {
	testPremierLeague(t, "#6caee0", munsell.LightBlue)
}
func TestPremierLeague_ManchesterCityAlt(t *testing.T) {
	testPremierLeague(t, "#333333", munsell.Black)
}

func TestPremierLeague_ManchesterUnited(t *testing.T) {
	testPremierLeague(t, "#d81920", munsell.Red)
}
func TestPremierLeague_ManchesterUnitedAlt(t *testing.T) {
	testPremierLeague(t, "#ffffff", munsell.White)
}

func TestPremierLeague_NewcastleUnited(t *testing.T) {
	testPremierLeague(t, "#383838", munsell.Black)
}
func TestPremierLeague_NewcastleUnitedAlt(t *testing.T) {
	testPremierLeague(t, "#00b6f1", munsell.LightBlue)
}

func TestPremierLeague_NorwichCity(t *testing.T) {
	testPremierLeague(t, "#ffff00", munsell.Yellow)
}

func TestPremierLeague_NottinghamForest(t *testing.T) {
	testPremierLeague(t, "#eb0024", munsell.Red)
}
func TestPremierLeague_NottinghamForestAlt(t *testing.T) {
	testPremierLeague(t, "#ffffff", munsell.White)
}

func TestPremierLeague_Southampton(t *testing.T) {
	testPremierLeague(t, "#d71920", munsell.Red)
}
func TestPremierLeague_SouthamptonAlt(t *testing.T) {
	testPremierLeague(t, "#ffffff", munsell.White)
}

func TestPremierLeague_TottenhamHotspur(t *testing.T) {
	testPremierLeague(t, "#f5f5f5", munsell.White)
}
func TestPremierLeague_TottenhamHotspurAlt(t *testing.T) {
	testPremierLeague(t, "#131f53", munsell.Blue)
}

func TestPremierLeague_Watford(t *testing.T) {
	testPremierLeague(t, "#ffff00", munsell.Yellow)
}

func TestPremierLeague_WestHamUnited(t *testing.T) {
	testPremierLeague(t, "#7d2c3b", munsell.Red)
}
func TestPremierLeague_WestHamUnitedAlt(t *testing.T) {
	testPremierLeague(t, "#f8d742", munsell.Yellow)
}

func TestPremierLeague_WolverhamptonWanderers(t *testing.T) {
	testPremierLeague(t, "#f9a01b", munsell.Orange)
}
func TestPremierLeague_WolverhamptonWanderersAlt(t *testing.T) {
	testPremierLeague(t, "#333333", munsell.Black)
}
