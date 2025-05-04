package framework

import (
	"image"
	"image/color"

	"github.com/KelleyTyler/GridTileEbit04_12/myPkgs/basic_geometry/coords"
	ui "github.com/KelleyTyler/GridTileEbit04_12/myPkgs/user_interface"
	"github.com/hajimehoshi/ebiten/v2"
)

/*
 */
func (gb *GameBoard) UI_INIT() {
	gb.Button_Panel.Init([]string{"gameboard_load_save_panel", "PRIMITIVE"}, gb.UI_Backend, nil, coords.CoordInts{X: 0, Y: 0}, coords.CoordInts{X: 204, Y: 138 + 68})
	gb.Button_Panel_Label.Init([]string{"gameboard_panel_label", "GAMEBOARD"}, gb.UI_Backend, nil, coords.CoordInts{X: 0, Y: 0}, coords.CoordInts{X: 204, Y: 16})
	gb.Button_Panel_Label.TextAlignMode = 10
	gb.Button_Panel_Label.Redraw()

	gb.Load_Map_Button.Init([]string{"r_map_btn", "LOAD\nMAP"}, gb.UI_Backend, nil, coords.CoordInts{X: 136, Y: 34}, coords.CoordInts{X: 64, Y: 32})
	gb.Save_Map_Button.Init([]string{"r_map_btn", "SAVE\nMAP"}, gb.UI_Backend, nil, coords.CoordInts{X: 4, Y: 34}, coords.CoordInts{X: 64, Y: 32})
	gb.NumSelect_ResetNumber.Init([]string{"n_map_btn", "mazeGen00"}, gb.UI_Backend, nil, coords.CoordInts{X: 70, Y: 34}, coords.CoordInts{X: 64, Y: 32})
	gb.NumSelect_ResetNumber.SetVals(10, 1, 0, 10, 0)
	gb.New_Map_Button.Init([]string{"n_map_btn", "New Map"}, gb.UI_Backend, nil, coords.CoordInts{X: 70, Y: 68}, coords.CoordInts{X: 64, Y: 32})
	gb.NumSelect_MapSize_X.Init([]string{"n_map_btn", "Map X"}, gb.UI_Backend, nil, coords.CoordInts{X: 4, Y: 68}, coords.CoordInts{X: 64, Y: 32}) //coords.CoordInts{X: 4, Y: 72}, coords.CoordInts{X: 68, Y: 36}
	gb.NumSelect_MapSize_X.SetVals(len(gb.IMat[0]), 1, 8, gb.UI_Backend.Settings.GameBoardXMax, 0)
	gb.NumSelect_MapSize_Y.Init([]string{"n_map_btn", "Map Y"}, gb.UI_Backend, nil, coords.CoordInts{X: 136, Y: 68}, coords.CoordInts{X: 64, Y: 32})
	gb.NumSelect_MapSize_Y.SetVals(len(gb.IMat), 1, 8, gb.UI_Backend.Settings.GameBoardYMax, 0)
	gb.NumSelect_TileSize_X.Init([]string{"n_map_btn", "TileX"}, gb.UI_Backend, nil, coords.CoordInts{X: 4, Y: 68 + 34}, coords.CoordInts{X: 64, Y: 32})
	gb.NumSelect_TileSize_X.SetVals(8, 1, 1, 64, 0)
	gb.Reset_Map_Btn.Init([]string{"r_map_btn", "ResetMap"}, gb.UI_Backend, nil, coords.CoordInts{X: 70, Y: 68 + 34}, coords.CoordInts{X: 64, Y: 32})

	gb.NumSelect_TileSize_Y.Init([]string{"n_map_btn", "TileY"}, gb.UI_Backend, nil, coords.CoordInts{X: 136, Y: 68 + 34}, coords.CoordInts{X: 64, Y: 32})
	gb.NumSelect_TileSize_Y.SetVals(8, 1, 1, 64, 0)

	gb.NumSelect_Tile_Margin_X.Init([]string{"n_map_btn", "TileMX"}, gb.UI_Backend, nil, coords.CoordInts{X: 4, Y: 68 + 68}, coords.CoordInts{X: 64, Y: 32})
	gb.NumSelect_Tile_Margin_X.SetVals(0, 1, 0, 16, 0)
	gb.Redraw_Tiles_Button.Init([]string{"r_map_btn", "Redraw\nTiles"}, gb.UI_Backend, nil, coords.CoordInts{X: 70, Y: 68 + 68}, coords.CoordInts{X: 64, Y: 32})
	gb.NumSelect_Tile_Margin_Y.Init([]string{"n_map_btn", "TileMY"}, gb.UI_Backend, nil, coords.CoordInts{X: 136, Y: 68 + 68}, coords.CoordInts{X: 64, Y: 32})
	gb.NumSelect_Tile_Margin_Y.SetVals(0, 1, 0, 16, 0)

	gb.Window_Save.Init([]string{"window_save", "SAVE WINDOW"}, gb.UI_Backend, nil, coords.CoordInts{X: 55, Y: 150}, coords.CoordInts{X: 256, Y: 144})
	gb.Window_Save.Redraw()
	gb.Window_Load.Init([]string{"window_load", "LOAD WINDOW"}, gb.UI_Backend, nil, coords.CoordInts{X: 55, Y: 150}, coords.CoordInts{X: 256, Y: 144})
	gb.Window_Load.Redraw()
	//Init([]string{"window_test", "TEST\n WINDOW"}, gb.UI_Backend, nil, coords.CoordInts{X: 55, Y: 150}, coords.CoordInts{X: 256, Y: 144})
	gb.Test_Window_Button.Init([]string{"btn_test_window", "Test\nWindow"}, gb.UI_Backend, nil, coords.CoordInts{X: 70, Y: 68 + 68 + 34}, coords.CoordInts{X: 64, Y: 32})
	gb.Window_Test.Init([]string{"window_test", "TEST WINDOW"}, gb.UI_Backend, nil, coords.CoordInts{X: 55, Y: 150}, coords.CoordInts{X: 256, Y: 128 + 32})
	gb.Perspective_Test_Button.Init([]string{"btn_test_perspective", "Test\nPerspective"}, gb.UI_Backend, nil, coords.CoordInts{X: 4, Y: 34}, coords.CoordInts{X: 64, Y: 32})
	// gb.Perspective_Test_Button.Btn_Type = 10
	gb.Set_Parent_Panel(&gb.Button_Panel)
	gb.Button_Panel.Redraw()
}

/**/
func (gb *GameBoard) SetParents(parent ui.UI_Object) {
	gb.Button_Panel.Init_Parents(parent)

	gb.mazeTool.UI_Init(parent, 170+38)
	gb.drawTool.InitUI(parent, 306+38)
	gb.pfindTest.UI_Init(parent, 442+38)
	//--------------------------------
	// gb.Window_Save.Init_Parents(parent)
}

/**/
func (gb *GameBoard) Set_Parent_Panel(parent ui.UI_Object) {
	gb.Button_Panel_Label.Init_Parents(parent)
	gb.Load_Map_Button.Init_Parents(parent)
	gb.Save_Map_Button.Init_Parents(parent)
	gb.NumSelect_ResetNumber.Init_Parents(parent)
	//-----------------------
	gb.New_Map_Button.Init_Parents(parent)
	gb.Reset_Map_Btn.Init_Parents(parent) //MazeGenBtn00
	//------
	gb.NumSelect_MapSize_X.Init_Parents(parent)
	gb.NumSelect_MapSize_Y.Init_Parents(parent)
	gb.NumSelect_TileSize_Y.Init_Parents(parent)
	gb.NumSelect_TileSize_X.Init_Parents(parent)
	gb.NumSelect_Tile_Margin_X.Init_Parents(parent)
	gb.NumSelect_Tile_Margin_Y.Init_Parents(parent)
	gb.Redraw_Tiles_Button.Init_Parents(parent)
	gb.Test_Window_Button.Init_Parents(parent)
	gb.Perspective_Test_Button.Init_Parents(&gb.Window_Test.Prim)
	//--------------------------------
	// gb.Window_Save.Init_Parents(parent)
}

/*
This attempts to draw with perspective;
*/
func (gb *GameBoard) PerpsecitveDraw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}

	op.GeoM.Reset()
	op.GeoM.Translate(float64(gb.BoardOptions.BoardPosition.X), float64(gb.BoardOptions.BoardPosition.Y))
	op.GeoM.Scale(1.0, 1.0)
	gb.Img.Fill(color.RGBA{20, 20, 20, 255})
	gb.Img.DrawImage(gb.Board_Buffer_Img, op)
	gb.Img.DrawImage(gb.Board_Overlay_Buffer_Img, op)
	w, h := gb.Board_Buffer_Img.Bounds().Dx(), gb.Board_Buffer_Img.Bounds().Dy()
	for i := 0; i < h; i++ {
		op.GeoM.Reset()

		// Move the image's center to the upper-left corner.
		op.GeoM.Translate(-float64(w)/2, -float64(h)/2)

		// Scale each lines and adjust the position.
		lineW := w + i*3/4
		x := -float64(lineW) / float64(w) / 2
		op.GeoM.Scale(float64(lineW)/float64(w), 1)
		op.GeoM.Translate(x, float64(i))

		// Move the image's center to the screen's center.
		op.GeoM.Translate(float64(gb.Position.X), float64(gb.Position.Y))

		screen.DrawImage(gb.Img.SubImage(image.Rect(0, i, w, i+1)).(*ebiten.Image), op)
	}

}
