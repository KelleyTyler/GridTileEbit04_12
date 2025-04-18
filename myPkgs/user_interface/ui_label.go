package user_interface

import (
	"fmt"

	"github.com/KelleyTyler/GridTileEbit04_12/myPkgs/basic_geometry/coords"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

/*
	type UI_Object_Style struct {
		LabelColor       color.Color
		PanelColor       color.Color
		BorderColor      color.Color
		BorderThickness  float32
		TextColor        color.Color

		TextSize         float32

		Internal_Margins [4]uint8

		WillScroll       bool

}
*/

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

type UI_Label struct {
	Position      coords.CoordInts
	ParentPos     *coords.CoordInts
	Dimensions    coords.CoordInts
	Style         *UI_Object_Style
	TextAlignMode int
	Backend       *UI_Backend
	id, Text      string
	Img           *ebiten.Image
	Parent        UI_Object
}

func (lbl *UI_Label) Init_00(backend *UI_Backend, label string, pos, dim coords.CoordInts, parentPosition *coords.CoordInts) {
	lbl.Style = &backend.Style
	lbl.Backend = backend
	lbl.Position = pos
	lbl.Dimensions = dim
	lbl.Img = ebiten.NewImage(dim.X, dim.Y)
	lbl.Text = label
	lbl.TextAlignMode = 0
	lbl.Redraw()
}
func (lbl *UI_Label) Init(idLabels []string, backend *UI_Backend, style *UI_Object_Style, pos, dim coords.CoordInts) error {

	lbl.Backend = backend
	if style != nil {
		lbl.Style = &lbl.Backend.Style
	} else {
		lbl.Style = &backend.Style
	}

	lbl.Position = pos
	lbl.Dimensions = dim
	lbl.Img = ebiten.NewImage(dim.X, dim.Y)
	if len(idLabels) > 1 {
		lbl.Text = idLabels[1]
		lbl.id = idLabels[0]
	} else {
		lbl.Text = "null"
		lbl.id = "lbl_id_00"
	}
	lbl.TextAlignMode = 0

	lbl.Redraw()
	return nil
}
func (lbl *UI_Label) Init_Parents(Parent UI_Object) error {
	lbl.Parent = Parent
	lbl.Parent.AddChild(lbl)
	lbl.Parent.Redraw()
	if lbl.Text == "null" {
		lbl.Text = Parent.GetID()
	}
	return nil
}
func (lbl *UI_Label) Init_Parents_spec(Parent UI_Object) error {
	lbl.Parent = Parent
	// lbl.Parent.AddChild(lbl)
	lbl.Parent.Redraw()
	if lbl.Text == "null" {
		lbl.Text = Parent.GetID()
	}
	return nil
}

func (lbl *UI_Label) Redraw() {
	lbl.Img.Fill(lbl.Style.BorderColor)
	borderThick := lbl.Backend.Style.BorderThickness
	vector.DrawFilledRect(lbl.Img, borderThick, borderThick, float32(lbl.Dimensions.X)-(borderThick*2), float32(lbl.Dimensions.Y)-(borderThick*2), lbl.Backend.Style.ButtonColor0[0], true)

	scaler := 1.0
	tops := &text.DrawOptions{}
	// temper := fmt.Sprintf("-%s-", lbl.Text)
	if lbl.TextAlignMode == 10 {
		// temper = fmt.Sprintf("x-%s-x", lbl.Text)
		tops.GeoM.Translate(float64(lbl.Dimensions.X/2)*scaler, float64(lbl.Dimensions.Y/2)*scaler)
		tops.GeoM.Scale(1/scaler, 1/scaler)
		tops.ColorScale.ScaleWithColor(lbl.Style.TextColor[0])
		tops.LineSpacing = float64(10) * scaler
		tops.PrimaryAlign = text.AlignCenter
		tops.SecondaryAlign = text.AlignCenter
	} else {
		// tops.GeoM.Translate(float64(lbl.Dimensions.X/2)*scaler, float64(lbl.Dimensions.Y/2)*scaler)
		tops.GeoM.Translate(float64(8)*scaler, float64(8)*scaler)
		tops.GeoM.Scale(1/scaler, 1/scaler)
		tops.ColorScale.ScaleWithColor(lbl.Style.TextColor[0])
		tops.LineSpacing = float64(10) * scaler
	}
	text.Draw(lbl.Img, lbl.Text, lbl.Backend.Btn_Text_Reg, tops) //Btn_Text_Reg
	// if lbl.Parent != nil {
	// 	lbl.Parent.Redraw()
	// }
}
func (lbl *UI_Label) Draw(screen *ebiten.Image) error {
	ops := ebiten.DrawImageOptions{}
	ops.GeoM.Reset()
	ops.GeoM.Translate(float64(lbl.Position.X), float64(lbl.Position.Y))
	screen.DrawImage(lbl.Img, &ops)
	return nil
}

func (lbl *UI_Label) Update() error {
	return nil
}                                              //--
func (lbl *UI_Label) Update_Unactive() error   { return nil }        //
func (lbl *UI_Label) Update_Any() (any, error) { return false, nil } //
func (lbl *UI_Label) Update_Ret_State_Redraw_Status() (uint8, bool, error) {
	return 0, false, nil
}
func (lbl *UI_Label) Update_Ret_State_Redraw_Status_Mport(Mouse_Pos_X, Mouse_Pos_Y, mode int) (uint8, bool, error) {
	return 0, false, nil
}
func (lbl *UI_Label) GetState() uint8        { return 0 }                 //
func (lbl *UI_Label) ToString() string       { return "This is a Label" } //
func (lbl *UI_Label) IsInit() bool           { return false }             //
func (lbl *UI_Label) GetID() string          { return lbl.id }            //
func (lbl *UI_Label) GetType() string        { return "UI_Label" }        //
func (lbl *UI_Label) IsCursorInBounds() bool { return false }             //
func (lbl *UI_Label) IsCursorInBounds_MousePort(Mouse_Pos_X, Mouse_Pos_Y, mode int) bool {
	return false
}                                                  //
func (lbl *UI_Label) GetPosition_Int() (int, int)  { return lbl.Position.X, lbl.Position.Y } //
func (lbl *UI_Label) GetNumber_Children() int      { return 0 }                              //
func (lbl *UI_Label) GetChild(index int) UI_Object { return nil }                            //
func (lbl *UI_Label) AddChild(child UI_Object) error {
	return fmt.Errorf("ERROR NOT POSSIBLE")
}
func (lbl *UI_Label) RemoveChild(index int) error {
	return fmt.Errorf("ERROR NOT POSSIBLE")
}
func (lbl *UI_Label) GetParent() UI_Object { return nil } //
func (lbl *UI_Label) HasParent() bool      { return lbl.Parent != nil }
