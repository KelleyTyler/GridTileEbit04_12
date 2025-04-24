package user_interface

import (
	"fmt"
	"log"

	"github.com/KelleyTyler/GridTileEbit04_12/myPkgs/basic_geometry/coords"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

/*

The idea here is to have buttons that are 'compound' or 'combo' buttons

that is number selection buttons, drop down menus, etc.
*/

type UI_Num_Select struct {
	Parent                                                 UI_Object
	Position, Dimensions                                   coords.CoordInts
	Backend                                                *UI_Backend
	Style                                                  *UI_Object_Style
	obj_id, text                                           string
	State                                                  uint8
	Image                                                  *ebiten.Image
	ImageUpdate                                            bool
	init                                                   bool
	IsActive, IsVisible                                    bool
	R_Button                                               UI_Button
	M_Button                                               UI_Button
	L_Button                                               UI_Button
	Label                                                  UI_Label
	CurrValue, MinValue, MaxValue, DefaultValue, IterValue int
	MiddleButtonMode                                       uint8
}

/*
 */
func (nSelect *UI_Num_Select) Init(idLabels []string, backend *UI_Backend, style *UI_Object_Style, Position, Dimensions coords.CoordInts) error {
	nSelect.obj_id = idLabels[0]
	nSelect.text = idLabels[1]
	nSelect.Dimensions = Dimensions
	nSelect.Position = Position
	nSelect.Backend = backend
	if style != nil {
		nSelect.Style = style
	} else {
		nSelect.Style = &nSelect.Backend.Style
	}
	nSelect.State = 0
	nSelect.Image = ebiten.NewImage(Dimensions.X, Dimensions.Y)
	dimY := (nSelect.Dimensions.Y / 2)
	thick := int(nSelect.Style.BorderThickness)
	btnwidth := 16
	nSelect.L_Button.Init([]string{"lbtn", "<"}, backend, nil, coords.CoordInts{X: (thick / 2), Y: dimY - (thick / 2)}, coords.CoordInts{X: btnwidth, Y: dimY})
	nSelect.M_Button.Init([]string{"lbtn", "000"}, backend, nil, coords.CoordInts{X: btnwidth + (thick / 2), Y: dimY - (thick / 2)}, coords.CoordInts{X: nSelect.Dimensions.X - ((btnwidth * 2) + (thick)), Y: dimY})
	nSelect.R_Button.Init([]string{"lbtn", ">"}, backend, nil, coords.CoordInts{X: nSelect.Dimensions.X - (btnwidth + thick - 1), Y: dimY - (thick / 2)}, coords.CoordInts{X: btnwidth, Y: dimY})
	nSelect.Label.Init([]string{"lbtn", idLabels[1]}, backend, nil, coords.CoordInts{X: thick / 2, Y: thick / 2}, coords.CoordInts{X: nSelect.Dimensions.X - (thick), Y: dimY})
	nSelect.Label.TextAlignMode = 10
	nSelect.L_Button.Init_Parents(nSelect)
	nSelect.M_Button.Init_Parents(nSelect)
	nSelect.R_Button.Init_Parents(nSelect)

	nSelect.Label.Init_Parents(nSelect)
	nSelect.Label.Redraw()
	nSelect.SetVals(0, 1, -10, 10, 0)
	// nSelect.R_Button.Init_00(backend, "->", coords.CoordInts{X: nSelect.Dimensions.X - 16, Y: 0}, coords.CoordInts{X: 16, Y: 16}, 0, nSelect)
	//-------Setting up Image
	nSelect.Redraw()
	nSelect.ImageUpdate = true

	//------Finishing Up
	if !nSelect.init {
		nSelect.init = true
	}
	return nil
}

/*
 */
func (nSelect *UI_Num_Select) Init_Parents(parent UI_Object) error {
	nSelect.Parent = parent
	nSelect.Parent.AddChild(nSelect)
	nSelect.Redraw()
	nSelect.Parent.Redraw()
	return nil
}

/*
 */
func (nSelect *UI_Num_Select) SetVals(defVal, interator, min, max int, middlebtn uint8) {
	nSelect.MaxValue = max
	nSelect.MinValue = min
	nSelect.DefaultValue = defVal
	nSelect.CurrValue = defVal
	nSelect.IterValue = interator
	nSelect.MiddleButtonMode = middlebtn
}

/*
 */
func (nSelect *UI_Num_Select) Draw(screen *ebiten.Image) error {
	ops := ebiten.DrawImageOptions{}
	scale := 1.0
	ops.GeoM.Reset()
	ops.GeoM.Translate(float64(nSelect.Position.X)*scale, float64(nSelect.Position.Y)*scale)
	screen.DrawImage(nSelect.Image, &ops)
	return nil
}

/*
 */
func (nSelect *UI_Num_Select) Redraw() {

	nSelect.M_Button.Redraw()
	// nSelect.Image.Fill(color.RGBA{255, 0, 0, 255}) //nSelect.Style.BorderColor
	nSelect.Image.Fill(nSelect.Style.BorderColor) //

	lineThick := nSelect.Style.BorderThickness
	// vector.DrawFilledRect(nSelect.Image, lineThick, lineThick, float32(nSelect.Dimensions.X)-(2*lineThick), float32(nSelect.Dimensions.Y)-(2*lineThick), color.RGBA{255, 255, 0, 255}, true) //nSelect.Style.PanelColor
	vector.DrawFilledRect(nSelect.Image, lineThick, lineThick, float32(nSelect.Dimensions.X)-(2*lineThick), float32(nSelect.Dimensions.Y)-(2*lineThick), nSelect.Style.PanelColor, true) //

	nSelect.L_Button.Draw(nSelect.Image)
	nSelect.M_Button.Draw(nSelect.Image)
	nSelect.R_Button.Draw(nSelect.Image)
	nSelect.Label.Draw(nSelect.Image)
}

/*
 */
func (nSelect *UI_Num_Select) Update() error {
	nSelect.State = 0
	Mouse_Pos_X, Mouse_Pos_Y := ebiten.CursorPosition()
	// var err error
	state0, to_redraw0, err0 := nSelect.L_Button.Update_Ret_State_Redraw_Status_Mport(Mouse_Pos_X, Mouse_Pos_Y, 0)
	if err0 != nil {
		log.Fatal(err0)
	}

	if state0 == 2 {
		// fmt.Printf("BTN L is Clicked %d %t\n", nSelect.CurrValue, nSelect.CurrValue != nSelect.DefaultValue)
		if nSelect.CurrValue > nSelect.MinValue {
			nSelect.CurrValue -= nSelect.IterValue
		}

	}
	state1, to_redraw1, err1 := nSelect.M_Button.Update_Ret_State_Redraw_Status_Mport(Mouse_Pos_X, Mouse_Pos_Y, 0)
	if err1 != nil {
		log.Fatal(err1)
	}

	if state1 == 2 {
		// fmt.Printf("BTN M is Clicked %d  %d %d %t\n", nSelect.CurrValue, nSelect.IterValue, nSelect.MinValue, nSelect.CurrValue != nSelect.DefaultValue)
		if nSelect.MiddleButtonMode == 0 {
			if nSelect.CurrValue != nSelect.DefaultValue {
				nSelect.CurrValue = nSelect.DefaultValue
			}
		} else {
			nSelect.State = 2
		}
	}
	state2, to_redraw2, err2 := nSelect.R_Button.Update_Ret_State_Redraw_Status_Mport(Mouse_Pos_X, Mouse_Pos_Y, 0)
	if err2 != nil {
		log.Fatal(err2)
	}

	if state2 == 2 {

		if nSelect.CurrValue < nSelect.MaxValue {
			nSelect.CurrValue = nSelect.CurrValue + nSelect.IterValue
			//fmt.Printf("BTN R is Clicked %d %t\n", nSelect.CurrValue, nSelect.CurrValue < nSelect.MaxValue)
		}
	}

	// if nSelect.IsMovable {
	// 	if nSelect.IsCursorInBounds() {

	// 	}
	// }
	if to_redraw0 || to_redraw1 || to_redraw2 {
		nSelect.M_Button.Label = fmt.Sprintf("%d", nSelect.CurrValue)
		nSelect.Redraw()
	}
	return nil
}

/*
 */
func (nSelect *UI_Num_Select) Update_Unactive() error {

	return nil
}

/*
This will return false; Use Only Sparingly!
*/
func (nSelect *UI_Num_Select) Update_Any() (any, error) {
	return false, nil
}

/*
 */
func (nSelect *UI_Num_Select) Update_Ret_State_Redraw_Status() (uint8, bool, error) {
	nSelect.State = 0
	Mouse_Pos_X, Mouse_Pos_Y := ebiten.CursorPosition()
	// var err error
	state0, to_redraw0, err0 := nSelect.L_Button.Update_Ret_State_Redraw_Status_Mport(Mouse_Pos_X, Mouse_Pos_Y, 0)
	if err0 != nil {
		log.Fatal(err0)
	}

	if state0 == 2 {
		// fmt.Printf("BTN L is Clicked %d %t\n", nSelect.CurrValue, nSelect.CurrValue != nSelect.DefaultValue)

		if nSelect.CurrValue > nSelect.MinValue {
			nSelect.CurrValue -= nSelect.IterValue
		}
	}
	state1, to_redraw1, err1 := nSelect.M_Button.Update_Ret_State_Redraw_Status_Mport(Mouse_Pos_X, Mouse_Pos_Y, 0)
	if err1 != nil {
		log.Fatal(err1)
	}

	if state1 == 2 {
		// fmt.Printf("BTN M is Clicked %d  %d %d %t\n", nSelect.CurrValue, nSelect.IterValue, nSelect.MinValue, nSelect.CurrValue != nSelect.DefaultValue)
		if nSelect.MiddleButtonMode == 0 {
			if nSelect.CurrValue != nSelect.DefaultValue {
				nSelect.CurrValue = nSelect.DefaultValue
			}
		} else {
			nSelect.State = 2
		}
	}
	state2, to_redraw2, err2 := nSelect.R_Button.Update_Ret_State_Redraw_Status_Mport(Mouse_Pos_X, Mouse_Pos_Y, 0)
	if err2 != nil {
		log.Fatal(err2)
	}

	if state2 == 2 {

		if nSelect.CurrValue < nSelect.MaxValue {
			nSelect.CurrValue = nSelect.CurrValue + nSelect.IterValue
			//fmt.Printf("BTN R is Clicked %d %t\n", nSelect.CurrValue, nSelect.CurrValue < nSelect.MaxValue)
		}
	}

	// if nSelect.IsMovable {
	// 	if nSelect.IsCursorInBounds() {

	// 	}
	// }
	export_redraw := false
	if to_redraw0 || to_redraw1 || to_redraw2 {
		nSelect.M_Button.Label = fmt.Sprintf("%03d", nSelect.CurrValue)
		nSelect.Redraw()
		export_redraw = true
	}

	return nSelect.State, export_redraw, nil
}

/*
 */
func (nSelect *UI_Num_Select) SetPosition(Position coords.CoordInts) { nSelect.Position = Position }

/*
Update_Ret_State_Redraw_Status
*/
func (nSelect *UI_Num_Select) Update_Ret_State_Redraw_Status_Mport(Mouse_Pos_X, Mouse_Pos_Y, mode int) (uint8, bool, error) {
	nSelect.State = 0
	state0, to_redraw0, err0 := nSelect.L_Button.Update_Ret_State_Redraw_Status_Mport(Mouse_Pos_X, Mouse_Pos_Y, mode)
	if err0 != nil {
		log.Fatal(err0)
	}

	if state0 == 2 {
		// fmt.Printf("BTN L is Clicked %d %t\n", nSelect.CurrValue, nSelect.CurrValue != nSelect.DefaultValue)

		if nSelect.CurrValue > nSelect.MinValue {
			nSelect.CurrValue -= nSelect.IterValue
		}
	}
	state1, to_redraw1, err1 := nSelect.M_Button.Update_Ret_State_Redraw_Status_Mport(Mouse_Pos_X, Mouse_Pos_Y, mode)
	if err1 != nil {
		log.Fatal(err1)
	}

	if state1 == 2 {
		// fmt.Printf("BTN M is Clicked %d  %d %d %t\n", nSelect.CurrValue, nSelect.IterValue, nSelect.MinValue, nSelect.CurrValue != nSelect.DefaultValue)
		if nSelect.MiddleButtonMode == 0 {
			if nSelect.CurrValue != nSelect.DefaultValue {
				nSelect.CurrValue = nSelect.DefaultValue
			}
		} else {
			nSelect.State = 2
		}
	}
	state2, to_redraw2, err2 := nSelect.R_Button.Update_Ret_State_Redraw_Status_Mport(Mouse_Pos_X, Mouse_Pos_Y, mode)
	if err2 != nil {
		log.Fatal(err2)
	}

	if state2 == 2 {

		if nSelect.CurrValue < nSelect.MaxValue {
			nSelect.CurrValue = nSelect.CurrValue + nSelect.IterValue
			//fmt.Printf("BTN R is Clicked %d %t\n", nSelect.CurrValue, nSelect.CurrValue < nSelect.MaxValue)
		}
	}

	// if nSelect.IsMovable {
	// 	if nSelect.IsCursorInBounds() {

	// 	}
	// }
	export_redraw := false
	if to_redraw0 || to_redraw1 || to_redraw2 {
		nSelect.M_Button.Label = fmt.Sprintf("%d", nSelect.CurrValue)
		nSelect.Redraw()
		export_redraw = true
	}

	return nSelect.State, export_redraw, nil
}

/*
This returns the state of the object
*/
func (nSelect *UI_Num_Select) GetState() uint8 { return nSelect.State }

/*
this returns a basic to string message
*/
func (nSelect *UI_Num_Select) ToString() string {
	strngOut := fmt.Sprintf("UI_Object nSelect:%s\n\tPositon %s\t", nSelect.obj_id, nSelect.Position.ToString())
	strngOut += fmt.Sprintf("\tDimensions: %s\n", nSelect.Dimensions.ToString())
	return strngOut
}

/*
 */
func (nSelect *UI_Num_Select) IsCursorInBounds() bool {
	if nSelect.IsActive && nSelect.IsVisible {
		cX, cY := ebiten.CursorPosition()
		var x0, y0, x1, y1 int

		if nSelect.Parent != nil {
			px, py := nSelect.Parent.GetPosition_Int()
			x0 = nSelect.Position.X + px
			y0 = nSelect.Position.Y + py
			x1 = nSelect.Position.X + nSelect.Dimensions.X + px
			y1 = nSelect.Position.Y + nSelect.Dimensions.Y + py
			// x0 = nSelect.Position.X + nSelect.ParentPos.X
			// y0 = nSelect.Position.Y + nSelect.ParentPos.X
			// x1 = nSelect.Position.X + nSelect.ParentPos.X + nSelect.Dimensions.X
			// y1 = nSelect.Position.Y + nSelect.ParentPos.Y + nSelect.Dimensions.Y
		} else {
			x0 = nSelect.Position.X
			y0 = nSelect.Position.Y
			x1 = nSelect.Position.X + nSelect.Dimensions.X
			y1 = nSelect.Position.Y + nSelect.Dimensions.Y
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
func (nSelect *UI_Num_Select) IsCursorInBounds_MousePort(Mouse_Pos_X, Mouse_Pos_Y, mode int) bool {
	if nSelect.IsActive && nSelect.IsVisible && mode == 0 {
		cX, cY := Mouse_Pos_X, Mouse_Pos_Y
		//mode stuff
		var x0, y0, x1, y1 int

		if nSelect.Parent != nil {
			px, py := nSelect.Parent.GetPosition_Int()
			x0 = nSelect.Position.X + px
			y0 = nSelect.Position.Y + py
			x1 = nSelect.Position.X + nSelect.Dimensions.X + px
			y1 = nSelect.Position.Y + nSelect.Dimensions.Y + py
		} else {
			x0 = nSelect.Position.X
			y0 = nSelect.Position.Y
			x1 = nSelect.Position.X + nSelect.Dimensions.X
			y1 = nSelect.Position.Y + nSelect.Dimensions.Y
		}
		return (cX > x0 && cX < x1) && (cY > y0 && cY < y1)
	}
	//mode stuff
	return false
}

/*
 */
func (nSelect *UI_Num_Select) GetPosition_Int() (int, int) {
	xx := nSelect.Position.X
	yy := nSelect.Position.Y
	if nSelect.Parent != nil {
		px, py := nSelect.Parent.GetPosition_Int()
		xx += px
		yy += py
	}
	return xx, yy
}

/*
this confirms the object is initilaized
*/
func (nSelect *UI_Num_Select) IsInit() bool { return nSelect.init }

/*
this gets the object ID
*/
func (nSelect *UI_Num_Select) GetID() string { return nSelect.obj_id }

/*
This returns a string specifying the objects type
*/
func (nSelect *UI_Num_Select) GetType() string { return "UI_Object nSelect" }

/*
 */
func (nSelect *UI_Num_Select) GetNumber_Children() int { return 0 }

/*
 */
func (nSelect *UI_Num_Select) GetChild(index int) UI_Object { return nil }

/*
 */
func (nSelect *UI_Num_Select) AddChild(child UI_Object) error { return nil }

/**/
func (nSelect *UI_Num_Select) RemoveChild(index int) error { return nil }

/**/
func (nSelect *UI_Num_Select) HasParent() bool { return nSelect.Parent != nil }

/**/
func (nSelect *UI_Num_Select) GetParent() UI_Object { return nSelect.Parent }
