package main

import (

	// 	testpkg "github.com/KelleyTyler/GridTileEbit04_12/myPkgs/testPkg"
	"log"

	"github.com/KelleyTyler/GridTileEbit04_12/myPkgs/framework"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	game := framework.GetNewGame()
	// ebiten.SetWindowSize(256, 256)
	ebiten.SetWindowSize(game.G_Setting.WindowSizeX, game.G_Setting.WindowSizeY)
	ebiten.SetWindowTitle("GTE_04_12")
	ebiten.SetVsyncEnabled(true)
	//ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
