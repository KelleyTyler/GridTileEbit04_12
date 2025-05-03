package matrix

import (
	"fmt"
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
	//--------------------------------
	/*
		BONUS OPTIONS:
		SubDivisions (basically allowing the grid to be divided up into little bitty segments)

	*/

	SubDiv_00_Size_X, SubDiv_00_Size_Y     uint8
	SubDiv_00_Offset_X, SubDiv_00_Offset_Y uint8
	SubDiv_00_Line_thickness               float32
	SubDiv_00_Line_Colors                  color.Color
	SubDiv_00_Show                         bool
}

/**/
func (Opts Integer_Matrix_Ebiten_DrawOptions) GetOptsPtr() (optsPtr *Integer_Matrix_Ebiten_DrawOptions) {
	return &Opts
}

/**/
func (Opts Integer_Matrix_Ebiten_DrawOptions) ToString() (outString string) {
	outString = fmt.Sprintf("Tile Size: %3d, %3d TileSpacing %3d %3d", Opts.TileSize.X, Opts.TileSize.Y, Opts.TileSpacing.X, Opts.TileSpacing.Y)
	return outString
}

/**/
func (Opts Integer_Matrix_Ebiten_DrawOptions) GetOptsByValue() (optsByValue Integer_Matrix_Ebiten_DrawOptions) {
	return Opts
}

/**/
func (Opts *Integer_Matrix_Ebiten_DrawOptions) Update_Opts_From_Argument(UpOps Integer_Matrix_Ebiten_DrawOptions) {
	Opts.TileSize = UpOps.TileSize
	Opts.TileSpacing = UpOps.TileSpacing
	Opts.BoardMargin = UpOps.BoardMargin
	Opts.BoardPosition = UpOps.BoardPosition
}

/*
	This Returns the Most Basic Form Of these, likely useful for the many 'draw grid with lines' requests;

x0,y0 is top left, (on the screen anyway; so down->right is the way the grid works)
x1,y1 is bottom right
*/
func (options *Integer_Matrix_Ebiten_DrawOptions) Modify_Coord_For_Options(coord coords.CoordInts) (x0, y0, x1, y1 float64) {
	// TileCoord := coords.CoordInts{X: options.TileSize.X * coord.X, Y: options.TileSize.Y * coord.Y}
	x0 = float64(options.TileSize.X*coord.X) + float64(options.TileSpacing.X*coord.X) + float64(options.BoardMargin.X)
	// float_X0_Spacing := float64(options.TileSpacing.X * coord.X)
	y0 = float64(options.TileSize.Y*coord.Y) + float64(options.TileSpacing.Y*coord.Y) + float64(options.BoardMargin.Y)
	// float_Y0_Spacing := float64(options.TileSpacing.Y * coord.Y)
	//-------------------------------------------------------
	x1 = x0 + float64(options.TileSize.X)
	// float_X1_Spacing := float64(options.TileSpacing.X * coord.X)
	y1 = y0 + float64(options.TileSize.Y)
	// float_Y1_Spacing := float64(options.TileSpacing.Y * coord.Y)
	return x0, y0, x1, y1
}

/*
this gets the center of the coord;
*/
func (options *Integer_Matrix_Ebiten_DrawOptions) GetCoordCenter(coord coords.CoordInts) (xC, yC float64) {
	x0, y0, x1, y1 := options.Modify_Coord_For_Options(coord)

	xC = (x0 + x1) / 2.0
	yC = (y0 + y1) / 2.0

	return xC, yC
}

/* This is the Advanced Version */
func (options *Integer_Matrix_Ebiten_DrawOptions) Modify_Coord_For_Options_ADVCED(coord coords.CoordInts) (x0, y0, x1, y1 float64) {

	var bonus_X float64 = 0.0
	var bonus_Y float64 = 0.0
	if coord.X > int(options.SubDiv_00_Offset_X) {
		x_pos_corrected_for_sDiv_offset := coord.X - int(options.SubDiv_00_Offset_X)
		num_of_S_Div_Behind := 0
		if options.SubDiv_00_Size_X > 0 {
			num_of_S_Div_Behind = x_pos_corrected_for_sDiv_offset / int(options.SubDiv_00_Size_X)
		}
		bonus_X = float64(num_of_S_Div_Behind) * float64(options.SubDiv_00_Line_thickness)
	}
	if coord.Y > int(options.SubDiv_00_Offset_Y) {
		y_pos_corrected_for_sDiv_offset := coord.Y - int(options.SubDiv_00_Offset_Y)
		num_of_S_Div_Behind := 0
		if options.SubDiv_00_Size_Y > 0 {
			num_of_S_Div_Behind = y_pos_corrected_for_sDiv_offset / int(options.SubDiv_00_Size_Y)
		}
		bonus_Y = float64(num_of_S_Div_Behind) * float64(options.SubDiv_00_Line_thickness)
	}
	// TileCoord := coords.CoordInts{X: options.TileSize.X * coord.X, Y: options.TileSize.Y * coord.Y}
	x0 = float64(options.TileSize.X*coord.X) + float64(options.TileSpacing.X*coord.X) + float64(options.BoardMargin.X) + bonus_X
	// float_X0_Spacing := float64(options.TileSpacing.X * coord.X)
	y0 = float64(options.TileSize.Y*coord.Y) + float64(options.TileSpacing.Y*coord.Y) + float64(options.BoardMargin.Y) + bonus_Y
	// float_Y0_Spacing := float64(options.TileSpacing.Y * coord.Y)
	//-------------------------------------------------------
	x1 = x0 + float64(options.TileSize.X)
	// float_X1_Spacing := float64(options.TileSpacing.X * coord.X)
	y1 = y0 + float64(options.TileSize.Y)
	// float_Y1_Spacing := float64(options.TileSpacing.Y * coord.Y)
	return x0, y0, x1, y1
}

/*
ShowTileLines     []bool
	TileLineColors    []color.Color
	TileLineThickness []float32
	AABody            bool
	AALines           bool
*/

/**/
func (imat IntegerMatrix2D) DrawAGridTile_With_Lines(screen *ebiten.Image, coord coords.CoordInts, clr0 color.Color, options *Integer_Matrix_Ebiten_DrawOptions) {

	x0, y0, x1, y1 := options.Modify_Coord_For_Options(coord)
	// xc, yc := options.GetCoordCenter(coord)
	vector.DrawFilledRect(screen, float32(x0), float32(y0), float32(options.TileSize.X), float32(options.TileSize.Y), clr0, options.AABody)
	// TileCoord := coords.CoordInts{X: options.TileSize.X * coord.X, Y: options.TileSize.Y * coord.Y}
	// TileSpace := CoordInts{X: options.TileSpacing.X * coord.X, Y: options.TileSpacing.Y * coord.Y}
	if options.ShowTileLines[0] {
		vector.StrokeRect(screen, float32(x0), float32(y0), float32(options.TileSize.X), float32(options.TileSize.Y), options.TileLineThickness[0], options.TileLineColors[0], options.AALines)
	}
	if options.ShowTileLines[1] {
		vector.StrokeLine(screen, float32(x0), float32(y0), float32(x1), float32(y1), options.TileLineThickness[0], options.TileLineColors[1], options.AALines)
	}
	if options.ShowTileLines[2] {
		// vector.DrawFilledCircle(screen, float32(xc), float32(yc), 1.0, color.Black, true)
		vector.StrokeLine(screen, float32(x0), float32(y1), float32(x1), float32(y0), options.TileLineThickness[0], options.TileLineColors[2], options.AALines)
	}
}

/**/
func (imat IntegerMatrix2D) DrawCoordListWithLines(screen *ebiten.Image, cList coords.CoordList, colors []color.Color, options Integer_Matrix_Ebiten_DrawOptions) {
	if len(colors) == 1 {
		for _, a := range cList {
			imat.DrawAGridTile_With_Lines(screen, a, colors[0], &options)
		}
	}
}

/**/
func (imat IntegerMatrix2D) DrawFullGridFromColors(screen *ebiten.Image, colors []color.Color, showBoardOL bool, options *Integer_Matrix_Ebiten_DrawOptions) {
	test1X := ((len(imat[0]) * options.TileSize.X) + (len(imat[0]) * options.TileSpacing.X)) + options.BoardMargin.X
	test1Y := ((len(imat) * options.TileSize.Y) + (len(imat) * options.TileSpacing.Y)) + options.BoardMargin.Y
	for y, _ := range imat {

		for x, b := range imat[y] {

			//----Draw Lines Here??
			x0, y0, _, _ := options.Modify_Coord_For_Options(coords.CoordInts{X: x, Y: y})

			if b < len(colors) {
				// vector.DrawFilledRect(screen, float32((options.TileSize.X*x)+(options.TileSpacing.X*x)+options.BoardMargin.X), float32((options.TileSize.Y*y)+(options.TileSpacing.Y*y)+options.BoardMargin.Y), float32(options.TileSize.X), float32(options.TileSize.Y), colors[b], options.AABody)
				vector.DrawFilledRect(screen, float32(x0), float32(y0), float32(options.TileSize.X), float32(options.TileSize.Y), colors[b], options.AABody)

			} else {
				// vector.DrawFilledRect(screen, float32((options.TileSize.X*x)+(options.TileSpacing.X*x)+options.BoardMargin.X), float32((options.TileSize.Y*y)+(options.TileSpacing.Y*y)+options.BoardMargin.Y), float32(options.TileSize.X), float32(options.TileSize.Y), colors[0], options.AABody)
				vector.DrawFilledRect(screen, float32(x0), float32(y0), float32(options.TileSize.X), float32(options.TileSize.Y), colors[0], options.AABody)

			}
			// color0 := color.RGBA{12, 12, 12, 100} //color.Black
			if options.ShowTileLines[0] {
				vector.StrokeRect(screen, float32(x0), float32(y0), float32(options.TileSize.X), float32(options.TileSize.Y), options.TileLineThickness[0], options.TileLineColors[0], options.AALines)

				// vector.StrokeRect(screen, float32((options.TileSize.X*x)+(options.TileSpacing.X*x)+options.BoardMargin.X), float32((options.TileSize.Y*y)+(options.TileSpacing.Y*y)+options.BoardMargin.Y), float32(options.TileSize.X), float32(options.TileSize.Y), options.TileLineThickness[0], options.TileLineColors[0], options.AALines)
			}
		}
	}
	if options.SubDiv_00_Show {
		imat.DrawSubLines(screen, options)

	}
	//vector.StrokeRect(screen, float32(OffsetX-0), float32(OffsetY-0), float32(test1X-0), float32(test1Y+0), 2.0, color.RGBA{210, 153, 100, 255}, true) //0, 179, 100, 255
	if showBoardOL {
		buffernum00 := 4
		buffernum01 := 2 * buffernum00
		vector.StrokeRect(screen, float32(options.BoardMargin.X-buffernum00), float32(options.BoardMargin.Y-buffernum00), float32(test1X-options.BoardMargin.X-options.TileSpacing.X+buffernum01), float32(test1Y-options.BoardMargin.Y-options.TileSpacing.Y+buffernum01), 2.0, color.White, true) //color.RGBA{0, 50, 50, 255}
	}
}

/*
vLine_x0, vLine_y0, _, _ := options.Modify_Coord_For_Options(coords.CoordInts{X: 0, Y: y})
vLine_x1, vLine_y1, _, _ := options.Modify_Coord_For_Options(coords.CoordInts{X: len(imat[y]), Y: y}) //-float32(options.TileSpacing.X)
vector.StrokeLine(screen, float32(vLine_x0), float32(vLine_y0), float32(vLine_x1), float32(vLine_y1), 1.0, options.SubDiv_00_Line_Colors, true)
vector.StrokeLine(screen, float32(vLine_x0), float32(vLine_y0)-float32(options.TileSpacing.Y), float32(vLine_x1), float32(vLine_y1)-float32(options.TileSpacing.Y), 1.0, options.SubDiv_00_Line_Colors, true)
*/
func (imat IntegerMatrix2D) DrawSubLines(screen *ebiten.Image, options *Integer_Matrix_Ebiten_DrawOptions) {
	for y := range len(imat) {
		if y%int(options.SubDiv_00_Size_Y) == 0 {
			vLine_x0, vLine_y0, _, _ := options.Modify_Coord_For_Options(coords.CoordInts{X: 0, Y: y})
			vLine_x1, vLine_y1, _, _ := options.Modify_Coord_For_Options(coords.CoordInts{X: len(imat[y]), Y: y}) //-float32(options.TileSpacing.X)
			vector.StrokeLine(screen, float32(vLine_x0), float32(vLine_y0), float32(vLine_x1), float32(vLine_y1), 1.0, options.SubDiv_00_Line_Colors, true)
			vector.StrokeLine(screen, float32(vLine_x0), float32(vLine_y0)-float32(options.TileSpacing.Y), float32(vLine_x1), float32(vLine_y1)-float32(options.TileSpacing.Y), 1.0, options.SubDiv_00_Line_Colors, true)

		}
	}

	for x := range len(imat[0]) {
		if x%int(options.SubDiv_00_Size_Y) == 0 {
			vLine_x0, vLine_y0, _, _ := options.Modify_Coord_For_Options(coords.CoordInts{X: x, Y: 0})
			vLine_x1, vLine_y1, _, _ := options.Modify_Coord_For_Options(coords.CoordInts{X: x, Y: len(imat)})
			vector.StrokeLine(screen, float32(vLine_x0), float32(vLine_y0), float32(vLine_x1), float32(vLine_y1), 1.0, options.SubDiv_00_Line_Colors, true)
			vector.StrokeLine(screen, float32(vLine_x0)-float32(options.TileSpacing.X), float32(vLine_y0), float32(vLine_x1)-float32(options.TileSpacing.X), float32(vLine_y1), 1.0, options.SubDiv_00_Line_Colors, true)

		}
	}

	_, vLine_y0, vLine_x0, _ := options.Modify_Coord_For_Options(coords.CoordInts{X: len(imat[0]) - 1, Y: 0})
	_, _, vLine_x1, vLine_y1 := options.Modify_Coord_For_Options(coords.CoordInts{X: len(imat[0]) - 1, Y: len(imat) - 1})
	vector.StrokeLine(screen, float32(vLine_x0), float32(vLine_y0), float32(vLine_x1), float32(vLine_y1), 1.0, options.SubDiv_00_Line_Colors, true) //options.SubDiv_00_Line_Colors

	vLine_x0, _, _, vLine_y0 = options.Modify_Coord_For_Options(coords.CoordInts{X: 0, Y: len(imat) - 1})
	_, _, vLine_x1, vLine_y1 = options.Modify_Coord_For_Options(coords.CoordInts{X: len(imat[0]) - 1, Y: len(imat) - 1})
	vector.StrokeLine(screen, float32(vLine_x0), float32(vLine_y0), float32(vLine_x1), float32(vLine_y1), 1.0, options.SubDiv_00_Line_Colors, true) //options.SubDiv_00_Line_Colors
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
		//log.Printf("0-->%6d,%6d\t%6d,%6d\t rec start: %6d %6d\t %6d %6d \n", Raw_Mouse_X, Raw_Mouse_Y, mXo, mYo, mXi_01, mYi_01, mXi_02, mYi_02)
		if (((mXo) > (mXi_01)) && (mXo) < mXi_02) && ((mYo) > (mYi_01) && mYo < mYi_02) {
			isOnTile = true
		}
	}
	return mXi, mYi, isOnTile
}

/**/
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
		//log.Printf("0-->%6d,%6d\t%6d,%6d\t rec start: %6d %6d\t %6d %6d \n", Raw_Mouse_X, Raw_Mouse_Y, mXo, mYo, mXi_01, mYi_01, mXi_02, mYi_02)
		if (((mXo) > (mXi_01)) && (mXo) < mXi_02) && ((mYo) > (mYi_01) && mYo < mYi_02) {
			isOnTile = true
		}
	}
	return mXi, mYi, isOnTile
}

*/

// func (imat IntegerMatrix2D) DrawAGridTile_With_Lines_Strings(screen *ebiten.Image, coord coords.CoordInts, strng string, clr0 color.Color, options *Integer_Matrix_Ebiten_DrawOptions) {
// 	vector.DrawFilledRect(screen, float32((options.TileSize.X*coord.X)+(options.TileSpacing.X*coord.X)+options.BoardMargin.X), float32((options.TileSize.Y*coord.Y)+(options.TileSpacing.Y*coord.Y)+options.BoardMargin.Y), float32(options.TileSize.X), float32(options.TileSize.Y), clr0, options.AABody)
// 	TileCoord := coords.CoordInts{X: options.TileSize.X * coord.X, Y: options.TileSize.Y * coord.Y}
// 	// TileSpace := CoordInts{X: options.TileSpacing.X * coord.X, Y: options.TileSpacing.Y * coord.Y}
// 	if options.ShowTileLines[0] {
// 		vector.StrokeRect(screen, float32((options.TileSize.X*coord.X)+(options.TileSpacing.X*coord.X)+options.BoardMargin.X), float32((options.TileSize.Y*coord.Y)+(options.TileSpacing.Y*coord.Y)+options.BoardMargin.Y), float32(options.TileSize.X), float32(options.TileSize.Y), options.TileLineThickness[0], options.TileLineColors[0], options.AALines)
// 	}
// 	if options.ShowTileLines[1] {
// 		vector.StrokeLine(screen, float32((TileCoord.X)+(options.TileSpacing.X*coord.X)+options.BoardMargin.X), float32((TileCoord.Y)+(options.TileSpacing.Y*coord.Y)+options.BoardMargin.Y), float32((options.TileSize.X*coord.X)+(options.TileSpacing.X*coord.X)+options.BoardMargin.X+options.TileSize.X), float32((options.TileSize.Y*coord.Y)+(options.TileSpacing.Y*coord.Y)+options.BoardMargin.Y+options.TileSize.Y), options.TileLineThickness[0], options.TileLineColors[1], options.AALines)
// 	}
// 	if options.ShowTileLines[2] {
// 		vector.StrokeLine(screen, float32((TileCoord.X)+(options.TileSpacing.X*coord.X)+options.BoardMargin.X), float32((TileCoord.Y)+(options.TileSpacing.Y*coord.Y)+options.BoardMargin.Y+options.TileSize.Y), float32((options.TileSize.X*coord.X)+(options.TileSpacing.X*coord.X)+options.BoardMargin.X+options.TileSize.X), float32((options.TileSize.Y*coord.Y)+(options.TileSpacing.Y*coord.Y)+options.BoardMargin.Y), options.TileLineThickness[0], options.TileLineColors[2], options.AALines)
// 	}
// }
