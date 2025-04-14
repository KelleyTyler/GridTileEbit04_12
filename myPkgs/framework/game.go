package framework

import (
	"fmt"

	coords "github.com/KelleyTyler/GridTileEbit04_12/myPkgs/basic_geometry/coords"
	settings "github.com/KelleyTyler/GridTileEbit04_12/myPkgs/settingsconfig"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Game struct {
	debugmsg  string
	Board     GameBoard
	G_Setting settings.GameSettings
}

func GetNewGame() *Game {
	game := Game{}
	game.G_Setting = settings.GetSettingsFromJSON()
	game.debugmsg = ""
	gBoardSize := coords.CoordInts{X: game.G_Setting.GameBoardX, Y: game.G_Setting.GameBoardY}
	gBoardTileSize := coords.CoordInts{X: game.G_Setting.GameBoardTileX, Y: game.G_Setting.GameBoardTileY}
	gBoardTileSpacing := coords.CoordInts{X: 2, Y: 2}
	game.Board.Init(coords.CoordInts{X: 168, Y: 8}, coords.CoordInts{X: 0, Y: 0}, coords.CoordInts{X: 0, Y: 0}, gBoardSize, gBoardTileSize, gBoardTileSpacing)
	return &game
}

func (g *Game) Update() error {
	g.debugmsg = fmt.Sprintf("FPS: %6.2f TPS:%6.2f\n", ebiten.ActualFPS(), ebiten.ActualTPS())
	g.Board.Update()

	if inpututil.IsKeyJustPressed(ebiten.KeyM) {
		g.Board.MazeGenPassthrough()
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, g.debugmsg)
	g.Board.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.G_Setting.ScreenResX, g.G_Setting.ScreenResY
}
