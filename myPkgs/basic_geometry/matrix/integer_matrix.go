package matrix

import (
	"fmt"

	coords "github.com/KelleyTyler/GridTileEbit04_12/myPkgs/basic_geometry/coords"
	misc "github.com/KelleyTyler/GridTileEbit04_12/myPkgs/misc"
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

/*
 */
func (imat IntegerMatrix2D) SetValAtCoord(c coords.CoordInts, value int) bool {
	if imat.IsValidCoords(c) {
		imat[c.Y][c.X] = value
		return true
	} else {
		return false
	}
}
func (imat IntegerMatrix2D) SetValAtCoord_Filtered(c coords.CoordInts, value int, filter []int) (bool, int) {
	if imat.IsValidCoords(c) {
		tempval := imat.GetValueOnCoord(c)
		isTrue, retVal := misc.IsNumInIntArrayAndWhat_boolInt(tempval, filter)
		if !isTrue {
			imat[c.Y][c.X] = value
			return true, -1
		} else {
			return false, retVal
		}
	} else {
		return false, -10
	}
}
func (imat IntegerMatrix2D) GetStrings() []string {
	sy, sx := imat.GetSize()
	outstrng := []string{fmt.Sprintf("INTEGER MATRIX 2D: %d, %d", sy, sx)}
	for i := 1; i < sy+1; i++ {
		outstrng = append(outstrng, "")
		for j := range sx {
			outstrng[i] += fmt.Sprintf("[%3d]", imat[i-1][j])
		}
		// outstrng[i]
	}
	return outstrng
}
func (imat IntegerMatrix2D) GetStrings_withCoordList(c coords.CoordList) []string {
	sy, sx := imat.GetSize()
	outstrng := []string{fmt.Sprintf("INTEGER MATRIX 2D: %d, %d", sy, sx)}
	for i := 1; i < sy+1; i++ {
		outstrng = append(outstrng, "")
		for j := range sx {
			if (c.IfListContains(coords.CoordInts{X: j, Y: i - 1})) {
				outstrng[i] += fmt.Sprintf("[%3s]", "-X-")
			} else {
				outstrng[i] += fmt.Sprintf("[%3d]", imat[i-1][j])
			}

		}
		// outstrng[i]
	}
	return outstrng
}
func (imat IntegerMatrix2D) PrintStrings() {
	strngs := imat.GetStrings()
	for i, s := range strngs {
		fmt.Printf("%d\t %s\n", i-1, s)
	}
}
func (imat IntegerMatrix2D) PrintStringsWithCoorList(c coords.CoordList) {
	strngs := imat.GetStrings_withCoordList(c)
	strngs0 := imat.GetStrings()

	for i, s := range strngs {
		fmt.Printf("%d\t %s\t%s\n", i, strngs0[i], s)
	}
}
func (imat IntegerMatrix2D) GetSize() (yy int, xx int) {
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
func (imat IntegerMatrix2D) IsValidCoordsWithinMargins(margins [4]uint, c coords.CoordInts) bool {
	ylim, xlim := imat.GetSize()
	con := !(c.X < int(margins[3]) || c.X > xlim-int(margins[1]+1)) && !(c.Y < int(margins[0]) || c.Y > ylim-int(margins[2]+1))
	// if !con {
	// 	fmt.Printf("%s is not in margins", c.ToString())
	// }
	return con
}
func (imat IntegerMatrix2D) GetValueOnCoord(c coords.CoordInts) int {
	if imat.IsValidCoords(c) {
		return imat[c.Y][c.X]
	} else {
		return -1
	}
}

func (imat *IntegerMatrix2D) ClearMatrix_To(value int) {
	yy, xx := imat.GetSize()
	for i := range yy {
		for j := range xx {
			imat.SetValAtCoord(coords.CoordInts{X: j, Y: i}, value)
		}
	}
}

func (imat IntegerMatrix2D) GetNeighborsAndValues_Cardinal(c coords.CoordInts, margins [4]uint) ([4]coords.CoordInts, [4]int) {
	outlist := [4]coords.CoordInts{}
	valList := [4]int{}
	var temp coords.CoordInts
	if imat.IsValidCoords(c) {
		temp = c.AddToReturn(coords.CoordInts{X: 0, Y: -1})
		outlist[0] = temp //north
		if imat.IsValidCoordsWithinMargins(margins, temp) {
			valList[0] = imat.GetValueOnCoord(outlist[0])
		} else {
			valList[0] = -1
		}

		temp = c.AddToReturn(coords.CoordInts{X: 1, Y: 0})
		outlist[1] = temp //east
		if imat.IsValidCoordsWithinMargins(margins, temp) {
			valList[1] = imat.GetValueOnCoord(outlist[1])
		} else {
			valList[1] = -1
		}

		temp = c.AddToReturn(coords.CoordInts{X: 0, Y: 1})
		outlist[2] = temp //south
		if imat.IsValidCoordsWithinMargins(margins, temp) {
			valList[2] = imat.GetValueOnCoord(outlist[2])
		} else {
			valList[2] = -1
		}

		temp = c.AddToReturn(coords.CoordInts{X: -1, Y: 0})
		outlist[3] = temp //west
		if imat.IsValidCoordsWithinMargins(margins, temp) {
			valList[3] = imat.GetValueOnCoord(outlist[3])
			//fmt.Printf("for 3: %s\t the value on coord is: %d\n", temp.ToString(), imat.GetValueOnCoord(outlist[3]))
		} else {
			valList[3] = -1
		}
	}
	return outlist, valList
}
func (imat IntegerMatrix2D) GetNeighborsAndValues_8(c coords.CoordInts, margins [4]uint) ([8]coords.CoordInts, [8]int) {
	outlist := [8]coords.CoordInts{}
	valList := [8]int{}
	var temp coords.CoordInts
	if imat.IsValidCoords(c) {
		temp = c.AddToReturn(coords.CoordInts{X: 0, Y: -1})
		outlist[0] = temp //north
		if imat.IsValidCoords(temp) {
			valList[0] = imat.GetValueOnCoord(outlist[0])
		} else {
			valList[0] = -1
		}
		temp = c.AddToReturn(coords.CoordInts{X: 1, Y: -1})
		outlist[1] = temp //northeast
		if imat.IsValidCoords(temp) {
			valList[1] = imat.GetValueOnCoord(outlist[1])
		} else {
			valList[1] = -1
		}

		temp = c.AddToReturn(coords.CoordInts{X: 1, Y: 0})
		outlist[2] = temp //east
		if imat.IsValidCoords(temp) {
			valList[2] = imat.GetValueOnCoord(outlist[2])
		} else {
			valList[2] = -1
		}
		temp = c.AddToReturn(coords.CoordInts{X: 1, Y: 1})
		outlist[3] = temp //northeast
		if imat.IsValidCoords(temp) {
			valList[3] = imat.GetValueOnCoord(outlist[3])
		} else {
			valList[3] = -1
		}
		temp = c.AddToReturn(coords.CoordInts{X: 0, Y: 1})
		outlist[4] = temp //south
		if imat.IsValidCoords(temp) {
			valList[4] = imat.GetValueOnCoord(outlist[4])
		} else {
			valList[4] = -1
		}
		temp = c.AddToReturn(coords.CoordInts{X: -1, Y: 1})
		outlist[5] = temp //south
		if imat.IsValidCoords(temp) {
			valList[5] = imat.GetValueOnCoord(outlist[5])
		} else {
			valList[5] = -1
		}
		temp = c.AddToReturn(coords.CoordInts{X: -1, Y: 0})
		outlist[6] = temp //west
		if imat.IsValidCoords(temp) {
			valList[6] = imat.GetValueOnCoord(outlist[6])
		} else {
			valList[6] = -1
		}
		temp = c.AddToReturn(coords.CoordInts{X: -1, Y: -1})
		outlist[7] = temp //north west
		if imat.IsValidCoords(temp) {
			valList[7] = imat.GetValueOnCoord(outlist[7])
		} else {
			valList[7] = -1
		}
	}
	return outlist, valList
}

func (imat IntegerMatrix2D) GetValidNeighbors4_no_Order(c coords.CoordInts) []coords.CoordInts {
	outlist := []coords.CoordInts{}
	var temp coords.CoordInts
	if imat.IsValidCoords(c) {
		temp = c.AddToReturn(coords.CoordInts{X: 0, Y: -1})
		if imat.IsValidCoords(temp) {
			outlist = append(outlist, temp) //north
		}

		temp = c.AddToReturn(coords.CoordInts{X: 1, Y: 0})
		if imat.IsValidCoords(temp) {
			outlist = append(outlist, temp) //east
		}

		temp = c.AddToReturn(coords.CoordInts{X: 0, Y: 1})
		if imat.IsValidCoords(temp) {
			outlist = append(outlist, temp) //south
		}

		temp = c.AddToReturn(coords.CoordInts{X: -1, Y: 0})
		if imat.IsValidCoords(temp) {
			outlist = append(outlist, temp) //west
		}
	}
	return outlist
}
func (imat IntegerMatrix2D) GetNeighborsAndValues_unordered(c coords.CoordInts) ([]coords.CoordInts, []int) {
	outlist := []coords.CoordInts{}
	valList := []int{}
	var temp coords.CoordInts
	if imat.IsValidCoords(c) {
		temp = c.AddToReturn(coords.CoordInts{X: 0, Y: -1})
		//north
		if imat.IsValidCoords(temp) {
			outlist = append(outlist, temp)
			valList = append(valList, imat.GetValueOnCoord(temp))
		}
		//northeast
		temp = c.AddToReturn(coords.CoordInts{X: 1, Y: -1})

		if imat.IsValidCoords(temp) {
			outlist = append(outlist, temp)
			valList = append(valList, imat.GetValueOnCoord(temp))
		}
		//east
		temp = c.AddToReturn(coords.CoordInts{X: 1, Y: 0})

		if imat.IsValidCoords(temp) {
			outlist = append(outlist, temp)
			valList = append(valList, imat.GetValueOnCoord(temp))
		}
		//southeast
		temp = c.AddToReturn(coords.CoordInts{X: 1, Y: 1})
		if imat.IsValidCoords(temp) {
			outlist = append(outlist, temp)
			valList = append(valList, imat.GetValueOnCoord(temp))
		}
		//south
		temp = c.AddToReturn(coords.CoordInts{X: 0, Y: 1})
		if imat.IsValidCoords(temp) {
			outlist = append(outlist, temp)
			valList = append(valList, imat.GetValueOnCoord(temp))
		}
		temp = c.AddToReturn(coords.CoordInts{X: -1, Y: 1})
		if imat.IsValidCoords(temp) {
			outlist = append(outlist, temp)
			valList = append(valList, imat.GetValueOnCoord(temp))
		}
		temp = c.AddToReturn(coords.CoordInts{X: -1, Y: 0})
		if imat.IsValidCoords(temp) {
			outlist = append(outlist, temp)
			valList = append(valList, imat.GetValueOnCoord(temp))
		}
		//
		temp = c.AddToReturn(coords.CoordInts{X: -1, Y: -1})
		if imat.IsValidCoords(temp) {
			outlist = append(outlist, temp)
			valList = append(valList, imat.GetValueOnCoord(temp))
		}
	}
	return outlist, valList
}
