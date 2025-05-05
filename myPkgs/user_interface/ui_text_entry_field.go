package user_interface

import (
	"image/color"
	"log"
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
func (ui_tef *UI_TextEntryField) Init(idLabels []string, backend *UI_Backend, style *UI_Object_Style, Position, Dimensions coords.CoordInts) error {
	ui_tef.Position = Position
	ui_tef.Dimensions = Dimensions
	ui_tef.Backend = backend
	ui_tef.maxLines = 1
	ui_tef.Image = ebiten.NewImage(Dimensions.X, Dimensions.Y)
	ui_tef.Text_Input_Image = ebiten.NewImage(Dimensions.X-2, Dimensions.Y-2)
	if style != nil {
		ui_tef.Style = style
	} else {
		ui_tef.Style = &ui_tef.Backend.Style
	}
	ui_tef.IsVisible = true
	// ui_tef.Redraw()
	return nil
}

/**/
func (ui_tef *UI_TextEntryField) Init_Parents(Parent UI_Object) error {
	ui_tef.Parent = Parent
	ui_tef.Parent.AddChild(ui_tef)
	ui_tef.Redraw()
	ui_tef.Parent.Redraw()
	return nil
}

/**/
func (ui_tef *UI_TextEntryField) Draw(screen *ebiten.Image) error {
	// ui_tef.Redraw()
	if ui_tef.IsVisible {
		ops := ebiten.DrawImageOptions{}
		scale := 1.0
		ops.GeoM.Reset()
		ops.GeoM.Translate(float64(ui_tef.Position.X)*scale, float64(ui_tef.Position.Y)*scale)
		ui_tef.Predraw()
		screen.DrawImage(ui_tef.Image, &ops)

	}
	return nil
} //--
/**/
func (ui_tef *UI_TextEntryField) Redraw() {

	ui_tef.Image.Fill(color.Black)
	vector.DrawFilledRect(ui_tef.Image, 2, 2, float32(ui_tef.Dimensions.X-4), float32(ui_tef.Dimensions.Y-4), ui_tef.Style.PanelColor, true)

} //--

func (ui_tef *UI_TextEntryField) Predraw() {
	//log.Printf("PREDRAW")
	// ui_tef.Image.Fill(color.RGBA{255, 255, 255, 255})
	ui_tef.Redraw()

	scaler := 1.0 //1.75
	tops := &text.DrawOptions{}
	tops.GeoM.Reset()
	// tops.GeoM.Translate(float64(ui_tef.Position.X+2)*scaler, float64(ui_tef.Position.Y)*scaler)
	tops.GeoM.Translate(float64(2)*scaler, float64(2)*scaler)

	tops.GeoM.Scale(1/scaler, 1/scaler)
	tops.ColorScale.ScaleWithColor(color.Black)
	tops.LineSpacing = float64(20)
	t := ui_tef.DataField
	if ui_tef.counter%60 < 30 {
		t += "_"
		ui_tef.counter = 0
	}
	text.Draw(ui_tef.Image, t, *ui_tef.Backend.GetTextFace(0, 20), tops)
}

/**/
func (ui_tef *UI_TextEntryField) Get_String_Output() (strng string) {

	strng = ui_tef.DataField

	ui_tef.DataField = ""
	log.Printf("DATAFIELD : %s\n", ui_tef.DataField)
	return strng
}

/*
This used to be "clear" but in the interests of making this whole shitshow make sense I've changed it to detoggle; then again i might change it back because it could make a bit more sense;
*/
func (ui_tef *UI_TextEntryField) Detoggle() {
	// ui_tef.IsActive = false
	// ui_tef.IsActive = false
	ui_tef.DataField = ""

}

/*
 */
func (ui_tef *UI_TextEntryField) Close() {
	ui_tef.IsActive = false
	ui_tef.IsVisible = false
	ui_tef.Detoggle()
}

/*
 */
func (ui_tef *UI_TextEntryField) Open() {
	ui_tef.IsActive = true
	ui_tef.IsVisible = true
	ui_tef.Detoggle()
}

/*
 */
func (ui_tef *UI_TextEntryField) Update() error {
	if ui_tef.IsCursorInBounds() {

	}

	if ui_tef.IsActive {
		// ui_tef.counter++
		ui_tef.Data = ebiten.AppendInputChars(ui_tef.Data[:0])
		ui_tef.DataField += string(ui_tef.Data)

		ss := strings.Split(ui_tef.DataField, "\n")
		// if len(ss) > ui_tef.maxLines {
		// 	ui_tef.DataStrng = strings.Join(ss[len(ss)-ui_tef.maxLines:], "\n")
		// } else if len(ui_tef.DataStrng) > 10 {
		// 	log.Printf("%d \n", len(ss))
		// }
		if len(ss) > int(ui_tef.maxLines) {
			ui_tef.DataField = strings.Join(ss[len(ss)-int(ui_tef.maxLines):], "\n")
		}

		if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
			ui_tef.DataField += "\n"
		}
		if inpututil.IsKeyJustPressed(ebiten.KeyBackspace) {
			if len(ui_tef.DataField) >= 1 {
				ui_tef.DataField = ui_tef.DataField[:len(ui_tef.DataField)-1]
			}
		}
		ui_tef.counter++
	}
	return nil
} //--
/**/
func (ui_tef *UI_TextEntryField) Update_Unactive() error { return nil } //
/**/
func (ui_tef *UI_TextEntryField) Update_Any() (any, error)       { return 0, nil }
func (ui_tef *UI_TextEntryField) AddChild(child UI_Object) error { return nil }

/**/
func (ui_tef UI_TextEntryField) Update_Ret_State_Redraw_Status() (uint8, bool, error) {
	return 0, false, nil
}

/**/
func (ui_tef *UI_TextEntryField) Update_Ret_State_Redraw_Status_Mport(Mouse_Pos_X, Mouse_Pos_Y, mode int) (uint8, bool, error) {
	return 0, false, nil
}

/**/
func (ui_tef *UI_TextEntryField) IsCursorInBounds() bool {
	Mouse_Pos_X, Mouse_Pos_Y := ebiten.CursorPosition()
	return ui_tef.IsCursorInBounds_MousePort(Mouse_Pos_X, Mouse_Pos_Y, 10)
}

/**/
func (ui_tef *UI_TextEntryField) IsCursorInBounds_MousePort(Mouse_Pos_X, Mouse_Pos_Y, mode int) bool {
	if ui_tef.IsVisible {
		var x0, y0, x1, y1 int

		if ui_tef.Parent != nil {
			px, py := ui_tef.Parent.GetPosition_Int()
			x0 = ui_tef.Position.X + px
			y0 = ui_tef.Position.Y + py
			x1 = ui_tef.Position.X + ui_tef.Dimensions.X + px
			y1 = ui_tef.Position.Y + ui_tef.Dimensions.Y + py
			if mode == 10 {
				x3, y3 := ui_tef.Parent.Get_Internal_Position_Int()
				x0 += x3
				x1 += x3
				y0 += y3
				y1 += y3
				if !ui_tef.Parent.IsCursorInBounds_MousePort(Mouse_Pos_X, Mouse_Pos_Y, 10) {
					return false
				}
			}
		} else {
			x0 = ui_tef.Position.X
			y0 = ui_tef.Position.Y
			x1 = ui_tef.Position.X + ui_tef.Dimensions.X
			y1 = ui_tef.Position.Y + ui_tef.Dimensions.Y
		}
		return (Mouse_Pos_X > x0 && Mouse_Pos_X < x1) && (Mouse_Pos_Y > y0 && Mouse_Pos_Y < y1)
	}
	return false
}

/*

 */
/**/
func (ui_tef *UI_TextEntryField) GetPosition_Int() (int, int) {
	if ui_tef.Parent != nil {
		xx, yy := ui_tef.Parent.GetPosition_Int()
		return ui_tef.Position.X + xx, ui_tef.Position.Y + yy
	} else {
		return ui_tef.Position.X, ui_tef.Position.Y
	}
}

/**/
func (ui_tef *UI_TextEntryField) Get_Internal_Position_Int() (x_pos int, y_pos int) {
	return ui_tef.GetPosition_Int()
}

/**/
func (ui_tef *UI_TextEntryField) Get_Internal_Dimensions_Int() (x_pos int, y_pos int) {
	x_pos, y_pos = ui_tef.GetPosition_Int()
	return x_pos, y_pos
}

/**/
func (ui_tef *UI_TextEntryField) SetPosition_Int(x_point, y_point int) {
	ui_tef.Position = coords.CoordInts{X: x_point, Y: y_point}
}

/**/
func (ui_tef *UI_TextEntryField) GetDimensions_Int() (int, int) {
	return ui_tef.Dimensions.X, ui_tef.Dimensions.Y
}

/**/
func (ui_tef *UI_TextEntryField) SetDimensions_Int(x_point, y_point int) {
	ui_tef.Dimensions = coords.CoordInts{X: x_point, Y: y_point}
	//---Redraw The S
}

/**/
func (ui_tef *UI_TextEntryField) GetState() uint8          { return 0 }
func (ui_tef *UI_TextEntryField) ToString() string         { return "" }
func (ui_tef *UI_TextEntryField) IsInit() bool             { return false }
func (ui_tef *UI_TextEntryField) GetID() string            { return "" }
func (ui_tef *UI_TextEntryField) GetType() string          { return "Text Entry Field" }
func (ui_tef *UI_TextEntryField) GetNumber_Children() int  { return 0 }
func (ui_tef *UI_TextEntryField) GetChild(n int) UI_Object { return nil }
func (ui_tef *UI_TextEntryField) RemoveChild(n int) error  { return nil }
func (ui_tef *UI_TextEntryField) GetParent() UI_Object     { return ui_tef.Parent }
func (ui_tef *UI_TextEntryField) HasParent() bool          { return ui_tef.Parent != nil }

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
