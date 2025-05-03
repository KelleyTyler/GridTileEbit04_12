package framework

import (
	"github.com/KelleyTyler/GridTileEbit04_12/myPkgs/basic_geometry/coords"
	ui "github.com/KelleyTyler/GridTileEbit04_12/myPkgs/user_interface"
)

/*
 */
func (gb *GameBoard) UI_INIT() {

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
	gb.Window_Test.Init([]string{"window_test", "TEST WINDOW"}, gb.UI_Backend, nil, coords.CoordInts{X: 55, Y: 150}, coords.CoordInts{X: 256, Y: 256 + 32})

}

/**/
func (gb *GameBoard) SetParents(parent ui.UI_Object) {
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

	gb.mazeTool.UI_Init(parent, 170+38)
	gb.drawTool.InitUI(parent, 306+38)
	gb.pfindTest.UI_Init(parent, 442+38)
	gb.Test_Window_Button.Init_Parents(parent)
	//--------------------------------
	// gb.Window_Save.Init_Parents(parent)
}
