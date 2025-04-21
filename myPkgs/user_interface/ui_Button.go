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

func (btn *UI_Button) Init_00(bckend *UI_Backend, labl string, pos, dimen coords.CoordInts, bType uint8, parent UI_Object) {
	btn.Backend = bckend
	// btn.Parentpos = parent
	btn.Position = pos
	btn.Dimensions = dimen
	btn.Btn_Type = bType
	// btn.ActionPtr = action
	btn.State = 0
	btn.BtnImg = ebiten.NewImage(dimen.X, dimen.Y)
	btn.IsActive = true
	btn.IsVisible = true
	btn.IsToggled = false
	btn.Label = labl
	btn.ToRedraw = true
	btn.Redraw()

	btn.isInit = true
	// btn.Colors=
}
func (btn *UI_Button) Init(label []string, bckend *UI_Backend, style *UI_Object_Style, pos, dimen coords.CoordInts) error {
	btn.Backend = bckend
	btn.Position = pos
	btn.Dimensions = dimen
	btn.Btn_Type = 0
	// btn.ActionPtr = action
	btn.State = 0
	btn.BtnImg = ebiten.NewImage(dimen.X, dimen.Y)
	btn.IsActive = true
	btn.IsVisible = true
	btn.IsToggled = false
	btn.id = label[0]
	btn.Label = label[1]
	btn.ToRedraw = true
	btn.Redraw()
	btn.isInit = true

	// btn.Colors=
	return nil
}
func (btn *UI_Button) Init_Parents(parent UI_Object) error {
	btn.Parent = parent
	parent.AddChild(btn)
	//fmt.Printf("BUTTON ADDING PARENT %t \n", btn.HasParent())
	btn.Redraw()
	btn.Parent.Redraw()
	return nil
}

func (btn UI_Button) Redraw() {
	// if  {
	// 	 //color.RGBA{125, 125, 125, 255}
	// }
	btn.BtnImg.Fill(btn.Backend.Style.BorderColor)
	borderThick := btn.Backend.Style.BorderThickness
	if btn.Btn_Type == 10 && btn.IsToggled {
		vector.DrawFilledRect(btn.BtnImg, borderThick, borderThick, float32(btn.Dimensions.X)-(borderThick*2), float32(btn.Dimensions.Y)-(borderThick*2), btn.Backend.Style.ButtonColor1[btn.State], true)
	} else {
		vector.DrawFilledRect(btn.BtnImg, borderThick, borderThick, float32(btn.Dimensions.X)-(borderThick*2), float32(btn.Dimensions.Y)-(borderThick*2), btn.Backend.Style.ButtonColor0[btn.State], true)
	}
	//------text
	scaler := 1.0
	tops := &text.DrawOptions{}

	tops.GeoM.Translate(float64(btn.Dimensions.X/2)*scaler, float64(btn.Dimensions.Y/2)*scaler)
	tops.GeoM.Scale(1/scaler, 1/scaler)
	tops.ColorScale.ScaleWithColor(color.White)
	tops.LineSpacing = float64(10) * scaler
	tops.PrimaryAlign = text.AlignCenter
	tops.SecondaryAlign = text.AlignCenter
	temp := btn.Label
	//temp += fmt.Sprintf("\n%d", btn.State)
	text.Draw(btn.BtnImg, temp, btn.Backend.Btn_Text_Reg, tops)
	// if btn.Parent != nil {
	// 	btn.Parent.Redraw()
	// }
}

func (btn UI_Button) Draw(screen *ebiten.Image) error {

	scaler := 1.0
	ops := ebiten.DrawImageOptions{}
	ops.GeoM.Reset()
	ops.GeoM.Translate(float64(btn.Position.X)*scaler, float64(btn.Position.Y)*scaler)
	screen.DrawImage(btn.BtnImg, &ops)
	return nil
}
func (btn *UI_Button) Update() error {
	// fmt.Printf("btn tick\n")
	// xx, yy := ebiten.CursorPosition()
	// q := btn.Position.X
	// r := btn.Position.X
	// s := 0
	// t := 0
	// if btn.HasParent() {
	// 	s, t = btn.Parent.GetPosition_Int()
	// }
	// if ebiten.IsKeyPressed(ebiten.Key0) {
	// 	fmt.Printf("HAH %s %d %d %t %d %d \n", btn.Label, xx, yy, btn.HasParent(), q+s, r+t)
	// }
	if btn.IsCursorInBounds() {
		if btn.Btn_Type == 10 {
			if inpututil.IsMouseButtonJustPressed(ebiten.MouseButton0) {
				btn.IsToggled = !btn.IsToggled

				btn.State = 2
				btn.Backend.PlaySound(3)
			} else {
				btn.State = 1
			}
			btn.ToRedraw = true

		} else {
			if inpututil.IsMouseButtonJustPressed(ebiten.MouseButton0) {
				btn.State = 2
				btn.Backend.PlaySound(1)
				btn.Redraw()
			} else {
				if btn.State != 1 {
					// if btn.State != 2 {
					// 	btn.Backend.PlaySound(1)
					// }
					btn.ToRedraw = true
					btn.State = 1
				}
			}
		}
	} else {
		if btn.State > 0 {
			btn.State -= 1
		}
	}

	if btn.ToRedraw {
		btn.Redraw()
		// fmt.Printf("REDRAW\n")
		btn.ToRedraw = false
	}
	return nil
}
func (btn *UI_Button) Update_Any() (any, error) {

	if btn.IsCursorInBounds() {
		if btn.Btn_Type == 10 {
			if inpututil.IsMouseButtonJustPressed(ebiten.MouseButton0) {
				btn.IsToggled = !btn.IsToggled

				btn.State = 2
				btn.Backend.PlaySound(3)
			} else {
				btn.State = 1
			}
			btn.ToRedraw = true
			//btn.Backend.PlaySound(1)
		} else {
			if inpututil.IsMouseButtonJustPressed(ebiten.MouseButton0) {
				btn.State = 2
				btn.Backend.PlaySound(1)
				btn.Redraw()
				return true, nil
			} else {
				//btn.Backend.PlaySound(1)
				if btn.State != 1 {
					// if btn.State != 2 {
					// 	btn.Backend.PlaySound(1)
					// }
					btn.ToRedraw = true
					btn.State = 1
				}
			}

		}
	} else {
		if btn.State > 0 {
			btn.State = 0
			btn.ToRedraw = true

		}
	}

	if btn.ToRedraw {
		btn.Redraw()
		// fmt.Printf("REDRAW\n")
		btn.ToRedraw = false
	}

	return btn.IsToggled, nil
}

func (btn *UI_Button) Update_Unactive() error {
	if btn.State > 0 {
		btn.State = 0
		btn.ToRedraw = true

	}
	return nil
}
func (btn *UI_Button) Update_Ret_State_Redraw_Status() (uint8, bool, error) {
	to_redraw := false
	if btn.IsCursorInBounds() {
		if btn.Btn_Type == 10 {
			if inpututil.IsMouseButtonJustPressed(ebiten.MouseButton0) {
				btn.IsToggled = !btn.IsToggled

				btn.State = 2
				btn.Backend.PlaySound(3)
			} else {
				btn.State = 1
			}
			btn.ToRedraw = true
			to_redraw = true
			//btn.Backend.PlaySound(1)
		} else {
			if inpututil.IsMouseButtonJustPressed(ebiten.MouseButton0) {
				btn.State = 2
				btn.Backend.PlaySound(1)
				btn.Redraw()
				return btn.GetState(), true, nil
			} else {
				//btn.Backend.PlaySound(1)
				if btn.State != 1 {
					// if btn.State != 2 {
					// 	btn.Backend.PlaySound(1)
					// }
					btn.ToRedraw = true
					to_redraw = true
					btn.State = 1
				}
			}

		}
	} else {
		if btn.State > 0 {
			btn.State = 0
			btn.ToRedraw = true
			to_redraw = true
		}
	}

	if btn.ToRedraw {
		btn.Redraw()
		// fmt.Printf("REDRAW\n")
		btn.ToRedraw = false
	}
	// return btn.IsToggled, nil
	return btn.GetState(), to_redraw, nil
}

/*
This is a nice compromise in many respects;
*/
func (btn *UI_Button) Update_Ret_State_Redraw_Status_Mport(Mouse_Pos_X, Mouse_Pos_Y, mode int) (uint8, bool, error) {
	to_redraw := false

	if btn.IsCursorInBounds_MousePort(Mouse_Pos_X, Mouse_Pos_Y, mode) {
		if btn.Btn_Type == 10 {
			if inpututil.IsMouseButtonJustPressed(ebiten.MouseButton0) {
				btn.IsToggled = !btn.IsToggled

				btn.State = 2
				btn.Backend.PlaySound(3)
			} else {
				btn.State = 1
			}
			btn.ToRedraw = true
			to_redraw = true
			//btn.Backend.PlaySound(1)
		} else {
			if inpututil.IsMouseButtonJustPressed(ebiten.MouseButton0) {
				btn.State = 2
				btn.Backend.PlaySound(1)
				btn.Redraw()
				return btn.GetState(), true, nil
			} else {
				//btn.Backend.PlaySound(1)
				if btn.State != 1 {
					// if btn.State != 2 {
					// 	btn.Backend.PlaySound(1)
					// }
					btn.ToRedraw = true
					btn.State = 1
					to_redraw = true

				}
			}

		}
	} else {
		if btn.State > 0 {
			btn.State = 0
			btn.ToRedraw = true
			to_redraw = true
		}
	}

	if btn.ToRedraw {
		btn.Redraw()
		// fmt.Printf("REDRAW\n")
		btn.ToRedraw = false
	}
	return btn.GetState(), to_redraw, nil
}
func (btn *UI_Button) IsCursorInBounds() bool {
	if btn.IsVisible && btn.IsActive {
		Mouse_Pos_X, Mouse_Pos_Y := ebiten.CursorPosition()
		x0, y0 := btn.Position.X, btn.Position.Y
		x1, y1 := btn.Position.X+btn.Dimensions.X, btn.Position.Y+btn.Dimensions.Y
		if btn.Parent != nil {
			x2, y2 := btn.Parent.GetPosition_Int()
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
func (btn *UI_Button) IsCursorInBounds_MousePort(Mouse_Pos_X, Mouse_Pos_Y, mode int) bool {
	if btn.IsVisible && btn.IsActive {
		// xx, yy := ebiten.CursorPosition()
		x0, y0 := btn.Position.X, btn.Position.Y
		x1, y1 := btn.Position.X+btn.Dimensions.X, btn.Position.Y+btn.Dimensions.Y
		if btn.Parent != nil {
			x2, y2 := btn.Parent.GetPosition_Int()
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

func (btn *UI_Button) DeToggle() {
	// fmt.Printf("DETOGGLE %t\n", btn.IsToggled)
	if btn.IsToggled {
		btn.IsToggled = false
		btn.State = 0
		btn.Redraw()
		if btn.Parent != nil {
			// fmt.Printf("%t------DEACTIVATE\n", btn.Parent != nil)
			btn.Parent.Redraw()
		}
	}
}
func (btn *UI_Button) GetState() uint8 {
	if btn.Btn_Type == 10 {
		if btn.IsToggled {
			return 2
		} else {
			return btn.State
		}
	} else {
		return btn.State
	}
} //
func (btn *UI_Button) ToString() string {
	strng := fmt.Sprintf("BUTTON: %s \n", btn.id)
	return strng
}
func (btn *UI_Button) IsInit() bool                 { return false }                          //
func (btn *UI_Button) GetID() string                { return btn.id }                         //
func (btn *UI_Button) GetType() string              { return "UI_Button" }                    //
func (btn *UI_Button) GetPosition_Int() (int, int)  { return btn.Position.X, btn.Position.Y } //
func (btn *UI_Button) GetNumber_Children() int      { return 0 }                              //
func (btn *UI_Button) GetChild(index int) UI_Object { return nil }                            //
func (btn *UI_Button) AddChild(child UI_Object) error {
	return fmt.Errorf("ERROR NOT POSSIBLE")
}
func (btn *UI_Button) RemoveChild(index int) error {
	return fmt.Errorf("ERROR NOT POSSIBLE")
}
func (btn *UI_Button) GetParent() UI_Object { return nil } //
func (btn *UI_Button) HasParent() bool      { return btn.Parent != nil }
