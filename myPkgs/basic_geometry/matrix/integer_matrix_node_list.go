package matrix

import (
	"fmt"
	"slices"

	"github.com/KelleyTyler/GridTileEbit04_12/myPkgs/basic_geometry/coords"
	"github.com/KelleyTyler/GridTileEbit04_12/myPkgs/misc"
)

//type IntMatNodeList []ImatNode
//alternatively lets try some slices;

/*
 */
func NodeList_ContainsPoint(point coords.CoordInts, S []*ImatNode) bool {
	return slices.ContainsFunc(S, func(E *ImatNode) bool {
		return E.Position.IsEqual(point)
	})
}

/*
 */
func NodeList_SortByFValue_Ascending(list []*ImatNode, start, target coords.CoordInts) {
	slices.SortFunc(list, func(a, e *ImatNode) int {
		aVal := a.GetFValue()
		eVal := e.GetFValue()
		if a == e {
			return 0
		} else if aVal < eVal {
			return -1
		} else if aVal > eVal {
			return 1
		} else {
			return 0
		}
	})
}

func NodeList_RemoveDuplicates_ToReturn(inList []*ImatNode) (outList []*ImatNode) {
	outList = make([]*ImatNode, len(inList))
	copy(outList, inList)
	//----------------
	return outList
}

/*
 */
func NodeList_SortByFValue_Ascending_toReturn(list []*ImatNode, start, target coords.CoordInts) (outlist []*ImatNode) {
	outlist = make([]*ImatNode, len(list))
	copy(outlist, list)

	slices.SortFunc(outlist, func(a, e *ImatNode) int {
		aVal := a.GetFValue()
		eVal := e.GetFValue()
		if a == e {
			return 0
		} else if aVal < eVal {
			return -1
		} else if aVal > eVal {
			return 1
		} else {
			return 0
		}
	})
	return
}

/*
returns up to 8 nodes;
distances are assigned based on hypotenuse distance;
*/
func NodeList_GetNeighbors_8_Filtered_Hypentenuse(node *ImatNode, start, target coords.CoordInts, imat *IntegerMatrix2D, floors, walls []int, margins [4]uint) (retlist []*ImatNode) {
	retlist = make([]*ImatNode, 0)
	//templist_int
	templist_coord, _ := imat.GetNeighborsAndValues_8(node.Position, margins)
	for _, c := range templist_coord {
		if node.Prev != nil {
			if !c.IsEqualTo(node.Prev.Position) {
				if !misc.IsNumInIntArray(imat.GetValueOnCoord(c), walls) {
					temp := GetNode(start, c, target, *imat, node)
					retlist = append(retlist, &temp)
				}
			}

		}

	}
	return retlist
}

/*
returns up to 8 nodes;
distances are assigned based on hypotenuse distance;
*/
func NodeList_GetNeighbors_4_Filtered_Hypentenuse(node *ImatNode, start, target coords.CoordInts, imat *IntegerMatrix2D, floors, walls []int, margins [4]uint) (retlist []*ImatNode) {
	retlist = make([]*ImatNode, 0)
	//templist_int
	templist_coord, _ := imat.GetNeighborsAndValues_Cardinal(node.Position, margins)
	for _, c := range templist_coord {
		if node.Prev != nil {
			if !c.IsEqualTo(node.Prev.Position) {
				if !misc.IsNumInIntArray(imat.GetValueOnCoord(c), walls) {
					temp := GetNode(start, c, target, *imat, node)
					retlist = append(retlist, &temp)
					fmt.Printf("%s IS GOOD\n", c.ToString())
				} else {
					fmt.Printf("%s has problem\n", c.ToString())
				}
			} else {
				fmt.Printf("%s is Equal\n", c.ToString())
			}

		} else {
			if !misc.IsNumInIntArray(imat.GetValueOnCoord(c), walls) {
				temp := GetNode(start, c, target, *imat, node)
				retlist = append(retlist, &temp)
				fmt.Printf("%s IS GOOD\n", c.ToString())

			} else {
				fmt.Printf("%s has problem\n", c.ToString())
			}
		}

	}
	return retlist
}

/*
 */
func NodeList_PopFromFront(inArray []*ImatNode) (reNode *ImatNode, retArray []*ImatNode) {
	retArray = make([]*ImatNode, 0)
	reNode = inArray[0]
	if len(inArray) > 1 {
		retArray = append(retArray, inArray[1:]...)
	}
	return reNode, retArray
}

func NodeList_Convert_Nodes_To_CoordList(inNode *ImatNode) (ret_coordList coords.CoordList) {
	ret_coordList = make(coords.CoordList, 0)
	tempTail := inNode.GetHead()

	for tempTail.Next != nil {
		ret_coordList = append(ret_coordList, tempTail.Position)
		tempTail = tempTail.Next
	}
	return ret_coordList
}

/*
 */
func NodeList_Sort_By_G_Value() {

}
