package matrix

import (
	"fmt"
	"log"
	"time"

	coords "github.com/KelleyTyler/GridTileEbit04_12/myPkgs/basic_geometry/coords"
	"github.com/KelleyTyler/GridTileEbit04_12/myPkgs/misc"
)

/**/
func Pathfind(start, target coords.CoordInts, imat IntegerMatrix2D, floors, walls []int, failconditions int, margins [4]uint) (path_is_found bool, path_out coords.CoordList) {
	//setup
	path_is_found = false
	path_out = make(coords.CoordList, 0)
	return path_is_found, path_out
}

/**/
func Pathfind_Phase1A(start, target coords.CoordInts, imat IntegerMatrix2D, floors, walls []int, margins [4]uint) (path_is_found bool, EndNode *ImatNode) {
	start_time := time.Now()
	path_is_found = false
	OpenList := make([]*ImatNode, 0)
	ClosedList := make([]*ImatNode, 0)
	BlockedList := make([]*ImatNode, 0)

	var max_fails int = 64 * 64 * 2
	var curr_fails int = 0
	var isFinished bool = false
	var closedList_LastLength int = 0
	var err error = nil
	log.Printf("PATHFINDER\n")

	EndNode = nil
	// justStart := true
	startNode := GetNodePTR(start, start, target, imat, nil)
	ClosedList = append(ClosedList, startNode)
	// ClosedList = NodeList_SortByFValue_Desc_toReturn(OpenList, start, target)
	temp := NodeList_GetNeighbors_4_Filtered_Hypentenuse(startNode, start, target, &imat, floors, walls, margins)
	if len(temp) > 0 {
		temp = NodeList_SortByFValue_Desc_toReturn(temp, start, target)
		NodeList_SortByFValue_DESC(temp, start, target)
		OpenList = append(OpenList, temp...)
		OpenList = NodeList_SortByFValue_Desc_toReturn(OpenList, start, target)
	} else {
		log.Printf("ERROR ERROR ERRROR")
		return false, nil
	}
	//beware pseudocode
	//ClosedList.ClosedList = ClosedList.append(Get_New_Node(start, start, end))//*pseudocode or not I'm not sure I even want this.
	// OpenList.append
	for !isFinished && curr_fails < max_fails {
		if len(OpenList) < 1 && len(ClosedList) > 1 { //!justStart
			// ClosedList = NodeList_SortByFValue_Desc_toReturn(ClosedList, start, target)
			temp := NodeList_GetNeighbors_4_Filtered_Hypentenuse(ClosedList[0], start, target, &imat, floors, walls, margins)
			temp = NodeList_FILTER_LIST(temp, ClosedList)
			temp = NodeList_FILTER_LIST(temp, OpenList)
			temp = NodeList_SortByFValue_Desc_toReturn(temp, start, target)

			OpenList = append(OpenList, temp...)
			OpenList = NodeList_SortByFValue_Desc_toReturn(OpenList, start, target)
			ClosedList = NodeList_SortByFValue_Desc_toReturn(OpenList, start, target)

		}
		OpenList, ClosedList, BlockedList, isFinished, EndNode, curr_fails, err = Pathfind_Phase1_Tick(start, target, OpenList, ClosedList, BlockedList, isFinished, curr_fails, max_fails, imat, floors, walls, margins)
		if err != nil {
			log.Fatal(fmt.Errorf("pathfinding error"))
		}

		if len(ClosedList) == closedList_LastLength {
			curr_fails++

		} else {
			closedList_LastLength = len(ClosedList)
		}
	}
	// dur := time.Since(start_time)
	// if isFinished {
	// 	log.Printf("FINISHED! TIME:%d FAILS: %d\tCLOSEDLIST: %d\tBLOCKED:%3d\n", dur.Milliseconds(), curr_fails, len(ClosedList), len(BlockedList))
	// 	if EndNode != nil {
	// 		EndNode.Set_Heads_Tails_On_Up()
	// 	}
	// 	path_is_found = true
	// } else {
	// 	log.Printf("FAILED! TIME:%d  FAILS: %d\t CLOSEDLIST: %d\t BLOCKED:%3d\n", dur.Milliseconds(), curr_fails, len(ClosedList), len(BlockedList))
	// 	var BLNode, CLNode *ImatNode
	// 	if len(BlockedList) > 0 {
	// 		BlockedList = NodeList_SortByHValue_Ascending_toReturn(BlockedList)
	// 		BLNode = BlockedList[0]
	// 	}
	// 	if len(ClosedList) > 0 {
	// 		ClosedList = NodeList_SortByHValue_Ascending_toReturn(ClosedList)
	// 		CLNode = ClosedList[0]
	// 	}

	// 	if CLNode != nil && BLNode != nil {
	// 		if CLNode.GetHValue() < BLNode.GetHValue() {
	// 			EndNode = CLNode
	// 		} else {
	// 			EndNode = BLNode
	// 		}
	// 	} else if CLNode != nil {
	// 		EndNode = CLNode
	// 	} else if BLNode != nil {
	// 		EndNode = BLNode
	// 	} else {
	// 		log.Printf("SOMETHING FUCKED\n")
	// 	}

	// 	return path_is_found, EndNode

	// }
	path_is_found, EndNode, _ = Pathfind_Phase1_Wrapup(OpenList, ClosedList, BlockedList, start, target, EndNode, curr_fails, start_time, isFinished)
	return path_is_found, EndNode
}

/*
This is to provide a Wrap-up for the code;
*/
func Pathfind_Phase1_Wrapup(openlist, closedlist, blockedlist []*ImatNode, start, target coords.CoordInts, EndNode *ImatNode, curr_fails int, start_time time.Time, Is_Finished bool) (isPathFound bool, output_node *ImatNode, err error) {
	err = nil
	output_node = EndNode
	dur := time.Since(start_time)
	log.Printf("FINISHED:%5t TIME:%010d FAILS: %05d\tOPENLIST: %05d\tCLOSEDLIST: %05d\tBLOCKED:%05d\n", Is_Finished, dur.Milliseconds(), curr_fails, len(openlist), len(closedlist), len(blockedlist))
	//---------------------- Prep Done
	if Is_Finished {
		if output_node != nil {
			output_node.Set_Heads_Tails_On_Up()
		}
	} else {
		var BLNode, CLNode *ImatNode
		if len(blockedlist) > 0 {
			blockedlist = NodeList_SortByHValue_Ascending_toReturn(blockedlist)
			BLNode = blockedlist[0]
		}
		if len(closedlist) > 0 {
			closedlist = NodeList_SortByHValue_Ascending_toReturn(closedlist)
			CLNode = closedlist[0]
		}

		if CLNode != nil && BLNode != nil {
			if CLNode.GetHValue() < BLNode.GetHValue() {
				output_node = CLNode
			} else {
				output_node = BLNode
			}
		} else if CLNode != nil {
			output_node = CLNode
		} else if BLNode != nil {
			output_node = BLNode
		} else {
			log.Printf("SOMETHING FUCKED\n")
		}
	}
	return Is_Finished, output_node, err
}

/**/
func Pathfind_Phase1_Tick(start, target coords.CoordInts, openlist, closedlist, blockedlist []*ImatNode, pathfound bool, curr_fails, max_fails int, imat IntegerMatrix2D, floors, walls []int, margins [4]uint) (oList, cList, bList []*ImatNode, pFound bool, p_Found_Node *ImatNode, fails int, err error) { //<--unsure if these are pass by reference or not
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
			//log.Printf("Point Q is %s\t", pointQ.Position.ToString())
			if pointQ.Position.IsEqual(target) {
				// log.Printf("\n\nHIT TARGET\n\n\n")
				p_Found_Node = pointQ
				pFound = true
				return oList, cList, bList, pFound, p_Found_Node, fails, err

			} else if imat.IsValidCoords(pointQ.Position) && !misc.IsNumInIntArray(imat.GetValueOnCoord(pointQ.Position), walls) {
				//log.Printf("Point Q is Valid %7.2f %7.2f\n", pointQ.GetFValue(), pointQ.GetHValue())
				temp_successors := NodeList_GetNeighbors_4_Filtered_Hypentenuse(pointQ, start, target, &imat, floors, walls, margins)
				//filter temp_successors into openlist
				temp_successors = NodeList_FILTER_LIST(temp_successors, oList)
				temp_successors = NodeList_FILTER_LIST(temp_successors, cList)
				successfulNodes := 0

				for _, suc := range temp_successors {
					if !NodeList_ContainsPoint(suc.Position, cList) && !suc.Position.IsEqual(pointQ.Position) {
						if !NodeList_ContainsPoint(suc.Position, bList) {
							oList = append(oList, suc)
							// log.Printf("ADDED SUCCESSFULLY! %s\n", suc.Position.ToString())
							successfulNodes++
						} else {
							// log.Printf("BLOCKED BLOCKED\n")
							// log.Printf("BLOCKED  %s\n", suc.Position.ToString())

							//remove it

						}
						// cList = append(cList, suc)

						// log.Printf("Success--------\n")
					}

				}
				if successfulNodes == 0 {
					// bList = append(bList, pointQ)
					// log.Printf("Point Q (%s) is invalid %7.2f %7.2f \t length of oList: %4d, clist:%4d, blist:%4d\n", pointQ.Position.ToString(), pointQ.GetFValue(), pointQ.GetHValue(), len(oList), len(cList), len(bList))
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
				// log.Printf("Point Q is NOT Valid \n")

			}

			oList = NodeList_SortByFValue_Desc_toReturn(oList, start, target) //f_value being sum of distance to and from
			oList = NodeList_RemoveDuplicates_ToReturn(oList)
			cList = NodeList_RemoveDuplicates_ToReturn(cList)
			//------NEEDS WORK
			oList, cList, bList, err = Pathfind_Phase1_Tick_Blocked_List_Manager(oList, cList, bList, start, target, &imat, floors, walls, margins)
			if err != nil {
				return oList, cList, bList, false, nil, fails, err
			}
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
					if !misc.IsNumInIntArray(imat.GetValueOnCoord(point.Position), walls) {
						numb++ // && !NodeList_ContainsPoint(point.Position, bList)
					}
				}
				if numb < 2 && !NodeList_ContainsPoint(nod.Position, bList) {
					bList = append(bList, nod)
					//remove from openlist
				}
			}
		}
		oList = NodeList_FILTER_LIST(oList, bList)
		if len(cList) > 10 {
			for _, nod := range cList {
				if !nod.Position.IsEqual(target) && !nod.Position.IsEqual(start) {
					temp := NodeList_GetNeighbors_4A_Filtered_Hypentenuse(nod, start, target, imat, floors, walls, margins)
					numb := 0
					for _, point := range temp {
						if !misc.IsNumInIntArray(imat.GetValueOnCoord(point.Position), walls) && !NodeList_ContainsPoint(point.Position, bList) {
							numb++ //&& !NodeList_ContainsPoint(point.Position, bList)
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
