package framework

import (
	"fmt"
	"image/color"

	"github.com/KelleyTyler/GridTileEbit04_12/myPkgs/basic_geometry/coords"
	mat "github.com/KelleyTyler/GridTileEbit04_12/myPkgs/basic_geometry/matrix"
	"github.com/KelleyTyler/GridTileEbit04_12/myPkgs/misc"
	ui "github.com/KelleyTyler/GridTileEbit04_12/myPkgs/user_interface"
	"github.com/hajimehoshi/ebiten/v2"
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
	dTool.NSelect_CircleRadius.SetVals(0, 1, 0, 10, 0)
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

type Pathfind_Tester struct {
	Start, Target                                                coords.CoordInts
	ClosedList, OpenList, BlockedList                            []*mat.ImatNode
	IsReady, HasStarted, IsFinished, IsStartInput, IsTargetInput bool

	EndNode *mat.ImatNode

	ShowOnScreen        bool
	ShowOnScreenOptions uint8
	GridOptions         mat.Integer_Matrix_Ebiten_DrawOptions
	IMat                *mat.IntegerMatrix2D

	UI_Backend *ui.UI_Backend

	max_fails, curr_fails int
	//------- User Interface and User Interface Buttons
	Button_Panel_Label                                       ui.UI_Label
	Button_Select_Points, Button_Pathfind_Tick, Button_Reset ui.UI_Button
	Button_Pathfind_Auto                                     ui.UI_Button

	tickCount int
}

func (pfind *Pathfind_Tester) Init(imat *mat.IntegerMatrix2D, backend *ui.UI_Backend, options *mat.Integer_Matrix_Ebiten_DrawOptions) {
	pfind.IMat = imat
	pfind.Reset()
	pfind.GridOptions = options.GetOptsByValue()
	pfind.GridOptions.TileLineColors = []color.Color{color.Black, color.Black, color.Black}
	pfind.UI_Backend = backend
}
func (pfind *Pathfind_Tester) Reset() {
	pfind.IsTargetInput = false
	pfind.IsStartInput = false
	pfind.IsFinished = false
	pfind.IsReady = false
	pfind.HasStarted = false
	pfind.ClosedList = make([]*mat.ImatNode, 0)
	pfind.OpenList = make([]*mat.ImatNode, 0)
	pfind.BlockedList = make([]*mat.ImatNode, 0)
	pfind.EndNode = nil
	pfind.max_fails = 100
	pfind.curr_fails = 0
	fmt.Printf("RESET PFIND TEST \n")
	pfind.tickCount = 0
}

func (pfind *Pathfind_Tester) UI_Init(parent ui.UI_Object, pfindBtnRow int) {
	pfind.Button_Panel_Label.Init([]string{"pfind_b_panel_label", "Pathfind"}, pfind.UI_Backend, nil, coords.CoordInts{X: 0, Y: pfindBtnRow}, coords.CoordInts{X: 204, Y: 32})
	pfind.Button_Panel_Label.TextAlignMode = 10
	pfind.Button_Panel_Label.Redraw()
	pfind.Button_Panel_Label.Init_Parents(parent)
	pfind.Button_Select_Points.Init([]string{"pfind_set_points", "SET\nPOINTS"}, pfind.UI_Backend, nil, coords.CoordInts{X: 4, Y: pfindBtnRow + 34}, coords.CoordInts{X: 64, Y: 32})
	pfind.Button_Select_Points.Btn_Type = 10
	pfind.Button_Select_Points.Init_Parents(parent)
	pfind.Button_Pathfind_Tick.Init([]string{"pfind_Tick", "TICK"}, pfind.UI_Backend, nil, coords.CoordInts{X: 70, Y: pfindBtnRow + 34}, coords.CoordInts{X: 64, Y: 32})
	pfind.Button_Pathfind_Tick.Init_Parents(parent)
	pfind.Button_Reset.Init([]string{"pfind_reset_points", "RESET"}, pfind.UI_Backend, nil, coords.CoordInts{X: 136, Y: pfindBtnRow + 34}, coords.CoordInts{X: 64, Y: 32})
	pfind.Button_Reset.Init_Parents(parent)

	pfind.Button_Pathfind_Auto.Init([]string{"pfind_Auto", "AUTO\nPathfind"}, pfind.UI_Backend, nil, coords.CoordInts{X: 4, Y: pfindBtnRow + 68}, coords.CoordInts{X: 64, Y: 32})
	pfind.Button_Pathfind_Auto.Btn_Type = 10
	pfind.Button_Pathfind_Auto.Init_Parents(parent)

}

func (pfind *Pathfind_Tester) OnValidMouseClickOnGameBoard(pos_X, pos_Y int) (overlay_change, board_change bool) {
	overlay_change = false
	board_change = false

	if pfind.Button_Select_Points.GetState() == 2 {
		pfind.InputCoords(coords.CoordInts{X: pos_X, Y: pos_Y})
		overlay_change = true
	}

	return overlay_change, board_change
}

func (pfind *Pathfind_Tester) Update_Passive() (overlay_change bool) {
	overlay_change = false
	if pfind.Button_Pathfind_Tick.GetState() == 2 {
		pfind.Process()
		overlay_change = true
	}
	if pfind.Button_Reset.GetState() == 2 {
		pfind.Reset()
		overlay_change = true
	}
	if pfind.Button_Pathfind_Auto.GetState() == 2 {
		pfind.Process()
		overlay_change = true
	}
	return overlay_change
}

func (pfind *Pathfind_Tester) InputCoords(coord coords.CoordInts) {
	if !pfind.IsStartInput {
		pfind.Start = coord

		fmt.Printf("START INPUT\n")
		pfind.IsStartInput = true
		pfind.UI_Backend.PlaySound(4)

	} else if !pfind.IsTargetInput {
		if !coord.IsEqual(pfind.Start) {
			pfind.Target = coord
			pfind.IsTargetInput = true
			pfind.IsReady = true
			pfind.IsFinished = false
			// pfind.EndNode = nil
			fmt.Printf("END INPUT\n")
			pfind.UI_Backend.PlaySound(4)

		}
	}
}

// -----replicating the process we've seen before;
func (pfind *Pathfind_Tester) Process() error {
	var err error

	// fmt.Printf("ENDNODE %t %t %t %t\n\n", pfind.EndNode != nil, pfind.HasStarted, pfind.IsFinished, pfind.IsReady)

	if pfind.IsReady {
		if !pfind.HasStarted {
			fmt.Printf("STARTING!\n")
			startnode := mat.GetNode(pfind.Start, pfind.Start, pfind.Target, *pfind.IMat, nil)
			pfind.ClosedList = append(pfind.ClosedList, &startnode)
			temp := mat.NodeList_GetNeighbors_4_Filtered_Hypentenuse(&startnode, pfind.Start, pfind.Target, pfind.IMat, []int{0, 1}, []int{9, 10}, [4]uint{1, 1, 1, 1})
			mat.NodeList_SortByFValue_DESC(temp, pfind.Start, pfind.Target)
			pfind.OpenList = append(pfind.OpenList, temp...)
			fmt.Printf("STARTUP! %d %d \n", len(pfind.ClosedList), len(pfind.OpenList))

			// pfind.OpenList, pfind.ClosedList, pfind.BlockedList, pfind.IsFinished, pfind.curr_fails, err = mat.Pathfind_Phase1_Tick(pfind.Start, pfind.Target, pfind.OpenList, pfind.ClosedList, pfind.BlockedList, pfind.IsFinished, pfind.curr_fails, pfind.max_fails, *pfind.IMat, []int{0, 1}, []int{9, 10}, [4]uint{1, 1, 1, 1})
			pfind.HasStarted = true
		} else {
			if !pfind.IsFinished {
				if len(pfind.OpenList) > 0 {
					// mat.NodeList_SortByHValue_Ascending(pfind.ClosedList)
					// mat.NodeList_SortByFValue_Ascending(pfind.ClosedList, pfind.Start, pfind.Target)
					pfind.ClosedList = mat.NodeList_SortByFValue_Desc_toReturn(pfind.ClosedList, pfind.Start, pfind.Target)
					// slices.Reverse(pfind.ClosedList)
					temp := mat.NodeList_GetNeighbors_4_Filtered_Hypentenuse(pfind.ClosedList[0], pfind.Start, pfind.Target, pfind.IMat, []int{0, 1}, []int{9, 10}, [4]uint{1, 1, 1, 1})
					temp = mat.NodeList_FILTER_LIST(temp, pfind.OpenList)
					temp = mat.NodeList_FILTER_LIST(temp, pfind.ClosedList)
					pfind.OpenList = append(pfind.OpenList, temp...)
					//fmt.Printf("TICK!----\n")

					// mat.NodeList_SortByFValue_Ascending(pfind.OpenList, pfind.Start, pfind.Target)
					// slices.Reverse(pfind.OpenList)
				}

				pfind.OpenList, pfind.ClosedList, pfind.BlockedList, pfind.IsFinished, pfind.EndNode, pfind.curr_fails, err = mat.Pathfind_Phase1_Tick(pfind.Start, pfind.Target, pfind.OpenList, pfind.ClosedList, pfind.BlockedList, pfind.IsFinished, pfind.curr_fails, pfind.max_fails, *pfind.IMat, []int{0, 1}, []int{9, 10}, [4]uint{1, 1, 1, 1})
				//fmt.Printf("RUNNING! %d %d \n", len(pfind.ClosedList), len(pfind.OpenList))
				pfind.tickCount++
			} else {
				fmt.Printf("FOUND! %t---- TICKS: %d\n\n", pfind.EndNode != nil, pfind.tickCount)
				// if pfind.EndNode != nil {

				// 	pfind.EndNode.Set_Heads_Tails_On_Up()
				// }
				pfind.EndNode.Set_Heads_Tails_On_Up() //<---- for some reason the whole thing fucking freezes if this is put into any kind of storage;
				if pfind.Button_Pathfind_Auto.GetState() == 2 {
					pfind.Button_Pathfind_Auto.DeToggle()
				}
			}
		}
	}
	return err
}
func (pfind *Pathfind_Tester) Draw(screen *ebiten.Image, drawOpts *mat.Integer_Matrix_Ebiten_DrawOptions) {
	pfind.GridOptions = *drawOpts
	pfind.GridOptions.TileLineColors = []color.Color{color.Black, color.Black, color.Black}

	if pfind.IsStartInput {
		pfind.IMat.DrawAGridTile_With_Lines(screen, pfind.Start, color.RGBA{128, 255, 0, 255}, &pfind.GridOptions)

	}
	if pfind.IsTargetInput {
		pfind.IMat.DrawAGridTile_With_Lines(screen, pfind.Target, color.RGBA{255, 128, 0, 255}, &pfind.GridOptions)
	}

	if len(pfind.ClosedList) > 1 {
		pfind.IMat.NodeList_DrawList(screen, pfind.ClosedList, []color.Color{color.RGBA{255, 24, 125, 255}, color.RGBA{200, 24, 100, 255}, color.RGBA{180, 24, 70, 255}}, &pfind.GridOptions)

		// for _, node := range pfind.ClosedList {
		// 	pfind.IMat.DrawAGridTile_With_Lines(screen, node.Position, color.RGBA{255, 24, 125, 255}, &pfind.GridOptions)

		// }
	}
	if len(pfind.OpenList) > 1 {
		// for _, node := range pfind.OpenList {
		// 	pfind.IMat.DrawAGridTile_With_Lines(screen, node.Position, color.RGBA{125, 128, 255, 255}, &pfind.GridOptions)

		// }
		pfind.IMat.NodeList_DrawList(screen, pfind.OpenList, []color.Color{color.RGBA{125, 128, 255, 255}, color.RGBA{100, 108, 222, 255}, color.RGBA{100, 108, 200, 255}}, &pfind.GridOptions)

	}
	if len(pfind.BlockedList) > 1 {
		pfind.IMat.NodeList_DrawList(screen, pfind.BlockedList, []color.Color{color.RGBA{75, 255, 75, 255}, color.RGBA{50, 255, 50, 255}, color.RGBA{25, 255, 25, 255}}, &pfind.GridOptions)
	}
	if pfind.EndNode != nil {
		// pfind.IMat.DrawAGridTile_With_Lines(screen, pfind.EndNode.Position, color.RGBA{255, 255, 255, 255}, &pfind.GridOptions)
		pfind.IMat.ImatNode_Draw(screen, pfind.EndNode, []color.Color{color.RGBA{255, 255, 255, 255}, color.RGBA{155, 155, 155, 255}, color.RGBA{75, 75, 75, 255}}, &pfind.GridOptions)
	}
}
