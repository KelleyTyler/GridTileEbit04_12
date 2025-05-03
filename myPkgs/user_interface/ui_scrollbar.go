package user_interface

import (
	"fmt"
	"log"

	"github.com/KelleyTyler/GridTileEbit04_12/myPkgs/basic_geometry/coords"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type UI_Scrollbar struct {
	Parent               UI_Object
	Position, Dimensions coords.CoordInts
	Backend              *UI_Backend
	Style                *UI_Object_Style
	obj_id, text         string
	State                uint8
	Image                *ebiten.Image
	ImageUpdate          bool
	init                 bool
	IsActive, IsVisible  bool
	R_Button             UI_Button
	M_Button             UI_Button
	L_Button             UI_Button
	// Label                                                  UI_Label
	CurrValue, MinValue, MaxValue, DefaultValue, IterValue int
	MiddleButtonMode                                       uint8
}

/*
 */
func (ui_scrollbar *UI_Scrollbar) Init(idLabels []string, backend *UI_Backend, style *UI_Object_Style, Position, Dimensions coords.CoordInts) error {
	ui_scrollbar.obj_id = idLabels[0]
	ui_scrollbar.text = idLabels[1]

	ui_scrollbar.Backend = backend
	if style != nil {
		ui_scrollbar.Style = style
	} else {
		ui_scrollbar.Style = &ui_scrollbar.Backend.Style
	}
	isVert := Dimensions.X < Dimensions.Y
	if isVert {
		ui_scrollbar.Dimensions = coords.CoordInts{X: Dimensions.X, Y: Dimensions.Y}
		ui_scrollbar.Position = Position
		ui_scrollbar.State = 0
		ui_scrollbar.Image = ebiten.NewImage(ui_scrollbar.Dimensions.X, ui_scrollbar.Dimensions.Y)
		// dimY := (ui_scrollbar.Dimensions.Y / 2)
		thick := int(ui_scrollbar.Style.BorderThickness) / 2
		btnwidth := Dimensions.X //
		ui_scrollbar.L_Button.Init([]string{"lbtn", "-"}, backend, nil, coords.CoordInts{X: (thick / 2), Y: ui_scrollbar.Dimensions.Y - (btnwidth + (thick / 2))}, coords.CoordInts{X: btnwidth, Y: btnwidth})
		// ui_scrollbar.M_Button.Init([]string{"lbtn", "000"}, backend, nil, coords.CoordInts{X: btnwidth + (thick / 2), Y: (thick / 2)}, coords.CoordInts{X: ui_scrollbar.Dimensions.X - ((btnwidth * 2) + (thick)), Y: btnwidth})
		ui_scrollbar.M_Button.Init([]string{"lbtn", "000"}, backend, nil, coords.CoordInts{X: (thick / 2), Y: btnwidth + (thick / 2)}, coords.CoordInts{X: btnwidth, Y: ui_scrollbar.Dimensions.Y - ((btnwidth * 2) + (thick / 2))})

		// ui_scrollbar.R_Button.Init([]string{"lbtn", "-"}, backend, nil, coords.CoordInts{X: ui_scrollbar.Dimensions.X - (btnwidth + thick - 1), Y: (thick / 2)}, coords.CoordInts{X: btnwidth, Y: btnwidth})
		ui_scrollbar.R_Button.Init([]string{"lbtn", "+"}, backend, nil, coords.CoordInts{X: (thick / 2), Y: (thick / 2)}, coords.CoordInts{X: btnwidth, Y: btnwidth})

		// ui_scrollbar.Label.Init([]string{"lbtn", idLabels[1]}, backend, nil, coords.CoordInts{X: thick / 2, Y: thick / 2}, coords.CoordInts{X: ui_scrollbar.Dimensions.X - (thick), Y: dimY})
		// ui_scrollbar.Label.TextAlignMode = 10
		ui_scrollbar.L_Button.Btn_Type = 20
		ui_scrollbar.R_Button.Btn_Type = 20
		ui_scrollbar.L_Button.Init_Parents(ui_scrollbar)
		ui_scrollbar.M_Button.Init_Parents(ui_scrollbar)
		ui_scrollbar.R_Button.Init_Parents(ui_scrollbar)
	} else {
		ui_scrollbar.Dimensions = Dimensions
		ui_scrollbar.Position = Position
		ui_scrollbar.State = 0
		ui_scrollbar.Image = ebiten.NewImage(ui_scrollbar.Dimensions.X, ui_scrollbar.Dimensions.Y)
		// dimY := (ui_scrollbar.Dimensions.Y / 2)
		thick := int(ui_scrollbar.Style.BorderThickness)
		btnwidth := Dimensions.Y
		ui_scrollbar.L_Button.Init([]string{"lbtn", "+"}, backend, nil, coords.CoordInts{X: (thick / 2), Y: (thick / 2)}, coords.CoordInts{X: btnwidth, Y: btnwidth})

		ui_scrollbar.M_Button.Init([]string{"lbtn", "000"}, backend, nil, coords.CoordInts{X: btnwidth + (thick / 2), Y: (thick / 2)}, coords.CoordInts{X: ui_scrollbar.Dimensions.X - ((btnwidth * 2) + (thick / 2)), Y: btnwidth})

		ui_scrollbar.R_Button.Init([]string{"lbtn", "-"}, backend, nil, coords.CoordInts{X: ui_scrollbar.Dimensions.X - (btnwidth + (thick / 2)), Y: (thick / 2)}, coords.CoordInts{X: btnwidth, Y: btnwidth})
		ui_scrollbar.L_Button.Btn_Type = 20
		ui_scrollbar.R_Button.Btn_Type = 20

		// ui_scrollbar.Label.Init([]string{"lbtn", idLabels[1]}, backend, nil, coords.CoordInts{X: thick / 2, Y: thick / 2}, coords.CoordInts{X: ui_scrollbar.Dimensions.X - (thick), Y: dimY})
		// ui_scrollbar.Label.TextAlignMode = 10
		ui_scrollbar.L_Button.Init_Parents(ui_scrollbar)
		ui_scrollbar.M_Button.Init_Parents(ui_scrollbar)
		ui_scrollbar.R_Button.Init_Parents(ui_scrollbar)
	}

	// ui_scrollbar.Label.Init_Parents(ui_scrollbar)
	// ui_scrollbar.Label.Redraw()
	// ui_scrollbar.SetVals(0, 1, -10, 10, 0)
	// ui_scrollbar.R_Button.Init_00(backend, "->", coords.CoordInts{X: ui_scrollbar.Dimensions.X - 16, Y: 0}, coords.CoordInts{X: 16, Y: 16}, 0, ui_scrollbar)
	//-------Setting up Image
	ui_scrollbar.Redraw()
	ui_scrollbar.ImageUpdate = true

	//------Finishing Up
	if !ui_scrollbar.init {
		ui_scrollbar.init = true
	}
	return nil
}

/*
 */
func (ui_scrollbar *UI_Scrollbar) Init_Parents(parent UI_Object) error {
	ui_scrollbar.Parent = parent
	ui_scrollbar.Parent.AddChild(ui_scrollbar)
	ui_scrollbar.Redraw()
	ui_scrollbar.Parent.Redraw()
	return nil
}

/*
 */
func (ui_scrollbar *UI_Scrollbar) SetVals(defVal, interator, min, max int, middlebtn uint8) {
	ui_scrollbar.MaxValue = max
	ui_scrollbar.MinValue = min
	ui_scrollbar.DefaultValue = defVal
	ui_scrollbar.CurrValue = defVal
	ui_scrollbar.IterValue = interator
	ui_scrollbar.MiddleButtonMode = middlebtn
}

/*
 */
func (ui_scrollbar *UI_Scrollbar) Draw(screen *ebiten.Image) error {
	ops := ebiten.DrawImageOptions{}
	scale := 1.0
	ops.GeoM.Reset()
	ops.GeoM.Translate(float64(ui_scrollbar.Position.X)*scale, float64(ui_scrollbar.Position.Y)*scale)
	screen.DrawImage(ui_scrollbar.Image, &ops)
	return nil
}

/*
 */
func (ui_scrollbar *UI_Scrollbar) Redraw() {

	ui_scrollbar.M_Button.Redraw()
	// ui_scrollbar.Image.Fill(color.RGBA{255, 0, 0, 255}) //ui_scrollbar.Style.BorderColor
	ui_scrollbar.Image.Fill(ui_scrollbar.Style.BorderColor) //

	lineThick := ui_scrollbar.Style.BorderThickness
	// vector.DrawFilledRect(ui_scrollbar.Image, lineThick, lineThick, float32(ui_scrollbar.Dimensions.X)-(2*lineThick), float32(ui_scrollbar.Dimensions.Y)-(2*lineThick), color.RGBA{255, 255, 0, 255}, true) //ui_scrollbar.Style.PanelColor
	vector.DrawFilledRect(ui_scrollbar.Image, lineThick, lineThick, float32(ui_scrollbar.Dimensions.X)-(2*lineThick), float32(ui_scrollbar.Dimensions.Y)-(2*lineThick), ui_scrollbar.Style.PanelColor, true) //

	ui_scrollbar.L_Button.Draw(ui_scrollbar.Image)
	ui_scrollbar.M_Button.Draw(ui_scrollbar.Image)
	ui_scrollbar.R_Button.Draw(ui_scrollbar.Image)
	// ui_scrollbar.Label.Draw(ui_scrollbar.Image)
}

/*
allows for the seleciton of middle button modes
*/
func (ui_scrollbar *UI_Scrollbar) SetMiddlebuttonMode(mode uint8, additionalLabels []string) {
	ui_scrollbar.MiddleButtonMode = mode
}

/*
 */
func (ui_scrollbar *UI_Scrollbar) Update() error {
	ui_scrollbar.State = 0
	Mouse_Pos_X, Mouse_Pos_Y := ebiten.CursorPosition()
	_, _, err := ui_scrollbar.Update_Ret_State_Redraw_Status_Mport(Mouse_Pos_X, Mouse_Pos_Y, 0)

	return err
}

/*
 */
func (ui_scrollbar *UI_Scrollbar) Update_Unactive() error {

	return nil
}

/*
This will return false; Use Only Sparingly!
*/
func (ui_scrollbar *UI_Scrollbar) Update_Any() (any, error) {
	return false, nil
}

/*
 */
func (ui_scrollbar *UI_Scrollbar) Update_Ret_State_Redraw_Status() (uint8, bool, error) {
	ui_scrollbar.State = 0
	Mouse_Pos_X, Mouse_Pos_Y := ebiten.CursorPosition()
	return ui_scrollbar.Update_Ret_State_Redraw_Status_Mport(Mouse_Pos_X, Mouse_Pos_Y, 0)
}

/*
 */
func (ui_scrollbar *UI_Scrollbar) SetPosition(Position coords.CoordInts) {
	ui_scrollbar.Position = Position
}

/**/
func (ui_scrollbar *UI_Scrollbar) SetPosition_Int(x_point, y_point int) {
	ui_scrollbar.Position = coords.CoordInts{X: x_point, Y: y_point}

}

/**/
func (ui_scrollbar *UI_Scrollbar) GetDimensions_Int() (int, int) {
	return ui_scrollbar.Dimensions.X, ui_scrollbar.Dimensions.Y
} //
/**/
func (ui_scrollbar *UI_Scrollbar) SetDimensions_Int(int, int) {

}

/*
Update_Ret_State_Redraw_Status
*/
func (ui_scrollbar *UI_Scrollbar) Update_Ret_State_Redraw_Status_Mport(Mouse_Pos_X, Mouse_Pos_Y, mode int) (uint8, bool, error) {
	ui_scrollbar.State = 0
	state0, to_redraw0, err0 := ui_scrollbar.L_Button.Update_Ret_State_Redraw_Status_Mport(Mouse_Pos_X, Mouse_Pos_Y, mode)
	if err0 != nil {
		log.Fatal(err0)
	}

	if state0 == 2 {
		// log.Printf("BTN L is Clicked %d %t\n", ui_scrollbar.CurrValue, ui_scrollbar.CurrValue != ui_scrollbar.DefaultValue)

		if ui_scrollbar.CurrValue > ui_scrollbar.MinValue {
			ui_scrollbar.CurrValue -= ui_scrollbar.IterValue
		}
	}
	state1, to_redraw1, err1 := ui_scrollbar.M_Button.Update_Ret_State_Redraw_Status_Mport(Mouse_Pos_X, Mouse_Pos_Y, mode)
	if err1 != nil {
		log.Fatal(err1)
	}

	if state1 == 2 {
		// log.Printf("BTN M is Clicked %d  %d %d %t\n", ui_scrollbar.CurrValue, ui_scrollbar.IterValue, ui_scrollbar.MinValue, ui_scrollbar.CurrValue != ui_scrollbar.DefaultValue)
		if ui_scrollbar.MiddleButtonMode == 0 {
			if ui_scrollbar.CurrValue != ui_scrollbar.DefaultValue {
				ui_scrollbar.CurrValue = ui_scrollbar.DefaultValue
			}
		} else {
			ui_scrollbar.State = 2
		}
	}
	state2, to_redraw2, err2 := ui_scrollbar.R_Button.Update_Ret_State_Redraw_Status_Mport(Mouse_Pos_X, Mouse_Pos_Y, mode)
	if err2 != nil {
		log.Fatal(err2)
	}

	if state2 == 2 {

		if ui_scrollbar.CurrValue < ui_scrollbar.MaxValue {
			ui_scrollbar.CurrValue = ui_scrollbar.CurrValue + ui_scrollbar.IterValue
			//log.Printf("BTN R is Clicked %d %t\n", ui_scrollbar.CurrValue, ui_scrollbar.CurrValue < ui_scrollbar.MaxValue)
		}
	}

	// if ui_scrollbar.IsMovable {
	// 	if ui_scrollbar.IsCursorInBounds() {

	// 	}
	// }
	export_redraw := false
	if to_redraw0 || to_redraw1 || to_redraw2 {
		ui_scrollbar.M_Button.Label = fmt.Sprintf("%03d", ui_scrollbar.CurrValue)
		ui_scrollbar.Redraw()
		export_redraw = true
	}

	return ui_scrollbar.State, export_redraw, nil
}

/*
This returns the state of the object
*/
func (ui_scrollbar *UI_Scrollbar) GetState() uint8 { return ui_scrollbar.State }

/*
this returns a basic to string message
*/
func (ui_scrollbar *UI_Scrollbar) ToString() string {
	strngOut := fmt.Sprintf("UI_Object ui_scrollbar:%s\n\tPositon %s\t", ui_scrollbar.obj_id, ui_scrollbar.Position.ToString())
	strngOut += fmt.Sprintf("\tDimensions: %s\n", ui_scrollbar.Dimensions.ToString())
	return strngOut
}

/*
 */
func (ui_scrollbar *UI_Scrollbar) IsCursorInBounds() bool {
	if ui_scrollbar.IsActive && ui_scrollbar.IsVisible {
		cX, cY := ebiten.CursorPosition()
		return ui_scrollbar.IsCursorInBounds_MousePort(cX, cY, 0)
	}
	return false
}

/*
Idea here is I don't want to waste time with having to get the cursor position when it possibly hasn't changed enough to matter;
This might be also a terrible idea overall I cannot tell quite yet

enter 0 for it to default
*/
func (ui_scrollbar *UI_Scrollbar) IsCursorInBounds_MousePort(Mouse_Pos_X, Mouse_Pos_Y, mode int) bool {
	if ui_scrollbar.IsActive && ui_scrollbar.IsVisible && mode == 0 {
		cX, cY := Mouse_Pos_X, Mouse_Pos_Y
		//mode stuff
		var x0, y0, x1, y1 int

		if ui_scrollbar.Parent != nil {
			px, py := ui_scrollbar.Parent.GetPosition_Int()
			x0 = ui_scrollbar.Position.X + px
			y0 = ui_scrollbar.Position.Y + py
			x1 = ui_scrollbar.Position.X + ui_scrollbar.Dimensions.X + px
			y1 = ui_scrollbar.Position.Y + ui_scrollbar.Dimensions.Y + py
		} else {
			x0 = ui_scrollbar.Position.X
			y0 = ui_scrollbar.Position.Y
			x1 = ui_scrollbar.Position.X + ui_scrollbar.Dimensions.X
			y1 = ui_scrollbar.Position.Y + ui_scrollbar.Dimensions.Y
		}
		return (cX > x0 && cX < x1) && (cY > y0 && cY < y1)
	}
	//mode stuff
	return false
}

/*
 */
func (ui_scrollbar *UI_Scrollbar) GetPosition_Int() (int, int) {
	xx := ui_scrollbar.Position.X
	yy := ui_scrollbar.Position.Y
	if ui_scrollbar.Parent != nil {
		px, py := ui_scrollbar.Parent.GetPosition_Int()
		xx += px
		yy += py
	}
	return xx, yy
}

/**/
func (ui_scrollbar *UI_Scrollbar) Get_Internal_Position_Int() (x_pos int, y_pos int) {
	x_pos, y_pos = 0, 0
	return x_pos, y_pos
}

/*
this confirms the object is initilaized
*/
func (ui_scrollbar *UI_Scrollbar) IsInit() bool { return ui_scrollbar.init }

/*
this gets the object ID
*/
func (ui_scrollbar *UI_Scrollbar) GetID() string { return ui_scrollbar.obj_id }

/*
This returns a string specifying the objects type
*/
func (ui_scrollbar *UI_Scrollbar) GetType() string { return "UI_Object ui_scrollbar" }

/*
 */
func (ui_scrollbar *UI_Scrollbar) GetNumber_Children() int { return 0 }

/*
 */
func (ui_scrollbar *UI_Scrollbar) GetChild(index int) UI_Object { return nil }

/*
 */
func (ui_scrollbar *UI_Scrollbar) AddChild(child UI_Object) error { return nil }

/**/
func (ui_scrollbar *UI_Scrollbar) RemoveChild(index int) error { return nil }

/**/
func (ui_scrollbar *UI_Scrollbar) HasParent() bool { return ui_scrollbar.Parent != nil }

/**/
func (ui_scrollbar *UI_Scrollbar) GetParent() UI_Object { return ui_scrollbar.Parent }

/**/
func (ui_scrollbar *UI_Scrollbar) Close() {}

/**/
func (ui_scrollbar *UI_Scrollbar) Open() {}

/**/
func (ui_scrollbar *UI_Scrollbar) Detoggle() {}
