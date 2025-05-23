package user_interface

import (
	"fmt"
	"image/color"
	"log"
	"math"

	"github.com/KelleyTyler/GridTileEbit04_12/myPkgs/basic_geometry/coords"
	settings "github.com/KelleyTyler/GridTileEbit04_12/myPkgs/settingsconfig"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type UI_Subwindow struct {
	Position                coords.CoordInts
	Dimensions              coords.CoordInts
	Close_Button            UI_Button
	Minimize_Button         UI_Button
	WindowLabel             UI_Label
	IsActive, IsVisible     bool
	IsResizable, IsMoveable bool
	PanelImage              *ebiten.Image
	Settings                *settings.GameSettings
	Backend                 *UI_Backend
	Style                   *UI_Object_Style
	WindowName              string

	//-------------------------Needs Content
}

func (sWin *UI_Subwindow) Init(backend *UI_Backend) {

}

func (sWin *UI_Subwindow) initImg() {
	// sWin.PanelImage
}

func (sWin *UI_Subwindow) Draw(screen *ebiten.Image) {

}

func (sWin *UI_Subwindow) Update() {

}

type UI_Window struct {
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

	// Button_Submit UI_Button

	// Textfield UI_TextEntryField
	// Scroller    UI_Scrollbar
	// Scroller_02 UI_Scrollbar

	Button_Thing_Zero UI_Button
	errorMsgString    string
	ScrollPane        UI_Scrollpane
	Prim              UI_Object_Primitive
	// tObject   UI_Object
}

func (ui_win *UI_Window) Init(idLabels []string, backend *UI_Backend, style *UI_Object_Style, Position, Dimensions coords.CoordInts) error {

	ui_win.obj_id = idLabels[0]
	ui_win.PanelName = idLabels[1]

	ui_win.Dimensions = Dimensions
	ui_win.Position = Position
	ui_win.BasePosition = Position
	ui_win.Backend = backend
	if style != nil {
		ui_win.Style = style
	} else {
		ui_win.Style = &ui_win.Backend.Style
	}
	ui_win.State = 0

	//-------Setting up Image
	ui_win.Image = ebiten.NewImage(Dimensions.X, Dimensions.Y)
	ui_win.Redraw()
	ui_win.ImageUpdate = true
	//------Finishing Up
	if !ui_win.init {
		ui_win.init = true
	}
	btnWidth := 16
	ui_win.Window_Label.Init([]string{"window_label00", idLabels[1]}, backend, nil, coords.CoordInts{X: 0, Y: 0}, coords.CoordInts{X: Dimensions.X, Y: btnWidth})
	ui_win.Window_Label.Init_Parents(ui_win)
	ui_win.Window_Label.TextAlignMode = 10
	ui_win.Window_Label.Redraw()
	ui_win.Redraw()

	ui_win.CloseButton.Init([]string{"window_close_button", "X"}, backend, nil, coords.CoordInts{X: Dimensions.X - btnWidth, Y: 0}, coords.CoordInts{X: btnWidth, Y: btnWidth})
	ui_win.CloseButton.Init_Parents(ui_win)
	ui_win.CloseButton.Redraw()
	ui_win.Redraw()

	// btnWidth_two := (btnWidth) - 2
	// btnWidth = 16
	//---------------------------------- This Needs To be 'dealt with' somehow; not sure 'how' but 'somehow'
	// ui_win.tObject = ui_win.textfield
	// ui_win.Textfield.Init([]string{"window_textfield", ""}, backend, nil, coords.CoordInts{X: 9, Y: Dimensions.Y - 76}, coords.CoordInts{X: Dimensions.X - 18, Y: ui_win.Window_Label.Dimensions.Y * 1})
	// ui_win.Textfield.Init_Parents(ui_win)
	// ui_win.Textfield.Redraw()
	// ui_win.Redraw()
	// ui_win.Scroller.Init([]string{"window_close_button", "X"}, backend, style, coords.CoordInts{X: 0, Y: 30}, coords.CoordInts{X: Dimensions.Y - 26, Y: 16})
	// ui_win.Scroller.Init([]string{"window_close_button", "X"}, backend, style, coords.CoordInts{X: Dimensions.X - btnWidth, Y: btnWidth_two}, coords.CoordInts{X: btnWidth, Y: Dimensions.Y - btnWidth_two - btnWidth})

	// ui_win.Scroller.SetVals(5, 1, -5, 10, 0)
	// ui_win.Scroller.Redraw()
	// ui_win.Scroller.Init_Parents(ui_win)
	// ui_win.Redraw()

	// ui_win.Scroller_02.Init([]string{"window_close_button", "X"}, backend, style, coords.CoordInts{X: 0, Y: Dimensions.Y - btnWidth}, coords.CoordInts{X: Dimensions.X - btnWidth, Y: btnWidth})
	// ui_win.Scroller_02.SetVals(5, 1, -5, 10, 0)

	// ui_win.Scroller_02.Redraw()
	// ui_win.Scroller_02.Init_Parents(ui_win)
	boarderThick := int(ui_win.Style.BorderThickness)
	ui_win.ScrollPane.Init([]string{"windowScrollPane", "Scroll pane"}, backend, style, coords.CoordInts{X: boarderThick, Y: btnWidth}, coords.CoordInts{X: Dimensions.X, Y: Dimensions.Y - btnWidth}) //Dimensions.X - btnWidth Dimensions.Y - btnWidth
	ui_win.ScrollPane.Init_Parents(ui_win)
	ui_win.Redraw()

	ui_win.Button_Thing_Zero.Init([]string{"window_submit_button", "Submit"}, backend, nil, coords.CoordInts{X: (Dimensions.X / 2) - 32, Y: Dimensions.Y - 40}, coords.CoordInts{X: 64, Y: 32})
	ui_win.Prim.Init([]string{"primitive_scroll_primitive", "PRIMITIVE"}, backend, style, coords.CoordInts{X: 0, Y: 0}, coords.CoordInts{X: Dimensions.X - 18, Y: ui_win.Dimensions.Y})
	ui_win.Prim.Init_Parents(&ui_win.ScrollPane)
	ui_win.Prim.IsActive = true
	ui_win.Prim.IsVisible = true

	ui_win.Button_Thing_Zero.Init_Parents(&ui_win.Prim)
	ui_win.Button_Thing_Zero.Redraw()
	ui_win.Button_Thing_Zero.IsActive = true
	ui_win.Button_Thing_Zero.IsVisible = true

	ui_win.Prim.Redraw()
	ui_win.ScrollPane.Redraw()
	// ui_win.Button_Submit.Init([]string{"window_submit_button", "Submit"}, backend, nil, coords.CoordInts{X: (Dimensions.X / 2) - 32, Y: Dimensions.Y - 40}, coords.CoordInts{X: 64, Y: 32})
	// ui_win.Button_Submit.Init_Parents(ui_win)
	// ui_win.Button_Submit.Redraw()
	// ui_win.Redraw()

	ui_win.State = 0
	ui_win.oldMouse = coords.CoordInts{X: 0, Y: 0}
	// ui_win.errorMsgString = "ERoor"
	return nil
}

func (ui_win *UI_Window) InitVersion() {

}

/*
btn.Parent = parent
parent.AddChild(btn)
//log.Printf("BUTTON ADDING PARENT %t \n", btn.HasParent())
btn.Redraw()
btn.Parent.Redraw()
return nil
*/
func (ui_win *UI_Window) Init_Parents(parent UI_Object) error {
	ui_win.Parent = parent
	ui_win.Parent.AddChild(ui_win)
	ui_win.Redraw()
	ui_win.Parent.Redraw()
	return nil
}

/*
 */
func (ui_win *UI_Window) Draw(screen *ebiten.Image) error {
	if ui_win.IsVisible {
		ops := ebiten.DrawImageOptions{}
		scale := 1.0
		ops.GeoM.Reset()
		ops.GeoM.Translate(float64(ui_win.Position.X)*scale, float64(ui_win.Position.Y)*scale)
		screen.DrawImage(ui_win.Image, &ops)
	}
	return nil
}

/**/
func (ui_win *UI_Window) Print_Error_Message(error_message string) {
	ui_win.errorMsgString = error_message
	ui_win.Redraw()
	ui_win.CloseButton.Detoggle()
	// ui_win.Button_Submit.DeToggle()
	// ui_win.Textfield.Clear()
	ui_win.State = 0
}

/**/
func (ui_win *UI_Window) Redraw() {
	ui_win.Image.Fill(ui_win.Style.BorderColor)
	lineThick := ui_win.Style.BorderThickness
	vector.DrawFilledRect(ui_win.Image, lineThick, lineThick, float32(ui_win.Dimensions.X)-lineThick*2, float32(ui_win.Dimensions.Y)-lineThick*2, ui_win.Style.PanelColor, true) //
	if len(ui_win.Children) > 0 {
		for i := 0; i < len(ui_win.Children); i++ {
			err := ui_win.Children[i].Draw(ui_win.Image)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
	scaler := 1.0
	tops := &text.DrawOptions{}

	tops.GeoM.Translate(float64(5)*scaler, (float64(ui_win.Dimensions.Y/2)-32)*scaler) //ui_win.Dimensions.X/2
	tops.GeoM.Scale(1/scaler, 1/scaler)
	tops.ColorScale.ScaleWithColor(color.RGBA{255, 0, 0, 255})
	tops.LineSpacing = float64(10) * scaler
	// tops.PrimaryAlign = text.AlignCenter
	// tops.SecondaryAlign = text.AlignCenter
	temp := ui_win.errorMsgString
	//temp += fmt.Sprintf("\n%d", btn.State)
	text.Draw(ui_win.Image, temp, ui_win.Backend.Btn_Text_Reg, tops)
}

/**/
func (ui_win *UI_Window) Update() error {
	_, _, err := ui_win.Update_Ret_State_Redraw_Status()
	return err
}

/**/
func (ui_win *UI_Window) Update_Ret_State_Redraw_Status() (retVal uint8, toRedraw bool, err error) {
	Mouse_Pos_X, Mouse_Pos_Y := ebiten.CursorPosition()
	err = nil
	retVal, toRedraw, err = ui_win.Update_Ret_State_Redraw_Status_Mport(Mouse_Pos_X, Mouse_Pos_Y, 0)
	// fmt.Printf("STATUS: %d %d\n", retVal, ui_win.State)
	return retVal, toRedraw, err

}

/**/
func (ui_win *UI_Window) Update_Ret_State_Redraw_Status_Mport(Mouse_Pos_X, Mouse_Pos_Y, mode int) (retVal uint8, toRedraw bool, err error) {
	if ui_win.IsActive {

		toRedraw = true
		// xx, yy := ebiten.CursorPosition()
		// ui_win.Textfield.IsActive = true
		// log.Printf("BEEP %3d %3d\n", xx, yy)
		ui_win.MouseMove(Mouse_Pos_X, Mouse_Pos_Y)
		if ui_win.IsCursorInBounds_Label(Mouse_Pos_X, Mouse_Pos_Y) && inpututil.IsMouseButtonJustPressed(ebiten.MouseButton0) {
			ui_win.IsMoving = true
			ui_win.oldMouse = coords.CoordInts{X: Mouse_Pos_X, Y: Mouse_Pos_Y} //Mouse_Pos_X, Mouse_Pos_Y
		}
		if inpututil.IsMouseButtonJustReleased(ebiten.MouseButton0) {
			ui_win.IsMoving = false
		}
		// else { ui_win.IsCursorInBounds_Label
		// 	log.Printf("%3d %3d\n", xx, yy)
		// } ui_win.CloseButton.GetState() == 2
		if ui_win.IsCursorInBounds_MousePort(Mouse_Pos_X, Mouse_Pos_Y, mode) {
			if state, redraw, _ := ui_win.CloseButton.Update_Ret_State_Redraw_Status_Mport(Mouse_Pos_X, Mouse_Pos_Y, 5); state == 2 {
				ui_win.State = 90
				// ui_win.IsActive = false
				// ui_win.IsVisible = false
				ui_win.IsMoving = false
				ui_win.oldMouse = coords.CoordInts{X: 0, Y: 0}
				if redraw {
					ui_win.Redraw()
				}
			}
			if state, _, _ := ui_win.Button_Thing_Zero.Update_Ret_State_Redraw_Status_Mport(Mouse_Pos_X, Mouse_Pos_Y, 10); state == 2 {
				x0, y0 := ui_win.Button_Thing_Zero.GetPosition_Int()
				x1, y1 := ui_win.Button_Thing_Zero.GetDimensions_Int()

				x2, y2 := ui_win.Button_Thing_Zero.Parent.GetPosition_Int()
				x3, y3 := ui_win.Button_Thing_Zero.Parent.Get_Internal_Position_Int()
				// x4, y4 := ui_win.Button_Thing_Zero.Parent.GetDimensions_Int()
				log.Printf("OUT OUT OUT:\n")
				x5, y5 := ui_win.Button_Thing_Zero.Parent.GetParent().Get_Internal_Position_Int()
				log.Printf("%14s:%3d,%3d\t%14s:%3d,%3d\n", "Position", x0, y0, "Dimensions", x1, y1)
				log.Printf("%14s:%3d,%3d\t%14s:%3d,%3d\t%14s:%3d,%3d\n", "Par_Pos", x2, y2, "par_Int_Pos", x3, y3, "par_par_int_pos", x5, y5)
				// log.Printf("%10s:%3d,%3d\n")
				// log.Printf("%10s:%3d,%3d\n")
			}
			// if state, redraw, _ := ui_win.Button_Submit.Update_Ret_State_Redraw_Status_Mport(Mouse_Pos_X, Mouse_Pos_Y, mode); state == 2 {
			// 	ui_win.State = 80
			// 	// ui_win.IsActive = false
			// 	// ui_win.IsVisible = false
			// 	ui_win.IsMoving = false
			// 	ui_win.oldMouse = coords.CoordInts{X: 0, Y: 0}
			// 	if redraw {
			// 		ui_win.Redraw()
			// 	}
			// }
			if len(ui_win.Children) > 0 {
				for i := 0; i < len(ui_win.Children); i++ {
					err := ui_win.Children[i].Update()
					if err != nil {
						log.Fatal(err)
					}
				}
			}
			// if ui_win.IsMovable {
			// 	if ui_win.IsCursorInBounds() {

			// 	}
			// }
			ui_win.Redraw()
		}
	}
	return ui_win.State, false, nil
}

/**/
func (ui_win *UI_Window) Update_Unactive() error {
	if len(ui_win.Children) > 0 {
		for i := 0; i < len(ui_win.Children); i++ {
			err := ui_win.Children[i].Update_Unactive()
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
func (ui_win *UI_Window) GetState() (out_state uint8) {
	if ui_win.IsActive && ui_win.IsVisible {
		out_state = ui_win.State
		ui_win.State = 0
	} else {
		ui_win.State = 90
	}
	return out_state
}

/**/
func (ui_win *UI_Window) MouseMove(Mouse_Pos_X, Mouse_Pos_Y int) {
	if ui_win.IsMoving {
		// x0, y0 := ebiten.CursorPosition()
		x0, y0 := Mouse_Pos_X, Mouse_Pos_Y
		x1 := ui_win.oldMouse.X
		y1 := ui_win.oldMouse.Y
		dx, dy := (x1 - x0), (y1 - y0)
		t0 := x1 == x0 && y1 == y0
		t1 := int(math.Abs(float64(dx))) < 2 && int(math.Abs(float64(dy))) < 2
		if t0 || t1 {
			ui_win.oldMouse = coords.CoordInts{X: x0, Y: y0}
		} else {
			x2, y2 := (x1-x0)/1, (y1-y0)/1
			//log.Printf("----- %d %d \n", x2, y2)
			ui_win.Position.Y -= y2
			ui_win.Position.X -= x2
			ui_win.oldMouse = coords.CoordInts{X: x0, Y: y0}
			// gb.BoardChanges = true
			// gb.BoardOverlayChanges = true
		}
	}
}

/*
this returns a basic to string message
*/
func (ui_win *UI_Window) ToString() string {
	strngOut := fmt.Sprintf("UI_Object ui_wintive:%s\n\tPositon %s\t", ui_win.obj_id, ui_win.Position.ToString())
	strngOut += fmt.Sprintf("\tDimensions: %s\n", ui_win.Dimensions.ToString())
	return strngOut
}

/*
this returns a basic to string message
*/
func (ui_win *UI_Window) Close() {
	ui_win.IsActive = false
	ui_win.IsVisible = false
	// ui_win.IsMoving = false
	ui_win.CloseButton.Detoggle()
	// ui_win.Button_Submit.DeToggle()
	// ui_win.Textfield.Clear()
	ui_win.State = 0
}

/*
this returns a basic to string message
*/
func (ui_win *UI_Window) Open() {
	ui_win.IsActive = true
	ui_win.IsVisible = true
	ui_win.Position = ui_win.BasePosition
	// ui_win.IsMoving = true
	ui_win.CloseButton.Detoggle()
	// ui_win.Button_Submit.DeToggle()
	// ui_win.Textfield.Clear()

	ui_win.State = 0

}
func (ui_win *UI_Window) Detoggle() {
	ui_win.Update_Unactive()
	ui_win.State = 0
}

/*
 */
func (ui_win *UI_Window) IsCursorInBounds() bool {
	if ui_win.IsActive && ui_win.IsVisible {
		cX, cY := ebiten.CursorPosition()
		var x0, y0, x1, y1 int

		if ui_win.Parent != nil {
			px, py := ui_win.Parent.GetPosition_Int()
			x0 = ui_win.Position.X + px
			y0 = ui_win.Position.Y + py
			x1 = ui_win.Position.X + ui_win.Dimensions.X + px
			y1 = ui_win.Position.Y + ui_win.Dimensions.Y + py
			// x0 = ui_win.Position.X + ui_win.ParentPos.X
			// y0 = ui_win.Position.Y + ui_win.ParentPos.X
			// x1 = ui_win.Position.X + ui_win.ParentPos.X + ui_win.Dimensions.X
			// y1 = ui_win.Position.Y + ui_win.ParentPos.Y + ui_win.Dimensions.Y
		} else {
			x0 = ui_win.Position.X
			y0 = ui_win.Position.Y
			x1 = ui_win.Position.X + ui_win.Dimensions.X
			y1 = ui_win.Position.Y + ui_win.Dimensions.Y
		}
		return (cX > x0 && cX < x1) && (cY > y0 && cY < y1)
	}
	return false
}

/*
 */
func (ui_win *UI_Window) IsCursorInBounds_Label(cX, cY int) bool {
	if ui_win.IsActive && ui_win.IsVisible {
		// cX, cY := ebiten.CursorPosition()
		var x0, y0, x1, y1 int

		if ui_win.Parent != nil {
			px, py := ui_win.Parent.GetPosition_Int()
			x0 = ui_win.Position.X + px
			y0 = ui_win.Position.Y + py
			x1 = ui_win.Position.X + ui_win.Dimensions.X + px - (2 + ui_win.CloseButton.Dimensions.X)
			y1 = ui_win.Position.Y + py + ui_win.Window_Label.Dimensions.Y
			// x0 = ui_win.Position.X + ui_win.ParentPos.X
			// y0 = ui_win.Position.Y + ui_win.ParentPos.X
			// x1 = ui_win.Position.X + ui_win.ParentPos.X + ui_win.Dimensions.X
			// y1 = ui_win.Position.Y + ui_win.ParentPos.Y + ui_win.Dimensions.Y
		} else {
			x0 = ui_win.Position.X
			y0 = ui_win.Position.Y
			x1 = ui_win.Position.X + ui_win.Dimensions.X
			y1 = ui_win.Position.Y + ui_win.Window_Label.Dimensions.Y
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
func (ui_win *UI_Window) IsCursorInBounds_MousePort(Mouse_Pos_X, Mouse_Pos_Y, mode int) bool {
	if ui_win.IsActive && ui_win.IsVisible {
		cX, cY := Mouse_Pos_X, Mouse_Pos_Y
		//mode stuff
		var x0, y0, x1, y1, x3, y3 int

		if ui_win.Parent != nil {
			px, py := ui_win.Parent.GetPosition_Int()
			x0 = ui_win.Position.X + px
			y0 = ui_win.Position.Y + py
			x1 = ui_win.Position.X + ui_win.Dimensions.X + px
			y1 = ui_win.Position.Y + ui_win.Dimensions.Y + py
			if mode == 10 {
				x3, y3 = ui_win.Parent.Get_Internal_Position_Int()
				x0 += x3
				x1 += x3
				y0 += y3
				y1 += y3
				if !ui_win.Parent.IsCursorInBounds_MousePort(Mouse_Pos_X, Mouse_Pos_Y, 0) {
					// log.Printf("OUT AT Subwindow\n")

					return false
				}
			}
		} else {
			x0 = ui_win.Position.X
			y0 = ui_win.Position.Y
			x1 = ui_win.Position.X + ui_win.Dimensions.X
			y1 = ui_win.Position.Y + ui_win.Dimensions.Y
		}
		// temp :=
		// if !temp {
		// 	log.Printf("OUT AT Subwindow 6 %3d %3d 1:%3d %3d\t2:%3d %3d3:%3d %3d\n", Mouse_Pos_X, Mouse_Pos_Y, x0, y0, x1, y1, x3, y3)
		// }
		return (cX > x0 && cX < x1) && (cY > y0 && cY < y1)
	}
	//mode stuff
	return false
}

/**/
func (ui_win *UI_Window) Get_Internal_Position_Int() (x_pos int, y_pos int) {
	x_pos, y_pos = ui_win.GetPosition_Int()
	return x_pos, y_pos
}

/**/
func (ui_win *UI_Window) GetPosition_Int() (int, int) {
	xx := ui_win.Position.X
	yy := ui_win.Position.Y
	if ui_win.Parent != nil {
		px, py := ui_win.Parent.GetPosition_Int()
		xx += px
		yy += py
	}
	return xx, yy
}

/**/
func (ui_win *UI_Window) SetPosition_Int(x_point, y_point int) {
	ui_win.Position = coords.CoordInts{X: x_point, Y: y_point}
}

/**/
func (ui_win *UI_Window) GetDimensions_Int() (int, int) {
	return ui_win.Dimensions.X, ui_win.Dimensions.Y
}

/**/
func (ui_win *UI_Window) SetDimensions_Int(x_point, y_point int) {
	ui_win.Dimensions = coords.CoordInts{X: x_point, Y: y_point}
	//---Redraw The S
}

/*
this confirms the object is initilaized
*/
func (ui_win *UI_Window) IsInit() bool {
	return ui_win.init
}

/*
this gets the object ID
*/
func (ui_win *UI_Window) GetID() string {
	return ui_win.obj_id
}

/*
This returns a string specifying the objects type
*/
func (ui_win *UI_Window) GetType() string {
	return "UI_Object ui_window"
}

func (ui_win *UI_Window) GetNumber_Children() int {
	return len(ui_win.Children)
}

func (ui_win *UI_Window) GetChild(index int) UI_Object {
	if len(ui_win.Children) > index {
		return ui_win.Children[index]
	} else {
		return nil
	}
}

func (ui_win *UI_Window) AddChild(child UI_Object) error {
	ui_win.Children = append(ui_win.Children, child)
	return nil
}

func (ui_win *UI_Window) RemoveChild(index int) error {
	// ui_win.Children = append(ui_win.Children, child)
	return nil
}

func (ui_win *UI_Window) HasParent() bool {
	return ui_win.Parent != nil
}

func (ui_win *UI_Window) GetParent() UI_Object {
	return ui_win.Parent
}
