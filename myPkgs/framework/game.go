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

type Game struct {
	debugmsg  string
	Board     GameBoard
	G_Setting settings.GameSettings
	Backend   ui.UI_Backend
	//--------------
	// btn00, btn01, btn02       ui.UI_Button
	// btn03, btn04, btn05       ui.UI_Button
	// btn06, btn07, btn08       ui.UI_Button
	// btn09, btn10, btn11       ui.UI_Button
	// btn12, btn13, btn14       ui.UI_Button
	// s_btn00, s_btn01, s_btn02 ui.UI_Button
	// s_btn03, s_btn04, s_btn05 ui.UI_Button
	lbl00, lbl01, lbl02 ui.UI_Label
	// btn00b                    ui.UI_Button
	obj0, obj1 ui.UI_Object
	primitive  ui.UI_Object_Primitive
	bPanel     ui.UI_ButtonPanel
	bnumScroll ui.UI_Num_Select
	// dbgui      dbgui_wrap.Basic_Debug_GUI_Window
}

func GetNewGame() *Game {
	game := Game{}
	game.G_Setting = settings.GetSettingsFromJSON()
	game.debugmsg = ""
	game.Backend = ui.GetUIBackend(&game.G_Setting, nil)
	gBoardSize := coords.CoordInts{X: game.G_Setting.GameBoardX, Y: game.G_Setting.GameBoardY}
	gBoardTileSize := coords.CoordInts{X: game.G_Setting.GameBoardTileX, Y: game.G_Setting.GameBoardTileY}
	gBoardTileSpacing := coords.CoordInts{X: game.G_Setting.GameBoardTile_Margin_X, Y: game.G_Setting.GameBoardTile_Margin_Y} //158
	game.Board.Init(&game.Backend, coords.CoordInts{X: 158, Y: 42}, coords.CoordInts{X: 4, Y: 4}, gBoardSize, gBoardTileSize, gBoardTileSpacing)
	num := game.Backend.Settings.ScreenResX - 212 //70 //-136
	// game.lbl00.Init(&game.Backend, "This Is Text", coords.CoordInts{X: num - 136, Y: 2}, coords.CoordInts{X: 200, Y: 32}, nil)
	// game.btn00b.Init(&game.Backend, "Save Map", coords.CoordInts{X: 32, Y: 32}, coords.CoordInts{X: 64, Y: 32}, 10, nil)
	game.bPanel.Init("Panel00", &game.Backend, coords.CoordInts{X: num, Y: 18}, coords.CoordInts{X: 208, Y: 600}, nil)
	// game.bPanel.AddButton("b1", coords.CoordInts{X: 4, Y: 34}, coords.CoordInts{X: 32, Y: 32}, 1)
	game.bPanel.AddButtons_Row_TypeA([]string{"b00", "b01", "b02", "b03", "b04", "b05"}, coords.CoordInts{X: 2, Y: 68}, coords.CoordInts{X: 32, Y: 32}, []uint8{0, 0, 0, 0, 10, 10})
	// game.InitGUI()
	game.primitive.Init([]string{"Primitive 00"}, &game.Backend, nil, coords.CoordInts{X: num, Y: 36}, coords.CoordInts{X: 208, Y: 256})
	// g.lbl02.Init(&game.Backend, "This Is Text", coords.CoordInts{X: num - 136, Y: Row7}, coords.CoordInts{X: 200, Y: 200}, nil)
	game.obj0 = &game.primitive
	game.lbl02.Init([]string{"lbl_02", "Primitve00"}, &game.Backend, nil, coords.CoordInts{X: 0, Y: 0}, coords.CoordInts{X: 208, Y: 32})
	game.primitive.Redraw()
	// game.primitive.Children = append(game.primitive.Children, &game.lbl02)
	game.lbl02.Init_Parents(game.obj0)
	game.obj1 = &game.lbl02
	// game.btn00.Init([]string{"ui_btn_00", "A"}, &game.Backend, nil, coords.CoordInts{X: 4, Y: 36}, coords.CoordInts{X: 32, Y: 32})
	// game.btn01.Init([]string{"ui_btn_01", "B"}, &game.Backend, nil, coords.CoordInts{X: 40, Y: 36}, coords.CoordInts{X: 32, Y: 32})
	// game.btn02.Init([]string{"ui_btn_02", "C"}, &game.Backend, nil, coords.CoordInts{X: 76, Y: 36}, coords.CoordInts{X: 32, Y: 32})
	// game.bnumScroll.Init([]string{"numScroll00", "SC00"}, &game.Backend, nil, coords.CoordInts{X: 4, Y: 72}, coords.CoordInts{X: 68, Y: 36})
	// game.btn00.Init_Parents(&game.primitive)
	// game.btn01.Init_Parents(&game.primitive)
	// game.btn02.Init_Parents(&game.primitive)
	game.Board.Load_Map_Button.Init_Parents(&game.primitive)
	game.Board.Save_Map_Button.Init_Parents(&game.primitive)

	game.Board.Redraw_Tiles_Button.Init_Parents(&game.primitive)
	game.Board.New_Map_Button.Init_Parents(&game.primitive)
	game.Board.Reset_Map_Btn.Init_Parents(&game.primitive)
	game.Board.NumSelect_MapSize_X.Init_Parents(&game.primitive)
	game.Board.NumSelect_MapSize_Y.Init_Parents(&game.primitive)
	game.Board.NumSelect_TileSize_X.Init_Parents(&game.primitive)
	game.Board.NumSelect_TileSize_Y.Init_Parents(&game.primitive)
	game.Board.NumSelect_Tile_Margin_X.Init_Parents(&game.primitive)
	game.Board.NumSelect_Tile_Margin_Y.Init_Parents(&game.primitive)
	game.Board.Maze_Gen_Button.Init_Parents(&game.primitive)
	game.Board.Maze_Gen_Button.Btn_Type = 10
	// game.btn01.Btn_Type = 10
	game.primitive.Redraw()

	return &game
}

func (g *Game) Update() error {
	g.debugmsg = fmt.Sprintf("FPS: %6.2f TPS:%6.2f\n", ebiten.ActualFPS(), ebiten.ActualTPS())
	g.Board.Update()
	// g.bPanel.Update()
	// if g.btn00b.Update_Ret() {
	// 	fmt.Printf("ACTIVATED ALMONDS \n")
	// }
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

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.G_Setting.ScreenResX, g.G_Setting.ScreenResY
}

// func (g *Game) UpdateGUI() {
// 	if g.btn00.Update_Ret() {
// 		fmt.Printf("Click\n")
// 	}
// 	if g.btn01.Update_Ret() {
// 		//fmt.Printf("Click2\n")
// 	}
// 	if g.btn02.Update_Ret() {
// 		//fmt.Printf("Click2\n")
// 	}
// 	if g.btn03.Update_Ret() {
// 		//fmt.Printf("Click2\n")
// 	}
// 	if g.btn04.Update_Ret() {
// 		//fmt.Printf("Click2\n")
// 	}
// 	if g.btn05.Update_Ret() {
// 		//fmt.Printf("Click2\n")
// 	}
// 	if g.btn06.Update_Ret() {
// 		fmt.Printf("Click\n")
// 	}
// 	if g.btn07.Update_Ret() {
// 		//fmt.Printf("Click2\n")
// 	}
// 	if g.btn08.Update_Ret() {
// 		//fmt.Printf("Click2\n")
// 	}
// 	if g.btn09.Update_Ret() {
// 		//fmt.Printf("Click2\n")
// 	}
// 	if g.btn10.Update_Ret() {
// 		//fmt.Printf("Click2\n")
// 	}
// 	if g.btn11.Update_Ret() {
// 		//fmt.Printf("Click2\n")
// 	}
// 	if g.btn12.Update_Ret() {
// 		//fmt.Printf("Click2\n")
// 	}
// 	if g.btn13.Update_Ret() {
// 		//fmt.Printf("Click2\n")
// 	}
// 	if g.btn14.Update_Ret() {
// 		//fmt.Printf("Click2\n")
// 	}
// }

// func (g *Game) DrawUI(screen *ebiten.Image) {
// 	g.btn00.Draw(screen)
// 	g.btn01.Draw(screen)
// 	g.btn02.Draw(screen)
// 	g.btn03.Draw(screen)
// 	g.btn04.Draw(screen)
// 	g.btn05.Draw(screen)
// 	g.btn06.Draw(screen)
// 	g.btn07.Draw(screen)
// 	g.btn08.Draw(screen)
// 	g.btn09.Draw(screen)
// 	g.btn10.Draw(screen)
// 	g.btn11.Draw(screen)
// 	g.btn12.Draw(screen)
// 	g.btn13.Draw(screen)
// 	g.btn14.Draw(screen)

// 	g.lbl00.Draw(screen)
// 	g.lbl01.Draw(screen)
// 	g.lbl02.Draw(screen)

// 	g.s_btn00.Draw(screen)
// 	g.s_btn01.Draw(screen)
// 	g.s_btn02.Draw(screen)
// 	g.s_btn03.Draw(screen)
// 	g.s_btn04.Draw(screen)
// 	g.s_btn05.Draw(screen)

// }

// func (g *Game) InitGUI() {
// 	num := g.Backend.Settings.ScreenResX - 70
// 	// game.lbl00.Init(&game.Backend, "This Is Text", coords.CoordInts{X: num - 136, Y: 2}, coords.CoordInts{X: 200, Y: 32}, nil)
// 	g.lbl00.Init_00(&g.Backend, "This Is Text", coords.CoordInts{X: 4, Y: 2}, coords.CoordInts{X: g.G_Setting.ScreenResX - 8, Y: 32}, nil)

// 	Row0 := 36
// 	g.btn00.Init(&g.Backend, "Save Map", coords.CoordInts{X: num - 136, Y: Row0}, coords.CoordInts{X: 64, Y: 32}, 0, nil)
// 	g.btn01.Init(&g.Backend, "Load Map", coords.CoordInts{X: num - 68, Y: Row0}, coords.CoordInts{X: 64, Y: 32}, 0, nil)
// 	g.btn02.Init(&g.Backend, "Clear Map", coords.CoordInts{X: num, Y: Row0}, coords.CoordInts{X: 64, Y: 32}, 0, nil)
// 	Row1 := Row0 + 34
// 	g.btn03.Init(&g.Backend, "Button03", coords.CoordInts{X: num - 136, Y: Row1}, coords.CoordInts{X: 64, Y: 32}, 0, nil)
// 	g.btn04.Init(&g.Backend, "Button04", coords.CoordInts{X: num - 68, Y: Row1}, coords.CoordInts{X: 64, Y: 32}, 0, nil)
// 	g.btn05.Init(&g.Backend, "Button05", coords.CoordInts{X: num, Y: Row1}, coords.CoordInts{X: 64, Y: 32}, 0, nil)
// 	Row2 := Row1 + 34
// 	g.btn06.Init(&g.Backend, "Button06", coords.CoordInts{X: num - 136, Y: Row2}, coords.CoordInts{X: 64, Y: 32}, 0, nil)
// 	g.btn07.Init(&g.Backend, "Button07", coords.CoordInts{X: num - 68, Y: Row2}, coords.CoordInts{X: 64, Y: 32}, 0, nil)
// 	g.btn08.Init(&g.Backend, "Button08", coords.CoordInts{X: num, Y: Row2}, coords.CoordInts{X: 64, Y: 32}, 0, nil)
// 	Row3 := Row2 + 34
// 	g.btn09.Init(&g.Backend, "Button09", coords.CoordInts{X: num - 136, Y: Row3}, coords.CoordInts{X: 64, Y: 32}, 0, nil)
// 	g.btn10.Init(&g.Backend, "Button10", coords.CoordInts{X: num - 68, Y: Row3}, coords.CoordInts{X: 64, Y: 32}, 0, nil)
// 	g.btn11.Init(&g.Backend, "Button11", coords.CoordInts{X: num, Y: Row3}, coords.CoordInts{X: 64, Y: 32}, 0, nil)
// 	Row4 := Row3 + 34
// 	g.btn12.Init(&g.Backend, "Button12", coords.CoordInts{X: num - 136, Y: Row4}, coords.CoordInts{X: 64, Y: 32}, 0, nil)
// 	g.btn13.Init(&g.Backend, "Button13", coords.CoordInts{X: num - 68, Y: Row4}, coords.CoordInts{X: 64, Y: 32}, 0, nil)
// 	g.btn14.Init(&g.Backend, "Button14", coords.CoordInts{X: num, Y: Row4}, coords.CoordInts{X: 64, Y: 32}, 0, nil)
// 	Row5 := Row4 + 34
// 	g.lbl01.Init_00(&g.Backend, "This Is Text", coords.CoordInts{X: num - 136, Y: Row5}, coords.CoordInts{X: 200, Y: 32}, nil)
// 	Row6 := Row5 + 34
// 	g.s_btn00.Init(&g.Backend, "A", coords.CoordInts{X: num - 136, Y: Row6}, coords.CoordInts{X: 32, Y: 32}, 0, nil)
// 	g.s_btn01.Init(&g.Backend, "B", coords.CoordInts{X: num - 136 + 33, Y: Row6}, coords.CoordInts{X: 32, Y: 32}, 0, nil)
// 	g.s_btn02.Init(&g.Backend, "C", coords.CoordInts{X: num - 136 + 33 + 33, Y: Row6}, coords.CoordInts{X: 32, Y: 32}, 0, nil)
// 	g.s_btn03.Init(&g.Backend, "D", coords.CoordInts{X: num - 136 + 33 + 36 + 33, Y: Row6}, coords.CoordInts{X: 32, Y: 32}, 0, nil)
// 	g.s_btn04.Init(&g.Backend, "E", coords.CoordInts{X: num - 136 + 33 + 36 + 33 + 33, Y: Row6}, coords.CoordInts{X: 32, Y: 32}, 0, nil)
// 	g.s_btn05.Init(&g.Backend, "F", coords.CoordInts{X: num - 136 + 33 + 36 + 33 + 33 + 33, Y: Row6}, coords.CoordInts{X: 32, Y: 32}, 0, nil)
// 	Row7 := Row6 + 34
// 	g.lbl02.Init_00(&g.Backend, "This Is Text", coords.CoordInts{X: num - 136, Y: Row7}, coords.CoordInts{X: 200, Y: 200}, nil)
// }
