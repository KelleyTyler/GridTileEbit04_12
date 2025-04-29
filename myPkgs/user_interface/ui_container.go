package user_interface

import (
	"github.com/KelleyTyler/GridTileEbit04_12/myPkgs/basic_geometry/coords"
	"github.com/hajimehoshi/ebiten/v2"
)

type UI_ButtonPanel struct {
	Position                    coords.CoordInts
	Dimensions                  coords.CoordInts
	ParentPos                   *coords.CoordInts
	Buttons                     []UI_Button //maybe pointers?
	ButtonOuts                  []bool
	IsVisible, IsActive         bool
	IsMoveable                  bool
	BackgroundImg               *ebiten.Image
	Style                       *UI_Object_Style
	btnspacing, internalmargins coords.CoordInts
	Backend                     *UI_Backend
	Label                       UI_Label
	//update_Image                bool
}

/**/
func (btnPanel *UI_ButtonPanel) Init(PanelName string, backend *UI_Backend, position, dimensions coords.CoordInts, ParentPos *coords.CoordInts) {
	btnPanel.Backend = backend
	btnPanel.Style = &backend.Style
	btnPanel.Position = position
	btnPanel.Dimensions = dimensions
	btnPanel.ParentPos = ParentPos
	btnPanel.BackgroundImg = ebiten.NewImage(dimensions.X, dimensions.Y)
	btnPanel.BackgroundImg.Fill(btnPanel.Style.PanelColor)
	btnPanel.Label = UI_Label{}
	btnPanel.Label.Init_00(backend, PanelName, coords.CoordInts{X: 0, Y: 0}, coords.CoordInts{X: dimensions.X, Y: 32}, &btnPanel.Position)
	btnPanel.internalmargins = coords.CoordInts{X: 4, Y: 4}
	btnPanel.btnspacing = coords.CoordInts{X: 1, Y: 1}

}

/**/
func (btnPanel *UI_ButtonPanel) AddButton(btnLabel string, relPos, dimen coords.CoordInts, btntype uint8) {
	tempbtn := UI_Button{}
	tempbtn.Init_00(btnPanel.Backend, btnLabel, relPos, dimen, uint8(btntype), nil)
	btnPanel.Buttons = append(btnPanel.Buttons, tempbtn)
	btnPanel.ButtonOuts = append(btnPanel.ButtonOuts, false)
}

/**/
func (btnPanel *UI_ButtonPanel) Update() {

	// for i, _ := range btnPanel.ButtonOuts {
	// 	btnPanel.ButtonOuts[i] = false
	// }
	// for i, btn := range btnPanel.Buttons {
	// 	btnPanel.ButtonOuts[i] = btn.Update_Ret()
	// }

	for i := 0; i < len(btnPanel.Buttons); i++ {

		// btnPanel.ButtonOuts[i] = btnPanel.Buttons[i].Update_Any()
	}
	btnPanel.Update_Drawing()
}

/**/
func (btnPanel *UI_ButtonPanel) Update_Drawing() {
	btnPanel.Label.Draw(btnPanel.BackgroundImg)
	for _, btn := range btnPanel.Buttons {
		btn.Draw(btnPanel.BackgroundImg)
	}
}

/**/
func (btnPanel *UI_ButtonPanel) Draw(screen *ebiten.Image) {

	ops := ebiten.DrawImageOptions{}
	ops.GeoM.Reset()
	scale := 1.0
	ops.GeoM.Translate(float64(btnPanel.Position.X)*scale, float64(btnPanel.Position.Y)*scale)
	screen.DrawImage(btnPanel.BackgroundImg, &ops)
}

/**/
func (btnPanel *UI_ButtonPanel) AddButtons_Row_TypeA(labels []string, InitialPosition, dimensions coords.CoordInts, Btntypes []uint8) {
	tempHSpace := dimensions.X + btnPanel.btnspacing.X

	for i := 0; i < len(labels); i++ {
		if i == 0 {
			btnPanel.AddButton(labels[i], InitialPosition, dimensions, Btntypes[i])
		} else {
			btnPanel.AddButton(labels[i], coords.CoordInts{X: InitialPosition.X + (tempHSpace * i), Y: InitialPosition.Y}, dimensions, Btntypes[i])
		}
	}
}
