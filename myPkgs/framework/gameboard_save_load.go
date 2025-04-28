package framework

import (
	"fmt"

	"github.com/KelleyTyler/GridTileEbit04_12/myPkgs/basic_geometry/coords"
)

/*
	Here I will Implement A System To Load And Save A Map
*/

func (gb *GameBoard) Load_Button_Pressed() {
	// fmt.Printf("LOAD MAP BUTTON PRESSED\n")
	if gb.GameBoard_UI_STATE != 30 {
		gb.GameBoard_UI_STATE = 30
		gb.Window_Save.IsActive = false
		gb.Window_Load.IsActive = true
		gb.Window_Load.IsVisible = true
		gb.Window_Save.IsVisible = false
		gb.Window_Load.State = 0
		gb.Window_Save.Textfield.Get_String_Output()

	} else {
		gb.GameBoard_UI_STATE = 10
		gb.Window_Save.IsActive = false
		gb.Window_Load.IsActive = false
		gb.Window_Save.IsVisible = false
		gb.Window_Load.IsVisible = false
		gb.Window_Save.State = 0
		gb.Window_Load.State = 0
		gb.Window_Load.Textfield.Get_String_Output()
		gb.Window_Save.Textfield.Get_String_Output()

	}
}

func (gb *GameBoard) Save_Button_Pressed() {
	if gb.GameBoard_UI_STATE != 40 {
		gb.GameBoard_UI_STATE = 40
		gb.Window_Save.IsActive = true
		gb.Window_Save.IsVisible = true
		gb.Window_Load.IsVisible = false
		gb.Window_Load.IsActive = false
		gb.Window_Save.State = 0
		gb.Window_Load.Textfield.Get_String_Output()

	} else {
		gb.GameBoard_UI_STATE = 10
		gb.Window_Save.IsActive = false
		gb.Window_Load.IsActive = false
		gb.Window_Save.IsVisible = false
		gb.Window_Load.IsVisible = false
		gb.Window_Save.State = 0
		gb.Window_Load.State = 0
		gb.Window_Load.Textfield.Get_String_Output()
		gb.Window_Save.Textfield.Get_String_Output()
		// fmt.Printf("SAVE MAP BUTTON PRESSED\n")

	}
}

/**/

func (gb *GameBoard) Save_A_File_Activate(file_name string) {
	var err error
	fmt.Printf("SAVE MAP BUTTON PRESSED %s\n", file_name)
	err = gb.IMat.Save_A_File(fmt.Sprintf("%s/%s", gb.SavePath, file_name))
	if err != nil {
		fmt.Printf("File Does Not Exist\n")
	}
}

/**/

func (gb *GameBoard) Load_A_File_Activate(file_name string) {
	var err error
	fmt.Printf("LOAD MAP BUTTON PRESSED %s\n", file_name)
	gb.IMat, err = gb.IMat.Load_A_File(fmt.Sprintf("%s/%s", gb.SavePath, file_name))
	if err != nil {
		fmt.Printf("File Does Not Exist\n")
	}
	gb.Redraw_Board_New_Params(coords.CoordInts{X: gb.NumSelect_TileSize_X.CurrValue, Y: gb.NumSelect_TileSize_Y.CurrValue}, coords.CoordInts{X: gb.NumSelect_Tile_Margin_X.CurrValue, Y: gb.NumSelect_Tile_Margin_Y.CurrValue})

}

/**/
func (gb *GameBoard) Save_Load_Update() {
	gb.Window_Save.Update()
	gb.Window_Load.Update()
	WS := gb.Window_Save.GetState()
	switch WS {
	case 90: //----Not visible
		gb.GameBoard_UI_STATE = 10
		gb.Window_Save.IsActive = false
		gb.Window_Save.IsVisible = false
		gb.Window_Save.State = 0
		// break
	case 30:
	case 20: //----Not visible
	case 10: //----Not visible
	default:
		if gb.Window_Save.Button_Submit.GetState() == 2 {
			gb.Save_A_File_Activate(gb.Window_Save.Textfield.Get_String_Output())

		}
	}

	WL := gb.Window_Load.GetState()
	switch WL {
	case 90: //----Not visible
		gb.GameBoard_UI_STATE = 10
		gb.Window_Load.IsActive = false
		gb.Window_Load.IsVisible = false
		gb.Window_Save.State = 0
		// break
	case 30:
	case 20: //----Not visible
	case 10: //----Not visible
	default:
		if gb.Window_Load.Button_Submit.GetState() == 2 {
			gb.Load_A_File_Activate(gb.Window_Load.Textfield.Get_String_Output())

		}

	}

}

// func (gb *GameBoard) Init_SaveMenu() {

// }
