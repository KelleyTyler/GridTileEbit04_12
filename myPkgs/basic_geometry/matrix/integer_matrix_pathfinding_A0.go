package matrix

import (
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
