package user_interface

import (
	"fmt"
	"image/color"
	"strings"

	"github.com/KelleyTyler/GridTileEbit04_12/myPkgs/basic_geometry/coords"
	settings "github.com/KelleyTyler/GridTileEbit04_12/myPkgs/settingsconfig"
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
	Position         coords.CoordInts
	Dimensions       coords.CoordInts
	Style            *UI_Object_Style
	Backend          *UI_Backend
	Image            *ebiten.Image
	Text_Input_Image *ebiten.Image

	Parent              UI_Object
	BG_color            color.Color
	TextColor           color.Color
	Data                []rune
	DataField           string
	IsActive, IsVisible bool
	WrapAround          bool
	TextSize            uint8
	MaxTextPerLine      uint16
	CurrentLine         uint16
	MaxLineOfText       uint16
	Scale               float32
	counter             uint16
	maxLines            uint8
}

/**/
func (tef *UI_TextEntryField) Init(idLabels []string, backend *UI_Backend, style *UI_Object_Style, Position, Dimensions coords.CoordInts) error {
	tef.Position = Position
	tef.Dimensions = Dimensions
	tef.Backend = backend
	tef.maxLines = 1
	tef.Image = ebiten.NewImage(Dimensions.X, Dimensions.Y)
	tef.Text_Input_Image = ebiten.NewImage(Dimensions.X-2, Dimensions.Y-2)
	if style != nil {
		tef.Style = style
	} else {
		tef.Style = &tef.Backend.Style
	}
	tef.IsVisible = true
	// tef.Redraw()
	return nil
}

/**/
func (tef *UI_TextEntryField) Init_Parents(Parent UI_Object) error {
	tef.Parent = Parent
	tef.Parent.AddChild(tef)
	tef.Redraw()
	tef.Parent.Redraw()
	return nil
}

/**/
func (tef *UI_TextEntryField) Draw(screen *ebiten.Image) error {
	// tef.Redraw()
	if tef.IsVisible {
		ops := ebiten.DrawImageOptions{}
		scale := 1.0
		ops.GeoM.Reset()
		ops.GeoM.Translate(float64(tef.Position.X)*scale, float64(tef.Position.Y)*scale)
		tef.Predraw()
		screen.DrawImage(tef.Image, &ops)

	}
	return nil
} //--
/**/
func (tef *UI_TextEntryField) Redraw() {

	tef.Image.Fill(color.Black)
	vector.DrawFilledRect(tef.Image, 2, 2, float32(tef.Dimensions.X-4), float32(tef.Dimensions.Y-4), tef.Style.PanelColor, true)

} //--

func (tef *UI_TextEntryField) Predraw() {
	//fmt.Printf("PREDRAW")
	// tef.Image.Fill(color.RGBA{255, 255, 255, 255})
	tef.Redraw()

	scaler := 1.0 //1.75
	tops := &text.DrawOptions{}
	tops.GeoM.Reset()
	// tops.GeoM.Translate(float64(tef.Position.X+2)*scaler, float64(tef.Position.Y)*scaler)
	tops.GeoM.Translate(float64(2)*scaler, float64(2)*scaler)

	tops.GeoM.Scale(1/scaler, 1/scaler)
	tops.ColorScale.ScaleWithColor(color.Black)
	tops.LineSpacing = float64(20)
	t := tef.DataField
	if tef.counter%60 < 30 {
		t += "_"
		tef.counter = 0
	}
	text.Draw(tef.Image, t, *tef.Backend.GetTextFace(0, 20), tops)
}

/**/
func (tef *UI_TextEntryField) Get_String_Output() (strng string) {

	strng = tef.DataField

	tef.DataField = ""
	fmt.Printf("DATAFIELD : %s\n", tef.DataField)
	return strng
}

/*
 */
func (tef *UI_TextEntryField) Clear() error {
	// tef.IsActive = false
	// tef.IsActive = false
	tef.DataField = ""

	return nil
}

/*
 */
func (tef *UI_TextEntryField) Update() error {
	if tef.IsCursorInBounds() {

	}

	if tef.IsActive {
		// tef.counter++
		tef.Data = ebiten.AppendInputChars(tef.Data[:0])
		tef.DataField += string(tef.Data)

		ss := strings.Split(tef.DataField, "\n")
		// if len(ss) > tef.maxLines {
		// 	tef.DataStrng = strings.Join(ss[len(ss)-tef.maxLines:], "\n")
		// } else if len(tef.DataStrng) > 10 {
		// 	fmt.Printf("%d \n", len(ss))
		// }
		if len(ss) > int(tef.maxLines) {
			tef.DataField = strings.Join(ss[len(ss)-int(tef.maxLines):], "\n")
		}

		if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
			tef.DataField += "\n"
		}
		if inpututil.IsKeyJustPressed(ebiten.KeyBackspace) {
			if len(tef.DataField) >= 1 {
				tef.DataField = tef.DataField[:len(tef.DataField)-1]
			}
		}
		tef.counter++
	}
	return nil
} //--
/**/
func (tef *UI_TextEntryField) Update_Unactive() error { return nil } //
/**/
func (tef *UI_TextEntryField) Update_Any() (any, error)       { return 0, nil }
func (tef *UI_TextEntryField) AddChild(child UI_Object) error { return nil }

/**/
func (tef UI_TextEntryField) Update_Ret_State_Redraw_Status() (uint8, bool, error) {
	return 0, false, nil
}

/**/
func (tef *UI_TextEntryField) Update_Ret_State_Redraw_Status_Mport(Mouse_Pos_X, Mouse_Pos_Y, mode int) (uint8, bool, error) {
	return 0, false, nil
}

/**/
func (tef *UI_TextEntryField) IsCursorInBounds() bool {
	if tef.IsVisible {
		cX, cY := ebiten.CursorPosition()
		var x0, y0, x1, y1 int

		if tef.Parent != nil {
			px, py := tef.Parent.GetPosition_Int()
			x0 = tef.Position.X + px
			y0 = tef.Position.Y + py
			x1 = tef.Position.X + tef.Dimensions.X + px
			y1 = tef.Position.Y + tef.Dimensions.Y + py
			// x0 = ui_win.Position.X + ui_win.ParentPos.X
			// y0 = ui_win.Position.Y + ui_win.ParentPos.X
			// x1 = ui_win.Position.X + ui_win.ParentPos.X + ui_win.Dimensions.X
			// y1 = ui_win.Position.Y + ui_win.ParentPos.Y + ui_win.Dimensions.Y
		} else {
			x0 = tef.Position.X
			y0 = tef.Position.Y
			x1 = tef.Position.X + tef.Dimensions.X
			y1 = tef.Position.Y + tef.Dimensions.Y
		}
		return (cX > x0 && cX < x1) && (cY > y0 && cY < y1)
	}
	return false
}

/**/
func (tef *UI_TextEntryField) IsCursorInBounds_MousePort(Mouse_Pos_X, Mouse_Pos_Y, mode int) bool {
	return false
}

/**/
func (tef *UI_TextEntryField) GetPosition_Int() (int, int) {
	if tef.Parent != nil {
		xx, yy := tef.Parent.GetPosition_Int()
		return tef.Position.X + xx, tef.Position.Y + yy
	} else {
		return tef.Position.X, tef.Position.Y
	}
}

func (tef *UI_TextEntryField) GetState() uint8          { return 0 }
func (tef *UI_TextEntryField) ToString() string         { return "" }
func (tef *UI_TextEntryField) IsInit() bool             { return false }
func (tef *UI_TextEntryField) GetID() string            { return "" }
func (tef *UI_TextEntryField) GetType() string          { return "Text Entry Field" }
func (tef *UI_TextEntryField) GetNumber_Children() int  { return 0 }
func (tef *UI_TextEntryField) GetChild(n int) UI_Object { return nil }
func (tef *UI_TextEntryField) RemoveChild(n int) error  { return nil }
func (tef *UI_TextEntryField) GetParent() UI_Object     { return tef.Parent }
func (tef *UI_TextEntryField) HasParent() bool          { return tef.Parent != nil }

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
