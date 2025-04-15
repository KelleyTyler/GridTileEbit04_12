package framework

import (
	"fmt"

	coords "github.com/KelleyTyler/GridTileEbit04_12/myPkgs/basic_geometry/coords"
	settings "github.com/KelleyTyler/GridTileEbit04_12/myPkgs/settingsconfig"
	ui "github.com/KelleyTyler/GridTileEbit04_12/myPkgs/userinterface"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Game struct {
	debugmsg  string
	Board     GameBoard
	G_Setting settings.GameSettings
	Backend   ui.UI_Backend
}

func GetNewGame() *Game {
	game := Game{}
	game.G_Setting = settings.GetSettingsFromJSON()
	game.debugmsg = ""
	game.Backend = ui.GetUIBackend(&game.G_Setting, nil)
	gBoardSize := coords.CoordInts{X: game.G_Setting.GameBoardX, Y: game.G_Setting.GameBoardY}
	gBoardTileSize := coords.CoordInts{X: game.G_Setting.GameBoardTileX, Y: game.G_Setting.GameBoardTileY}
	gBoardTileSpacing := coords.CoordInts{X: 2, Y: 2}
	game.Board.Init(&game.Backend, coords.CoordInts{X: 168, Y: 8}, coords.CoordInts{X: 4, Y: 4}, gBoardSize, gBoardTileSize, gBoardTileSpacing)
	return &game
}

func (g *Game) Update() error {
	g.debugmsg = fmt.Sprintf("FPS: %6.2f TPS:%6.2f\n", ebiten.ActualFPS(), ebiten.ActualTPS())
	g.Board.Update()

	if inpututil.IsKeyJustPressed(ebiten.KeyM) {
		g.Board.MazeGenPassthrough()
	}
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

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, g.debugmsg)
	g.Board.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.G_Setting.ScreenResX, g.G_Setting.ScreenResY
}
