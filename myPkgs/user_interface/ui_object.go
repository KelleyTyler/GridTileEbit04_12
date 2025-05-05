package user_interface

import (
	"fmt"
	"log"

	"github.com/KelleyTyler/GridTileEbit04_12/myPkgs/basic_geometry/coords"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

/*
coming up with a UI interface that is a generic object;

this is not where we store the Objects; rather it's where POINTERS to objects are stored;
*/
type UI_Object interface {
	// Init0() //initialize void;
	Init(idLabels []string, backend *UI_Backend, style *UI_Object_Style, Position, Dimensions coords.CoordInts) error //--
	Init_Parents(Parent UI_Object) error                                                                              //--
	Draw(screen *ebiten.Image) error                                                                                  //--
	Redraw()                                                                                                          //--
	Update() error                                                                                                    //--
	Update_Unactive() error                                                                                           //

	//Update_Any() (any, error) //
	Update_Ret_State_Redraw_Status() (uint8, bool, error)
	Update_Ret_State_Redraw_Status_Mport(Mouse_Pos_X, Mouse_Pos_Y, mode int) (uint8, bool, error)

	GetState() uint8                                                    //
	ToString() string                                                   //
	IsInit() bool                                                       //
	GetID() string                                                      //
	GetType() string                                                    //
	IsCursorInBounds() bool                                             //
	IsCursorInBounds_MousePort(Mouse_Pos_X, Mouse_Pos_Y, mode int) bool //
	Get_Internal_Position_Int() (x_pos int, y_pos int)
	Get_Internal_Dimensions_Int() (x_pos int, y_pos int)

	GetPosition_Int() (int, int)
	SetPosition_Int(int, int)      //
	GetDimensions_Int() (int, int) //
	/**/
	Close()
	/**/
	Open()
	/**/
	Detoggle()
	SetDimensions_Int(int, int)
	GetNumber_Children() int //

	GetChild(index int) UI_Object
	AddChild(child UI_Object) error //
	RemoveChild(index int) error
	GetParent() UI_Object //
	HasParent() bool      //
	// getType() string //might want to change this output to like an int or something using a golang equivelant of an enum;
}

/*
----- So what is the point of this and where do I go from here???
----- There's going to need to be an implementation of ebitendebugGui as a framework or something...
----	More importantly there's stuff that I'm finding a bit annoying with my own implementations here;
---- 	Might be worth it to just ditch this and go with EbitenGui as that seems to have at least somewhat of a start/good functionality;
----	Also might be worthwhile to go with DearImGui
--------------------------------------------------------
----	so some other ideas while I'm here;
----	split up UI_Object interface into 2-3 separate things;
----	"UI_Objects"
----	"UI_Widgets"//-->
	"UI_Containers"{

	}
*/

/*
Given the lack of inherentance involved in golang; this is just a guideline to follow;
*/
type UI_Object_Primitive struct {
	init                bool
	Position            coords.CoordInts
	Dimensions          coords.CoordInts
	IsVisible, IsActive bool
	State               uint8 //should be a const
	Backend             *UI_Backend
	obj_id              string
	Style               *UI_Object_Style
	Image               *ebiten.Image
	ImageUpdate         bool
	Parent              UI_Object   //not all these things will have this option
	Children            []UI_Object //not all derivatives will have this option

	// IsMovable bool
	// IsMoving  bool
	// oldMouse  coords.CoordInts
}

/*
UI_Object_Primitive.Init
*/
func (ui_panel *UI_Object_Primitive) Init(idLabels []string, backend *UI_Backend, style *UI_Object_Style, Position, Dimensions coords.CoordInts) error {
	ui_panel.obj_id = "ui_panelitive_UI_OBJ"
	ui_panel.Dimensions = Dimensions
	ui_panel.Position = Position
	ui_panel.Backend = backend
	if style != nil {
		ui_panel.Style = style
	} else {
		ui_panel.Style = &ui_panel.Backend.Style
	}
	ui_panel.State = 0

	//-------Setting up Image
	ui_panel.Image = ebiten.NewImage(Dimensions.X, Dimensions.Y)
	ui_panel.Redraw()
	ui_panel.ImageUpdate = true
	//------Finishing Up
	if !ui_panel.init {
		ui_panel.init = true
	}
	return nil
}

func (ui_panel *UI_Object_Primitive) Init_Parents(parent UI_Object) error {
	ui_panel.Parent = parent
	ui_panel.Parent.AddChild(ui_panel)
	ui_panel.Redraw()
	ui_panel.Parent.Redraw()
	return nil
}

func (ui_panel *UI_Object_Primitive) Draw(screen *ebiten.Image) error {
	ui_panel.Redraw()
	ops := ebiten.DrawImageOptions{}
	scale := 1.0
	ops.GeoM.Reset()
	ops.GeoM.Translate(float64(ui_panel.Position.X)*scale, float64(ui_panel.Position.Y)*scale)
	screen.DrawImage(ui_panel.Image, &ops)
	return nil
}
func (ui_panel *UI_Object_Primitive) Redraw() {
	ui_panel.Image.Fill(ui_panel.Style.BorderColor)
	lineThick := ui_panel.Style.BorderThickness
	vector.DrawFilledRect(ui_panel.Image, lineThick, lineThick, float32(ui_panel.Dimensions.X)-(2*lineThick), float32(ui_panel.Dimensions.Y)-(2*lineThick), ui_panel.Style.PanelColor, true)
	if len(ui_panel.Children) > 0 {
		for i := 0; i < len(ui_panel.Children); i++ {
			err := ui_panel.Children[i].Draw(ui_panel.Image)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}
func (ui_panel *UI_Object_Primitive) Update() error {
	xx, yy := ebiten.CursorPosition()
	ui_panel.Update_Ret_State_Redraw_Status_Mport(xx, yy, 0)
	// if len(ui_panel.Children) > 0 {
	// 	for i := 0; i < len(ui_panel.Children); i++ {
	// 		_, to_redraw, err := ui_panel.Children[i].Update_Ret_State_Redraw_Status()
	// 		if err != nil {
	// 			log.Fatal(err)
	// 		}
	// 		if to_redraw {
	// 			ui_panel.Children[i].Draw(ui_panel.Image)
	// 		}
	// 	}
	// }
	// if ui_panel.IsMovable {
	// 	if ui_panel.IsCursorInBounds() {

	// 	}
	// }
	// ui_panel.Redraw()
	return nil
}

func (ui_panel *UI_Object_Primitive) Update_Unactive() error {
	if len(ui_panel.Children) > 0 {
		for i := 0; i < len(ui_panel.Children); i++ {
			err := ui_panel.Children[i].Update_Unactive()
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
func (ui_panel *UI_Object_Primitive) Update_Any() (any, error) {
	return false, nil
}
func (ui_panel *UI_Object_Primitive) Update_Ret_State_Redraw_Status() (uint8, bool, error) {
	return ui_panel.State, false, nil
}
func (ui_panel *UI_Object_Primitive) Update_Ret_State_Redraw_Status_Mport(Mouse_Pos_X, Mouse_Pos_Y, mode int) (uint8, bool, error) {

	if mode == 0 {
		if len(ui_panel.Children) > 0 {
			for i := 0; i < len(ui_panel.Children); i++ {
				_, to_redraw, err := ui_panel.Children[i].Update_Ret_State_Redraw_Status_Mport(Mouse_Pos_X, Mouse_Pos_Y, 0)
				if err != nil {
					log.Fatal(err)
				}
				if to_redraw {
					ui_panel.Children[i].Draw(ui_panel.Image)
				}
			}
		}
		if ui_panel.IsCursorInBounds_MousePort(Mouse_Pos_X, Mouse_Pos_Y, mode) {
			if inpututil.IsMouseButtonJustPressed(ebiten.MouseButton0) {
				log.Printf("%d %d %t\n", Mouse_Pos_X, Mouse_Pos_Y, ui_panel.HasParent())

			}
		}
	}
	return ui_panel.State, false, nil
}

/*
This returns the state of the object
*/
func (ui_panel *UI_Object_Primitive) GetState() uint8 {
	return ui_panel.State
}

/*
this returns a basic to string message
*/
func (ui_panel *UI_Object_Primitive) ToString() string {
	strngOut := fmt.Sprintf("UI_Object Primtive:%s\n\tPositon %s\t", ui_panel.obj_id, ui_panel.Position.ToString())
	strngOut += fmt.Sprintf("\tDimensions: %s\n", ui_panel.Dimensions.ToString())
	return strngOut
}

/*
this confirms the object is initilaized
*/
func (ui_panel *UI_Object_Primitive) IsInit() bool {
	return ui_panel.init
}

/*
this gets the object ID
*/
func (ui_panel *UI_Object_Primitive) GetID() string {
	return ui_panel.obj_id
}

/*
This returns a string specifying the objects type
*/
func (ui_panel *UI_Object_Primitive) GetType() string {
	return "UI_Object Primitive"
}

/*
 */
func (ui_panel *UI_Object_Primitive) IsCursorInBounds() bool {
	if ui_panel.IsActive && ui_panel.IsVisible {
		cX, cY := ebiten.CursorPosition()
		var x0, y0, x1, y1 int

		if ui_panel.Parent != nil {
			px, py := ui_panel.Parent.GetPosition_Int()
			x0 = ui_panel.Position.X + px
			y0 = ui_panel.Position.Y + py
			x1 = ui_panel.Position.X + ui_panel.Dimensions.X + px
			y1 = ui_panel.Position.Y + ui_panel.Dimensions.Y + py
			// x0 = ui_panel.Position.X + ui_panel.ParentPos.X
			// y0 = ui_panel.Position.Y + ui_panel.ParentPos.X
			// x1 = ui_panel.Position.X + ui_panel.ParentPos.X + ui_panel.Dimensions.X
			// y1 = ui_panel.Position.Y + ui_panel.ParentPos.Y + ui_panel.Dimensions.Y
		} else {
			x0 = ui_panel.Position.X
			y0 = ui_panel.Position.Y
			x1 = ui_panel.Position.X + ui_panel.Dimensions.X
			y1 = ui_panel.Position.Y + ui_panel.Dimensions.Y
		}
		return (cX > x0 && cX < x1) && (cY > y0 && cY < y1)
	}
	return false
}

/**/
func (ui_panel *UI_Object_Primitive) Close() {}

/**/
func (ui_panel *UI_Object_Primitive) Open() {}

/**/
func (ui_panel *UI_Object_Primitive) Detoggle() {}

/*
Idea here is I don't want to waste time with having to get the cursor position when it possibly hasn't changed enough to matter;
This might be also a terrible idea overall I cannot tell quite yet

enter 0 for it to default
*/
func (ui_panel *UI_Object_Primitive) IsCursorInBounds_MousePort(Mouse_Pos_X, Mouse_Pos_Y, mode int) bool {
	if ui_panel.IsActive && ui_panel.IsVisible {
		cX, cY := Mouse_Pos_X, Mouse_Pos_Y
		//mode stuff
		var x0, y0, x1, y1 int

		if ui_panel.Parent != nil {
			px, py := ui_panel.Parent.GetPosition_Int()
			x0 = ui_panel.Position.X + px
			y0 = ui_panel.Position.Y + py
			x1 = ui_panel.Position.X + ui_panel.Dimensions.X + px
			y1 = ui_panel.Position.Y + ui_panel.Dimensions.Y + py
			if mode == 10 {
				x3, y3 := ui_panel.Parent.Get_Internal_Position_Int()
				x0 += x3
				x1 += x3
				y0 += y3
				y1 += y3
				if !ui_panel.Parent.IsCursorInBounds_MousePort(Mouse_Pos_X, Mouse_Pos_Y, 10) {

					return false
				}
			}
		} else {
			x0 = ui_panel.Position.X
			y0 = ui_panel.Position.Y
			x1 = ui_panel.Position.X + ui_panel.Image.Bounds().Dx()
			y1 = ui_panel.Position.Y + ui_panel.Image.Bounds().Dy()
		}
		return (cX > x0 && cX < x1) && (cY > y0 && cY < y1)
	}
	//mode stuff
	return false
}

/**/
func (ui_panel *UI_Object_Primitive) Get_Internal_Position_Int() (x_pos int, y_pos int) {
	if ui_panel.Parent != nil {
		x_pos, y_pos = ui_panel.Parent.Get_Internal_Position_Int()
	}
	return x_pos, y_pos
}

/**/
func (ui_panel *UI_Object_Primitive) Get_Internal_Dimensions_Int() (x_pos int, y_pos int) {
	x_pos, y_pos = ui_panel.GetPosition_Int()
	return x_pos, y_pos
}

/**/
func (ui_panel *UI_Object_Primitive) GetPosition_Int() (int, int) {
	xx := ui_panel.Position.X
	yy := ui_panel.Position.Y
	if ui_panel.Parent != nil {
		px, py := ui_panel.Parent.GetPosition_Int()
		xx += px
		yy += py
	}
	return xx, yy
}

/**/
func (ui_panel *UI_Object_Primitive) SetPosition_Int(pos_X, pos_Y int) {
	ui_panel.Position = coords.CoordInts{X: pos_X, Y: pos_Y}
}

/**/
func (ui_panel *UI_Object_Primitive) GetDimensions_Int() (int, int) {
	return ui_panel.Dimensions.X, ui_panel.Dimensions.Y
} //
/**/
func (ui_panel *UI_Object_Primitive) SetDimensions_Int(pos_X, pos_Y int) {
	ui_panel.Dimensions = coords.CoordInts{X: pos_X, Y: pos_Y}

}

func (ui_panel *UI_Object_Primitive) GetNumber_Children() int {
	return len(ui_panel.Children)
}
func (ui_panel *UI_Object_Primitive) GetChild(index int) UI_Object {
	if len(ui_panel.Children) > index {
		return ui_panel.Children[index]
	} else {
		return nil
	}
}
func (ui_panel *UI_Object_Primitive) AddChild(child UI_Object) error {
	ui_panel.Children = append(ui_panel.Children, child)
	return nil
}
func (ui_panel *UI_Object_Primitive) RemoveChild(index int) error {
	// ui_panel.Children = append(ui_panel.Children, child)
	return nil
}
func (ui_panel *UI_Object_Primitive) HasParent() bool {
	return ui_panel.Parent != nil
}
func (ui_panel *UI_Object_Primitive) GetParent() UI_Object {
	return ui_panel.Parent
}
