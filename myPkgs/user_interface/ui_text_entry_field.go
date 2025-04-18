package user_interface

import (
	"image/color"

	"github.com/KelleyTyler/GridTileEbit04_12/myPkgs/basic_geometry/coords"
	settings "github.com/KelleyTyler/GridTileEbit04_12/myPkgs/settingsconfig"
	"github.com/hajimehoshi/ebiten/v2"
)

type TextEntryField struct {
	Position            coords.CoordInts
	ParentPos           *coords.CoordInts
	Dimensions          coords.CoordInts
	Img, SubImg         *ebiten.Image
	BG_color            color.Color
	TextColor           color.Color
	DataField           string
	IsActive, IsVisible bool
	WrapAround          bool
	TextSize            uint8
	MaxTextPerLine      uint16
	CurrenLine          uint16
	MaxLineOfText       uint16
	Scale               float32
}

type UI_TextEntryWindow struct {
	Position                coords.CoordInts
	Dimensions              coords.CoordInts
	ParentPos               *coords.CoordInts
	CloseButton             UI_Button
	ClearButton             UI_Button
	SubmitButton            UI_Button
	TexField                TextEntryField
	IsActive, IsVisible     bool
	IsResizable, IsMoveable bool
	PanelImage              *ebiten.Image
	Settings                *settings.GameSettings
	Backend                 *UI_Backend
	Style                   *UI_Object_Style
	WindowLabel             string
	//-------------------------Needs Content

}
