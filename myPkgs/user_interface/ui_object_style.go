package user_interface

import "image/color"

type UI_Object_Style struct {
	LabelColor       color.Color
	PanelColor       color.Color
	BorderColor      color.Color
	ButtonColor0     []color.Color
	ButtonColor1     []color.Color
	BorderThickness  float32
	TextColor        []color.Color
	Internal_Margins [4]uint8
	TextSizes        []int
	TextAlignMode    int
}

/*
	This function returns a preselection of various styles
*/
func Get_UI_Object_Style(styleNumber int) (out_Style UI_Object_Style) {
	out_Style.BorderColor = color.RGBA{50, 80, 90, 255} //25 40 45 //50 80 90
	out_Style.PanelColor = color.RGBA{222, 229, 232, 255}
	out_Style.LabelColor = color.RGBA{76, 132, 151, 255}
	out_Style.TextColor = []color.Color{color.White, color.Black}
	out_Style.TextSizes = []int{10, 15, 20, 25}
	out_Style.BorderThickness = 3.0
	out_Style.ButtonColor0 = []color.Color{color.RGBA{75, 150, 150, 255}, color.RGBA{85, 170, 170, 255}, color.RGBA{95, 190, 190, 255}}
	out_Style.ButtonColor1 = []color.Color{color.RGBA{150, 130, 130, 255}, color.RGBA{150, 130, 145, 255}, color.RGBA{150, 130, 105, 255}}
	out_Style.TextAlignMode = 10
	return out_Style
}
