package matrix

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

/*
	The focus of this file is to have ebiten ends to work


*/

//------------
/**/
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
		if node.Prev != nil {
			x0, y0 := options.GetCoordCenter(node.Position)
			x1, y1 := options.GetCoordCenter(node.Prev.Position)
			vector.StrokeLine(screen, float32(x0), float32(y0), float32(x1), float32(y1), 3.0, color.RGBA{25, 25, 25, 255}, true)
			vector.StrokeLine(screen, float32(x0), float32(y0), float32(x1), float32(y1), 1.5, color.RGBA{255, 255, 25, 255}, true)

		}
		temp := node
		for temp.Prev != nil {
			x0, y0 := options.GetCoordCenter(temp.Position)
			x1, y1 := options.GetCoordCenter(temp.Prev.Position)
			vector.StrokeLine(screen, float32(x0), float32(y0), float32(x1), float32(y1), 3.0, color.RGBA{0, 255, 0, 255}, true)
			temp = temp.Prev
		}
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
		x0, y0 := options.GetCoordCenter(temp.Position)
		x1, y1 := options.GetCoordCenter(temp.Prev.Position)

		imat.DrawAGridTile_With_Lines(screen, temp.Position, colors[currColor], options)
		vector.StrokeLine(screen, float32(x0), float32(y0), float32(x1), float32(y1), 3.0, color.RGBA{255, 0, 0, 255}, true)
		if temp.Next != nil {
			x1, y1 = options.GetCoordCenter(temp.Next.Position)
			vector.StrokeLine(screen, float32(x0), float32(y0), float32(x1), float32(y1), 1.5, color.RGBA{255, 0, 0, 255}, true)

		}
		temp = temp.Prev
	}
	return err
}
