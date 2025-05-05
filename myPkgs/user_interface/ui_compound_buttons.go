package user_interface

import (
	"fmt"
	"log"

	"github.com/KelleyTyler/GridTileEbit04_12/myPkgs/basic_geometry/coords"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

/*

The idea here is to have buttons that are 'compound' or 'combo' buttons

that is number selection buttons, drop down menus, etc.
*/

type UI_Num_Select struct {
	Parent                                                 UI_Object
	Position, Dimensions                                   coords.CoordInts
	Backend                                                *UI_Backend
	Style                                                  *UI_Object_Style
	obj_id, text                                           string
	State                                                  uint8
	Image                                                  *ebiten.Image
	ImageUpdate                                            bool
	init                                                   bool
	IsActive, IsVisible                                    bool
	R_Button                                               UI_Button
	M_Button                                               UI_Button
	L_Button                                               UI_Button
	Label                                                  UI_Label
	CurrValue, MinValue, MaxValue, DefaultValue, IterValue int
	MiddleButtonMode                                       uint8
}

/*
 */
func (ui_num_select *UI_Num_Select) Init(idLabels []string, backend *UI_Backend, style *UI_Object_Style, Position, Dimensions coords.CoordInts) error {
	ui_num_select.obj_id = idLabels[0]
	ui_num_select.text = idLabels[1]
	ui_num_select.Dimensions = Dimensions
	ui_num_select.Position = Position
	ui_num_select.Backend = backend
	if style != nil {
		ui_num_select.Style = style
	} else {
		ui_num_select.Style = &ui_num_select.Backend.Style
	}
	ui_num_select.State = 0
	ui_num_select.Image = ebiten.NewImage(Dimensions.X, Dimensions.Y)
	dimY := (ui_num_select.Dimensions.Y / 2)
	thick := int(ui_num_select.Style.BorderThickness)
	btnwidth := 16
	ui_num_select.L_Button.Init([]string{"lbtn", "<"}, backend, nil, coords.CoordInts{X: (thick / 2), Y: dimY - (thick / 2)}, coords.CoordInts{X: btnwidth, Y: dimY})
	ui_num_select.M_Button.Init([]string{"lbtn", "000"}, backend, nil, coords.CoordInts{X: btnwidth + (thick / 2), Y: dimY - (thick / 2)}, coords.CoordInts{X: ui_num_select.Dimensions.X - ((btnwidth * 2) + (thick)), Y: dimY})
	ui_num_select.R_Button.Init([]string{"lbtn", ">"}, backend, nil, coords.CoordInts{X: ui_num_select.Dimensions.X - (btnwidth + thick - 1), Y: dimY - (thick / 2)}, coords.CoordInts{X: btnwidth, Y: dimY})
	ui_num_select.Label.Init([]string{"lbtn", idLabels[1]}, backend, nil, coords.CoordInts{X: thick / 2, Y: thick / 2}, coords.CoordInts{X: ui_num_select.Dimensions.X - (thick), Y: dimY})
	ui_num_select.Label.TextAlignMode = 10
	ui_num_select.L_Button.Init_Parents(ui_num_select)
	ui_num_select.M_Button.Init_Parents(ui_num_select)
	ui_num_select.R_Button.Init_Parents(ui_num_select)

	ui_num_select.Label.Init_Parents(ui_num_select)
	ui_num_select.Label.Redraw()
	ui_num_select.SetVals(0, 1, -10, 10, 0)
	// ui_num_select.R_Button.Init_00(backend, "->", coords.CoordInts{X: ui_num_select.Dimensions.X - 16, Y: 0}, coords.CoordInts{X: 16, Y: 16}, 0, ui_num_select)
	//-------Setting up Image
	ui_num_select.Redraw()
	ui_num_select.ImageUpdate = true

	//------Finishing Up
	if !ui_num_select.init {
		ui_num_select.init = true
	}
	return nil
}

/*
 */
func (ui_num_select *UI_Num_Select) Init_Parents(parent UI_Object) error {
	ui_num_select.Parent = parent
	ui_num_select.Parent.AddChild(ui_num_select)
	ui_num_select.Redraw()
	ui_num_select.Parent.Redraw()
	return nil
}

/*
 */
func (ui_num_select *UI_Num_Select) SetVals(defVal, interator, min, max int, middlebtn uint8) {
	ui_num_select.MaxValue = max
	ui_num_select.MinValue = min
	ui_num_select.DefaultValue = defVal
	ui_num_select.CurrValue = defVal
	ui_num_select.IterValue = interator
	ui_num_select.MiddleButtonMode = middlebtn
}

/*
 */
func (ui_num_select *UI_Num_Select) Draw(screen *ebiten.Image) error {
	ops := ebiten.DrawImageOptions{}
	scale := 1.0
	ops.GeoM.Reset()
	ops.GeoM.Translate(float64(ui_num_select.Position.X)*scale, float64(ui_num_select.Position.Y)*scale)
	screen.DrawImage(ui_num_select.Image, &ops)
	return nil
}

/*
 */
func (ui_num_select *UI_Num_Select) Redraw() {

	ui_num_select.M_Button.Redraw()
	// ui_num_select.Image.Fill(color.RGBA{255, 0, 0, 255}) //ui_num_select.Style.BorderColor
	ui_num_select.Image.Fill(ui_num_select.Style.BorderColor) //

	lineThick := ui_num_select.Style.BorderThickness
	// vector.DrawFilledRect(ui_num_select.Image, lineThick, lineThick, float32(ui_num_select.Dimensions.X)-(2*lineThick), float32(ui_num_select.Dimensions.Y)-(2*lineThick), color.RGBA{255, 255, 0, 255}, true) //ui_num_select.Style.PanelColor
	vector.DrawFilledRect(ui_num_select.Image, lineThick, lineThick, float32(ui_num_select.Dimensions.X)-(2*lineThick), float32(ui_num_select.Dimensions.Y)-(2*lineThick), ui_num_select.Style.PanelColor, true) //

	ui_num_select.L_Button.Draw(ui_num_select.Image)
	ui_num_select.M_Button.Draw(ui_num_select.Image)
	ui_num_select.R_Button.Draw(ui_num_select.Image)
	ui_num_select.Label.Draw(ui_num_select.Image)
}

/*
allows for the seleciton of middle button modes
*/
func (ui_num_select *UI_Num_Select) SetMiddlebuttonMode(mode uint8, additionalLabels []string) {
	ui_num_select.MiddleButtonMode = mode
}

/*
 */
func (ui_num_select *UI_Num_Select) Update() error {
	ui_num_select.State = 0
	Mouse_Pos_X, Mouse_Pos_Y := ebiten.CursorPosition()
	_, _, err := ui_num_select.Update_Ret_State_Redraw_Status_Mport(Mouse_Pos_X, Mouse_Pos_Y, 0)

	return err
}

/*
 */
func (ui_num_select *UI_Num_Select) Update_Unactive() error {

	return nil
}

/*
This will return false; Use Only Sparingly!
*/
func (ui_num_select *UI_Num_Select) Update_Any() (any, error) {
	return false, nil
}

/*
 */
func (ui_num_select *UI_Num_Select) Update_Ret_State_Redraw_Status() (uint8, bool, error) {
	ui_num_select.State = 0
	Mouse_Pos_X, Mouse_Pos_Y := ebiten.CursorPosition()
	return ui_num_select.Update_Ret_State_Redraw_Status_Mport(Mouse_Pos_X, Mouse_Pos_Y, 0)
}

/*
Update_Ret_State_Redraw_Status
*/
func (ui_num_select *UI_Num_Select) Update_Ret_State_Redraw_Status_Mport(Mouse_Pos_X, Mouse_Pos_Y, mode int) (uint8, bool, error) {
	ui_num_select.State = 0
	state0, to_redraw0, err0 := ui_num_select.L_Button.Update_Ret_State_Redraw_Status_Mport(Mouse_Pos_X, Mouse_Pos_Y, mode)
	if err0 != nil {
		log.Fatal(err0)
	}

	if state0 == 2 {
		// log.Printf("BTN L is Clicked %d %t\n", ui_num_select.CurrValue, ui_num_select.CurrValue != ui_num_select.DefaultValue)

		if ui_num_select.CurrValue > ui_num_select.MinValue {
			ui_num_select.CurrValue -= ui_num_select.IterValue
		}
	}
	state1, to_redraw1, err1 := ui_num_select.M_Button.Update_Ret_State_Redraw_Status_Mport(Mouse_Pos_X, Mouse_Pos_Y, mode)
	if err1 != nil {
		log.Fatal(err1)
	}

	if state1 == 2 {
		// log.Printf("BTN M is Clicked %d  %d %d %t\n", ui_num_select.CurrValue, ui_num_select.IterValue, ui_num_select.MinValue, ui_num_select.CurrValue != ui_num_select.DefaultValue)
		if ui_num_select.MiddleButtonMode == 0 {
			if ui_num_select.CurrValue != ui_num_select.DefaultValue {
				ui_num_select.CurrValue = ui_num_select.DefaultValue
			}
		} else {
			ui_num_select.State = 2
		}
	}
	state2, to_redraw2, err2 := ui_num_select.R_Button.Update_Ret_State_Redraw_Status_Mport(Mouse_Pos_X, Mouse_Pos_Y, mode)
	if err2 != nil {
		log.Fatal(err2)
	}

	if state2 == 2 {

		if ui_num_select.CurrValue < ui_num_select.MaxValue {
			ui_num_select.CurrValue = ui_num_select.CurrValue + ui_num_select.IterValue
			//log.Printf("BTN R is Clicked %d %t\n", ui_num_select.CurrValue, ui_num_select.CurrValue < ui_num_select.MaxValue)
		}
	}

	// if ui_num_select.IsMovable {
	// 	if ui_num_select.IsCursorInBounds() {

	// 	}
	// }
	export_redraw := false
	if to_redraw0 || to_redraw1 || to_redraw2 {
		ui_num_select.M_Button.Label = fmt.Sprintf("%03d", ui_num_select.CurrValue)
		ui_num_select.Redraw()
		export_redraw = true
	}

	return ui_num_select.State, export_redraw, nil
}

/*
This returns the state of the object
*/
func (ui_num_select *UI_Num_Select) GetState() uint8 { return ui_num_select.State }

/*
this returns a basic to string message
*/
func (ui_num_select *UI_Num_Select) ToString() string {
	strngOut := fmt.Sprintf("UI_Object ui_num_select:%s\n\tPositon %s\t", ui_num_select.obj_id, ui_num_select.Position.ToString())
	strngOut += fmt.Sprintf("\tDimensions: %s\n", ui_num_select.Dimensions.ToString())
	return strngOut
}

/*
 */
func (ui_num_select *UI_Num_Select) IsCursorInBounds() bool {
	if ui_num_select.IsActive && ui_num_select.IsVisible {
		cX, cY := ebiten.CursorPosition()
		return ui_num_select.IsCursorInBounds_MousePort(cX, cY, 0)
	}
	return false
}

/*
Idea here is I don't want to waste time with having to get the cursor position when it possibly hasn't changed enough to matter;
This might be also a terrible idea overall I cannot tell quite yet

enter 0 for it to default
*/
func (ui_num_select *UI_Num_Select) IsCursorInBounds_MousePort(Mouse_Pos_X, Mouse_Pos_Y, mode int) bool {
	if ui_num_select.IsActive && ui_num_select.IsVisible && (mode == 0 || mode == 10) {
		cX, cY := Mouse_Pos_X, Mouse_Pos_Y
		//mode stuff
		var x0, y0, x1, y1 int

		if ui_num_select.Parent != nil {
			px, py := ui_num_select.Parent.GetPosition_Int()
			x0 = ui_num_select.Position.X + px
			y0 = ui_num_select.Position.Y + py
			x1 = ui_num_select.Position.X + ui_num_select.Dimensions.X + px
			y1 = ui_num_select.Position.Y + ui_num_select.Dimensions.Y + py
			if mode == 10 || mode == 0 {
				x3, y3 := ui_num_select.Parent.Get_Internal_Position_Int()
				x0 += x3
				x1 += x3
				y0 += y3
				y1 += y3
				// if !ui_num_select.Parent.IsCursorInBounds_MousePort(Mouse_Pos_X, Mouse_Pos_Y, 10) {
				// 	return false
				// }
			}

		} else {
			x0 = ui_num_select.Position.X
			y0 = ui_num_select.Position.Y
			x1 = ui_num_select.Position.X + ui_num_select.Dimensions.X
			y1 = ui_num_select.Position.Y + ui_num_select.Dimensions.Y
		}
		return (cX > x0 && cX < x1) && (cY > y0 && cY < y1)
	}
	//mode stuff
	return false
}

/**/
func (ui_num_select *UI_Num_Select) Close() {}

/**/
func (ui_num_select *UI_Num_Select) Open() {}

/**/
func (ui_num_select *UI_Num_Select) Detoggle() {
	ui_num_select.L_Button.Detoggle()
	ui_num_select.M_Button.Detoggle()
	ui_num_select.R_Button.Detoggle()
}

/**/
func (ui_num_select *UI_Num_Select) Get_Internal_Position_Int() (x_pos int, y_pos int) {
	// x_pos, y_pos = ui_num_select.GetPosition_Int()
	if ui_num_select.Parent != nil {
		x_pos, y_pos = ui_num_select.Parent.Get_Internal_Position_Int()
	}
	return x_pos, y_pos
}

/**/
func (ui_num_select *UI_Num_Select) Get_Internal_Dimensions_Int() (x_pos int, y_pos int) {
	x_pos, y_pos = ui_num_select.GetPosition_Int()
	return x_pos, y_pos
}

/*
 */
func (ui_num_select *UI_Num_Select) GetPosition_Int() (int, int) {
	xx := ui_num_select.Position.X
	yy := ui_num_select.Position.Y
	if ui_num_select.Parent != nil {
		px, py := ui_num_select.Parent.GetPosition_Int()
		xx += px
		yy += py
	}
	return xx, yy
}

/**/
func (ui_num_select *UI_Num_Select) SetPosition_Int(X, Y int) {

}

/**/
func (ui_num_select *UI_Num_Select) GetDimensions_Int() (int, int) {
	return 0, 0
} //
/**/
func (ui_num_select *UI_Num_Select) SetDimensions_Int(int, int) {

}

/*
this confirms the object is initilaized
*/
func (ui_num_select *UI_Num_Select) IsInit() bool { return ui_num_select.init }

/*
this gets the object ID
*/
func (ui_num_select *UI_Num_Select) GetID() string { return ui_num_select.obj_id }

/*
This returns a string specifying the objects type
*/
func (ui_num_select *UI_Num_Select) GetType() string { return "UI_Object ui_num_select" }

/*
 */
func (ui_num_select *UI_Num_Select) GetNumber_Children() int { return 0 }

/*
 */
func (ui_num_select *UI_Num_Select) GetChild(index int) UI_Object { return nil }

/*
 */
func (ui_num_select *UI_Num_Select) AddChild(child UI_Object) error { return nil }

/**/
func (ui_num_select *UI_Num_Select) RemoveChild(index int) error { return nil }

/**/
func (ui_num_select *UI_Num_Select) HasParent() bool { return ui_num_select.Parent != nil }

/**/
func (ui_num_select *UI_Num_Select) GetParent() UI_Object { return ui_num_select.Parent }

/**/
