package framework

import (
	"github.com/KelleyTyler/GridTileEbit04_12/myPkgs/basic_geometry/coords"
	mat "github.com/KelleyTyler/GridTileEbit04_12/myPkgs/basic_geometry/matrix"
	"github.com/KelleyTyler/GridTileEbit04_12/myPkgs/misc"
	ui "github.com/KelleyTyler/GridTileEbit04_12/myPkgs/user_interface"
)

type Drawing_Tool struct {
	Points coords.CoordList
	// Point0, Point1 coords.CoordInts
	MaxPoints       int
	CurrPoints      int
	DisplaySettings mat.Integer_Matrix_Ebiten_DrawOptions
	Imat            *mat.IntegerMatrix2D

	UI_Backend *ui.UI_Backend

	// hasP1, hasP2 bool

	//User_Interface_Buttons
	Label_DrawTool                                                             ui.UI_Label
	Button_DrawPoint, Button_DrawLine, Button_DrawCircle, Button_DrawRectangle ui.UI_Button
	Button_StopMode                                                            ui.UI_Button
	NSelect_CircleRadius, NSelect_LineColor, NSelect_FillColor                 ui.UI_Num_Select
}

func (dtool *Drawing_Tool) Init(intMat *mat.IntegerMatrix2D, backend *ui.UI_Backend, maxPoints int, dsetting mat.Integer_Matrix_Ebiten_DrawOptions) {
	dtool.Points = make(coords.CoordList, 0)
	dtool.MaxPoints = maxPoints
	dtool.DisplaySettings = dsetting
	dtool.Imat = intMat
	dtool.UI_Backend = backend
}
func (dTool *Drawing_Tool) Clear() {
	dTool.Points = make(coords.CoordList, 0)
	dTool.CurrPoints = 0
}
func (dTool *Drawing_Tool) AddToWithInts(x, y int) {
	if dTool.CurrPoints < dTool.MaxPoints {
		dTool.Points = append(dTool.Points, coords.CoordInts{X: x, Y: y})
		dTool.CurrPoints++
	}
}

func (dTool *Drawing_Tool) GetFirstTwoPoints() (point0, point1 coords.CoordInts) {
	if len(dTool.Points) >= 2 {
		point0 = dTool.Points[0]
		point1 = dTool.Points[1]
	}
	return point0, point1
}
func (dTool *Drawing_Tool) DrawLineToGrid(value int) (ret bool) {
	ret = false
	if dTool.CurrPoints >= dTool.MaxPoints {
		c0, c1 := dTool.GetFirstTwoPoints()
		line := coords.BresenhamLine(c0, c1)
		//fmt.Printf("DRAWTOOL: %s to %s: points %d:%d length %d %d\n", c0.ToString(), c1.ToString(), dTool.CurrPoints, dTool.MaxPoints, len(dTool.Points), len(line))
		for _, a := range line {
			if dTool.Imat.IsValidCoords(a) {
				dTool.Imat.SetValAtCoord(a, value)
			}
		}
		ret = true
	}
	return
}

/*
 */
func (dTool *Drawing_Tool) InitUI(parent ui.UI_Object, yValue int) {
	dTool.Label_DrawTool.Init([]string{"n_draw_tool_lbl", "Drawing Tool"}, dTool.UI_Backend, nil, coords.CoordInts{X: 0, Y: yValue}, coords.CoordInts{X: 204, Y: 32})
	dTool.Label_DrawTool.TextAlignMode = 10
	dTool.Label_DrawTool.Redraw()
	dTool.Label_DrawTool.Init_Parents(parent)

	dTool.NSelect_LineColor.Init([]string{"nselect_line_color", "Line Color"}, dTool.UI_Backend, nil, coords.CoordInts{X: 4, Y: yValue + 34}, coords.CoordInts{X: 64, Y: 32})
	dTool.NSelect_LineColor.SetVals(0, 1, 0, 10, 0)
	dTool.NSelect_LineColor.Init_Parents(parent)

	dTool.NSelect_CircleRadius.Init([]string{"nselect_Circ_rad", "CircleRadius"}, dTool.UI_Backend, nil, coords.CoordInts{X: 70, Y: yValue + 34}, coords.CoordInts{X: 64, Y: 32})
	dTool.NSelect_CircleRadius.SetVals(0, 1, 0, 32, 0)
	dTool.NSelect_CircleRadius.Init_Parents(parent)

	dTool.NSelect_FillColor.Init([]string{"nselect_fill_color", "Fill Color"}, dTool.UI_Backend, nil, coords.CoordInts{X: 136, Y: yValue + 34}, coords.CoordInts{X: 64, Y: 32})
	dTool.NSelect_FillColor.SetVals(0, 1, 0, 10, 0)
	dTool.NSelect_FillColor.Init_Parents(parent)

	dTool.Button_DrawPoint.Init([]string{"n_draw_point_btn", "Draw\nPoint"}, dTool.UI_Backend, nil, coords.CoordInts{X: 4, Y: yValue + 68}, coords.CoordInts{X: 64, Y: 32})
	dTool.Button_DrawLine.Init([]string{"n_draw_point_btn", "Draw\nLine"}, dTool.UI_Backend, nil, coords.CoordInts{X: 70, Y: yValue + 68}, coords.CoordInts{X: 64, Y: 32})
	dTool.Button_DrawRectangle.Init([]string{"n_draw_rectangle_btn", "Draw\nRectangle"}, dTool.UI_Backend, nil, coords.CoordInts{X: 136, Y: yValue + 68}, coords.CoordInts{X: 64, Y: 32})
	//-------------------------------Row3
	dTool.Button_DrawCircle.Init([]string{"n_draw_circle_btn", "Draw\nCircle"}, dTool.UI_Backend, nil, coords.CoordInts{X: 4, Y: yValue + 102}, coords.CoordInts{X: 64, Y: 32})
	dTool.Button_StopMode.Init([]string{"n_draw_circle_btn", "Stop\nMode"}, dTool.UI_Backend, nil, coords.CoordInts{X: 70, Y: yValue + 102}, coords.CoordInts{X: 64, Y: 32})

	dTool.Button_DrawPoint.Btn_Type = 10
	dTool.Button_DrawLine.Btn_Type = 10
	dTool.Button_DrawRectangle.Btn_Type = 10
	dTool.Button_DrawCircle.Btn_Type = 10
	dTool.Button_StopMode.Btn_Type = 10
	dTool.Button_DrawPoint.Init_Parents(parent)
	dTool.Button_DrawLine.Init_Parents(parent)
	dTool.Button_DrawRectangle.Init_Parents(parent)
	dTool.Button_DrawCircle.Init_Parents(parent)
	dTool.Button_StopMode.Init_Parents(parent)

}

/*
This handles the results of clicking on the gameboard when a part of the UI is active
*/
func (dTool *Drawing_Tool) OnValidMouseClickOnGameBoard(pos_X, pos_Y int) (OverlayChange, BoardChange bool) {
	OverlayChange = false
	BoardChange = false

	if dTool.Button_DrawPoint.GetState() == 2 {
		// dTool.AddToWithInts(pos_X, pos_Y)
		val := dTool.NSelect_LineColor.CurrValue
		dTool.Imat.SetValAtCoord(coords.CoordInts{X: pos_X, Y: pos_Y}, val)
		BoardChange = true
		dTool.UI_Backend.PlaySound(4)

	}
	if dTool.Button_DrawLine.GetState() == 2 {
		dTool.AddToWithInts(pos_X, pos_Y)
		OverlayChange = true
		dTool.UI_Backend.PlaySound(4)

		if dTool.Button_StopMode.GetState() == 2 {
			if dTool.DrawLineToGrid_Stop_At_Walls(dTool.NSelect_LineColor.CurrValue, []int{9, 10}) {
				OverlayChange = true
				BoardChange = true
				dTool.Clear()
			}
		} else {
			if dTool.DrawLineToGrid(dTool.NSelect_LineColor.CurrValue) {
				BoardChange = true
				dTool.Clear()
			}
		}
	}
	if dTool.Button_DrawRectangle.GetState() == 2 {
		dTool.UI_Backend.PlaySound(4)

		dTool.AddToWithInts(pos_X, pos_Y)
		OverlayChange = true
		if dTool.DrawRectangleToGrid(dTool.NSelect_LineColor.CurrValue, dTool.NSelect_FillColor.CurrValue) {
			BoardChange = true
			dTool.Clear()
		}
	}
	if dTool.Button_DrawCircle.GetState() == 2 {
		dTool.UI_Backend.PlaySound(4)
		dTool.AddToWithInts(pos_X, pos_Y)
		if dTool.Button_StopMode.GetState() == 2 {
			if dTool.DrawCircleToGrid_Radius_Bresenham_Walls(dTool.NSelect_CircleRadius.CurrValue, dTool.NSelect_LineColor.CurrValue, []int{9, 10}) {
				OverlayChange = true
				BoardChange = true
				dTool.Clear()
			}

		} else {
			if dTool.DrawCircleToGrid_Radius(dTool.NSelect_CircleRadius.CurrValue, dTool.NSelect_LineColor.CurrValue) {
				OverlayChange = true
				BoardChange = true
				dTool.Clear()
			}

		}
		OverlayChange = true

	}
	return OverlayChange, BoardChange
}

func (dTool *Drawing_Tool) DrawRectangleToGrid(lineValue, fillvalue int) (ret bool) {
	ret = false
	if dTool.CurrPoints >= dTool.MaxPoints {
		c0, c1 := dTool.GetFirstTwoPoints()
		var x0, y0, x1, y1 int = c1.X, c1.Y, c0.X, c0.Y
		if c0.X < c1.X {
			x0 = c0.X
			x1 = c1.X
		} else {
			x0 = c1.X
			x1 = c0.X
		}
		if c0.Y < c1.Y {
			y0 = c0.Y
			y1 = c1.Y
		} else {
			y0 = c1.Y
			y1 = c0.Y
		}
		for i := y0; i <= y1; i++ {
			for j := x0; j <= x1; j++ {
				temp := coords.CoordInts{X: j, Y: i}
				if i == y0 || j == x0 || i == y1 || j == x1 {
					if dTool.Imat.IsValidCoords(temp) && lineValue != -1 {
						dTool.Imat.SetValAtCoord(temp, lineValue)
					}
				} else {
					if dTool.Imat.IsValidCoords(temp) && fillvalue != -1 {
						dTool.Imat.SetValAtCoord(temp, fillvalue)
					}
				}
			}
		}
		ret = true
	}
	return
}

func (dTool *Drawing_Tool) DrawCircleToGrid(value int) (ret bool) {
	ret = false
	if dTool.CurrPoints >= dTool.MaxPoints {
		c0, c1 := dTool.GetFirstTwoPoints()
		dist := c0.GetHypotenuseDistance_Int(c1)
		tList := c0.GetACirclePointsFromCenter(dist)

		for _, c := range tList {
			if dTool.Imat.IsValidCoords(c) {
				dTool.Imat.SetValAtCoord(c, value)
			}
		}

		ret = true
	}
	return
}

func (dTool *Drawing_Tool) DrawLineToGrid_Stop_At_Walls(value int, walls []int) (ret bool) { //c0, c1 coords.CoordInts,
	ret = false
	if dTool.CurrPoints >= dTool.MaxPoints {
		c0, c1 := dTool.GetFirstTwoPoints()

		line := coords.BresenhamLine(c0, c1)
		//fmt.Printf("DRAWTOOL: %s to %s: points %d:%d length %d %d\n", c0.ToString(), c1.ToString(), dTool.CurrPoints, dTool.MaxPoints, len(dTool.Points), len(line))
		for _, a := range line {
			if dTool.Imat.IsValidCoords(a) && !misc.IsNumInIntArray(dTool.Imat.GetValueOnCoord(a), walls) {

				dTool.Imat.SetValAtCoord(a, value)
			} else {
				break
			}
		}
		ret = true
	}
	return
}

func (dTool *Drawing_Tool) DrawCircleToGrid_Bresenham_Walls(value int, walls []int) (ret bool) {
	ret = false
	if dTool.CurrPoints >= dTool.MaxPoints {
		c0, c1 := dTool.GetFirstTwoPoints()
		dist := c0.GetHypotenuseDistance_Int(c1)
		for i := 0; i < dist+1; i++ {
			tList := c0.GetACirclePointsFromCenter(i)

			for _, c := range tList {
				// dTool.DrawLineToGrid_Stop_At_Walls_CoordInts(dTool.Imat, c0, c, value, walls)
				if dTool.Imat.IsValidCoords(c) {
					dTool.Imat.SetValAtCoord(c, value)
				}
			}
		}

		ret = true
	}
	return
}

func (dTool *Drawing_Tool) DrawCircleToGrid_Radius_Bresenham_Walls(radius, value int, walls []int) (ret bool) {
	ret = false
	if dTool.CurrPoints >= dTool.MaxPoints {
		c0, _ := dTool.GetFirstTwoPoints()
		// dist := c0.GetHypotenuseDistance_Int(c1)
		dist := radius
		for i := 0; i < dist+1; i++ {
			tList := c0.GetACirclePointsFromCenter(i)

			for _, c := range tList {
				dTool.DrawLineToGrid_Stop_At_Walls_CoordInts(dTool.Imat, c0, c, value, walls)
				// if dTool.Imat.IsValidCoords(c) {
				// 	dTool.Imat.SetValAtCoord(c, value)
				// }
			}
		}

		ret = true
	}
	return
}

func (dTool *Drawing_Tool) DrawCircleToGrid_Radius(radius, value int) (ret bool) {
	ret = false
	if dTool.CurrPoints >= dTool.MaxPoints {
		c0, _ := dTool.GetFirstTwoPoints()
		// dist := c0.GetHypotenuseDistance_Int(c1)
		dist := radius
		for i := 0; i < dist+1; i++ {
			tList := c0.GetACirclePointsFromCenter(i)

			for _, c := range tList {
				// dTool.DrawLineToGrid_CoordInts(dTool.Imat, c0, c, value)
				dTool.Imat.SetValAtCoord(c, value)
			}
		}

		ret = true
	}
	return
}
func (dTool *Drawing_Tool) DrawLineToGrid_Stop_At_Walls_CoordInts(imat *mat.IntegerMatrix2D, c0, c1 coords.CoordInts, value int, walls []int) (ret bool) { //c0, c1 coords.CoordInts,
	ret = false
	if dTool.CurrPoints >= dTool.MaxPoints {
		// c0, c1 := dTool.GetFirstTwoPoints()

		line := coords.BresenhamLine(c0, c1)
		//fmt.Printf("DRAWTOOL: %s to %s: points %d:%d length %d %d\n", c0.ToString(), c1.ToString(), dTool.CurrPoints, dTool.MaxPoints, len(dTool.Points), len(line))
		for _, a := range line {
			if dTool.Imat.IsValidCoords(a) && !misc.IsNumInIntArray(dTool.Imat.GetValueOnCoord(a), walls) {

				imat.SetValAtCoord(a, value)
			} else {
				break
			}
		}
		ret = true
	}
	return
}
func (dTool *Drawing_Tool) DrawLineToGrid_CoordInts(imat *mat.IntegerMatrix2D, c0, c1 coords.CoordInts, value int) (ret bool) { //c0, c1 coords.CoordInts,
	ret = false
	if dTool.CurrPoints >= dTool.MaxPoints {
		// c0, c1 := dTool.GetFirstTwoPoints()

		line := coords.BresenhamLine(c0, c1)
		//fmt.Printf("DRAWTOOL: %s to %s: points %d:%d length %d %d\n", c0.ToString(), c1.ToString(), dTool.CurrPoints, dTool.MaxPoints, len(dTool.Points), len(line))
		for _, a := range line {
			if dTool.Imat.IsValidCoords(a) {

				imat.SetValAtCoord(a, value)
			}
		}
		ret = true
	}
	return
}
