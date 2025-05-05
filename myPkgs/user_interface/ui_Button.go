package user_interface

import (
	"fmt"
	"image/color"

	"github.com/KelleyTyler/GridTileEbit04_12/myPkgs/basic_geometry/coords"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

/*
type UI_Object interface {
	// Init0() //initialize void;
	Init(idLabels []string, backend *UI_Backend, style *UI_Object_Style, Position, Dimensions coords.CoordInts) error //--
	Init_Parents(Parent UI_Object) error                                                                              //--
	Draw(screen *ebiten.Image) error                                                                                  //--
	Redraw()                                                                                                          //--
	Update() error                                                                                                    //--
	Update_Unactive() error                                                                                           //
	Update_Any() (any, error)                                                                                         //
	GetState() uint8                                                                                                  //
	ToString() string                                                                                                 //
	IsInit() bool                                                                                                     //
	GetID() string                                                                                                    //
	GetType() string                                                                                                  //
	IsCursorInBounds() bool                                                                                           //
	IsCursorInBounds_MousePort(Mouse_Pos_X, Mouse_Pos_Y, mode int) bool                                               //
	GetPosition_Int() (int, int)                                                                                      //
	GetNumber_Children() int                                                                                          //
	AddChild(UI_Object) error
	GetChild(index int) UI_Object                                                                                     //
	GetParent() UI_Object                                                                                             //
	HasParent() bool                                                                                                  //
	// getType() string //might want to change this output to like an int or something using a golang equivelant of an enum;
}
*/

/*
UI_Button_mk0
*/
type UI_Button struct {
	Backend *UI_Backend
	// GSettings  *settings.GameSettings
	Position coords.CoordInts
	// Parentpos  *coords.CoordInts //this changes the position of the thing depending on where it's parent is;
	Dimensions coords.CoordInts
	BtnImg     *ebiten.Image
	// ActionPtr                      *func() //<---big test if this works
	Btn_Type                       uint8
	State                          uint8
	Colors                         []color.Color
	IsActive, IsVisible, IsToggled bool
	id, Label                      string
	ToRedraw                       bool
	Parent                         UI_Object
	isInit                         bool
}

/**/
func (ui_button *UI_Button) Init_00(bckend *UI_Backend, labl string, pos, dimen coords.CoordInts, bType uint8, parent UI_Object) {
	ui_button.Backend = bckend
	// ui_button.Parentpos = parent
	ui_button.Position = pos
	ui_button.Dimensions = dimen
	ui_button.Btn_Type = bType
	// ui_button.ActionPtr = action
	ui_button.State = 0
	ui_button.BtnImg = ebiten.NewImage(dimen.X, dimen.Y)
	ui_button.IsActive = true
	ui_button.IsVisible = true
	ui_button.IsToggled = false
	ui_button.Label = labl
	ui_button.ToRedraw = true
	ui_button.Redraw()

	ui_button.isInit = true
	// ui_button.Colors=
}

/**/
func (ui_button *UI_Button) Init(label []string, bckend *UI_Backend, style *UI_Object_Style, pos, dimen coords.CoordInts) error {
	ui_button.Backend = bckend
	ui_button.Position = pos
	ui_button.Dimensions = dimen
	ui_button.Btn_Type = 0
	// ui_button.ActionPtr = action
	ui_button.State = 0
	ui_button.BtnImg = ebiten.NewImage(dimen.X, dimen.Y)
	ui_button.IsActive = true
	ui_button.IsVisible = true
	ui_button.IsToggled = false
	ui_button.id = label[0]
	ui_button.Label = label[1]
	ui_button.ToRedraw = true
	ui_button.Redraw()
	ui_button.isInit = true

	// ui_button.Colors=
	return nil
}

/**/
func (ui_button *UI_Button) Init_Parents(parent UI_Object) error {
	ui_button.Parent = parent
	parent.AddChild(ui_button)
	//log.Printf("BUTTON ADDING PARENT %t \n", ui_button.HasParent())
	ui_button.Redraw()
	ui_button.Parent.Redraw()
	return nil
}

/**/
func (ui_button UI_Button) Redraw() {
	// if  {
	// 	 //color.RGBA{125, 125, 125, 255}
	// }
	ui_button.BtnImg.Fill(ui_button.Backend.Style.BorderColor)
	borderThick := ui_button.Backend.Style.BorderThickness
	if ui_button.Btn_Type == 10 && ui_button.IsToggled {
		vector.DrawFilledRect(ui_button.BtnImg, borderThick, borderThick, float32(ui_button.Dimensions.X)-(borderThick*2), float32(ui_button.Dimensions.Y)-(borderThick*2), ui_button.Backend.Style.ButtonColor1[ui_button.State], true)
	} else {
		vector.DrawFilledRect(ui_button.BtnImg, borderThick, borderThick, float32(ui_button.Dimensions.X)-(borderThick*2), float32(ui_button.Dimensions.Y)-(borderThick*2), ui_button.Backend.Style.ButtonColor0[ui_button.State], true)
	}
	//------text
	scaler := 1.0
	tops := &text.DrawOptions{}

	tops.GeoM.Translate(float64(ui_button.Dimensions.X/2)*scaler, float64(ui_button.Dimensions.Y/2)*scaler)
	tops.GeoM.Scale(1/scaler, 1/scaler)
	tops.ColorScale.ScaleWithColor(color.White)
	tops.LineSpacing = float64(10) * scaler
	tops.PrimaryAlign = text.AlignCenter
	tops.SecondaryAlign = text.AlignCenter
	temp := ui_button.Label
	//temp += fmt.Sprintf("\n%d", ui_button.State)
	text.Draw(ui_button.BtnImg, temp, ui_button.Backend.Btn_Text_Reg, tops)
	// if ui_button.Parent != nil {
	// 	ui_button.Parent.Redraw()
	// }
}

/**/
func (ui_button UI_Button) Draw(screen *ebiten.Image) error {

	scaler := 1.0
	ops := ebiten.DrawImageOptions{}
	ops.GeoM.Reset()
	ops.GeoM.Translate(float64(ui_button.Position.X)*scaler, float64(ui_button.Position.Y)*scaler)
	screen.DrawImage(ui_button.BtnImg, &ops)
	return nil
}

/**/
func (ui_button *UI_Button) Update() error {
	xx, yy := ebiten.CursorPosition()
	_, _, err := ui_button.Update_Ret_State_Redraw_Status_Mport(xx, yy, 0)

	return err
}

// /*This is going to be removed when I revise UI_OBJECT's and similar things*/
// func (ui_button *UI_Button) Update_Any() (any, error) {

// 	if ui_button.IsCursorInBounds() {
// 		if ui_button.Btn_Type == 10 {
// 			if inpututil.IsMouseButtonJustPressed(ebiten.MouseButton0) {
// 				ui_button.IsToggled = !ui_button.IsToggled

// 				ui_button.State = 2
// 				ui_button.Backend.PlaySound(3)
// 			} else {
// 				ui_button.State = 1
// 			}
// 			ui_button.ToRedraw = true
// 			//ui_button.Backend.PlaySound(1)
// 		} else {
// 			if inpututil.IsMouseButtonJustPressed(ebiten.MouseButton0) {
// 				ui_button.State = 2
// 				ui_button.Backend.PlaySound(1)
// 				ui_button.Redraw()
// 				return true, nil
// 			} else {
// 				//ui_button.Backend.PlaySound(1)
// 				if ui_button.State != 1 {
// 					// if ui_button.State != 2 {
// 					// 	ui_button.Backend.PlaySound(1)
// 					// }
// 					ui_button.ToRedraw = true
// 					ui_button.State = 1
// 				}
// 			}

// 		}
// 	} else {
// 		if ui_button.State > 0 {
// 			ui_button.State = 0
// 			ui_button.ToRedraw = true

// 		}
// 	}

// 	if ui_button.ToRedraw {
// 		ui_button.Redraw()
// 		// log.Printf("REDRAW\n")
// 		ui_button.ToRedraw = false
// 	}

// 	return ui_button.IsToggled, nil
// }

/**/
func (ui_button *UI_Button) Update_Unactive() error {
	if ui_button.State > 0 {
		ui_button.State = 0
		ui_button.ToRedraw = true

	}
	return nil
}

/**/
func (ui_button *UI_Button) Update_Ret_State_Redraw_Status() (uint8, bool, error) {
	xx, yy := ebiten.CursorPosition()
	return ui_button.Update_Ret_State_Redraw_Status_Mport(xx, yy, 0)
}

/*
This is a nice compromise in many respects;
what I'm attempting to do here is ensure that the buttons will be redrawn properly and are not going to need to have additional measures placed on them;
*/
func (ui_button *UI_Button) Update_Ret_State_Redraw_Status_Mport(Mouse_Pos_X, Mouse_Pos_Y, mode int) (uint8, bool, error) {
	to_redraw := false

	if ui_button.IsCursorInBounds_MousePort(Mouse_Pos_X, Mouse_Pos_Y, mode) {
		if ui_button.Btn_Type == 10 {
			if inpututil.IsMouseButtonJustPressed(ebiten.MouseButton0) {
				ui_button.IsToggled = !ui_button.IsToggled

				ui_button.State = 2
				ui_button.Backend.PlaySound(3)
			} else {
				ui_button.State = 1
			}
			ui_button.ToRedraw = true
			to_redraw = true
			//ui_button.Backend.PlaySound(1)
		} else if ui_button.Btn_Type == 20 {
			if ebiten.IsMouseButtonPressed(ebiten.MouseButton0) {
				ui_button.State = 2
				// ui_button.Backend.PlaySound(1)
				ui_button.Redraw()
				return ui_button.GetState(), true, nil
			} else {
				//ui_button.Backend.PlaySound(1)
				if ui_button.State != 1 {
					// if ui_button.State != 2 {
					// 	ui_button.Backend.PlaySound(1)
					// }
					ui_button.ToRedraw = true
					ui_button.State = 1
					to_redraw = true

				}
			}

		} else {
			if inpututil.IsMouseButtonJustPressed(ebiten.MouseButton0) {
				ui_button.State = 2
				ui_button.Backend.PlaySound(1)
				ui_button.Redraw()
				return ui_button.GetState(), true, nil
			} else {
				//ui_button.Backend.PlaySound(1)
				if ui_button.State != 1 {
					// if ui_button.State != 2 {
					// 	ui_button.Backend.PlaySound(1)
					// }
					ui_button.ToRedraw = true
					ui_button.State = 1
					to_redraw = true

				}
			}

		}
	} else {
		if ui_button.State > 0 {
			ui_button.State = 0
			ui_button.ToRedraw = true
			to_redraw = true
		}
	}

	if ui_button.ToRedraw {
		ui_button.Redraw()
		// log.Printf("REDRAW\n")
		ui_button.ToRedraw = false
	}
	return ui_button.GetState(), to_redraw, nil
}

/**/
func (ui_button *UI_Button) IsCursorInBounds() bool {
	if ui_button.IsVisible && ui_button.IsActive {
		Mouse_Pos_X, Mouse_Pos_Y := ebiten.CursorPosition()

		return ui_button.IsCursorInBounds_MousePort(Mouse_Pos_X, Mouse_Pos_Y, 0)
		// x0, y0 := ui_button.Position.X, ui_button.Position.Y
		// x1, y1 := ui_button.Position.X+ui_button.Dimensions.X, ui_button.Position.Y+ui_button.Dimensions.Y
		// if ui_button.Parent != nil {
		// 	x2, y2 := ui_button.Parent.GetPosition_Int()
		// 	x0 += x2
		// 	y0 += y2
		// 	x1 += x2
		// 	y1 += y2
		// }
		// return (Mouse_Pos_X > x0 && Mouse_Pos_X < x1) && (Mouse_Pos_Y > y0 && Mouse_Pos_Y < y1)
	}
	return false
	// if()
}

/**/
func (ui_button *UI_Button) IsCursorInBounds_MousePort(Mouse_Pos_X, Mouse_Pos_Y, mode int) bool {
	if ui_button.IsVisible && ui_button.IsActive {
		// xx, yy := ebiten.CursorPosition()
		x0, y0 := ui_button.Position.X, ui_button.Position.Y
		x1, y1 := ui_button.Position.X+ui_button.Dimensions.X, ui_button.Position.Y+ui_button.Dimensions.Y
		if ui_button.Parent != nil {
			x2, y2 := ui_button.Parent.GetPosition_Int()
			x3, y3 := ui_button.Parent.Get_Internal_Position_Int()
			if mode == 10 || mode == 0 { //
				// log.Printf("P I P: %d,%d %d,%d\n", x2, y2, x3, y3)
				x2 += x3
				y2 += y3
				// if !ui_button.Parent.IsCursorInBounds_MousePort(Mouse_Pos_X, Mouse_Pos_Y, mode) {
				// 	return false
				// } else {
				// 	// if pp := ui_button.Parent.GetParent(); pp != nil {
				// 	// 	if !pp.IsCursorInBounds_MousePort(Mouse_Pos_X, Mouse_Pos_Y, 0) {
				// 	// 		return false
				// 	// 	}
				// 	// }
				// }
			}
			x0 += x2
			y0 += y2
			x1 += x2
			y1 += y2
		}
		return (Mouse_Pos_X > x0 && Mouse_Pos_X < x1) && (Mouse_Pos_Y > y0 && Mouse_Pos_Y < y1)
	}
	return false
	// if()
}

func (ui_button *UI_Button) Detoggle() {
	// log.Printf("DETOGGLE %t\n", ui_button.IsToggled)
	if ui_button.IsToggled {
		ui_button.IsToggled = false
		ui_button.State = 0
		ui_button.Redraw()
		if ui_button.Parent != nil {
			// log.Printf("%t------DEACTIVATE\n", ui_button.Parent != nil)
			ui_button.Parent.Redraw()
		}
	}
}

/**/
func (ui_button *UI_Button) Close() {

}

/**/
func (ui_button *UI_Button) Open() {

}

/**/
func (ui_button *UI_Button) GetState() uint8 {
	if ui_button.Btn_Type == 10 {
		if ui_button.IsToggled {
			return 2
		} else {
			return ui_button.State
		}
	} else {
		return ui_button.State
	}
}

/**/
func (ui_button *UI_Button) Get_Internal_Position_Int() (x_pos int, y_pos int) {
	x_pos, y_pos = ui_button.GetPosition_Int()
	return x_pos, y_pos
}

/**/
func (ui_button *UI_Button) Get_Internal_Dimensions_Int() (x_pos int, y_pos int) {
	x_pos, y_pos = ui_button.GetPosition_Int()
	return x_pos, y_pos
}

/**/
func (ui_button *UI_Button) SetPosition_Int(x_pos, y_pos int) {
	ui_button.Position = coords.CoordInts{X: x_pos, Y: y_pos}
}

/**/
func (ui_button *UI_Button) GetDimensions_Int() (x_pos, y_pos int) {
	return ui_button.Dimensions.X, ui_button.Dimensions.Y
} //
/**/
func (ui_button *UI_Button) SetDimensions_Int(x_pos, y_pos int) {
	ui_button.Dimensions = coords.CoordInts{X: x_pos, Y: y_pos}
}

/**/
func (ui_button *UI_Button) ToString() string {
	strng := fmt.Sprintf("BUTTON: %s \n", ui_button.id)
	return strng
}

/**/
func (ui_button *UI_Button) IsInit() bool { return false } //

/**/
func (ui_button *UI_Button) GetID() string { return ui_button.id } //
/**/
func (ui_button *UI_Button) GetType() string { return "UI_Button" } //
/**/
func (ui_button *UI_Button) GetPosition_Int() (int, int) {
	return ui_button.Position.X, ui_button.Position.Y
} //
/**/
func (ui_button *UI_Button) GetNumber_Children() int { return 0 } //
/**/
func (ui_button *UI_Button) GetChild(index int) UI_Object { return nil } //
/**/
func (ui_button *UI_Button) AddChild(child UI_Object) error {
	return fmt.Errorf("ERROR NOT POSSIBLE")
}

/**/
func (ui_button *UI_Button) RemoveChild(index int) error {
	return fmt.Errorf("ERROR NOT POSSIBLE")
}

/**/
func (ui_button *UI_Button) GetParent() UI_Object { return nil } //
/**/
func (ui_button *UI_Button) HasParent() bool { return ui_button.Parent != nil }
