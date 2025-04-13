package matrix

import (
	"fmt"

	coords "github.com/KelleyTyler/GridTileEbit04_12/myPkgs/basic_geometry/coords"
)

type IntegerMatrix2D [][]int

func makeIntegerMatrix(numRows, numColumns, intialValue int) IntegerMatrix2D {
	temp := make(IntegerMatrix2D, numRows)
	for i := range temp {
		temp[i] = make([]int, numColumns)
		for j := range temp[i] {
			temp[i][j] = intialValue
		}
	}
	return temp
}

func (imat *IntegerMatrix2D) Init(nRows, nColumns, initValue int) {
	*imat = makeIntegerMatrix(nRows, nColumns, initValue)
}

func (imat IntegerMatrix2D) GetStrings() []string {
	sy, sx := imat.GetSize()
	outstrng := []string{fmt.Sprintf("INTEGER MATRIX 2D: %d, %d", sy, sx)}
	for i := range sy {
		outstrng = append(outstrng, "")
		for j := range sx {
			outstrng[i] += fmt.Sprintf("[%3d]", imat[i][j])
		}
	}
	return outstrng
}

func (imat IntegerMatrix2D) GetSize() (int, int) {
	outRows, outCols := 0, 0
	if len(imat) > 0 {
		outRows = len(imat)
		if len(imat[0]) > 0 {
			outCols = len(imat[0])
		}
	}
	return outRows, outCols
}

func (imat IntegerMatrix2D) IsValidCoords(c coords.CoordInts) bool {
	ylim, xlim := imat.GetSize()
	return !(c.X < 0 || c.X > xlim-1) && !(c.Y < 0 || c.Y > ylim-1)
}

func (imat IntegerMatrix2D) GetValueOnCoord(c coords.CoordInts) int {
	if imat.IsValidCoords(c) {
		return imat[c.Y][c.X]
	} else {
		return -1
	}
}

func (imat IntegerMatrix2D) GetNeighbors() {}
