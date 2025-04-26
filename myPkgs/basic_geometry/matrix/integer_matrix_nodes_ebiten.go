package matrix

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

/*
	The focus of this file is to have ebiten ends to work


*/

//------------

func (imat *IntegerMatrix2D) NodeList_DrawList(screen *ebiten.Image, nodes []*ImatNode, colorArray []color.Color, options *Integer_Matrix_Ebiten_DrawOptions) (err error) {

	var maxColors int
	var colors []color.Color

	if colorArray != nil {
		maxColors = len(colorArray)
		colors = make([]color.Color, len(colorArray))
		copy(colors, colorArray)
	}
	if maxColors == 0 {
		colors = append(colors, color.RGBA{0, 128, 0, 255})
	}
	var currColor int = 0

	for _, node := range nodes {
		imat.DrawAGridTile_With_Lines(screen, node.Position, colors[currColor], options)
	}
	return err
}

/*
 */
func (imat *IntegerMatrix2D) ImatNode_Draw(screen *ebiten.Image, node *ImatNode, colorArray []color.Color, options *Integer_Matrix_Ebiten_DrawOptions) (err error) {
	temp := node.GetTail()

	maxColors := len(colorArray)
	colors := make([]color.Color, len(colorArray))
	copy(colors, colorArray)
	if maxColors == 0 {
		colors = append(colors, color.RGBA{0, 128, 0, 255})
	}
	var currColor int = 0
	for temp.Prev != nil {
		currColor++
		if currColor > maxColors-1 {
			currColor = 0
		}
		imat.DrawAGridTile_With_Lines(screen, temp.Position, colors[currColor], options)
		temp = temp.Prev
	}
	return err
}
