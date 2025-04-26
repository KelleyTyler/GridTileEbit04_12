package matrix

import (
	"fmt"
	"log"

	coords "github.com/KelleyTyler/GridTileEbit04_12/myPkgs/basic_geometry/coords"
	"github.com/KelleyTyler/GridTileEbit04_12/myPkgs/misc"
)

func Pathfind(start, target coords.CoordInts, imat IntegerMatrix2D, floors, walls []int, failconditions int, margins [4]uint) (path_is_found bool, path_out coords.CoordList) {
	//setup
	path_is_found = false
	path_out = make(coords.CoordList, 0)
	return path_is_found, path_out
}
func Pathfind_Phase1A(start, target coords.CoordInts, imat IntegerMatrix2D, floors, walls []int, margins [4]uint) (path_is_found bool, EndNode *ImatNode) {

	path_is_found = false
	OpenList := make([]*ImatNode, 0)
	ClosedList := make([]*ImatNode, 0)
	BlockedList := make([]*ImatNode, 0)

	var max_fails int = 10000
	var curr_fails int = 0
	var isFinished bool = false
	var err error = nil
	EndNode = nil
	justStart := true
	startNode := GetNodePTR(start, start, target, imat, nil)
	ClosedList = append(ClosedList, startNode)
	// ClosedList = NodeList_SortByFValue_Desc_toReturn(OpenList, start, target)
	temp := NodeList_GetNeighbors_4_Filtered_Hypentenuse(startNode, start, target, &imat, floors, walls, margins)
	// temp = NodeList_SortByFValue_Desc_toReturn(temp, start, target)
	NodeList_SortByFValue_DESC(temp, start, target)
	OpenList = append(OpenList, temp...)
	OpenList = NodeList_SortByFValue_Desc_toReturn(OpenList, start, target)
	//beware pseudocode
	//ClosedList.ClosedList = ClosedList.append(Get_New_Node(start, start, end))//*pseudocode or not I'm not sure I even want this.
	// OpenList.append
	for !isFinished && curr_fails < max_fails {
		if !justStart {
			ClosedList = NodeList_SortByFValue_Desc_toReturn(ClosedList, start, target)
			temp := NodeList_GetNeighbors_4_Filtered_Hypentenuse(ClosedList[0], start, target, &imat, floors, walls, margins)
			temp = NodeList_FILTER_LIST(temp, ClosedList)
			temp = NodeList_FILTER_LIST(temp, OpenList)
			temp = NodeList_SortByFValue_Desc_toReturn(temp, start, target)

			OpenList = append(OpenList, temp...)
			// OpenList = NodeList_SortByFValue_Desc_toReturn(OpenList, start, target)
			// ClosedList = NodeList_SortByFValue_Desc_toReturn(OpenList, start, target)

		} else {
			justStart = false
		}
		OpenList, ClosedList, BlockedList, isFinished, EndNode, curr_fails, err = Pathfind_Phase1_Tick(start, target, OpenList, ClosedList, BlockedList, isFinished, curr_fails, max_fails, imat, floors, walls, margins)
		if err != nil {
			log.Fatal(fmt.Errorf("pathfinding error"))
		}
		curr_fails++
	}

	if isFinished {
		fmt.Printf("FINISHED! %d\n", curr_fails)
		if EndNode != nil {
			EndNode.Set_Heads_Tails_On_Up()
		}

		// PotentialPaths = append(PotentialPaths, ClosedList...)
	}

	return path_is_found, EndNode
}
func Pathfind_Phase1_Tick(start, target coords.CoordInts, openlist, closedlist, blockedlist []*ImatNode, pathfound bool, curr_fails, max_fails int, imat IntegerMatrix2D, floors, walls []int, margins [4]uint) (oList, cList, bList []*ImatNode, pFound bool, pFoundNode *ImatNode, fails int, err error) { //<--unsure if these are pass by reference or not
	oList = make([]*ImatNode, len(openlist))
	cList = make([]*ImatNode, len(closedlist))
	bList = make([]*ImatNode, len(blockedlist))
	copy(oList, openlist)
	copy(cList, closedlist)
	copy(bList, blockedlist)
	pFound = pathfound
	fails = curr_fails
	// bList = NodeList_FILTER_LIST(bList, oList) //
	// bList = NodeList_FILTER_LIST(bList, cList) //

	err = nil
	oList = NodeList_SortByFValue_Desc_toReturn(oList, start, target)
	// slices.Reverse(oList)
	//pfound_phase2 = false //(the path needs to be found more)
	//---------------------- Prep Done
	var pointQ *ImatNode
	if len(oList) > 0 {
		pointQ, oList = NodeList_PopFromFront(oList)

		// if err != nil {
		// 	return oList, cList, bList, false, fails, err
		// }
		if pointQ != nil {
			//fmt.Printf("Point Q is %s\t", pointQ.Position.ToString())
			if pointQ.Position.IsEqual(target) {
				// fmt.Printf("\n\nHIT TARGET\n\n\n")
				pFoundNode = pointQ
				pFound = true
				return oList, cList, bList, pFound, pFoundNode, fails, err

			} else if imat.IsValidCoords(pointQ.Position) && !misc.IsNumInIntArray(imat.GetValueOnCoord(pointQ.Position), walls) {
				//fmt.Printf("Point Q is Valid %7.2f %7.2f\n", pointQ.GetFValue(), pointQ.GetHValue())
				temp_successors := NodeList_GetNeighbors_4_Filtered_Hypentenuse(pointQ, start, target, &imat, floors, walls, margins)
				//filter temp_successors into openlist
				temp_successors = NodeList_FILTER_LIST(temp_successors, oList)
				temp_successors = NodeList_FILTER_LIST(temp_successors, cList)
				successfulNodes := 0

				for _, suc := range temp_successors {
					if !NodeList_ContainsPoint(suc.Position, cList) && !suc.Position.IsEqual(pointQ.Position) {
						if !NodeList_ContainsPoint(suc.Position, bList) {
							oList = append(oList, suc)
							// fmt.Printf("ADDED SUCCESSFULLY! %s\n", suc.Position.ToString())
							successfulNodes++
						} else {
							// fmt.Printf("BLOCKED BLOCKED\n")
							// fmt.Printf("BLOCKED  %s\n", suc.Position.ToString())

							//remove it

						}
						// cList = append(cList, suc)

						// fmt.Printf("Success--------\n")
					}

				}
				if successfulNodes == 0 {
					// bList = append(bList, pointQ)
					// fmt.Printf("Point Q (%s) is invalid %7.2f %7.2f \t length of oList: %4d, clist:%4d, blist:%4d\n", pointQ.Position.ToString(), pointQ.GetFValue(), pointQ.GetHValue(), len(oList), len(cList), len(bList))
					// return oList, cList, bList, pFound, nil, fails, err

				} else {
					//filter pointQ into closedlist comparing it with any overlap and resolving the contradiction;

				}
				override_cList := true
				for _, parent := range cList {
					if parent.Position.IsEqual(pointQ.Position) {
						if pointQ.GetFValue() < parent.GetFValue() {
							override_cList = true
						} else {
							override_cList = false
						}
					}
				}
				if override_cList {
					if !pointQ.Position.IsEqual(start) {

						cList = append(cList, pointQ)
					}
				}
			} else {
				// fmt.Printf("Point Q is NOT Valid \n")

			}

			oList = NodeList_SortByFValue_Desc_toReturn(oList, start, target) //f_value being sum of distance to and from
			oList = NodeList_RemoveDuplicates_ToReturn(oList)
			cList = NodeList_RemoveDuplicates_ToReturn(cList)
			//------NEEDS WORK
			// oList, cList, bList, err = Pathfind_Phase1_Tick_Blocked_List_Manager(oList, cList, bList, start, target, &imat, floors, walls, margins)
			// if err != nil {
			// 	return oList, cList, bList, false, nil, fails, err
			// }
			oList = NodeList_FILTER_LIST(oList, bList)
			cList = NodeList_FILTER_LIST(cList, bList)

		}

	}

	return oList, cList, bList, pFound, nil, fails, err
}

/*
While this is supposed to manage the 'blocked list' it presently doesn't work very well at all.
*/
func Pathfind_Phase1_Tick_Blocked_List_Manager(openlist, closedlist, blockedlist []*ImatNode, start, target coords.CoordInts, imat *IntegerMatrix2D, floors, walls []int, margins [4]uint) (oList, cList, bList []*ImatNode, err error) {
	oList = make([]*ImatNode, len(openlist))
	cList = make([]*ImatNode, len(closedlist))
	bList = make([]*ImatNode, len(blockedlist))
	copy(oList, openlist)
	copy(cList, closedlist)
	copy(bList, blockedlist)

	//---------------------- Prep Done
	if len(oList) > 5 {
		for _, nod := range oList {
			if !nod.Position.IsEqual(target) && !nod.Position.IsEqual(start) {
				temp := NodeList_GetNeighbors_4A_Filtered_Hypentenuse(nod, start, target, imat, floors, walls, margins) //<-----problem
				numb := 0
				for _, point := range temp {
					if !misc.IsNumInIntArray(imat.GetValueOnCoord(point.Position), walls) && !NodeList_ContainsPoint(point.Position, bList) {
						numb++
					}
				}
				if numb < 1 && !NodeList_ContainsPoint(nod.Position, bList) {
					bList = append(bList, nod)
					//remove from openlist
				}
			}
		}
		oList = NodeList_FILTER_LIST(oList, bList)
		if len(cList) > 50 {
			for _, nod := range cList {
				if !nod.Position.IsEqual(target) && !nod.Position.IsEqual(start) {
					temp := NodeList_GetNeighbors_4A_Filtered_Hypentenuse(nod, start, target, imat, floors, walls, margins)
					numb := 0
					for _, point := range temp {
						if !misc.IsNumInIntArray(imat.GetValueOnCoord(point.Position), walls) && !NodeList_ContainsPoint(point.Position, bList) {
							numb++
						}
					}
					if numb < 2 && !NodeList_ContainsPoint(nod.Position, bList) {
						bList = append(bList, nod)
						//remove from openlist
					}
				}
			}
		}
		cList = NodeList_FILTER_LIST(cList, bList)
	}

	return oList, cList, bList, nil
}
