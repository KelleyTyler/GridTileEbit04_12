package coords

import (
	"fmt"
	"slices"
)

type CoordList []CoordInts

func (cList *CoordList) ToString() string {
	return fmt.Sprintf("Length:%d --- ", len(*cList))
}
func (cList *CoordList) ToStringEntirely() []string {
	strng := []string{}
	strng = append(strng, fmt.Sprintf("Length:%d", len(*cList)))
	for _, a := range *cList {
		strng = append(strng, a.ToString())
	}
	return strng
}

func (cList *CoordList) ToPrint() {
	fmt.Printf("%s\n", cList.ToString())
	for _, a := range *cList {
		fmt.Printf("%s\n", a.ToString())
	}

}

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
func (cList *CoordList) SortBy(val int, desc bool) {
	temp := make(CoordList, len(*cList))
	copy(temp, *cList)
	if len(temp) > 1 {
		for range temp {
			for i := 1; i < len(temp); i++ {
				var a0 int
				var b0 int
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

func (cList *CoordList) PushToFront(cord CoordInts) {
	temp0 := make(CoordList, len(*cList))
	copy(temp0, *cList)
	temp1 := make(CoordList, 0)
	temp1 = append(temp1, cord)
	temp1 = append(temp1, temp0...)
	copy(*cList, temp1)

}
func (cList *CoordList) PushToBack(cord CoordInts) {
	temp0 := make(CoordList, len(*cList))
	copy(temp0, *cList)

	temp0 = append(temp0, cord)
	*cList = make(CoordList, len(temp0))
	copy(*cList, temp0)

}

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

func (cList *CoordList) IfListContains(cord CoordInts) bool {
	outer := false
	for _, a := range *cList {
		if a.IsEqual(cord) {
			outer = true
		}
	}
	return outer
}
