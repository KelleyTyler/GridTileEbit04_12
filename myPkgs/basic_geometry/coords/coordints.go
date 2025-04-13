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
	cord.X += cord0.X
	cord.Y += cord0.Y
	return *cord
}

func (cord *CoordInts) IsEqualTo(cord0 CoordInts) bool {
	return cord.X == cord0.X && cord.Y == cord0.Y
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
