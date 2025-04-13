package framework

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
	debugmsg string
}

func (g *Game) Update() error {
	g.debugmsg = fmt.Sprintf("FPS: %6.2f TPS:%6.2f\n", ebiten.ActualFPS(), ebiten.ActualTPS())

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, g.debugmsg)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 480
}
