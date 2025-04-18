package user_interface

import (
	"fmt"
	"log"

	"github.com/KelleyTyler/GridTileEbit04_12/myPkgs/basic_geometry/coords"
	settings "github.com/KelleyTyler/GridTileEbit04_12/myPkgs/settingsconfig"
	"github.com/hajimehoshi/ebiten/v2"
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
	init                bool
	Position            coords.CoordInts
	Dimensions          coords.CoordInts
	Dimensions_Alt      coords.CoordInts
	IsVisible, IsActive bool
	State               uint8 //should be a const
	Backend             *UI_Backend
	obj_id, PanelName   string
	Style               *UI_Object_Style
	Image               *ebiten.Image
	ImageUpdate         bool
	Parent              UI_Object   //not all these things will have this option
	Children            []UI_Object //not all derivatives will have this option

	Window_Label UI_Label
	// IsMovable bool
	// IsMoving  bool
	// oldMouse  coords.CoordInts
}

func (ui_win *UI_Window) Init(idLabels []string, backend *UI_Backend, style *UI_Object_Style, Position, Dimensions coords.CoordInts) error {

	ui_win.obj_id = idLabels[0]
	ui_win.PanelName = idLabels[1]

	ui_win.Dimensions = Dimensions
	ui_win.Position = Position
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
	return nil
}

func (ui_win *UI_Window) Init_Parents(parent UI_Object) error {
	ui_win.Parent = parent
	return nil
}

func (ui_win *UI_Window) Draw(screen *ebiten.Image) error {
	ops := ebiten.DrawImageOptions{}
	scale := 1.0
	ops.GeoM.Reset()
	ops.GeoM.Translate(float64(ui_win.Position.X)*scale, float64(ui_win.Position.Y)*scale)
	screen.DrawImage(ui_win.Image, &ops)
	return nil
}
func (ui_win *UI_Window) Redraw() {
	ui_win.Image.Fill(ui_win.Style.BorderColor)
	lineThick := ui_win.Style.BorderThickness
	vector.DrawFilledRect(ui_win.Image, lineThick, lineThick, float32(ui_win.Dimensions.X)-lineThick, float32(ui_win.Dimensions.Y)-lineThick, ui_win.Style.PanelColor, true)
	if len(ui_win.Children) > 0 {
		for i := 0; i < len(ui_win.Children); i++ {
			err := ui_win.Children[i].Draw(ui_win.Image)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}
func (ui_win *UI_Window) Update() error {
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
	return nil
}

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
This will return false; Use Only Sparingly!
*/
func (ui_win *UI_Window) Update_Any() (any, error) {
	return false, nil
}

/*
This returns the state of the object
*/
func (ui_win *UI_Window) GetState() uint8 {
	return ui_win.State
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
Idea here is I don't want to waste time with having to get the cursor position when it possibly hasn't changed enough to matter;
This might be also a terrible idea overall I cannot tell quite yet

enter 0 for it to default
*/
func (ui_win *UI_Window) IsCursorInBounds_MousePort(Mouse_Pos_X, Mouse_Pos_Y, mode int) bool {
	if ui_win.IsActive && ui_win.IsVisible && mode == 0 {
		cX, cY := Mouse_Pos_X, Mouse_Pos_Y
		//mode stuff
		var x0, y0, x1, y1 int

		if ui_win.Parent != nil {
			px, py := ui_win.Parent.GetPosition_Int()
			x0 = ui_win.Position.X + px
			y0 = ui_win.Position.Y + py
			x1 = ui_win.Position.X + ui_win.Dimensions.X + px
			y1 = ui_win.Position.Y + ui_win.Dimensions.Y + py
		} else {
			x0 = ui_win.Position.X
			y0 = ui_win.Position.Y
			x1 = ui_win.Position.X + ui_win.Dimensions.X
			y1 = ui_win.Position.Y + ui_win.Dimensions.Y
		}
		return (cX > x0 && cX < x1) && (cY > y0 && cY < y1)
	}
	//mode stuff
	return false
}

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
