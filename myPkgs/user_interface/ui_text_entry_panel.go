package user_interface

import (
	"fmt"

	"github.com/KelleyTyler/GridTileEbit04_12/myPkgs/basic_geometry/coords"
	"github.com/hajimehoshi/ebiten/v2"
)

//tew *UI_TextEntryWindow

/*

	Ideally this should just be it's own 'panel window'
*/

type UI_Text_Entry_Panel struct {
	init                  bool
	IsActive, IsVisible   bool
	ui_obj_id, panel_name string
	Position, Dimensions  coords.CoordInts
	State                 uint8 //should be a const
	Style                 *UI_Object_Style
	Backend               *UI_Backend
	Image                 *ebiten.Image
	ImageUpdate           bool
	Parent                UI_Object
	//----------------------------
	Button_Submit, Button_Clear UI_Button
	Panel_Label                 UI_Label
	Textfield                   UI_TextEntryField
	Dialogue_Text               string
	error_message               string
}

func (ui_tep *UI_Text_Entry_Panel) Init(idLabels []string, backend *UI_Backend, style *UI_Object_Style, Position, Dimensions coords.CoordInts) error {

	ui_tep.ui_obj_id = idLabels[0]
	ui_tep.panel_name = idLabels[1]

	ui_tep.Dimensions = Dimensions
	ui_tep.Position = Position
	ui_tep.Backend = backend
	if style != nil {
		ui_tep.Style = style
	} else {
		ui_tep.Style = &ui_tep.Backend.Style
	}
	ui_tep.State = 0

	//-------Setting up Image
	ui_tep.Image = ebiten.NewImage(Dimensions.X, Dimensions.Y)
	ui_tep.Redraw()
	ui_tep.ImageUpdate = true
	//------Finishing Up
	if !ui_tep.init {
		ui_tep.init = true
	}
	ui_tep.Panel_Label.Init([]string{"ui_txt_entry_panel_label", idLabels[1]}, backend, nil, coords.CoordInts{X: 0, Y: 0}, coords.CoordInts{X: Dimensions.X, Y: 32})
	ui_tep.Panel_Label.Init_Parents(ui_tep)
	ui_tep.Panel_Label.TextAlignMode = 10
	ui_tep.Panel_Label.Redraw()
	ui_tep.Redraw()
	//---------------------------------- This Needs To be 'dealt with' somehow; not sure 'how' but 'somehow'
	// ui_tep.tObject = ui_tep.textfield
	ui_tep.Textfield.Init([]string{"ui_txt_entry_panel_textfield", ""}, backend, nil, coords.CoordInts{X: 9, Y: 48}, coords.CoordInts{X: Dimensions.X - 18, Y: ui_tep.Panel_Label.Dimensions.Y * 1})
	ui_tep.Textfield.Init_Parents(ui_tep)
	ui_tep.Textfield.Redraw()
	ui_tep.Redraw()
	ui_tep.Dialogue_Text = ""
	ui_tep.error_message = ""
	ui_tep.Button_Submit.Init([]string{"ui_txt_entry_panel_textfield", "Submit"}, backend, nil, coords.CoordInts{X: (Dimensions.X / 2) - 32, Y: Dimensions.Y - 40}, coords.CoordInts{X: 64, Y: 32})
	ui_tep.Button_Submit.Init_Parents(ui_tep)
	ui_tep.Button_Submit.Redraw()
	ui_tep.Redraw()

	ui_tep.State = 0
	// ui_tep.oldMouse = coords.CoordInts{X: 0, Y: 0}
	return nil
}

/**/
func (ui_tep *UI_Text_Entry_Panel) Init_Parents(parent UI_Object) error {

	ui_tep.Parent = parent
	ui_tep.Parent.AddChild(ui_tep)
	ui_tep.Redraw()
	ui_tep.Parent.Redraw()
	return nil
}

/**/
func (ui_tep *UI_Text_Entry_Panel) Draw(screen *ebiten.Image) error {
	if ui_tep.IsActive && ui_tep.IsVisible {
		ops := ebiten.DrawImageOptions{}
		scale := 1.0
		ops.GeoM.Reset()
		ops.GeoM.Translate(float64(ui_tep.Position.X)*scale, float64(ui_tep.Position.Y)*scale)
		screen.DrawImage(ui_tep.Image, &ops)
	}
	return nil
}

/*
 */
func (ui_tep *UI_Text_Entry_Panel) Redraw() {
	ui_tep.Image.Fill(ui_tep.Style.PanelColor)

}

/**/
func (ui_tep *UI_Text_Entry_Panel) Update() error {

	return nil
}

/**/
func (ui_tep *UI_Text_Entry_Panel) Update_Unactive() error { return nil }

/**/
func (ui_tep *UI_Text_Entry_Panel) Update_Ret_State_Redraw_Status() (uint8, bool, error) {
	xx, yy := ebiten.CursorPosition()

	return ui_tep.Update_Ret_State_Redraw_Status_Mport(xx, yy, 0)
}

/**/
func (ui_tep *UI_Text_Entry_Panel) Print_Error_Message(error_message string) {
	ui_tep.error_message = error_message
	ui_tep.Button_Submit.DeToggle()
	ui_tep.Textfield.Clear()
	ui_tep.Redraw()
	ui_tep.State = 0
}

/**/
func (ui_tep *UI_Text_Entry_Panel) Update_Ret_State_Redraw_Status_Mport(Mouse_Pos_X, Mouse_Pos_Y, mode int) (uint8, bool, error) {
	var end_state uint8 = 0
	var end_redraw bool = false
	var end_err error = nil
	if ui_tep.IsActive && ui_tep.IsVisible {
		if ui_tep.IsCursorInBounds_MousePort(Mouse_Pos_X, Mouse_Pos_Y, mode) {
			//-------This is Where We Have It:
			if state, redraw, _ := ui_tep.Button_Submit.Update_Ret_State_Redraw_Status_Mport(Mouse_Pos_X, Mouse_Pos_Y, mode); state == 2 {
				if redraw {
					ui_tep.Redraw()
					if !end_redraw {
						end_redraw = true
					}
				}
				if ui_tep.Textfield.DataField != "" {
					end_state = 80
				}
			}
		}
	} else {
		//else make sure the parent is visible;
	}
	return end_state, end_redraw, end_err
}

/**/
func (ui_tep *UI_Text_Entry_Panel) GetState() uint8 {
	return ui_tep.State
}

/**/
func (ui_tep *UI_Text_Entry_Panel) ToString() string {
	return fmt.Sprintf("ui_tep_PANE; AT %3d,%3d", ui_tep.Position.X, ui_tep.Position.Y)
}

/**/
func (ui_tep *UI_Text_Entry_Panel) IsInit() bool { return ui_tep.init }

/**/
func (ui_tep *UI_Text_Entry_Panel) GetID() string { return ui_tep.ui_obj_id }

/**/
func (ui_tep *UI_Text_Entry_Panel) GetType() string { return "ui_tep_pane" }

// /**/
func (ui_tep *UI_Text_Entry_Panel) IsCursorInBounds() bool {
	if ui_tep.IsActive && ui_tep.IsVisible {
		mouse_Pos_X, mouse_Pos_Y := ebiten.CursorPosition()
		var x0, y0, x1, y1 int

		if ui_tep.Parent != nil {
			px, py := ui_tep.Parent.GetPosition_Int()
			x0 = ui_tep.Position.X + px
			y0 = ui_tep.Position.Y + py
			x1 = ui_tep.Position.X + ui_tep.Image.Bounds().Dx() + px
			y1 = ui_tep.Position.Y + ui_tep.Image.Bounds().Dy() + py
			// x0 = prim.Position.X + prim.ParentPos.X
			// y0 = prim.Position.Y + prim.ParentPos.X
			// x1 = prim.Position.X + prim.ParentPos.X + prim.Dimensions.X
			// y1 = prim.Position.Y + prim.ParentPos.Y + prim.Dimensions.Y
		} else {
			x0 = ui_tep.Position.X
			y0 = ui_tep.Position.Y
			x1 = ui_tep.Position.X + ui_tep.Image.Bounds().Dx()
			y1 = ui_tep.Position.Y + ui_tep.Image.Bounds().Dy()
		}
		return (mouse_Pos_X > x0 && mouse_Pos_X < x1) && (mouse_Pos_Y > y0 && mouse_Pos_Y < y1)
	}
	return false
}

/**/
func (ui_tep *UI_Text_Entry_Panel) IsCursorInBounds_MousePort(Mouse_Pos_X, Mouse_Pos_Y, mode int) bool {
	if ui_tep.IsActive && ui_tep.IsVisible {
		var x0, y0, x1, y1 int

		if ui_tep.Parent != nil {
			px, py := ui_tep.Parent.GetPosition_Int()
			x0 = ui_tep.Position.X + px
			y0 = ui_tep.Position.Y + py
			x1 = ui_tep.Position.X + ui_tep.Image.Bounds().Dx() + px
			y1 = ui_tep.Position.Y + ui_tep.Image.Bounds().Dy() + py
			// x0 = prim.Position.X + prim.ParentPos.X
			// y0 = prim.Position.Y + prim.ParentPos.X
			// x1 = prim.Position.X + prim.ParentPos.X + prim.Dimensions.X
			// y1 = prim.Position.Y + prim.ParentPos.Y + prim.Dimensions.Y
		} else {
			x0 = ui_tep.Position.X
			y0 = ui_tep.Position.Y
			x1 = ui_tep.Position.X + ui_tep.Image.Bounds().Dx()
			y1 = ui_tep.Position.Y + ui_tep.Image.Bounds().Dy()
		}
		return (Mouse_Pos_X > x0 && Mouse_Pos_X < x1) && (Mouse_Pos_Y > y0 && Mouse_Pos_Y < y1)
	}
	return false
}

/**/
func (ui_tep *UI_Text_Entry_Panel) GetPosition_Int() (int, int) {
	return ui_tep.Position.X, ui_tep.Position.Y
}

/**/
func (ui_tep *UI_Text_Entry_Panel) SetPosition_Int(x_point, y_point int) {
	ui_tep.Position = coords.CoordInts{X: x_point, Y: y_point}
}

/**/
func (ui_tep *UI_Text_Entry_Panel) GetDimensions_Int() (int, int) {
	return ui_tep.Dimensions.X, ui_tep.Dimensions.Y
}

/**/
func (ui_tep *UI_Text_Entry_Panel) SetDimensions_Int(x_point, y_point int) {
	ui_tep.Dimensions = coords.CoordInts{X: x_point, Y: y_point}
	//---Redraw The S
}

/**/
func (ui_tep *UI_Text_Entry_Panel) GetNumber_Children() int { return 0 }

/**/
func (ui_tep *UI_Text_Entry_Panel) GetChild(index int) UI_Object { return nil }

/**/
func (ui_tep *UI_Text_Entry_Panel) AddChild(child UI_Object) error { return nil }

/**/
func (ui_tep *UI_Text_Entry_Panel) RemoveChild(index int) error { return nil }

/**/
func (ui_tep *UI_Text_Entry_Panel) GetParent() UI_Object {
	return ui_tep.Parent
}

/**/
func (ui_tep *UI_Text_Entry_Panel) HasParent() bool {
	return ui_tep.Parent != nil
}
