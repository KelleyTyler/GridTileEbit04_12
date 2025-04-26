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
	//temp_Nodes := make([]*ImatNode, 0)
	//------
	//phase 1: ---> get an array of []nodes, or even a list of []nodes of potential paths (if they exist)
	//path_is_found02, temp_Nodes := Pathfind_Phase1(start, target, imat, floors, walls, margins)
	//path_is_found = path_is_found02

	//------
	//phase 2: ---> sort through paths and find the shortest one;
	if path_is_found {
		//path_out = Pathfind_Phase2A(tempNodes, start, target, imat, ... )
	} else {
		//reasoning it might be good to gets a set of paths
		switch failconditions {
		default:
			//path_out = Pathfind_Phase2B(tempNodes, start,target,imat,...)
		}
	}

	return path_is_found, path_out
}
func Pathfind_Phase1(start, target coords.CoordInts, imat IntegerMatrix2D, floors, walls []int, margins [4]uint) (path_is_found bool, PotentialPaths []*ImatNode) {

	path_is_found = false
	OpenList := make([]*ImatNode, 0)
	ClosedList := make([]*ImatNode, 0)
	BlockedList := make([]*ImatNode, 0)
	PotentialPaths = make([]*ImatNode, 0)
	var max_fails int = 100
	var curr_fails int = 0
	var isFinished bool = false
	var err error = nil
	var EndNode *ImatNode = nil
	//beware pseudocode
	//ClosedList.ClosedList = ClosedList.append(Get_New_Node(start, start, end))//*pseudocode or not I'm not sure I even want this.
	// OpenList.append
	for !isFinished {
		OpenList, ClosedList, BlockedList, isFinished, EndNode, curr_fails, err = Pathfind_Phase1_Tick(start, target, OpenList, ClosedList, BlockedList, isFinished, curr_fails, max_fails, imat, floors, walls, margins)
		if err != nil {
			log.Fatal(fmt.Errorf("pathfinding error"))
		}
	}

	if isFinished {
		if EndNode != nil {
			EndNode.Set_Heads_Tails_On_Up()
		}

		// PotentialPaths = append(PotentialPaths, ClosedList...)
	}

	return path_is_found, PotentialPaths
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
			if pointQ.Position.IsEqualTo(target) {
				//fmt.Printf("HIT TARGET\n")
				pFoundNode = pointQ
				pFound = true
				return oList, cList, bList, pFound, pFoundNode, fails, err

			} else if imat.IsValidCoords(pointQ.Position) && !misc.IsNumInIntArray(imat.GetValueOnCoord(pointQ.Position), walls) {
				//fmt.Printf("Point Q is Valid %7.2f %7.2f\n", pointQ.GetFValue(), pointQ.GetHValue())
				temp_successors := NodeList_GetNeighbors_4_Filtered_Hypentenuse(pointQ, start, target, &imat, floors, walls, margins)
				//filter temp_successors into openlist
				temp_successors = NodeList_FILTER_LIST(temp_successors, oList)
				temp_successors = NodeList_FILTER_LIST(temp_successors, cList)
				for _, suc := range temp_successors {
					if !NodeList_ContainsPoint(suc.Position, cList) && !NodeList_ContainsPoint(suc.Position, bList) {
						// cList = append(cList, suc)
						oList = append(oList, suc)

						// fmt.Printf("Success--------\n")
					}

				}

				//filter pointQ into closedlist comparing it with any overlap and resolving the contradiction;
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
				fmt.Printf("Point Q is NOT Valid \n")

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
func Pathfind_Phase1_Tick_Blocked_List_Manager(openlist, closedlist, blockedlist []*ImatNode, start, target coords.CoordInts, imat *IntegerMatrix2D, floors, walls []int, margins [4]uint) (oList, cList, bList []*ImatNode, err error) {
	oList = make([]*ImatNode, len(openlist))
	cList = make([]*ImatNode, len(closedlist))
	bList = make([]*ImatNode, len(blockedlist))
	copy(oList, openlist)
	copy(cList, closedlist)
	copy(bList, blockedlist)

	//---------------------- Prep Done
	for _, nod := range oList {
		if !nod.Position.IsEqual(target) && !nod.Position.IsEqual(start) {
			temp := NodeList_GetNeighbors_4_Filtered_Hypentenuse(nod, start, target, imat, floors, walls, margins)
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
	for _, nod := range cList {
		if !nod.Position.IsEqual(target) && !nod.Position.IsEqual(start) {
			temp := NodeList_GetNeighbors_4_Filtered_Hypentenuse(nod, start, target, imat, floors, walls, margins)
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
	cList = NodeList_FILTER_LIST(cList, bList)
	return oList, cList, bList, nil
}
