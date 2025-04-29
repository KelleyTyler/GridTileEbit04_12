package matrix

import (
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
	// ZEROS := 0

	slices.SortFunc(list, func(a, e *ImatNode) int {
		aVal := a.GetFValue()
		eVal := e.GetFValue()
		if aVal < eVal {
			return 1
		} else if aVal > eVal {
			return -1
		} else {
			a_HVal := a.GetHValue()
			e_HVal := e.GetHValue()
			if a_HVal > e_HVal {
				return 1
			} else if a_HVal < e_HVal {
				return -1
			} else {
				// ZEROS++
				return 0
			}
		}
	})
	// fmt.Printf("ZEROS: %d --------\n", ZEROS)

}

/**/
func NodeList_SortByFValue_DESC(list []*ImatNode, start, target coords.CoordInts) {
	// ZEROS := 0
	slices.SortFunc(list, func(a, e *ImatNode) int {
		aVal := a.GetFValue()
		eVal := e.GetFValue()
		if aVal > eVal {
			return 1
		} else if aVal < eVal {
			return -1
		} else {
			a_HVal := a.GetHValue()
			e_HVal := e.GetHValue()
			if a_HVal > e_HVal {
				return 1
			} else if a_HVal < e_HVal {
				return -1
			} else {
				// ZEROS++
				// fmt.Printf("--------\n")
				return 0
			}
		}
	})
	// fmt.Printf("ZEROS: %d --------\n", ZEROS)

}

/**/
func NodeList_SortByHValue_Ascending(list []*ImatNode) {
	slices.SortFunc(list, func(a, e *ImatNode) int {
		aVal := a.GetHValue()
		eVal := e.GetHValue()
		if aVal < eVal {
			return 1
		} else if aVal > eVal {
			return -1
		} else {
			return 0
		}
	})
}

/**/
func NodeList_RemoveDuplicates_ToReturn(inList []*ImatNode) (outList []*ImatNode) {
	outList = make([]*ImatNode, len(inList))
	copy(outList, inList)
	//----------------
	NodeList_SortByHValue_Ascending(inList)
	posmap := make(map[coords.CoordInts]float64)
	for _, a := range inList {
		val := posmap[a.Position]
		if val == 0 {
			posmap[a.Position] = (a.GetHValue())
		} else {
			if val > (a.GetHValue()) {
				outList = NodeList_RemoveByPosition_Return(outList, a.Position)
				outList = append(outList, a)
				posmap[a.Position] = (a.GetHValue())
			}
		}
	}

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
		if aVal < eVal {
			return 1
		} else if aVal > eVal {
			return -1
		} else {
			/*
				a_HVal := a.GetHValue()
				e_HVal := e.GetHValue()
				if a_HVal > e_HVal {
					return 1
				} else if a_HVal < e_HVal {
					return -1
				} else {
					return 0
				}
			*/
			return 0
		}
	})
	return outlist
}

/*
This function sorts the slice by
*/
func NodeList_SortByFValue_Desc_toReturn(list []*ImatNode, start, target coords.CoordInts) (outlist []*ImatNode) {
	outlist = make([]*ImatNode, len(list))
	copy(outlist, list)

	slices.SortFunc(outlist, func(a, e *ImatNode) int {
		aVal := a.GetFValue()
		eVal := e.GetFValue()
		if aVal > eVal {
			return 1
		} else if aVal < eVal {
			return -1
		} else {
			//after much trial and error:
			a_HVal := a.GetHValue()
			e_HVal := e.GetHValue()
			if a_HVal > e_HVal {
				return 1
			} else if a_HVal < e_HVal {
				return -1
			} else {
				return 0
			}
		}
	})
	return outlist
}

/**/
func NodeList_SortByHValue_Ascending_toReturn(list []*ImatNode) (outlist []*ImatNode) {
	outlist = make([]*ImatNode, len(list))
	copy(outlist, list)

	slices.SortFunc(outlist, func(a, e *ImatNode) int {
		aVal := a.GetHValue()
		eVal := e.GetHValue()
		if aVal < eVal {
			return -1
		} else if aVal > eVal {
			return 1
		} else {
			return 0
		}
	})
	return outlist
}

/**/
func NodeList_FILTER_LIST(in_list_00, in_list01 []*ImatNode) (outlist_00 []*ImatNode) {
	// outlist_00 = make([]*ImatNode, len(in_list_00))
	// copy(outlist_00, in_list_00)
	outlist_00 = make([]*ImatNode, 0)

	for _, node := range in_list_00 {
		if !NodeList_ContainsPoint(node.Position, in_list01) {
			outlist_00 = append(outlist_00, node)
		}
	}

	return outlist_00
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
			if !c.IsEqual(node.Prev.Position) {
				if !misc.IsNumInIntArray(imat.GetValueOnCoord(c), walls) {
					temp := GetNode(start, c, target, *imat, node)
					retlist = append(retlist, &temp)
				}
			} else if node.Position.IsEqual(start) {
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
returns up to 4 nodes;
distances are assigned based on hypotenuse distance;
*/
func NodeList_GetNeighbors_4_Filtered_Hypentenuse(node *ImatNode, start, target coords.CoordInts, imat *IntegerMatrix2D, floors, walls []int, margins [4]uint) (retlist []*ImatNode) {
	retlist = make([]*ImatNode, 0)
	//templist_int
	templist_coord, _ := imat.GetNeighborsAndValues_Cardinal(node.Position, margins)
	for _, c := range templist_coord {
		if node.Prev != nil {
			if !c.IsEqual(node.Prev.Position) {
				if !misc.IsNumInIntArray(imat.GetValueOnCoord(c), walls) {
					temp := GetNode(start, c, target, *imat, node)
					temp.Target = target
					retlist = append(retlist, &temp)

					//fmt.Printf("%s IS GOOD\n", c.ToString())
				} else {
					//fmt.Printf("%s has problem\n", c.ToString())
				}
			} else {
				//fmt.Printf("%s is Equal\n", c.ToString())
			}

		} else {
			if !misc.IsNumInIntArray(imat.GetValueOnCoord(c), walls) {
				temp := GetNode(start, c, target, *imat, node)
				temp.Target = target
				retlist = append(retlist, &temp)

				//fmt.Printf("%s IS GOOD\n", c.ToString())

			} else {
				//fmt.Printf("%s has problem\n", c.ToString())
			}
		}

	}
	return retlist
}

/*
returns 4 nodes;
distances are assigned based on hypotenuse distance;
*/
func NodeList_GetNeighbors_4A_Filtered_Hypentenuse(node *ImatNode, start, target coords.CoordInts, imat *IntegerMatrix2D, floors, walls []int, margins [4]uint) (retlist [4]*ImatNode) {
	// retlist
	//templist_int
	templist_coord, vals := imat.GetNeighborsAndValues_Cardinal(node.Position, margins)
	for i, c := range templist_coord {
		temp := GetNode(start, c, target, *imat, node)
		temp.ValueOnGrid = vals[i]
		temp.Target = target
		retlist[i] = &temp
	}
	return retlist
}

/*
 */
func NodeList_PopFromFront(inArray []*ImatNode) (reNode *ImatNode, retArray []*ImatNode) {
	retArray = make([]*ImatNode, 0)
	// NodeList_SortByHValue_Ascending(inArray)
	reNode = inArray[0]
	if len(inArray) > 1 {
		retArray = append(retArray, inArray[1:]...)
	}
	return reNode, retArray
}

/**/
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

/*
 */
func NodeList_RemoveByPosition_Return(inArray []*ImatNode, position coords.CoordInts) (retArray []*ImatNode) {
	retArray = make([]*ImatNode, 0)
	for _, a := range inArray {
		if !a.Position.IsEqual(position) {
			retArray = append(retArray, a)
		}
	}
	return retArray
}

// func NodesTest() {
// 	nodes_00 := make([]*ImatNode, 0)
// 	node1:=
// }
