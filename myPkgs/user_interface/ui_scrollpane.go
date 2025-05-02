package user_interface

import (
	"fmt"
	"image/color"
	"log"

	"github.com/KelleyTyler/GridTileEbit04_12/myPkgs/basic_geometry/coords"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

/*
Init(idLabels []string, backend *UI_Backend, style *UI_Object_Style, Position, Dimensions coords.CoordInts) error //--
	Init_Parents(Parent UI_Object) error                                                                              //--
	Draw(screen *ebiten.Image) error                                                                                  //--
	Redraw()                                                                                                          //--
	Update() error                                                                                                    //--
	Update_Unactive() error                                                                                           //

	//Update_Any() (any, error) //
	Update_Ret_State_Redraw_Status() (uint8, bool, error)
	Update_Ret_State_Redraw_Status_Mport(Mouse_Pos_X, Mouse_Pos_Y, mode int) (uint8, bool, error)

	GetState() uint8
	ToString() string
	IsInit() bool
	GetID() string
	GetType() string
	IsCursorInBounds() bool
	IsCursorInBounds_MousePort(Mouse_Pos_X, Mouse_Pos_Y, mode int) bool
	GetPosition_Int() (int, int)
	GetNumber_Children() int
	GetChild(index int) UI_Object
	AddChild(child UI_Object) error
	RemoveChild(index int) error
	GetParent() UI_Object
	HasParent() bool

*/
/*
	The goal here is a pane that can scroll;
*/

/*
This is a UI Pane that has a scroll-menu/thing
*/
type UI_Scrollpane struct {
	Position, Internal_Position     coords.CoordInts
	Dim_Default, Dim_2              coords.CoordInts
	Parent                          UI_Object
	Backend                         *UI_Backend
	Style                           *UI_Object_Style
	Image, DisplayImage, UnderImage *ebiten.Image

	ui_obj_id       string
	pane_name       string
	scrollbar_width uint8
	// scrollbuttonSize                 [2]uint8
	ShowLabel, Scrollable, Resizable bool
	IsActive, IsVisible              bool
	State                            uint8

	Scrollbar_Vertical   UI_Scrollbar
	Scrollbar_Horizontal UI_Scrollbar
	Children             []UI_Object
}

/**/
func (ui_scroll *UI_Scrollpane) Init(idLabels []string, backend *UI_Backend, style *UI_Object_Style, Position, Dimensions coords.CoordInts) error {
	ui_scroll.ui_obj_id = idLabels[0]
	ui_scroll.pane_name = idLabels[1]
	ui_scroll.Backend = backend
	if style != nil {
		ui_scroll.Style = style
	} else {
		ui_scroll.Style = &backend.Style
	}
	ui_scroll.Children = make([]UI_Object, 0)
	ui_scroll.scrollbar_width = 16
	modified := int(2*ui_scroll.Style.BorderThickness) + int(ui_scroll.scrollbar_width)
	ui_scroll.Image = ebiten.NewImage(Dimensions.X+modified, Dimensions.Y+modified)
	ui_scroll.DisplayImage = ebiten.NewImage(Dimensions.X, Dimensions.Y)
	ui_scroll.UnderImage = ebiten.NewImage(Dimensions.X*2, Dimensions.Y*2)
	log.Printf("IMAGE INTIIALIZATION \n IMAGE: %5t\t %d %d ", ui_scroll.Image != nil, ui_scroll.Image.Bounds().Dx(), ui_scroll.Image.Bounds().Dy())
	log.Printf(" UNDERIMAGE: %5t\t %d %d ", ui_scroll.UnderImage != nil, ui_scroll.UnderImage.Bounds().Dx(), ui_scroll.Image.Bounds().Dy())
	log.Printf(" DisplayImage: %5t\t %d %d ", ui_scroll.DisplayImage != nil, ui_scroll.DisplayImage.Bounds().Dx(), ui_scroll.Image.Bounds().Dy())
	ui_scroll.Scrollbar_Vertical.Init([]string{"scrollbar_vertical", "X"}, backend, style, coords.CoordInts{X: Dimensions.X - int(ui_scroll.scrollbar_width), Y: 0}, coords.CoordInts{X: int(ui_scroll.scrollbar_width), Y: Dimensions.Y - int(ui_scroll.scrollbar_width)*2})
	ui_scroll.Scrollbar_Vertical.SetVals(5, 1, -5, 10, 0)

	log.Printf(" CHILDREN ---- ")

	ui_scroll.Scrollbar_Horizontal.Init([]string{"scrollbar_horizontal", "X"}, backend, style, coords.CoordInts{X: 0, Y: Dimensions.Y - int(ui_scroll.scrollbar_width)}, coords.CoordInts{X: Dimensions.X - int(ui_scroll.scrollbar_width), Y: int(ui_scroll.scrollbar_width)})
	ui_scroll.Scrollbar_Horizontal.SetVals(5, 1, -5, 10, 0)
	ui_scroll.Scrollbar_Vertical.Init_Parents(ui_scroll)
	ui_scroll.Scrollbar_Horizontal.Init_Parents(ui_scroll)
	ui_scroll.Scrollbar_Vertical.Redraw()
	ui_scroll.Scrollbar_Horizontal.Redraw()

	ui_scroll.Redraw()
	return nil
}

/*
this is where basic images are going to be created and positioned;
*/
func (ui_scroll *UI_Scrollpane) Init_2(Display_Dimensions, Under_Image_Dimensions coords.CoordInts) error {

	return nil
}

/**/
func (ui_scroll *UI_Scrollpane) Init_Parents(parent UI_Object) error {
	ui_scroll.Parent = parent
	ui_scroll.Parent.AddChild(ui_scroll)
	ui_scroll.Redraw()
	ui_scroll.Parent.Redraw()
	return nil
}

/**/
func (ui_scroll *UI_Scrollpane) Draw(screen *ebiten.Image) error {
	if ui_scroll.IsVisible {
		ops := ebiten.DrawImageOptions{}
		scale := 1.0
		ops.GeoM.Reset()
		ops.GeoM.Translate(float64(ui_scroll.Position.X)*scale, float64(ui_scroll.Position.Y)*scale)
		screen.DrawImage(ui_scroll.Image, &ops)
	}
	return nil
}

/*
 */
func (ui_scroll *UI_Scrollpane) Redraw_Under() {
	ui_scroll.UnderImage.Fill(ui_scroll.Style.BorderColor)
	lineThick := ui_scroll.Style.BorderThickness
	vector.DrawFilledRect(ui_scroll.Image, lineThick, lineThick, float32(ui_scroll.Dim_Default.X)-lineThick*2, float32(ui_scroll.Dim_Default.Y)-lineThick*2, ui_scroll.Style.PanelColor, true) //
	if len(ui_scroll.Children) > 0 {
		for i, _ := range ui_scroll.Children {
			ui_scroll.Children[i].Draw(ui_scroll.UnderImage)
		}
	}
}

/*
 */
func (ui_scroll *UI_Scrollpane) Redraw() {
	// ui_scroll.DisplayImage.Fill(ui_scroll.Style.BorderColor)
	ui_scroll.DisplayImage.Fill(color.RGBA{255, 0, 0, 255})
	ui_scroll.Redraw_Under()
	scale := 1.0
	opts := ebiten.DrawImageOptions{}
	opts.GeoM.Translate(float64(ui_scroll.Internal_Position.X)/scale, float64(ui_scroll.Internal_Position.X)/scale)
	ui_scroll.DisplayImage.DrawImage(ui_scroll.UnderImage, &opts)
	ui_scroll.Scrollbar_Vertical.Draw(ui_scroll.DisplayImage)
	ui_scroll.Scrollbar_Horizontal.Draw(ui_scroll.DisplayImage)
	opts.GeoM.Reset()
	ui_scroll.Image.DrawImage(ui_scroll.DisplayImage, &opts)

	log.Printf("%d %d\n", ui_scroll.Scrollbar_Horizontal.CurrValue, ui_scroll.Scrollbar_Vertical.CurrValue)

	// ui_scroll.DisplayImage.Draw
}

/**/
func (ui_scroll *UI_Scrollpane) Update() error {
	if ui_scroll.IsActive && ui_scroll.IsVisible {
		Mouse_Pos_X, Mouse_Pos_Y := ebiten.CursorPosition()
		ui_scroll.Update_Ret_State_Redraw_Status_Mport(Mouse_Pos_X, Mouse_Pos_Y, 0)
	}
	return nil
}

/**/
func (ui_scroll *UI_Scrollpane) Update_Unactive() error { return nil }

/**/
func (ui_scroll *UI_Scrollpane) Update_Ret_State_Redraw_Status() (uint8, bool, error) {
	if ui_scroll.IsActive && ui_scroll.IsVisible {
		Mouse_Pos_X, Mouse_Pos_Y := ebiten.CursorPosition()
		// Mouspos-internalpos
		return ui_scroll.Update_Ret_State_Redraw_Status_Mport(Mouse_Pos_X, Mouse_Pos_Y, 0)
	}
	return 0, false, nil
}

/**/
func (ui_scroll *UI_Scrollpane) Update_Ret_State_Redraw_Status_Mport(Mouse_Pos_X, Mouse_Pos_Y, mode int) (uint8, bool, error) {

	if ui_scroll.IsCursorInBounds_MousePort(Mouse_Pos_X, Mouse_Pos_Y, mode) {
		// var status uint8
		// var trfals bool
		// var err error
		if len(ui_scroll.Children) > 0 {
			for i, _ := range ui_scroll.Children {
				name := ui_scroll.Children[i].GetID()
				ui_scroll.Children[i].Update_Ret_State_Redraw_Status_Mport(Mouse_Pos_X-ui_scroll.Internal_Position.X, Mouse_Pos_Y-ui_scroll.Internal_Position.Y, 0)
				if name == "scrollbar_vertical" {
					ui_scroll.Internal_Position.Y = ui_scroll.Scrollbar_Vertical.CurrValue
				} else if name == "scrollbar_horizontal" {
					ui_scroll.Internal_Position.X = ui_scroll.Scrollbar_Horizontal.CurrValue

				}
			}
		}
		log.Printf("%d %d\n", ui_scroll.Scrollbar_Horizontal.CurrValue, ui_scroll.Scrollbar_Vertical.CurrValue)
		ui_scroll.Redraw_Under()
		ui_scroll.Redraw()
	}
	return 0, false, nil
}

/**/
func (ui_scroll *UI_Scrollpane) GetState() uint8 {
	return ui_scroll.State
}

/**/
func (ui_scroll *UI_Scrollpane) ToString() string {
	return fmt.Sprintf("UI_SCROLL_PANE; AT %3d,%3d", ui_scroll.Position.X, ui_scroll.Position.Y)
}

/**/
func (ui_scroll *UI_Scrollpane) IsInit() bool { return true }

/**/
func (ui_scroll *UI_Scrollpane) GetID() string { return ui_scroll.ui_obj_id }

/**/
func (ui_scroll *UI_Scrollpane) GetType() string { return "ui_scroll_pane" }

// /**/
func (ui_scroll *UI_Scrollpane) IsCursorInBounds() bool {
	if ui_scroll.IsActive && ui_scroll.IsVisible {
		mouse_Pos_X, mouse_Pos_Y := ebiten.CursorPosition()
		return ui_scroll.IsCursorInBounds_MousePort(mouse_Pos_X, mouse_Pos_Y, 0)
	}
	return false
}

/**/
func (ui_scroll *UI_Scrollpane) IsCursorInBounds_MousePort(Mouse_Pos_X, Mouse_Pos_Y, mode int) bool {
	if ui_scroll.IsActive && ui_scroll.IsVisible {
		var x0, y0, x1, y1 int

		if ui_scroll.Parent != nil {
			px, py := ui_scroll.Parent.GetPosition_Int()
			x0 = ui_scroll.Position.X + px
			y0 = ui_scroll.Position.Y + py
			x1 = ui_scroll.Position.X + ui_scroll.DisplayImage.Bounds().Dx() + px
			y1 = ui_scroll.Position.Y + ui_scroll.DisplayImage.Bounds().Dy() + py
			// x0 = prim.Position.X + prim.ParentPos.X
			// y0 = prim.Position.Y + prim.ParentPos.X
			// x1 = prim.Position.X + prim.ParentPos.X + prim.Dimensions.X
			// y1 = prim.Position.Y + prim.ParentPos.Y + prim.Dimensions.Y
		} else {
			x0 = ui_scroll.Position.X
			y0 = ui_scroll.Position.Y
			x1 = ui_scroll.Position.X + ui_scroll.DisplayImage.Bounds().Dx()
			y1 = ui_scroll.Position.Y + ui_scroll.DisplayImage.Bounds().Dy()
		}
		return (Mouse_Pos_X > x0 && Mouse_Pos_X < x1) && (Mouse_Pos_Y > y0 && Mouse_Pos_Y < y1)
	}
	return false
}

/**/
func (ui_scroll *UI_Scrollpane) GetPosition_Int() (int, int) {
	return ui_scroll.Position.X, ui_scroll.Position.Y
}

/**/
func (ui_scroll *UI_Scrollpane) SetPosition_Int(x_point, y_point int) {
	ui_scroll.Position = coords.CoordInts{X: x_point, Y: y_point}
}

/**/
func (ui_scroll *UI_Scrollpane) GetDimensions_Int() (int, int) {
	return ui_scroll.Dim_Default.X, ui_scroll.Dim_Default.Y
}

/**/
func (ui_scroll *UI_Scrollpane) SetDimensions_Int(x_point, y_point int) {
	ui_scroll.Dim_Default = coords.CoordInts{X: x_point, Y: y_point}
	//---Redraw The S
}

/**/
func (ui_scroll *UI_Scrollpane) GetNumber_Children() int { return len(ui_scroll.Children) }

/**/
func (ui_scroll *UI_Scrollpane) GetChild(index int) UI_Object { return ui_scroll.Children[index] }

/**/
func (ui_scroll *UI_Scrollpane) AddChild(child UI_Object) error {

	ui_scroll.Children = append(ui_scroll.Children, child)
	log.Printf("UI_SCROLL CHILDREN %d\n", len(ui_scroll.Children))
	return nil
}

/**/
func (ui_scroll *UI_Scrollpane) RemoveChild(index int) error { return nil }

/**/
func (ui_scroll *UI_Scrollpane) GetParent() UI_Object {
	return ui_scroll.Parent
}

/**/
func (ui_scroll *UI_Scrollpane) HasParent() bool {
	return ui_scroll.Parent != nil
}
