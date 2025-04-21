package framework

import (
	"github.com/KelleyTyler/GridTileEbit04_12/myPkgs/basic_geometry/coords"
	mat "github.com/KelleyTyler/GridTileEbit04_12/myPkgs/basic_geometry/matrix"
	"github.com/KelleyTyler/GridTileEbit04_12/myPkgs/misc"
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
func (dTool *Drawing_Tool) DrawRectangleToGrid(value int) (ret bool) {
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
				if dTool.Imat.IsValidCoords(temp) {
					dTool.Imat.SetValAtCoord(temp, value)
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
