package matrix

import (
	"fmt"
	"log"
	"time"

	coords "github.com/KelleyTyler/GridTileEbit04_12/myPkgs/basic_geometry/coords"
	misc "github.com/KelleyTyler/GridTileEbit04_12/myPkgs/misc"
)

/*
	attempting to make the pathfinding algorithms more efficient;
*/

/*
While this is supposed to manage the 'blocked list' it presently doesn't work very well at all.
*/
func Pathfind_Blocked_List_Manager(openlist, closedlist, blockedlist []*ImatNode, start, target coords.CoordInts, imat *IntegerMatrix2D, floors, walls []int, margins [4]uint) (oList, cList, bList []*ImatNode, err error) {
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

/**/
func Pathfind_Phase_1A_0(start, target coords.CoordInts, imat IntegerMatrix2D, floors, walls []int, margins [4]uint) (path_is_found bool, EndNode *ImatNode) {
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
