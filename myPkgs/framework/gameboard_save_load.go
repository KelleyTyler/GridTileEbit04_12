package framework

import "fmt"

/*
	Here I will Implement A System To Load And Save A Map
*/

func (gb *GameBoard) Load_Button_Pressed() {
	fmt.Printf("LOAD MAP BUTTON PRESSED\n")
	if gb.GameBoard_UI_STATE != 30 {
		gb.GameBoard_UI_STATE = 30
		gb.Window_Save.IsActive = false
		gb.Window_Load.IsActive = true
		gb.Window_Load.IsVisible = true
		gb.Window_Save.IsVisible = false

	} else {
		gb.GameBoard_UI_STATE = 10
		gb.Window_Save.IsActive = false
		gb.Window_Load.IsActive = false
		gb.Window_Save.IsVisible = false
		gb.Window_Load.IsVisible = false
	}
}

func (gb *GameBoard) Save_Button_Pressed() {
	fmt.Printf("SAVE MAP BUTTON PRESSED\n")
	if gb.GameBoard_UI_STATE != 40 {
		gb.GameBoard_UI_STATE = 40
		gb.Window_Save.IsActive = true
		gb.Window_Save.IsVisible = true
		gb.Window_Load.IsVisible = false

		gb.Window_Load.IsActive = false
	} else {
		gb.GameBoard_UI_STATE = 10
		gb.Window_Save.IsActive = false
		gb.Window_Load.IsActive = false
		gb.Window_Save.IsVisible = false
		gb.Window_Load.IsVisible = false

		// fmt.Printf("SAVE MAP BUTTON PRESSED\n")

	}
}
func (gb *GameBoard) Save_Load_Update() {
	gb.Window_Save.Update()

	gb.Window_Load.Update()

}

// func (gb *GameBoard) Init_SaveMenu() {

// }
