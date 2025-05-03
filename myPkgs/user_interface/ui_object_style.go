package user_interface

import "image/color"

type UI_Object_Style struct {
	LabelColor       color.Color
	PanelColor       color.Color
	BorderColor      color.Color
	ButtonColor0     []color.Color
	ButtonColor1     []color.Color
	ButtonColor2     []color.Color
	BorderThickness  float32
	TextColor        []color.Color
	Internal_Margins [4]uint8
	TextSizes        []int
	TextAlignMode    int

	Panel_Child_Internal_Border_Offset [4]uint8 //the buffer between the border of the parent and the child element
	Child_Buffer                       [4]uint8 //the buffer between child elements
}

/*
type UI_Object_Type uint

const (

)
*/

// var (
// 	COLOR_RED         color.RGBA = color.RGBA{255, 0, 0, 255}
// 	COLOR_ORANGE      color.RGBA = color.RGBA{255, 128, 0, 255}
// 	COLOR_YELLOW      color.RGBA = color.RGBA{255, 255, 0, 255}
// 	COLOR_LIGHT_GREEN color.RGBA = color.RGBA{128, 255, 0, 255}
// 	COLOR_GREEN       color.RGBA = color.RGBA{0, 255, 0, 255}
// 	COLOR_LIGHT_TEAL  color.RGBA = color.RGBA{0, 255, 128, 255}
// 	COLOR_CYAN        color.RGBA = color.RGBA{0, 255, 255, 255}
// 	COLOR_AZURE       color.RGBA = color.RGBA{0, 128, 255, 255}
// 	COLOR_BLUE        color.RGBA = color.RGBA{0, 0, 255, 255}
// 	COLOR_VIOLET      color.RGBA = color.RGBA{128, 0, 255, 255}
// 	COLOR_MAGENTA     color.RGBA = color.RGBA{255, 0, 255, 255}

// 	COLOR_DARK_BROWN color.RGBA = color.RGBA{82, 50, 0, 255}
// 	COLOR_DARK_SAND  color.RGBA = color.RGBA{97, 78, 48, 255}
// 	COLOR_LIGHT_SAND color.RGBA = color.RGBA{202, 171, 121, 255}
// )

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
	out_Style.ButtonColor1 = []color.Color{color.RGBA{150, 130, 130, 255}, color.RGBA{165, 130, 145, 255}, color.RGBA{175, 130, 105, 255}}
	out_Style.ButtonColor2 = []color.Color{color.RGBA{150, 130, 130, 255}, color.RGBA{150, 130, 145, 255}, color.RGBA{150, 130, 105, 255}}
	out_Style.TextAlignMode = 10
	out_Style.Panel_Child_Internal_Border_Offset = [4]uint8{4, 4, 4, 4}
	out_Style.Child_Buffer = [4]uint8{4, 4, 4, 4}

	return out_Style
}

/**/
func Get_Color_Array(leng int, color_0 color.RGBA) (ColorAr []color.Color) {

	for i := 0; i < leng; i++ {
		temp0 := float64(i) / float64(leng)
		tR := uint8(float64(color_0.R) * temp0)
		tG := uint8(float64(color_0.G) * temp0)
		tB := uint8(float64(color_0.B) * temp0)
		tA := uint8(float64(color_0.A) * temp0)
		ColorAr = append(ColorAr, color.RGBA{tR, tG, tB, tA})
	}
	return ColorAr
}
