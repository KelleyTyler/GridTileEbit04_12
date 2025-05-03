package coords

import (
	"fmt"
	"math"
	"slices"
)

type CoordList []CoordInts

/**/
func (cList *CoordList) ToString() string {
	return fmt.Sprintf("Length:%d --- ", len(*cList))
}

/**/
func (cList *CoordList) ToStringEntirely() []string {
	strng := []string{}
	strng = append(strng, fmt.Sprintf("Length:%d", len(*cList)))
	for _, a := range *cList {
		strng = append(strng, a.ToString())
	}
	return strng
}

/**/
func (cList *CoordList) ToPrint() {
	fmt.Printf("%s\n", cList.ToString())
	for _, a := range *cList {
		fmt.Printf("%s\n", a.ToString())
	}

}

/**/
func (cList CoordList) GetLength() int {
	return len(cList)

}

/**/
func (cList *CoordList) SortByX() {
	temp := make(CoordList, len(*cList))
	copy(temp, *cList)
	if len(temp) > 1 {
		for range temp {
			for i := 1; i < len(temp); i++ {
				if temp[i].X > temp[i-1].X {
					tempval := temp[i]
					temp[i] = temp[i-1]
					temp[i-1] = tempval
				}
			}
		}
	}
	copy(*cList, temp)
}

/**/
func (cList *CoordList) SortBy(val int, desc bool) {
	temp := make(CoordList, len(*cList))
	copy(temp, *cList)
	if len(temp) > 1 {
		for range temp {
			for i := 1; i < len(temp); i++ {
				var a0 int
				var b0 int
				// var a1, b1 int
				switch val {
				case 0: //just x;
					a0 = temp[i].X
					b0 = temp[i-1].X
				case 1: //X Then Y;
					a0 = temp[i].X
					b0 = temp[i-1].X
					if a0 == b0 {
						a0 = temp[i].Y
						b0 = temp[i-1].Y
					}
				case 2: //just Y;
					a0 = temp[i].Y
					b0 = temp[i-1].Y
				case 3: //Y_ThenX;
					a0 = temp[i].Y
					b0 = temp[i-1].Y
					if a0 == b0 {
						a0 = temp[i].X
						b0 = temp[i-1].X
					}
				default:
					a0 = temp[i].X
					b0 = temp[i-1].X
				}
				if desc {
					if a0 > b0 {
						tempval := temp[i]
						temp[i] = temp[i-1]
						temp[i-1] = tempval
					}
				} else {
					if a0 < b0 {
						tempval := temp[i]
						temp[i] = temp[i-1]
						temp[i-1] = tempval
					}
				}
			}
		}
	}
	copy(*cList, temp)
}

// func (cList *CoordList) Length() int {
// 	return len(*cList)
// }
/**/
func (cList *CoordList) PopFromFront() CoordInts {
	temp := make(CoordList, len(*cList))
	copy(temp, *cList)

	if len(*cList) > 1 {
		tempval := temp[0]
		temp = slices.Delete(temp, 0, 1)
		copy(*cList, temp)
		return tempval
	} else {
		tempval := temp[0]
		*cList = make(CoordList, 0)
		return tempval
	}

}

/**/
func (cList *CoordList) PushToFront(cord CoordInts) {
	temp0 := make(CoordList, len(*cList))
	copy(temp0, *cList)
	temp1 := make(CoordList, 0)
	temp1 = append(temp1, cord)
	temp1 = append(temp1, temp0...)
	copy(*cList, temp1)

}

/**/
func (cList *CoordList) PushToBack(cord CoordInts) {
	temp0 := make(CoordList, len(*cList))
	copy(temp0, *cList)

	temp0 = append(temp0, cord)
	*cList = make(CoordList, len(temp0))
	copy(*cList, temp0)

}

/**/
func (cList *CoordList) RemoveFromIndex(index int) {
	if index < len(*cList) {
		temp := make(CoordList, 0)
		// copy(temp, *cList)
		for i, a := range *cList {
			if i != index {
				temp = append(temp, a)
			}
		}
		*cList = make(CoordList, len(temp))
		copy(*cList, temp)
		// if index < cList.Length()-1 {
		// 	temp := make(CoordList, cList.Length())
		// 	copy(temp, *cList)
		// 	temp = slices.Delete(temp, index, index+1)
		// 	copy(*cList, temp)
		// } else {
		// 	temp := make(CoordList, cList.Length())
		// 	copy(temp, *cList)
		// 	temp = slices.Delete(temp, index, index+1)
		// 	copy(*cList, temp)
		// }
	}
}

/**/
func (cList *CoordList) RemoveCoord(coord CoordInts) {
	if len(*cList) > 0 {
		temp := make(CoordList, 0)
		// copy(temp, *cList)
		for _, a := range *cList {
			if !a.IsEqual(coord) {
				temp = append(temp, a)
			}
		}
		*cList = make(CoordList, len(temp))
		copy(*cList, temp)
	}
}

/**/
func (cList *CoordList) RemoveDuplicates() {
	if len(*cList) > 0 {
		temp := make(CoordList, 0)
		//copy(temp, *cList)
		//temp2 := make(CoordList, 0)
		seenMap := make(map[CoordInts]bool)
		for range len(*cList) {
			for _, a := range *cList {
				if !seenMap[a] {
					seenMap[a] = true
					temp = append(temp, a)
				}
			}
		}
		// fmt.Printf("")
		*cList = make(CoordList, len(temp))
		copy(*cList, temp)
	}
}

/**/
func (cList *CoordList) IfListContains(cord CoordInts) bool {
	outer := false
	for _, a := range *cList {
		if a.IsEqual(cord) {
			outer = true
		}
	}
	return outer
}

/**/
func (cList *CoordList) FlipOrder() {
	if len(*cList) > 0 {
		temp := make(CoordList, len(*cList))
		copy(temp, *cList)
		//flipping

		slices.Reverse(temp)
		copy(*cList, temp)
	}
}

// func (cList CoordList) IfListContains(cord CoordInts) bool {
// 	outer := false
// 	if len(cList) > 100 {
// 		for i := 0; i < len(cList)/2; i++ {
// 			a := cList[i]
// 			b := cList[len(cList)-(i+1)]
// 			if a.IsEqualTo(cord) {
// 				outer = true
// 			}
// 			if b.IsEqual(cord) {
// 				outer = true
// 			}
// 		}
// 	} else {
// 		for _, a := range cList {
// 			if a.IsEqual(cord) {
// 				outer = true
// 			}
// 		}
// 	}
// 	return outer
// }
/**/
func BresenhamLine(c1 CoordInts, c2 CoordInts) CoordList {
	//outList := make(CoordList, 0)
	var outList CoordList
	if math.Abs(float64(c2.Y)-float64(c1.Y)) < math.Abs(float64(c2.X)-float64(c1.X)) {
		if c1.X > c2.X {
			//fmt.Printf("BRESENHAM:%16s \n", "Low Inverted")
			outList = BresenhamLine_Low(c2, c1)
			// outList = StraightDiagonal(c2, c2)

			outList.FlipOrder()
		} else {
			//fmt.Printf("BRESENHAM:%16s \n", "Low Regular")
			// outList = StraightDiagonal(c1, c2)
			outList = BresenhamLine_Low(c1, c2)
		}
	} else {
		if c1.Y > c2.Y {
			//fmt.Printf("BRESENHAM:%16s \n", "High Inverted")
			outList = BresenhamLine_High(c2, c1)
			// outList = StraightDiagonal(c1, c2)
			outList.FlipOrder()
		} else {
			//fmt.Printf("BRESENHAM:%16s \n", "High Regular")
			outList = BresenhamLine_High(c1, c2)
		}
	}
	if !outList[0].IsEqual(c1) {
		outList = append(outList, c1)
		//outList.PushToFront(c1)
	}
	if !outList[len(outList)-1].IsEqual(c2) {
		outList = append(outList, c2)
	}
	return outList
}

/**/
func BresenhamLine_Low(c1 CoordInts, c2 CoordInts) CoordList {
	outList := make(CoordList, 0)
	dx := (c2.X - c1.X)
	dy := (c2.Y - c1.Y)
	yi := 1
	if dy < 0 {
		yi = -1
		dy = -dy
	}
	D := (2 * dy) - dx
	y := (c1.Y)
	//fmt.Printf("\tdx,dy:%3d,%3d\n\tyi:%3d y:%d\n\tInitial D:%d\n", dx, dy, yi, y, D)
	//need a conditional here;
	for x := c1.X; x < c2.X; x++ {
		//add to array here

		if D > 0 {
			y = y + yi
			D = D + (2 * (dy - dx))
		} else {
			D = D + (2 * dy)
		}
		//fmt.Printf("\n\t x:%3d y:%3d D:%3d\n", x, y, D)
		outList = append(outList, CoordInts{X: x, Y: int(y)})
	}
	outList = append(outList, c2)
	return outList
}

/**/
func BresenhamLine_High(c1 CoordInts, c2 CoordInts) CoordList {
	outList := make(CoordList, 0)
	dx := (c2.X - c1.X)
	dy := (c2.Y - c1.Y)
	xi := 1
	if dx < 0 {
		xi = -1
		dx = -dx
	}
	D := (2 * dx) - dy
	x := (c1.X)

	for y := c1.Y; y < c2.Y; y++ {
		//Add to array here;
		if D > 0 {
			x = x + xi
			D = D + (2 * (dx - dy))
		} else {
			D = D + (2 * dx)
		}
		outList = append(outList, CoordInts{X: int(x), Y: y})
	}
	outList = append(outList, c2)
	return outList
}

/*
void bresenham(int x1, int y1, int x2, int y2)

	{
	    int m_new = 2 * (y2 - y1);
	    int slope_error_new = m_new - (x2 - x1);
	    for (int x = x1, y = y1; x <= x2; x++) {
	        cout << "(" << x << "," << y << ")\n";

	        // Add slope to increment angle formed
	        slope_error_new += m_new;

	        // Slope error reached limit, time to
	        // increment y and update slope error.
	        if (slope_error_new >= 0) {
	            y++;
	            slope_error_new -= 2 * (x2 - x1);
	        }
	    }
	}
*/
func Straight_Diagonal(c1 CoordInts, c2 CoordInts) (outlist CoordList) {
	outlist = make(CoordList, 0)
	x1 := c1.X
	y1 := c1.Y
	x2 := c2.X
	y2 := c2.Y
	m_new := (2 * (y2 - y1))
	slope_error_new := m_new - (x2 - x1)
	y := y1
	for x := x1; x <= x2; x++ {
		outlist = append(outlist, CoordInts{X: x, Y: y})
		slope_error_new += m_new
		if slope_error_new >= 0 {
			y++
			slope_error_new -= (2*x2 - x1)
		}
	}
	return outlist
}

func (center *CoordInts) GetACirclePointsFromCenter(radius int) CoordList {
	tempList := make(CoordList, 0)
	P := 1 - radius
	x := radius
	y := 0

	for x > y {
		y++
		if P <= 0 {
			P = P + 2*y + 1
		} else {
			x--
			P = P + 2*y - 2*x + 1
		}

		tempList = append(tempList, center.GetACirclePointsSUB(x, y, radius)...)
		if x < y {
			break
		}
	}
	tempList = append(tempList, CoordInts{X: center.X + radius, Y: center.Y})
	tempList = append(tempList, CoordInts{X: center.X, Y: center.Y + radius})
	tempList = append(tempList, CoordInts{X: center.X - radius, Y: center.Y})
	tempList = append(tempList, CoordInts{X: center.X, Y: center.Y - radius})
	tempList.RemoveDuplicates()
	return tempList
}

/**/
func (center *CoordInts) GetACirclePointsSUB(x, y, radius int) CoordList {
	tempList := make(CoordList, 0)
	temp_01A := *center
	temp_01A.X += x
	temp_01A.Y += y

	temp_01B := *center
	temp_01B.X -= x
	temp_01B.Y += y

	temp_02A := *center
	temp_02B := *center
	temp_02A.X += x
	temp_02A.Y -= y
	//-----------------
	temp_02B.X -= x
	temp_02B.Y -= y
	//----------------------
	tempList = append(tempList, temp_01A)
	tempList = append(tempList, temp_01B)
	tempList = append(tempList, temp_02A)
	tempList = append(tempList, temp_02B)

	if x != y { //
		tempList = append(tempList, CoordInts{X: center.X + y, Y: center.Y + x})
		tempList = append(tempList, CoordInts{X: center.X - y, Y: center.Y + x})
		tempList = append(tempList, CoordInts{X: center.X + y, Y: center.Y - x})
		tempList = append(tempList, CoordInts{X: center.X - y, Y: center.Y - x})
		// tempList = append(tempList, CoordInts{X: center.Y - x, Y: center.X + x})
		// tempList = append(tempList, CoordInts{X: center.X - x, Y: center.X - y})
		// tempList = append(tempList, CoordInts{X: center.Y + x, Y: center.X - x})

	}
	return tempList
}
