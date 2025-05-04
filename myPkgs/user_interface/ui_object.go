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
func (prim *UI_Object_Primitive) Init(idLabels []string, backend *UI_Backend, style *UI_Object_Style, Position, Dimensions coords.CoordInts) error {
	prim.obj_id = "primitive_UI_OBJ"
	prim.Dimensions = Dimensions
	prim.Position = Position
	prim.Backend = backend
	if style != nil {
		prim.Style = style
	} else {
		prim.Style = &prim.Backend.Style
	}
	prim.State = 0

	//-------Setting up Image
	prim.Image = ebiten.NewImage(Dimensions.X, Dimensions.Y)
	prim.Redraw()
	prim.ImageUpdate = true
	//------Finishing Up
	if !prim.init {
		prim.init = true
	}
	return nil
}

func (prim *UI_Object_Primitive) Init_Parents(parent UI_Object) error {
	prim.Parent = parent
	prim.Parent.AddChild(prim)
	prim.Redraw()
	prim.Parent.Redraw()
	return nil
}

func (prim *UI_Object_Primitive) Draw(screen *ebiten.Image) error {
	prim.Redraw()
	ops := ebiten.DrawImageOptions{}
	scale := 1.0
	ops.GeoM.Reset()
	ops.GeoM.Translate(float64(prim.Position.X)*scale, float64(prim.Position.Y)*scale)
	screen.DrawImage(prim.Image, &ops)
	return nil
}
func (prim *UI_Object_Primitive) Redraw() {
	prim.Image.Fill(prim.Style.BorderColor)
	lineThick := prim.Style.BorderThickness
	vector.DrawFilledRect(prim.Image, lineThick, lineThick, float32(prim.Dimensions.X)-(2*lineThick), float32(prim.Dimensions.Y)-(2*lineThick), prim.Style.PanelColor, true)
	if len(prim.Children) > 0 {
		for i := 0; i < len(prim.Children); i++ {
			err := prim.Children[i].Draw(prim.Image)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}
func (prim *UI_Object_Primitive) Update() error {
	xx, yy := ebiten.CursorPosition()
	prim.Update_Ret_State_Redraw_Status_Mport(xx, yy, 0)
	// if len(prim.Children) > 0 {
	// 	for i := 0; i < len(prim.Children); i++ {
	// 		_, to_redraw, err := prim.Children[i].Update_Ret_State_Redraw_Status()
	// 		if err != nil {
	// 			log.Fatal(err)
	// 		}
	// 		if to_redraw {
	// 			prim.Children[i].Draw(prim.Image)
	// 		}
	// 	}
	// }
	// if prim.IsMovable {
	// 	if prim.IsCursorInBounds() {

	// 	}
	// }
	// prim.Redraw()
	return nil
}

func (prim *UI_Object_Primitive) Update_Unactive() error {
	if len(prim.Children) > 0 {
		for i := 0; i < len(prim.Children); i++ {
			err := prim.Children[i].Update_Unactive()
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
func (prim *UI_Object_Primitive) Update_Any() (any, error) {
	return false, nil
}
func (prim *UI_Object_Primitive) Update_Ret_State_Redraw_Status() (uint8, bool, error) {
	return prim.State, false, nil
}
func (prim *UI_Object_Primitive) Update_Ret_State_Redraw_Status_Mport(Mouse_Pos_X, Mouse_Pos_Y, mode int) (uint8, bool, error) {

	if mode == 0 {
		if len(prim.Children) > 0 {
			for i := 0; i < len(prim.Children); i++ {
				_, to_redraw, err := prim.Children[i].Update_Ret_State_Redraw_Status_Mport(Mouse_Pos_X, Mouse_Pos_Y, 0)
				if err != nil {
					log.Fatal(err)
				}
				if to_redraw {
					prim.Children[i].Draw(prim.Image)
				}
			}
		}
		if prim.IsCursorInBounds_MousePort(Mouse_Pos_X, Mouse_Pos_Y, mode) {
			if inpututil.IsMouseButtonJustPressed(ebiten.MouseButton0) {
				log.Printf("%d %d %t\n", Mouse_Pos_X, Mouse_Pos_Y, prim.HasParent())

			}
		}
	}
	return prim.State, false, nil
}

/*
This returns the state of the object
*/
func (prim *UI_Object_Primitive) GetState() uint8 {
	return prim.State
}

/*
this returns a basic to string message
*/
func (prim *UI_Object_Primitive) ToString() string {
	strngOut := fmt.Sprintf("UI_Object Primtive:%s\n\tPositon %s\t", prim.obj_id, prim.Position.ToString())
	strngOut += fmt.Sprintf("\tDimensions: %s\n", prim.Dimensions.ToString())
	return strngOut
}

/*
this confirms the object is initilaized
*/
func (prim *UI_Object_Primitive) IsInit() bool {
	return prim.init
}

/*
this gets the object ID
*/
func (prim *UI_Object_Primitive) GetID() string {
	return prim.obj_id
}

/*
This returns a string specifying the objects type
*/
func (prim *UI_Object_Primitive) GetType() string {
	return "UI_Object Primitive"
}

/*
 */
func (prim *UI_Object_Primitive) IsCursorInBounds() bool {
	if prim.IsActive && prim.IsVisible {
		cX, cY := ebiten.CursorPosition()
		var x0, y0, x1, y1 int

		if prim.Parent != nil {
			px, py := prim.Parent.GetPosition_Int()
			x0 = prim.Position.X + px
			y0 = prim.Position.Y + py
			x1 = prim.Position.X + prim.Dimensions.X + px
			y1 = prim.Position.Y + prim.Dimensions.Y + py
			// x0 = prim.Position.X + prim.ParentPos.X
			// y0 = prim.Position.Y + prim.ParentPos.X
			// x1 = prim.Position.X + prim.ParentPos.X + prim.Dimensions.X
			// y1 = prim.Position.Y + prim.ParentPos.Y + prim.Dimensions.Y
		} else {
			x0 = prim.Position.X
			y0 = prim.Position.Y
			x1 = prim.Position.X + prim.Dimensions.X
			y1 = prim.Position.Y + prim.Dimensions.Y
		}
		return (cX > x0 && cX < x1) && (cY > y0 && cY < y1)
	}
	return false
}

/**/
func (prim *UI_Object_Primitive) Close() {}

/**/
func (prim *UI_Object_Primitive) Open() {}

/**/
func (prim *UI_Object_Primitive) Detoggle() {}

/*
Idea here is I don't want to waste time with having to get the cursor position when it possibly hasn't changed enough to matter;
This might be also a terrible idea overall I cannot tell quite yet

enter 0 for it to default
*/
func (prim *UI_Object_Primitive) IsCursorInBounds_MousePort(Mouse_Pos_X, Mouse_Pos_Y, mode int) bool {
	if prim.IsActive && prim.IsVisible {
		cX, cY := Mouse_Pos_X, Mouse_Pos_Y
		//mode stuff
		var x0, y0, x1, y1 int

		if prim.Parent != nil {
			px, py := prim.Parent.GetPosition_Int()
			x0 = prim.Position.X + px
			y0 = prim.Position.Y + py
			x1 = prim.Position.X + prim.Dimensions.X + px
			y1 = prim.Position.Y + prim.Dimensions.Y + py
			if mode == 10 {
				x3, y3 := prim.Parent.Get_Internal_Position_Int()
				x0 += x3
				x1 += x3
				y0 += y3
				y1 += y3
				if !prim.Parent.IsCursorInBounds_MousePort(Mouse_Pos_X, Mouse_Pos_Y, 10) {

					return false
				}
			}
		} else {
			x0 = prim.Position.X
			y0 = prim.Position.Y
			x1 = prim.Position.X + prim.Image.Bounds().Dx()
			y1 = prim.Position.Y + prim.Image.Bounds().Dy()
		}
		return (cX > x0 && cX < x1) && (cY > y0 && cY < y1)
	}
	//mode stuff
	return false
}

/**/
func (prim *UI_Object_Primitive) Get_Internal_Position_Int() (x_pos int, y_pos int) {
	if prim.Parent != nil {
		x_pos, y_pos = prim.Parent.Get_Internal_Position_Int()
	}
	return x_pos, y_pos
}

/**/
func (prim *UI_Object_Primitive) GetPosition_Int() (int, int) {
	xx := prim.Position.X
	yy := prim.Position.Y
	if prim.Parent != nil {
		px, py := prim.Parent.GetPosition_Int()
		xx += px
		yy += py
	}
	return xx, yy
}

/**/
func (prim *UI_Object_Primitive) SetPosition_Int(pos_X, pos_Y int) {
	prim.Position = coords.CoordInts{X: pos_X, Y: pos_Y}
}

/**/
func (prim *UI_Object_Primitive) GetDimensions_Int() (int, int) {
	return prim.Dimensions.X, prim.Dimensions.Y
} //
/**/
func (prim *UI_Object_Primitive) SetDimensions_Int(pos_X, pos_Y int) {
	prim.Dimensions = coords.CoordInts{X: pos_X, Y: pos_Y}

}

func (prim *UI_Object_Primitive) GetNumber_Children() int {
	return len(prim.Children)
}
func (prim *UI_Object_Primitive) GetChild(index int) UI_Object {
	if len(prim.Children) > index {
		return prim.Children[index]
	} else {
		return nil
	}
}
func (prim *UI_Object_Primitive) AddChild(child UI_Object) error {
	prim.Children = append(prim.Children, child)
	return nil
}
func (prim *UI_Object_Primitive) RemoveChild(index int) error {
	// prim.Children = append(prim.Children, child)
	return nil
}
func (prim *UI_Object_Primitive) HasParent() bool {
	return prim.Parent != nil
}
func (prim *UI_Object_Primitive) GetParent() UI_Object {
	return prim.Parent
}
