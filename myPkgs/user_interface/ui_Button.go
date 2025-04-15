package user_interface

import (
	"image/color"

	"github.com/KelleyTyler/GridTileEbit04_12/myPkgs/basic_geometry/coords"
	"github.com/hajimehoshi/ebiten"
)

/*
UI_Button_mk0
*/
type UI_Button struct {
	Backend *UI_Backend
	// GSettings  *settings.GameSettings
	Position            coords.CoordInts
	Parentpos           *coords.CoordInts //this changes the position of the thing depending on where it's parent is;
	Dimensions          coords.CoordInts
	ActionPtr           *func() //<---big test if this works
	Btn_Type            int
	State               int
	Colors              []color.Color
	IsActive, IsVisible bool
}

func (btn UI_Button) Init(bckend *UI_Backend, pos, dimen coords.CoordInts, bType int, parent *coords.CoordInts, action *func()) {
	btn.Backend = bckend
	btn.Parentpos = parent
	btn.Position = pos
	btn.Dimensions = dimen
	btn.Btn_Type = bType
	btn.ActionPtr = action
	// btn.Colors=
}

func (btn UI_Button) Draw() {

}
func (btn *UI_Button) Update() {
	if btn.IsActive && btn.IsVisible {

	}
}
func (btn *UI_Button) IsCursorInBounds() bool {
	if btn.IsVisible && btn.IsActive {
		xx, yy := ebiten.CursorPosition()
		if btn.Parentpos != nil {
			return !(xx > btn.Position.X+btn.Parentpos.X && xx < btn.Position.X+btn.Parentpos.X+btn.Dimensions.X) && !(yy > btn.Position.Y+btn.Parentpos.Y && yy < btn.Position.Y+btn.Parentpos.Y+btn.Dimensions.Y)

		} else {
			return !(xx > btn.Position.X && xx < btn.Position.X+btn.Dimensions.X) && !(yy > btn.Position.Y && yy < btn.Position.Y+btn.Dimensions.Y)
		}
	}
	return false
	// if()
}
