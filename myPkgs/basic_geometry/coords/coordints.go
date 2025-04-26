package coords

import (
	"fmt"
	"math"
)

type CoordInts struct {
	X, Y int
}

func (cord *CoordInts) ToString() string {
	return fmt.Sprintf("X: %3d, Y: %3d", cord.X, cord.Y)
}
func (cord *CoordInts) Add(cord0 CoordInts) {
	cord.X += cord0.X
	cord.Y += cord0.Y
}
func (cord *CoordInts) AddToReturn(cord0 CoordInts) CoordInts {
	temp := *cord
	temp.X += cord0.X
	temp.Y += cord0.Y
	return temp
}

func (cord CoordInts) IsEqual(cord0 CoordInts) bool {
	return cord.X == cord0.X && cord.Y == cord0.Y
}

func (cord *CoordInts) GetDifferences_TwoInts(cord0 CoordInts) (int, int) {
	xx := cord0.X - cord.X
	yy := cord0.Y - cord.Y
	return xx, yy
}
func (cord *CoordInts) GetManhattanDistance_Int(cord0 CoordInts) int {
	xx, yy := cord.GetDifferences_TwoInts(cord0)
	if xx < 0 {
		xx = xx * -1
	}
	if yy < 0 {
		yy = yy * -1
	}
	return xx + yy
}
func (cord *CoordInts) GetHypotenuseDistance_Int(cord0 CoordInts) int {
	xx, yy := cord.GetDifferences_TwoInts(cord0)
	return int(math.Sqrt(float64(math.Pow(float64(xx), 2) + math.Pow(float64(yy), 2))))
}

/*
returns the hypotenuse distance
*/
func (cord *CoordInts) GetHypotenuseDistance_Float(cord0 CoordInts) float64 {
	xx, yy := cord.GetDifferences_TwoInts(cord0)
	return math.Sqrt(float64(math.Pow(float64(xx), 2) + math.Pow(float64(yy), 2)))
}

/*
returns X,Y as two floating points
*/
func (cord CoordInts) GetValuesFloats() (x, y float64) {
	x = float64(cord.X)
	y = float64(cord.Y)
	return x, y
}
