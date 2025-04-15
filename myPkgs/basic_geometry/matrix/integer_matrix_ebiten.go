package matrix

import (
	"image/color"

	coords "github.com/KelleyTyler/GridTileEbit04_12/myPkgs/basic_geometry/coords"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

/*
	this is to allow for a connection to the ebitengine frontened;
*/

type Integer_Matrix_Ebiten_DrawOptions struct {
	TileSize    coords.CoordInts
	TileSpacing coords.CoordInts

	BoardMargin   coords.CoordInts
	BoardPosition coords.CoordInts
	//------------------------------------

	ShowTileLines     []bool
	TileLineColors    []color.Color
	TileLineThickness []float32
	AABody            bool
	AALines           bool
}

/*
ShowTileLines     []bool
	TileLineColors    []color.Color
	TileLineThickness []float32
	AABody            bool
	AALines           bool
*/

func (imat IntegerMatrix2D) DrawAGridTile_With_Lines(screen *ebiten.Image, coord coords.CoordInts, clr0 color.Color, options *Integer_Matrix_Ebiten_DrawOptions) {
	vector.DrawFilledRect(screen, float32((options.TileSize.X*coord.X)+(options.TileSpacing.X*coord.X)+options.BoardMargin.X), float32((options.TileSize.Y*coord.Y)+(options.TileSpacing.Y*coord.Y)+options.BoardMargin.Y), float32(options.TileSize.X), float32(options.TileSize.Y), clr0, options.AABody)
	TileCoord := coords.CoordInts{X: options.TileSize.X * coord.X, Y: options.TileSize.Y * coord.Y}
	// TileSpace := CoordInts{X: options.TileSpacing.X * coord.X, Y: options.TileSpacing.Y * coord.Y}
	if options.ShowTileLines[0] {
		vector.StrokeRect(screen, float32((options.TileSize.X*coord.X)+(options.TileSpacing.X*coord.X)+options.BoardMargin.X), float32((options.TileSize.Y*coord.Y)+(options.TileSpacing.Y*coord.Y)+options.BoardMargin.Y), float32(options.TileSize.X), float32(options.TileSize.Y), options.TileLineThickness[0], options.TileLineColors[0], options.AALines)
	}
	if options.ShowTileLines[1] {
		vector.StrokeLine(screen, float32((TileCoord.X)+(options.TileSpacing.X*coord.X)+options.BoardMargin.X), float32((TileCoord.Y)+(options.TileSpacing.Y*coord.Y)+options.BoardMargin.Y), float32((options.TileSize.X*coord.X)+(options.TileSpacing.X*coord.X)+options.BoardMargin.X+options.TileSize.X), float32((options.TileSize.Y*coord.Y)+(options.TileSpacing.Y*coord.Y)+options.BoardMargin.Y+options.TileSize.Y), options.TileLineThickness[0], options.TileLineColors[1], options.AALines)
	}
	if options.ShowTileLines[2] {
		vector.StrokeLine(screen, float32((TileCoord.X)+(options.TileSpacing.X*coord.X)+options.BoardMargin.X), float32((TileCoord.Y)+(options.TileSpacing.Y*coord.Y)+options.BoardMargin.Y+options.TileSize.Y), float32((options.TileSize.X*coord.X)+(options.TileSpacing.X*coord.X)+options.BoardMargin.X+options.TileSize.X), float32((options.TileSize.Y*coord.Y)+(options.TileSpacing.Y*coord.Y)+options.BoardMargin.Y), options.TileLineThickness[0], options.TileLineColors[2], options.AALines)
	}
}
func (imat IntegerMatrix2D) DrawCoordListWithLines(screen *ebiten.Image, cList coords.CoordList, colors []color.Color, options Integer_Matrix_Ebiten_DrawOptions) {
	if len(colors) == 1 {
		for _, a := range cList {
			imat.DrawAGridTile_With_Lines(screen, a, colors[0], &options)
		}
	}
}

func (imat IntegerMatrix2D) DrawFullGridFromColors(screen *ebiten.Image, colors []color.Color, showBoardOL bool, options *Integer_Matrix_Ebiten_DrawOptions) {
	test1X := ((len(imat[0]) * options.TileSize.X) + (len(imat[0]) * options.TileSpacing.X)) + options.BoardMargin.X
	test1Y := ((len(imat) * options.TileSize.Y) + (len(imat) * options.TileSpacing.Y)) + options.BoardMargin.Y
	for y, _ := range imat {
		for x, b := range imat[y] {
			if b < len(colors) {
				vector.DrawFilledRect(screen, float32((options.TileSize.X*x)+(options.TileSpacing.X*x)+options.BoardMargin.X), float32((options.TileSize.Y*y)+(options.TileSpacing.Y*y)+options.BoardMargin.Y), float32(options.TileSize.X), float32(options.TileSize.Y), colors[b], options.AABody)

			} else {
				vector.DrawFilledRect(screen, float32((options.TileSize.X*x)+(options.TileSpacing.X*x)+options.BoardMargin.X), float32((options.TileSize.Y*y)+(options.TileSpacing.Y*y)+options.BoardMargin.Y), float32(options.TileSize.X), float32(options.TileSize.Y), colors[0], options.AABody)

			}
			// color0 := color.RGBA{12, 12, 12, 100} //color.Black
			if options.ShowTileLines[0] {
				vector.StrokeRect(screen, float32((options.TileSize.X*x)+(options.TileSpacing.X*x)+options.BoardMargin.X), float32((options.TileSize.Y*y)+(options.TileSpacing.Y*y)+options.BoardMargin.Y), float32(options.TileSize.X), float32(options.TileSize.Y), options.TileLineThickness[0], options.TileLineColors[0], options.AALines)
			}
		}
	}
	//vector.StrokeRect(screen, float32(OffsetX-0), float32(OffsetY-0), float32(test1X-0), float32(test1Y+0), 2.0, color.RGBA{210, 153, 100, 255}, true) //0, 179, 100, 255
	if showBoardOL {
		vector.StrokeRect(screen, float32(options.BoardMargin.X-3), float32(options.BoardMargin.Y-3), float32(test1X-options.BoardMargin.X-options.TileSpacing.X+6), float32(test1Y-options.BoardMargin.Y-options.TileSpacing.Y+6), 2.0, color.Black, true) //color.RGBA{0, 50, 50, 255}

	}
}

/*
this translates a mouseclick onto the screen;
*/
func (imat IntegerMatrix2D) GetCoordOfMouseEvent(Raw_Mouse_X, Raw_Mouse_Y, offsetX, offsetY int, options Integer_Matrix_Ebiten_DrawOptions) (mXi int, mYi int, isOnTile bool) {
	mXi = -1
	mYi = -1
	isOnTile = false
	adj_MouseX := Raw_Mouse_X - options.BoardPosition.X
	adj_MouseY := Raw_Mouse_Y - options.BoardPosition.Y
	test1X := ((len(imat[0]) * options.TileSize.X) + (len(imat[0]) * options.TileSpacing.X)) + offsetX
	test1Y := ((len(imat) * options.TileSize.Y) + (len(imat) * options.TileSpacing.Y)) + offsetY
	//----
	if ((adj_MouseX > offsetX) && adj_MouseX < test1X-options.TileSpacing.X) && (adj_MouseY > offsetY && adj_MouseY < test1Y-options.TileSpacing.Y) {
		var mX = float32(adj_MouseX-offsetX) / float32(options.TileSize.X+options.TileSpacing.X) //float32(test1X)
		var mY = float32(adj_MouseY-offsetY) / float32(options.TileSize.Y+options.TileSpacing.Y) //float32(test1Y)
		mXi = int(mX)
		mYi = int(mY)
		mXo, mYo := (adj_MouseX - offsetX), (adj_MouseY - offsetY)
		mXi_01 := (options.TileSize.X * mXi) + (mXi * options.TileSpacing.X)
		mYi_01 := (options.TileSize.Y * mYi) + (mYi * options.TileSpacing.Y)

		mXi_02 := (options.TileSize.X * mXi) + (mXi * options.TileSpacing.X) + options.TileSize.X
		mYi_02 := (options.TileSize.Y * mYi) + (mYi * options.TileSpacing.Y) + options.TileSize.Y
		//fmt.Printf("0-->%6d,%6d\t%6d,%6d\t rec start: %6d %6d\t %6d %6d \n", Raw_Mouse_X, Raw_Mouse_Y, mXo, mYo, mXi_01, mYi_01, mXi_02, mYi_02)
		if (((mXo) > (mXi_01)) && (mXo) < mXi_02) && ((mYo) > (mYi_01) && mYo < mYi_02) {
			isOnTile = true
		}
	}
	return mXi, mYi, isOnTile
}
func (imat IntegerMatrix2D) GetCursorBounds(options Integer_Matrix_Ebiten_DrawOptions) (int, int) {
	test1X := ((len(imat[0]) * options.TileSize.X) + (len(imat[0]) * options.TileSpacing.X)) + options.BoardPosition.X
	test1Y := ((len(imat) * options.TileSize.Y) + (len(imat) * options.TileSpacing.Y)) + options.BoardPosition.Y
	return test1X, test1Y
}

/*
func (imat IntMatrix) GetCoordOfMouseEvent(Raw_Mouse_X int, Raw_Mouse_Y int, OffsetX int, OffsetY int, options.TileSize.X int, options.TileSize.Y int, options.TileSpacing.X int, GapY int) (int, int, bool) {
	test1X := ((len(imat[0]) * options.TileSize.X) + (len(imat[0]) * options.TileSpacing.X)) + OffsetX
	test1Y := ((len(imat) * options.TileSize.Y) + (len(imat) * GapY)) + OffsetY
	var mXi = -1
	var mYi = -1
	var isOnTile = false
	if (Raw_Mouse_X > OffsetX && Raw_Mouse_X < test1X-options.TileSpacing.X) && (Raw_Mouse_Y > OffsetY && Raw_Mouse_Y < test1Y-GapY) {
		var mX = float32(Raw_Mouse_X-OffsetX) / float32(options.TileSize.X+options.TileSpacing.X) //float32(test1X)
		var mY = float32(Raw_Mouse_Y-OffsetY) / float32(options.TileSize.Y+GapY) //float32(test1Y)
		mXi = int(mX)
		mYi = int(mY)
		mXo, mYo := (Raw_Mouse_X - OffsetX), (Raw_Mouse_Y - OffsetY)
		mXi_01 := (options.TileSize.X * mXi) + (mXi * options.TileSpacing.X)
		mYi_01 := (options.TileSize.Y * mYi) + (mYi * GapY)

		mXi_02 := (options.TileSize.X * mXi) + (mXi * options.TileSpacing.X) + options.TileSize.X
		mYi_02 := (options.TileSize.Y * mYi) + (mYi * GapY) + options.TileSize.Y
		//fmt.Printf("0-->%6d,%6d\t%6d,%6d\t rec start: %6d %6d\t %6d %6d \n", Raw_Mouse_X, Raw_Mouse_Y, mXo, mYo, mXi_01, mYi_01, mXi_02, mYi_02)
		if (((mXo) > (mXi_01)) && (mXo) < mXi_02) && ((mYo) > (mYi_01) && mYo < mYi_02) {
			isOnTile = true
		}
	}
	return mXi, mYi, isOnTile
}

*/
