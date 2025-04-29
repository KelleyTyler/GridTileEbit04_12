package framework

import (
	"fmt"

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
}

/**/
func GetNewGame() *Game {
	game := Game{}
	game.G_Setting = settings.GetSettingsFromJSON()
	game.debugmsg = ""
	game.Backend = ui.GetUIBackend(&game.G_Setting, nil)
	gBoardSize := coords.CoordInts{X: game.G_Setting.GameBoardX, Y: game.G_Setting.GameBoardY}
	gBoardTileSize := coords.CoordInts{X: game.G_Setting.GameBoardTileX, Y: game.G_Setting.GameBoardTileY}
	gBoardTileSpacing := coords.CoordInts{X: game.G_Setting.GameBoardTile_Margin_X, Y: game.G_Setting.GameBoardTile_Margin_Y} //158
	num := game.Backend.Settings.ScreenResX - 208                                                                             //70 //-136
	game.primitive.Init([]string{"Primitive 00"}, &game.Backend, nil, coords.CoordInts{X: num, Y: 4}, coords.CoordInts{X: 204, Y: 632})
	game.lbl00.Init([]string{"lbl_02", "Primitve00"}, &game.Backend, nil, coords.CoordInts{X: 0, Y: 0}, coords.CoordInts{X: 204, Y: 32})

	game.lbl00.TextAlignMode = 10
	game.lbl00.Redraw()
	game.Board.Init(&game.Backend, &game.primitive, coords.CoordInts{X: 158, Y: 42}, coords.CoordInts{X: 4, Y: 4}, gBoardSize, gBoardTileSize, gBoardTileSpacing)

	game.primitive.Redraw()
	game.lbl00.Init_Parents(&game.primitive)
	// game.Board.SetParents(&game.primitive)
	game.primitive.Redraw()

	return &game
}

/**/
func (g *Game) Update() error {
	g.debugmsg = fmt.Sprintf("FPS: %6.2f TPS:%6.2f\n", ebiten.ActualFPS(), ebiten.ActualTPS())
	g.Board.Update()
	g.primitive.Update()
	// if g.bPanel.ButtonOuts[0] {
	// 	fmt.Printf("HAH 00\n")
	// }
	// if g.bPanel.ButtonOuts[4] {
	// 	fmt.Printf("HAH 04\n")
	// g.bnumScroll.Update()
	// }
	// if inpututil.IsKeyJustPressed(ebiten.KeyM) {
	// 	g.Board.MazeGenPassthrough()
	// }
	// if inpututil.IsKeyJustPressed(ebiten.Key1) {
	// 	g.Backend.PlaySound(0)
	// }
	// if inpututil.IsKeyJustPressed(ebiten.Key2) {
	// 	g.Backend.PlaySound(1)
	// }
	// if inpututil.IsKeyJustPressed(ebiten.Key3) {
	// 	g.Backend.PlaySound(2)
	// }
	// if inpututil.IsKeyJustPressed(ebiten.Key4) {
	// 	g.Backend.PlaySound(3)
	// }

	// if inpututil.IsKeyJustPressed(ebiten.Key5) {
	// 	g.Backend.PlaySound(4)
	// }

	// if inpututil.IsKeyJustPressed(ebiten.Key6) {
	// 	g.Backend.PlaySound(5)
	// }

	// if inpututil.IsKeyJustPressed(ebiten.Key7) {
	// 	g.Backend.PlaySound(6)

	// }
	return nil
}

/**/
func (g *Game) Draw(screen *ebiten.Image) {

	//screen.Fill(color.White)
	g.Board.Draw(screen)
	// g.bPanel.Draw(screen)
	g.primitive.Draw(screen)

	// g.bnumScroll.Draw(screen)

	// g.lbl02.Draw(screen)
	// g.btn00b.Draw(screen)
	// g.DrawUI(screen)
	ebitenutil.DebugPrint(screen, g.debugmsg)
}

/**/
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.G_Setting.ScreenResX, g.G_Setting.ScreenResY
}
