package framework

import (
	"fmt"

	"github.com/KelleyTyler/GridTileEbit04_12/myPkgs/basic_geometry/coords"
)

/*
	Here I will Implement A System To Load And Save A Map
*/

func (gb *GameBoard) Load_Button_Pressed() {

	if gb.GameBoard_UI_STATE != 30 {
		gb.GameBoard_UI_STATE = 30
		// gb.Window_Save.Close()
		gb.Window_Load.Open()

	} else {
		gb.GameBoard_UI_STATE = 10
		gb.Window_Save.Close()
		gb.Window_Load.Close()

	}
	fmt.Printf("LOAD MAP BUTTON PRESSED %d\n", gb.GameBoard_UI_STATE)
}

func (gb *GameBoard) Save_Button_Pressed() {

	if gb.GameBoard_UI_STATE != 40 {
		gb.GameBoard_UI_STATE = 40
		gb.Window_Save.Open()
		// gb.Window_Load.Close()

	} else {
		gb.GameBoard_UI_STATE = 10
		gb.Window_Save.Close()
		gb.Window_Load.Close()

	}
	fmt.Printf("SAVE MAP BUTTON PRESSED:%d\n", gb.GameBoard_UI_STATE)
}

/**/

func (gb *GameBoard) Save_A_File_Activate(file_name string) {
	var err error
	fmt.Printf("SAVE MAP BUTTON PRESSED %s\n", file_name)
	err = gb.IMat.Save_A_File(fmt.Sprintf("%s/%s", gb.SavePath, file_name))
	if err != nil {
		fmt.Printf("ERROR\n")
		gb.GameBoard_UI_STATE = 10
		gb.Window_Save.Close()
		gb.Window_Load.Close()
	} else {
		gb.GameBoard_UI_STATE = 10
		gb.Window_Save.Close()
		gb.Window_Load.Close()
	}
	gb.GameBoard_UI_STATE = 10

}

/**/

func (gb *GameBoard) Load_A_File_Activate(file_name string) {
	// var err error
	fmt.Printf("LOAD MAP BUTTON PRESSED %s\n", file_name)
	temp, err := gb.IMat.Load_A_File(fmt.Sprintf("%s/%s", gb.SavePath, file_name))
	if err != nil {
		fmt.Printf("File Does Not Exist\n")
		gb.GameBoard_UI_STATE = 10
		gb.Window_Save.Close()
		gb.Window_Load.Close()
	} else {
		gb.IMat = temp
		gb.GameBoard_UI_STATE = 10
		gb.Window_Save.Close()
		gb.Window_Load.Close()

		gb.Redraw_Board_New_Params(coords.CoordInts{X: gb.NumSelect_TileSize_X.CurrValue, Y: gb.NumSelect_TileSize_Y.CurrValue}, coords.CoordInts{X: gb.NumSelect_Tile_Margin_X.CurrValue, Y: gb.NumSelect_Tile_Margin_Y.CurrValue})
	}
	gb.GameBoard_UI_STATE = 10

}

/**/
func (gb *GameBoard) Save_Load_Update() {
	gb.Window_Save.Update()
	gb.Window_Load.Update()
	WS := gb.Window_Save.GetState()
	switch WS {
	case 90: //----Not visible
		gb.GameBoard_UI_STATE = 10
		gb.Window_Save.Close()
	case 80:
		gb.Save_A_File_Activate(gb.Window_Save.Textfield.Get_String_Output())

		// break
	case 30:
	case 20: //----Not visible
	case 10: //----Not visible
		break
	default:

	}

	WL := gb.Window_Load.GetState()
	switch WL {
	case 90: //----Not visible
		gb.GameBoard_UI_STATE = 10
		gb.Window_Load.Close()
		// break
	case 80:
		gb.Load_A_File_Activate(gb.Window_Load.Textfield.Get_String_Output())
	case 30:
	case 20: //----Not visible
	case 10: //----Not visible
		break
	default:
		// if gb.Window_Load.Button_Submit.GetState() == 2 {

		// }

	}

}

// func (gb *GameBoard) Init_SaveMenu() {

// }
