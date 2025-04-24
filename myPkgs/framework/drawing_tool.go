package framework

import (
	"fmt"
	"image/color"
	"slices"

	"github.com/KelleyTyler/GridTileEbit04_12/myPkgs/basic_geometry/coords"
	mat "github.com/KelleyTyler/GridTileEbit04_12/myPkgs/basic_geometry/matrix"
	"github.com/KelleyTyler/GridTileEbit04_12/myPkgs/misc"
	"github.com/hajimehoshi/ebiten/v2"
)

type Drawing_Tool struct {
	Points coords.CoordList
	// Point0, Point1 coords.CoordInts
	MaxPoints       int
	CurrPoints      int
	DisplaySettings mat.Integer_Matrix_Ebiten_DrawOptions
	Imat            *mat.IntegerMatrix2D
	// hasP1, hasP2 bool
}

func (dtool *Drawing_Tool) Init(intMat *mat.IntegerMatrix2D, maxPoints int, dsetting mat.Integer_Matrix_Ebiten_DrawOptions) {
	dtool.Points = make(coords.CoordList, 0)
	dtool.MaxPoints = maxPoints
	dtool.DisplaySettings = dsetting
	dtool.Imat = intMat
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

type Pathfind_Tester struct {
	Start, Target                                                coords.CoordInts
	ClosedList, OpenList, BlockedList                            []*mat.ImatNode
	IsReady, HasStarted, IsFinished, IsStartInput, IsTargetInput bool

	ShowOnScreen          bool
	ShowOnScreenOptions   uint8
	GridOptions           mat.Integer_Matrix_Ebiten_DrawOptions
	IMat                  *mat.IntegerMatrix2D
	max_fails, curr_fails int
}

func (pfind *Pathfind_Tester) Init(imat *mat.IntegerMatrix2D, options mat.Integer_Matrix_Ebiten_DrawOptions) {
	pfind.IMat = imat
	pfind.Reset()
	pfind.GridOptions = options
	pfind.GridOptions.TileLineColors = []color.Color{color.Black, color.Black, color.Black}

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
	pfind.max_fails = 100
	pfind.curr_fails = 0
	fmt.Printf("RESET PFIND TEST \n")
}

func (pfind *Pathfind_Tester) InputCoords(coord coords.CoordInts) {
	if !pfind.IsStartInput {
		pfind.Start = coord

		fmt.Printf("START INPUT\n")
		pfind.IsStartInput = true
	} else if !pfind.IsTargetInput {
		if !coord.IsEqual(pfind.Start) {
			pfind.Target = coord
			pfind.IsTargetInput = true
			pfind.IsReady = true
			pfind.IsFinished = false
			fmt.Printf("END INPUT\n")

		}
	}
}

// -----replicating the process we've seen before;
func (pfind *Pathfind_Tester) Process() error {
	var err error
	if pfind.IsReady {
		if !pfind.HasStarted {
			fmt.Printf("STARTING!\n")
			startnode := mat.GetNode(pfind.Start, pfind.Start, pfind.Target, *pfind.IMat, nil)
			pfind.ClosedList = append(pfind.ClosedList, &startnode)
			temp := mat.NodeList_GetNeighbors_4_Filtered_Hypentenuse(&startnode, pfind.Start, pfind.Target, pfind.IMat, []int{0, 1}, []int{9, 10}, [4]uint{1, 1, 1, 1})
			pfind.OpenList = append(pfind.OpenList, temp...)
			fmt.Printf("STARTUP! %d %d \n", len(pfind.ClosedList), len(pfind.OpenList))

			// pfind.OpenList, pfind.ClosedList, pfind.BlockedList, pfind.IsFinished, pfind.curr_fails, err = mat.Pathfind_Phase1_Tick(pfind.Start, pfind.Target, pfind.OpenList, pfind.ClosedList, pfind.BlockedList, pfind.IsFinished, pfind.curr_fails, pfind.max_fails, *pfind.IMat, []int{0, 1}, []int{9, 10}, [4]uint{1, 1, 1, 1})
			pfind.HasStarted = true
		} else {
			if len(pfind.OpenList) < 1 {
				mat.NodeList_SortByFValue_Ascending(pfind.ClosedList, pfind.Start, pfind.Target)
				// slices.Reverse(pfind.ClosedList)
				temp := mat.NodeList_GetNeighbors_4_Filtered_Hypentenuse(pfind.ClosedList[0], pfind.Start, pfind.Target, pfind.IMat, []int{0, 1}, []int{9, 10}, [4]uint{1, 1, 1, 1})
				pfind.OpenList = append(pfind.OpenList, temp...)
				mat.NodeList_SortByFValue_Ascending(pfind.OpenList, pfind.Start, pfind.Target)
				slices.Reverse(pfind.OpenList)
			}

			pfind.OpenList, pfind.ClosedList, pfind.BlockedList, pfind.IsFinished, pfind.curr_fails, err = mat.Pathfind_Phase1_Tick(pfind.Start, pfind.Target, pfind.OpenList, pfind.ClosedList, pfind.BlockedList, pfind.IsFinished, pfind.curr_fails, pfind.max_fails, *pfind.IMat, []int{0, 1}, []int{9, 10}, [4]uint{1, 1, 1, 1})
			fmt.Printf("RUNNING! %d %d \n", len(pfind.ClosedList), len(pfind.OpenList))
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
		for _, node := range pfind.ClosedList {
			pfind.IMat.DrawAGridTile_With_Lines(screen, node.Position, color.RGBA{255, 24, 125, 255}, &pfind.GridOptions)

		}
	}
	if len(pfind.OpenList) > 1 {
		for _, node := range pfind.OpenList {
			pfind.IMat.DrawAGridTile_With_Lines(screen, node.Position, color.RGBA{125, 128, 255, 255}, &pfind.GridOptions)

		}
	}
}
