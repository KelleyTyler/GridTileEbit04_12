package framework

import (
	"fmt"
	"image/color"

	coords "github.com/KelleyTyler/GridTileEbit04_12/myPkgs/basic_geometry/coords"
	settings "github.com/KelleyTyler/GridTileEbit04_12/myPkgs/settingsconfig"
	ui "github.com/KelleyTyler/GridTileEbit04_12/myPkgs/user_interface"

	// "github.com/ebitengine/debugui"
	// dbgui_wrap "github.com/KelleyTyler/GridTileEbit04_12/myPkgs/db_gui_wrap"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

/**/
type Game struct {
	debugmsg  string
	Board     GameBoard
	G_Setting settings.GameSettings
	Backend   ui.UI_Backend
	//--------------

	lbl00 ui.UI_Label
	// btn00b                    ui.UI_Button
	primitive ui.UI_Object_Primitive
	init      bool
}

/**/
func GetNewGame() *Game {
	game := Game{}
	game.G_Setting = settings.GetSettingsFromJSON()
	game.debugmsg = ""
	game.Backend = ui.GetUIBackend(&game.G_Setting, nil)
	// gBoardSize := coords.CoordInts{X: game.G_Setting.GameBoardX, Y: game.G_Setting.GameBoardY}
	// gBoardTileSize := coords.CoordInts{X: game.G_Setting.GameBoardTileX, Y: game.G_Setting.GameBoardTileY}
	// gBoardTileSpacing := coords.CoordInts{X: game.G_Setting.GameBoardTile_Margin_X, Y: game.G_Setting.GameBoardTile_Margin_Y} //158
	// num := game.Backend.Settings.ScreenResX - 208                                                                             //70 //-136
	// game.primitive.Init([]string{"Primitive 00"}, &game.Backend, nil, coords.CoordInts{X: num, Y: 4}, coords.CoordInts{X: 204, Y: 632})
	// game.lbl00.Init([]string{"lbl_02", "Primitve00"}, &game.Backend, nil, coords.CoordInts{X: 0, Y: 0}, coords.CoordInts{X: 204, Y: 32})

	// game.lbl00.TextAlignMode = 10
	// game.lbl00.Redraw()
	// game.Board.Init(&game.Backend, &game.primitive, coords.CoordInts{X: 158, Y: 42}, coords.CoordInts{X: 4, Y: 4}, gBoardSize, gBoardTileSize, gBoardTileSpacing)

	// game.primitive.Redraw()
	// game.lbl00.Init_Parents(&game.primitive)
	// // game.Board.SetParents(&game.primitive)
	// game.primitive.Redraw()

	return &game
}

/**/
func (g *Game) Init() error {
	// game := Game{}
	// g.G_Setting = settings.GetSettingsFromJSON()
	// g.debugmsg = ""
	// g.Backend = ui.GetUIBackend(&g.G_Setting, nil)
	gBoardSize := coords.CoordInts{X: g.G_Setting.GameBoardX, Y: g.G_Setting.GameBoardY}
	gBoardTileSize := coords.CoordInts{X: g.G_Setting.GameBoardTileX, Y: g.G_Setting.GameBoardTileY}
	gBoardTileSpacing := coords.CoordInts{X: g.G_Setting.GameBoardTile_Margin_X, Y: g.G_Setting.GameBoardTile_Margin_Y} //158
	num := g.Backend.Settings.ScreenResX - 208                                                                          //70 //-136
	g.primitive.Init([]string{"Primitive 00"}, &g.Backend, nil, coords.CoordInts{X: num, Y: 4}, coords.CoordInts{X: 204, Y: 632})
	g.lbl00.Init([]string{"lbl_02", "Primitve00"}, &g.Backend, nil, coords.CoordInts{X: 0, Y: 0}, coords.CoordInts{X: 204, Y: 32})

	g.lbl00.TextAlignMode = 10
	g.lbl00.Redraw()
	g.Board.Init(&g.Backend, &g.primitive, coords.CoordInts{X: 158, Y: 42}, coords.CoordInts{X: 4, Y: 4}, gBoardSize, gBoardTileSize, gBoardTileSpacing)

	g.primitive.Redraw()
	g.lbl00.Init_Parents(&g.primitive)
	// game.Board.SetParents(&game.primitive)
	g.primitive.Redraw()
	g.init = true
	// return &game
	// ebiten.SetWindowSize(g.G_Setting.WindowSizeX, g.G_Setting.WindowSizeY)
	return nil
}

/**/
func (g *Game) Update() error {
	if g.init {
		g.debugmsg = fmt.Sprintf("FPS: %6.2f TPS:%6.2f\n", ebiten.ActualFPS(), ebiten.ActualTPS())
		g.Board.Update()
		g.primitive.Update()
	} else {
		go g.Init()
	}
	return nil
}

/**/
func (g *Game) Draw(screen *ebiten.Image) {

	//screen.Fill(color.White)
	if g.init {
		g.Board.Draw(screen)
		// g.bPanel.Draw(screen)
		g.primitive.Draw(screen)

		ebitenutil.DebugPrint(screen, g.debugmsg)
	} else {
		screen.Fill(color.RGBA{25, 50, 188, 255})
	}

}

/**/
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.G_Setting.ScreenResX, g.G_Setting.ScreenResY
}
