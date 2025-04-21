package user_interface

import (
	"image/color"

	"github.com/KelleyTyler/GridTileEbit04_12/myPkgs/basic_geometry/coords"
	settings "github.com/KelleyTyler/GridTileEbit04_12/myPkgs/settingsconfig"
	"github.com/hajimehoshi/ebiten/v2"
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

		Update_Any() (any, error) //
		Update_Ret_State_Redraw_Status() (uint8, bool, error)
		Update_Ret_State_Redraw_Status_Mport(Mouse_Pos_X, Mouse_Pos_Y, mode int) (uint8, bool, error)

		GetState() uint8                                                    //
		ToString() string                                                   //
		IsInit() bool                                                       //
		GetID() string                                                      //
		GetType() string                                                    //
		IsCursorInBounds() bool                                             //
		IsCursorInBounds_MousePort(Mouse_Pos_X, Mouse_Pos_Y, mode int) bool //
		GetPosition_Int() (int, int)                                        //
		GetNumber_Children() int                                            //
		GetChild(index int) UI_Object
		AddChild(child UI_Object) error //
		RemoveChild(index int) error
		GetParent() UI_Object //
		HasParent() bool      //
		// getType() string //might want to change this output to like an int or something using a golang equivelant of an enum;
	}
*/
type UI_TextEntryField struct {
	Position            coords.CoordInts
	Dimensions          coords.CoordInts
	Style               *UI_Object_Style
	Backend             *UI_Backend
	Img, SubImg         *ebiten.Image
	Parent              UI_Object
	BG_color            color.Color
	TextColor           color.Color
	DataField           string
	IsActive, IsVisible bool
	WrapAround          bool
	TextSize            uint8
	MaxTextPerLine      uint16
	CurrentLine         uint16
	MaxLineOfText       uint16
	Scale               float32
}

func (tef *UI_TextEntryField) Init(idLabels []string, backend *UI_Backend, style *UI_Object_Style, Position, Dimensions coords.CoordInts) error {
	tef.Position = Position
	tef.Dimensions = Dimensions
	tef.Backend = backend
	if style != nil {
		tef.Style = style
	} else {
		tef.Style = &tef.Backend.Style
	}
	return nil
}                                                                  //--
func (tef *UI_TextEntryField) Init_Parents(Parent UI_Object) error { return nil } //--
func (tef *UI_TextEntryField) Draw(screen *ebiten.Image) error     { return nil } //--
func (tef *UI_TextEntryField) Redraw()                             {}             //--
func (tef *UI_TextEntryField) Update() error                       { return nil } //--
func (tef *UI_TextEntryField) Update_Unactive() error              { return nil } //

func (tef *UI_TextEntryField) Update_Any() (any, error) { return 0, nil }
func (tef *UI_TextEntryField) Update_Ret_State_Redraw_Status() (uint8, bool, error) {
	return 0, false, nil
}
func (tef *UI_TextEntryField) Update_Ret_State_Redraw_Status_Mport(Mouse_Pos_X, Mouse_Pos_Y, mode int) (uint8, bool, error) {
	return 0, false, nil
}

func (tef *UI_TextEntryField) IsCursorInBounds() bool {
	return false
} //
func (tef *UI_TextEntryField) IsCursorInBounds_MousePort(Mouse_Pos_X, Mouse_Pos_Y, mode int) bool {
	return false
}

func (tef *UI_TextEntryField) GetPosition_Int() (int, int) {
	if tef.Parent != nil {
		xx, yy := tef.Parent.GetPosition_Int()
		return tef.Position.X + xx, tef.Position.Y + yy
	} else {
		return tef.Position.X, tef.Position.Y
	}
}

func (tef *UI_TextEntryField) GetState() uint8         { return 0 }
func (tef *UI_TextEntryField) ToString() string        { return "" }
func (tef *UI_TextEntryField) IsInit() bool            { return false }
func (tef *UI_TextEntryField) GetID() string           { return "" }
func (tef *UI_TextEntryField) GetType() string         { return "Text Entry Field" }
func (tef *UI_TextEntryField) GetNumber_Children() int { return 0 }
func (tef *UI_TextEntryField) GetChild() UI_Object     { return nil }
func (tef *UI_TextEntryField) RemoveChild() error      { return nil }
func (tef *UI_TextEntryField) GetParent() UI_Object    { return tef.Parent }
func (tef *UI_TextEntryField) HasParent() bool         { return tef.Parent != nil }

type UI_TextEntryWindow struct {
	Position                coords.CoordInts
	Dimensions              coords.CoordInts
	Parent                  UI_Object
	CloseButton             UI_Button
	ClearButton             UI_Button
	SubmitButton            UI_Button
	TexField                UI_TextEntryField
	IsActive, IsVisible     bool
	IsResizable, IsMoveable bool
	PanelImage              *ebiten.Image
	Settings                *settings.GameSettings
	Backend                 *UI_Backend
	Style                   *UI_Object_Style
	WindowLabel             string
	//-------------------------Needs Content

}

func (tew *UI_TextEntryWindow) Init(idLabels []string, backend *UI_Backend, style *UI_Object_Style, Position, Dimensions coords.CoordInts) error {
	return nil
}                                                                   //--
func (tew *UI_TextEntryWindow) Init_Parents(Parent UI_Object) error { return nil } //--
func (tew *UI_TextEntryWindow) Draw(screen *ebiten.Image) error     { return nil } //--
func (tew *UI_TextEntryWindow) Redraw()                             {}             //--
func (tew *UI_TextEntryWindow) Update() error                       { return nil } //--
func (tew *UI_TextEntryWindow) Update_Unactive() error              { return nil } //

func (tew *UI_TextEntryWindow) Update_Any() (any, error) { return 0, nil }
func (tew *UI_TextEntryWindow) Update_Ret_State_Redraw_Status() (uint8, bool, error) {
	return 0, false, nil
}
func (tew *UI_TextEntryWindow) Update_Ret_State_Redraw_Status_Mport(Mouse_Pos_X, Mouse_Pos_Y, mode int) (uint8, bool, error) {
	return 0, false, nil
}

func (tew *UI_TextEntryWindow) IsCursorInBounds() bool {
	return false
} //
func (tew *UI_TextEntryWindow) IsCursorInBounds_MousePort(Mouse_Pos_X, Mouse_Pos_Y, mode int) bool {
	return false
}

func (tew *UI_TextEntryWindow) GetPosition_Int() (int, int) {
	if tew.Parent != nil {
		xx, yy := tew.Parent.GetPosition_Int()
		return tew.Position.X + xx, tew.Position.Y + yy
	} else {
		return tew.Position.X, tew.Position.Y
	}
}
func (tew *UI_TextEntryWindow) GetState() uint8         { return 0 }
func (tew *UI_TextEntryWindow) ToString() string        { return "" }
func (tew *UI_TextEntryWindow) IsInit() bool            { return false }
func (tew *UI_TextEntryWindow) GetID() string           { return "" }
func (tew *UI_TextEntryWindow) GetType() string         { return "Text Entry Window" }
func (tew *UI_TextEntryWindow) GetNumber_Children() int { return 0 }
func (tew *UI_TextEntryWindow) GetChild() UI_Object     { return nil }
func (tew *UI_TextEntryWindow) RemoveChild() error      { return nil }
func (tew *UI_TextEntryWindow) GetParent() UI_Object    { return tew.Parent }
func (tew *UI_TextEntryWindow) HasParent() bool         { return tew.Parent != nil }
