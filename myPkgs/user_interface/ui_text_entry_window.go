package user_interface

import (
	"fmt"
	"image/color"
	"log"
	"math"

	"github.com/KelleyTyler/GridTileEbit04_12/myPkgs/basic_geometry/coords"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type UI_Text_Entry_Window_00 struct {
	init                   bool
	Position, BasePosition coords.CoordInts
	Dimensions             coords.CoordInts
	Dimensions_Alt         coords.CoordInts
	IsVisible, IsActive    bool
	State                  uint8 //should be a const
	Status                 uint8
	Backend                *UI_Backend
	obj_id, PanelName      string
	Style                  *UI_Object_Style
	Image                  *ebiten.Image
	ImageUpdate            bool
	Parent                 UI_Object   //not all these things will have this option
	Children               []UI_Object //not all derivatives will have this option

	Window_Label UI_Label
	CloseButton  UI_Button
	// IsMovable bool
	IsMoving bool
	oldMouse coords.CoordInts

	Button_Submit UI_Button

	Textfield UI_TextEntryField

	errorMsgString string
	// tObject   UI_Object
}

func (ui_tew_00 *UI_Text_Entry_Window_00) Init(idLabels []string, backend *UI_Backend, style *UI_Object_Style, Position, Dimensions coords.CoordInts) error {

	ui_tew_00.obj_id = idLabels[0]
	ui_tew_00.PanelName = idLabels[1]

	ui_tew_00.Dimensions = Dimensions
	ui_tew_00.Position = Position
	ui_tew_00.BasePosition = Position
	ui_tew_00.Backend = backend
	if style != nil {
		ui_tew_00.Style = style
	} else {
		ui_tew_00.Style = &ui_tew_00.Backend.Style
	}
	ui_tew_00.State = 0

	//-------Setting up Image
	ui_tew_00.Image = ebiten.NewImage(Dimensions.X, Dimensions.Y)
	ui_tew_00.Redraw()
	ui_tew_00.ImageUpdate = true
	//------Finishing Up
	if !ui_tew_00.init {
		ui_tew_00.init = true
	}
	ui_tew_00.Window_Label.Init([]string{"window_label00", idLabels[1]}, backend, nil, coords.CoordInts{X: 0, Y: 0}, coords.CoordInts{X: Dimensions.X, Y: 32})
	ui_tew_00.Window_Label.Init_Parents(ui_tew_00)
	ui_tew_00.Window_Label.TextAlignMode = 10
	ui_tew_00.Window_Label.Redraw()
	ui_tew_00.Redraw()

	ui_tew_00.CloseButton.Init([]string{"window_close_button", "X"}, backend, nil, coords.CoordInts{X: Dimensions.X - 32, Y: 0}, coords.CoordInts{X: 32, Y: 32})
	ui_tew_00.CloseButton.Init_Parents(ui_tew_00)
	ui_tew_00.CloseButton.Redraw()
	ui_tew_00.Redraw()
	//---------------------------------- This Needs To be 'dealt with' somehow; not sure 'how' but 'somehow'
	// ui_tew_00.tObject = ui_tew_00.textfield
	ui_tew_00.Textfield.Init([]string{"window_textfield", ""}, backend, nil, coords.CoordInts{X: 9, Y: Dimensions.Y - 76}, coords.CoordInts{X: Dimensions.X - 18, Y: ui_tew_00.Window_Label.Dimensions.Y * 1})
	ui_tew_00.Textfield.Init_Parents(ui_tew_00)
	ui_tew_00.Textfield.Redraw()
	ui_tew_00.Redraw()

	ui_tew_00.Button_Submit.Init([]string{"window_submit_button", "Submit"}, backend, nil, coords.CoordInts{X: (Dimensions.X / 2) - 32, Y: Dimensions.Y - 40}, coords.CoordInts{X: 64, Y: 32})
	ui_tew_00.Button_Submit.Init_Parents(ui_tew_00)
	ui_tew_00.Button_Submit.Redraw()
	ui_tew_00.Redraw()

	ui_tew_00.State = 0
	ui_tew_00.oldMouse = coords.CoordInts{X: 0, Y: 0}
	// ui_tew_00.errorMsgString = "ERoor"
	return nil
}

func (ui_tew_00 *UI_Text_Entry_Window_00) InitVersion() {

}

/*
btn.Parent = parent
parent.AddChild(btn)
//log.Printf("BUTTON ADDING PARENT %t \n", btn.HasParent())
btn.Redraw()
btn.Parent.Redraw()
return nil
*/
func (ui_tew_00 *UI_Text_Entry_Window_00) Init_Parents(parent UI_Object) error {
	ui_tew_00.Parent = parent
	ui_tew_00.Parent.AddChild(ui_tew_00)
	ui_tew_00.Redraw()
	ui_tew_00.Parent.Redraw()
	return nil
}

/*
 */
func (ui_tew_00 *UI_Text_Entry_Window_00) Draw(screen *ebiten.Image) error {
	if ui_tew_00.IsVisible {
		ops := ebiten.DrawImageOptions{}
		scale := 1.0
		ops.GeoM.Reset()
		ops.GeoM.Translate(float64(ui_tew_00.Position.X)*scale, float64(ui_tew_00.Position.Y)*scale)
		screen.DrawImage(ui_tew_00.Image, &ops)
	}
	return nil
}

/**/
func (ui_tew_00 *UI_Text_Entry_Window_00) Print_Error_Message(error_message string) {
	ui_tew_00.errorMsgString = error_message
	ui_tew_00.Redraw()
	ui_tew_00.CloseButton.DeToggle()
	ui_tew_00.Button_Submit.DeToggle()
	ui_tew_00.Textfield.Clear()
	ui_tew_00.State = 0
}

/**/
func (ui_tew_00 *UI_Text_Entry_Window_00) Redraw() {
	ui_tew_00.Image.Fill(ui_tew_00.Style.BorderColor)
	lineThick := ui_tew_00.Style.BorderThickness
	vector.DrawFilledRect(ui_tew_00.Image, lineThick, lineThick, float32(ui_tew_00.Dimensions.X)-lineThick*2, float32(ui_tew_00.Dimensions.Y)-lineThick*2, ui_tew_00.Style.PanelColor, true) //
	if len(ui_tew_00.Children) > 0 {
		for i := 0; i < len(ui_tew_00.Children); i++ {
			err := ui_tew_00.Children[i].Draw(ui_tew_00.Image)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
	scaler := 1.0
	tops := &text.DrawOptions{}

	tops.GeoM.Translate(float64(5)*scaler, (float64(ui_tew_00.Dimensions.Y/2)-32)*scaler) //ui_tew_00.Dimensions.X/2
	tops.GeoM.Scale(1/scaler, 1/scaler)
	tops.ColorScale.ScaleWithColor(color.RGBA{255, 0, 0, 255})
	tops.LineSpacing = float64(10) * scaler
	// tops.PrimaryAlign = text.AlignCenter
	// tops.SecondaryAlign = text.AlignCenter
	temp := ui_tew_00.errorMsgString
	//temp += fmt.Sprintf("\n%d", btn.State)
	text.Draw(ui_tew_00.Image, temp, ui_tew_00.Backend.Btn_Text_Reg, tops)
}

/**/
func (ui_tew_00 *UI_Text_Entry_Window_00) Update() error {
	_, _, err := ui_tew_00.Update_Ret_State_Redraw_Status()
	return err
}

/**/
func (ui_tew_00 *UI_Text_Entry_Window_00) Update_Ret_State_Redraw_Status() (retVal uint8, toRedraw bool, err error) {
	Mouse_Pos_X, Mouse_Pos_Y := ebiten.CursorPosition()
	err = nil
	retVal, toRedraw, err = ui_tew_00.Update_Ret_State_Redraw_Status_Mport(Mouse_Pos_X, Mouse_Pos_Y, 0)
	return retVal, toRedraw, err

}

/**/
func (ui_tew_00 *UI_Text_Entry_Window_00) Update_Ret_State_Redraw_Status_Mport(Mouse_Pos_X, Mouse_Pos_Y, mode int) (retVal uint8, toRedraw bool, err error) {
	if ui_tew_00.IsActive {

		toRedraw = true
		// xx, yy := ebiten.CursorPosition()
		ui_tew_00.Textfield.IsActive = true
		// log.Printf("BEEP %3d %3d\n", xx, yy)
		ui_tew_00.MouseMove(Mouse_Pos_X, Mouse_Pos_Y)
		if ui_tew_00.IsCursorInBounds_Label(Mouse_Pos_X, Mouse_Pos_Y) && inpututil.IsMouseButtonJustPressed(ebiten.MouseButton0) {
			ui_tew_00.IsMoving = true
			ui_tew_00.oldMouse = coords.CoordInts{X: Mouse_Pos_X, Y: Mouse_Pos_Y} //Mouse_Pos_X, Mouse_Pos_Y
		}
		if inpututil.IsMouseButtonJustReleased(ebiten.MouseButton0) {
			ui_tew_00.IsMoving = false
		}
		// else { ui_tew_00.IsCursorInBounds_Label
		// 	log.Printf("%3d %3d\n", xx, yy)
		// } ui_tew_00.CloseButton.GetState() == 2
		if ui_tew_00.IsCursorInBounds_MousePort(Mouse_Pos_X, Mouse_Pos_Y, mode) {
			if state, redraw, _ := ui_tew_00.CloseButton.Update_Ret_State_Redraw_Status_Mport(Mouse_Pos_X, Mouse_Pos_Y, mode); state == 2 {
				ui_tew_00.State = 90
				// ui_tew_00.IsActive = false
				// ui_tew_00.IsVisible = false
				ui_tew_00.IsMoving = false
				ui_tew_00.oldMouse = coords.CoordInts{X: 0, Y: 0}
				if redraw {
					ui_tew_00.Redraw()
				}
			}

			if state, redraw, _ := ui_tew_00.Button_Submit.Update_Ret_State_Redraw_Status_Mport(Mouse_Pos_X, Mouse_Pos_Y, mode); state == 2 {
				ui_tew_00.State = 80
				// ui_tew_00.IsActive = false
				// ui_tew_00.IsVisible = false
				ui_tew_00.IsMoving = false
				ui_tew_00.oldMouse = coords.CoordInts{X: 0, Y: 0}
				if redraw {
					ui_tew_00.Redraw()
				}
			}
			if len(ui_tew_00.Children) > 0 {
				for i := 0; i < len(ui_tew_00.Children); i++ {
					err := ui_tew_00.Children[i].Update()
					if err != nil {
						log.Fatal(err)
					}
				}
			}
			// if ui_tew_00.IsMovable {
			// 	if ui_tew_00.IsCursorInBounds() {

			// 	}
			// }
			ui_tew_00.Redraw()
		}
	}
	return 0, false, nil
}

/**/
func (ui_tew_00 *UI_Text_Entry_Window_00) Update_Unactive() error {
	if len(ui_tew_00.Children) > 0 {
		for i := 0; i < len(ui_tew_00.Children); i++ {
			err := ui_tew_00.Children[i].Update_Unactive()
			if err != nil {
				log.Fatal(err)
			}
		}
	}
	return nil
}

/*
This returns the state of the object
*/
func (ui_tew_00 *UI_Text_Entry_Window_00) GetState() (out_state uint8) {
	if ui_tew_00.IsActive && ui_tew_00.IsVisible {
		out_state = ui_tew_00.State
		ui_tew_00.State = 0
	} else {
		ui_tew_00.State = 90
	}
	return out_state
}

/**/
func (ui_tew_00 *UI_Text_Entry_Window_00) MouseMove(Mouse_Pos_X, Mouse_Pos_Y int) {
	if ui_tew_00.IsMoving {
		// x0, y0 := ebiten.CursorPosition()
		x0, y0 := Mouse_Pos_X, Mouse_Pos_Y
		x1 := ui_tew_00.oldMouse.X
		y1 := ui_tew_00.oldMouse.Y
		dx, dy := (x1 - x0), (y1 - y0)
		t0 := x1 == x0 && y1 == y0
		t1 := int(math.Abs(float64(dx))) < 2 && int(math.Abs(float64(dy))) < 2
		if t0 || t1 {
			ui_tew_00.oldMouse = coords.CoordInts{X: x0, Y: y0}
		} else {
			x2, y2 := (x1-x0)/1, (y1-y0)/1
			//log.Printf("----- %d %d \n", x2, y2)
			ui_tew_00.Position.Y -= y2
			ui_tew_00.Position.X -= x2
			ui_tew_00.oldMouse = coords.CoordInts{X: x0, Y: y0}
			// gb.BoardChanges = true
			// gb.BoardOverlayChanges = true
		}
	}
}

/*
this returns a basic to string message
*/
func (ui_tew_00 *UI_Text_Entry_Window_00) ToString() string {
	strngOut := fmt.Sprintf("UI_Object ui_tew_00tive:%s\n\tPositon %s\t", ui_tew_00.obj_id, ui_tew_00.Position.ToString())
	strngOut += fmt.Sprintf("\tDimensions: %s\n", ui_tew_00.Dimensions.ToString())
	return strngOut
}

/*
this returns a basic to string message
*/
func (ui_tew_00 *UI_Text_Entry_Window_00) Close() {
	ui_tew_00.IsActive = false
	ui_tew_00.IsVisible = false
	// ui_tew_00.IsMoving = false
	ui_tew_00.CloseButton.DeToggle()
	ui_tew_00.Button_Submit.DeToggle()
	ui_tew_00.Textfield.Clear()
	ui_tew_00.State = 0
}

/*
this returns a basic to string message
*/
func (ui_tew_00 *UI_Text_Entry_Window_00) Open() {
	ui_tew_00.IsActive = true
	ui_tew_00.IsVisible = true
	ui_tew_00.Position = ui_tew_00.BasePosition
	// ui_tew_00.IsMoving = true
	ui_tew_00.CloseButton.DeToggle()
	ui_tew_00.Button_Submit.DeToggle()
	ui_tew_00.Textfield.Clear()
	ui_tew_00.State = 0

}

/*
 */
func (ui_tew_00 *UI_Text_Entry_Window_00) IsCursorInBounds() bool {
	if ui_tew_00.IsActive && ui_tew_00.IsVisible {
		cX, cY := ebiten.CursorPosition()
		var x0, y0, x1, y1 int

		if ui_tew_00.Parent != nil {
			px, py := ui_tew_00.Parent.GetPosition_Int()
			x0 = ui_tew_00.Position.X + px
			y0 = ui_tew_00.Position.Y + py
			x1 = ui_tew_00.Position.X + ui_tew_00.Dimensions.X + px
			y1 = ui_tew_00.Position.Y + ui_tew_00.Dimensions.Y + py
			// x0 = ui_tew_00.Position.X + ui_tew_00.ParentPos.X
			// y0 = ui_tew_00.Position.Y + ui_tew_00.ParentPos.X
			// x1 = ui_tew_00.Position.X + ui_tew_00.ParentPos.X + ui_tew_00.Dimensions.X
			// y1 = ui_tew_00.Position.Y + ui_tew_00.ParentPos.Y + ui_tew_00.Dimensions.Y
		} else {
			x0 = ui_tew_00.Position.X
			y0 = ui_tew_00.Position.Y
			x1 = ui_tew_00.Position.X + ui_tew_00.Dimensions.X
			y1 = ui_tew_00.Position.Y + ui_tew_00.Dimensions.Y
		}
		return (cX > x0 && cX < x1) && (cY > y0 && cY < y1)
	}
	return false
}

/*
 */
func (ui_tew_00 *UI_Text_Entry_Window_00) IsCursorInBounds_Label(cX, cY int) bool {
	if ui_tew_00.IsActive && ui_tew_00.IsVisible {
		// cX, cY := ebiten.CursorPosition()
		var x0, y0, x1, y1 int

		if ui_tew_00.Parent != nil {
			px, py := ui_tew_00.Parent.GetPosition_Int()
			x0 = ui_tew_00.Position.X + px
			y0 = ui_tew_00.Position.Y + py
			x1 = ui_tew_00.Position.X + ui_tew_00.Dimensions.X + px - (2 + ui_tew_00.CloseButton.Dimensions.X)
			y1 = ui_tew_00.Position.Y + py + ui_tew_00.Window_Label.Dimensions.Y
			// x0 = ui_tew_00.Position.X + ui_tew_00.ParentPos.X
			// y0 = ui_tew_00.Position.Y + ui_tew_00.ParentPos.X
			// x1 = ui_tew_00.Position.X + ui_tew_00.ParentPos.X + ui_tew_00.Dimensions.X
			// y1 = ui_tew_00.Position.Y + ui_tew_00.ParentPos.Y + ui_tew_00.Dimensions.Y
		} else {
			x0 = ui_tew_00.Position.X
			y0 = ui_tew_00.Position.Y
			x1 = ui_tew_00.Position.X + ui_tew_00.Dimensions.X
			y1 = ui_tew_00.Position.Y + ui_tew_00.Window_Label.Dimensions.Y
		}
		return (cX > x0 && cX < x1) && (cY > y0 && cY < y1)
	}
	return false
}

/*
Idea here is I don't want to waste time with having to get the cursor position when it possibly hasn't changed enough to matter;
This might be also a terrible idea overall I cannot tell quite yet

enter 0 for it to default
*/
func (ui_tew_00 *UI_Text_Entry_Window_00) IsCursorInBounds_MousePort(Mouse_Pos_X, Mouse_Pos_Y, mode int) bool {
	if ui_tew_00.IsActive && ui_tew_00.IsVisible && mode == 0 {
		cX, cY := Mouse_Pos_X, Mouse_Pos_Y
		//mode stuff
		var x0, y0, x1, y1 int

		if ui_tew_00.Parent != nil {
			px, py := ui_tew_00.Parent.GetPosition_Int()
			x0 = ui_tew_00.Position.X + px
			y0 = ui_tew_00.Position.Y + py
			x1 = ui_tew_00.Position.X + ui_tew_00.Dimensions.X + px
			y1 = ui_tew_00.Position.Y + ui_tew_00.Dimensions.Y + py
		} else {
			x0 = ui_tew_00.Position.X
			y0 = ui_tew_00.Position.Y
			x1 = ui_tew_00.Position.X + ui_tew_00.Dimensions.X
			y1 = ui_tew_00.Position.Y + ui_tew_00.Dimensions.Y
		}
		return (cX > x0 && cX < x1) && (cY > y0 && cY < y1)
	}
	//mode stuff
	return false
}

func (ui_tew_00 *UI_Text_Entry_Window_00) GetPosition_Int() (int, int) {
	xx := ui_tew_00.Position.X
	yy := ui_tew_00.Position.Y
	if ui_tew_00.Parent != nil {
		px, py := ui_tew_00.Parent.GetPosition_Int()
		xx += px
		yy += py
	}
	return xx, yy
}

/**/
func (ui_tew_00 *UI_Text_Entry_Window_00) SetPosition_Int(x_point, y_point int) {
	ui_tew_00.Position = coords.CoordInts{X: x_point, Y: y_point}
}

/**/
func (ui_tew_00 *UI_Text_Entry_Window_00) GetDimensions_Int() (int, int) {
	return ui_tew_00.Dimensions.X, ui_tew_00.Dimensions.Y
}

/**/
func (ui_tew_00 *UI_Text_Entry_Window_00) SetDimensions_Int(x_point, y_point int) {
	ui_tew_00.Dimensions = coords.CoordInts{X: x_point, Y: y_point}
	//---Redraw The Thing;
}

/*
this confirms the object is initilaized
*/
func (ui_tew_00 *UI_Text_Entry_Window_00) IsInit() bool {
	return ui_tew_00.init
}

/*
this gets the object ID
*/
func (ui_tew_00 *UI_Text_Entry_Window_00) GetID() string {
	return ui_tew_00.obj_id
}

/*
This returns a string specifying the objects type
*/
func (ui_tew_00 *UI_Text_Entry_Window_00) GetType() string {
	return "UI_Object UI_Text_Entry_Window_00"
}

func (ui_tew_00 *UI_Text_Entry_Window_00) GetNumber_Children() int {
	return len(ui_tew_00.Children)
}

func (ui_tew_00 *UI_Text_Entry_Window_00) GetChild(index int) UI_Object {
	if len(ui_tew_00.Children) > index {
		return ui_tew_00.Children[index]
	} else {
		return nil
	}
}

func (ui_tew_00 *UI_Text_Entry_Window_00) AddChild(child UI_Object) error {
	ui_tew_00.Children = append(ui_tew_00.Children, child)
	return nil
}

func (ui_tew_00 *UI_Text_Entry_Window_00) RemoveChild(index int) error {
	// ui_tew_00.Children = append(ui_tew_00.Children, child)
	return nil
}

func (ui_tew_00 *UI_Text_Entry_Window_00) HasParent() bool {
	return ui_tew_00.Parent != nil
}

func (ui_tew_00 *UI_Text_Entry_Window_00) GetParent() UI_Object {
	return ui_tew_00.Parent
}
