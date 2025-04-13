package main

import (

	// 	testpkg "github.com/KelleyTyler/GridTileEbit04_12/myPkgs/testPkg"
	"log"

	"github.com/KelleyTyler/GridTileEbit04_12/myPkgs/framework"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	game := &framework.Game{}
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("GTE_04_12")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
