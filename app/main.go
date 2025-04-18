package main

import (

	// 	testpkg "github.com/KelleyTyler/GridTileEbit04_12/myPkgs/testPkg"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct{}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "GTE_04_12")
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("GTE_04_12")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
