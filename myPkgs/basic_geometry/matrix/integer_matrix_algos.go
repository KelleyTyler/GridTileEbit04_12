package matrix

import (
	"log"
	"math/rand"

	coords "github.com/KelleyTyler/GridTileEbit04_12/myPkgs/basic_geometry/coords"
	"github.com/KelleyTyler/GridTileEbit04_12/myPkgs/misc"
)

/**/
func (imat *IntegerMatrix2D) BasicDecay_Step(value_to_set_as, fails int, FrontierList coords.CoordList, filterfor []int, margins [4]uint) (int, coords.CoordList) {
	tempFrontier := make(coords.CoordList, len(FrontierList))
	copy(tempFrontier, FrontierList)
	for _, c := range FrontierList {
		frustration := true
		tempFrontier2 := make(coords.CoordList, 0)
		t := imat.SetValAtCoord(c, value_to_set_as)
		if !t {
			log.Printf("\nFAILURE TO SET VALUE\n")
		}
		tempcoordList, tempValslist := imat.GetNeighborsAndValues_Cardinal(c, margins)
		if tempValslist[0] != -1 && tempValslist[0] != value_to_set_as { //&& imat.IsValidCoordsWithinMargins(margins, tempcoordList[0])
			tempFrontier2 = append(tempFrontier2, tempcoordList[0])
			frustration = false
		}
		if tempValslist[1] != -1 && tempValslist[1] != value_to_set_as { //&& imat.IsValidCoordsWithinMargins(margins, tempcoordList[1])
			tempFrontier2 = append(tempFrontier2, tempcoordList[1])
			frustration = false
		}
		if tempValslist[2] != -1 && tempValslist[2] != value_to_set_as { //&& imat.IsValidCoordsWithinMargins(margins, tempcoordList[2])
			tempFrontier2 = append(tempFrontier2, tempcoordList[2])
			frustration = false
		}
		if tempValslist[3] != -1 && tempValslist[3] != value_to_set_as { //&& imat.IsValidCoordsWithinMargins(margins, tempcoordList[3])
			tempFrontier2 = append(tempFrontier2, tempcoordList[3])
			frustration = false
		}
		if !frustration {
			tempFrontier = append(tempFrontier, tempFrontier2...)
			tempFrontier.RemoveCoord(c)
		} else {
			fails++
			log.Printf("FAILS Temp.len = %3d %3d\n", len(tempFrontier), len(tempFrontier2))
			tempFrontier.RemoveCoord(c)
		}

	}
	tempFrontier.RemoveDuplicates()
	return fails, tempFrontier
}

/**/
func (imat *IntegerMatrix2D) PrimLike_Maze_Algorithm_Random(fails, maxfails int, FrontierList coords.CoordList, floorVals, wallVals, filterFor []int, margins [4]uint, culldiagonals bool) (coords.CoordList, int) {
	randInt := rand.Intn(len(FrontierList))

	tempList, failsout := imat.PrimLike_Maze_Algorithm_Step(randInt, fails, maxfails, FrontierList, floorVals, wallVals, filterFor, margins, culldiagonals)
	// c := FrontierList[randInt]
	return tempList, failsout
}

/*
filterfor should default to []int{-1, 1, 4}
*/
func (imat *IntegerMatrix2D) PrimLike_Maze_Algorithm_Step(FCL_Num, fails, maxfails int, FrontierList coords.CoordList, floorVals, wallVals, filterForPre []int, margins [4]uint, culldiagonals bool) (coords.CoordList, int) {
	temp := make(coords.CoordList, len(FrontierList))
	copy(temp, FrontierList)
	fails_out := fails
	filterFor := filterForPre
	filterFor = append(filterFor, floorVals...)
	filterFor = append(filterFor, wallVals...)
	//------
	if len(temp) > 0 {
		c := temp[FCL_Num]
		frustration := true
		// templist, tempar := imat.GetNeighborsAndValues_8(c, margins)
		templist, tempar := imat.GetNeighborsAndValues_8(c, margins)

		if imat.PrimMazeGenCell_CheckingRules02(c, filterFor, margins) {
			nNBool := !misc.IsNumInIntArray(tempar[0], filterFor)
			nEBool := !misc.IsNumInIntArray(tempar[1], filterFor)
			eEBool := !misc.IsNumInIntArray(tempar[2], filterFor)
			sEBool := !misc.IsNumInIntArray(tempar[3], filterFor)
			sSBool := !misc.IsNumInIntArray(tempar[4], filterFor)

			sWBool := !misc.IsNumInIntArray(tempar[5], filterFor)
			wWBool := !misc.IsNumInIntArray(tempar[6], filterFor)
			nWBool := !misc.IsNumInIntArray(tempar[7], filterFor)

			if culldiagonals {
				if nNBool && nEBool && nWBool { //tempar[0] != -1 && tempar[0] != 1 && tempar[0] != 4
					temp.PushToBack(templist[0])
					frustration = false
				}
				if eEBool && nEBool && sEBool { //tempar[2] != -1 && tempar[2] != 1 && tempar[2] != 4
					temp.PushToBack(templist[2])

					frustration = false
				}
				if sSBool && sEBool && sWBool { //tempar[4] != -1 && tempar[4] != 1 && tempar[4] != 4
					temp.PushToBack(templist[4])
					frustration = false
				}
				if wWBool && nWBool && sWBool { //tempar[6] != -1 && tempar[6] != 1 && tempar[6] != 4
					temp.PushToBack(templist[6]) //nEBool && sEBool
					frustration = false
				}
				temp.RemoveCoord(templist[1])
				temp.RemoveCoord(templist[3])
				temp.RemoveCoord(templist[5])
				temp.RemoveCoord(templist[7])
			} else {
				if nNBool && nEBool && nWBool { //tempar[0] != -1 && tempar[0] != 1 && tempar[0] != 4
					temp.PushToBack(templist[0]) //(sEBool || sWBool)&& (eEBool || wWBool)
					frustration = false
				}
				if eEBool && nEBool && sEBool { //tempar[2] != -1 && tempar[2] != 1 && tempar[2] != 4
					temp.PushToBack(templist[2]) //(sWBool || nWBool) && (nNBool || sSBool)

					frustration = false
				}
				if sSBool && sEBool && sWBool { //tempar[4] != -1 && tempar[4] != 1 && tempar[4] != 4
					temp.PushToBack(templist[4]) //(nEBool || nWBool)&& (eEBool || wWBool)
					frustration = false
				}
				if wWBool && nWBool && sWBool { //tempar[6] != -1 && tempar[6] != 1 && tempar[6] != 4
					temp.PushToBack(templist[6]) //(nEBool || sEBool)&& (nNBool || sSBool)
					frustration = false
				}
			}
		} else {
			imat.SetValAtCoord(c, wallVals[0]) //----
			temp.RemoveCoord(c)
		}
		if !frustration {
			temp.RemoveCoord(c)
			imat.SetValAtCoord(c, floorVals[0])
			fails_out = 0

		} else {
			fails_out++
			if fails_out > maxfails {
				temp.RemoveCoord(c)
				// mazeM.Imat.SetValAtCoord(c, 4)
			}
		}

	}
	temp.RemoveDuplicates()
	return temp, fails_out
}

/*
imported from previous verison;
default filter should be []int{-1, 1, 4}
*/
func (imat *IntegerMatrix2D) PrimMazeGenCell_CheckingRules(cord coords.CoordInts, filter []int, margin [4]uint) bool {
	// _, tempAr, _ := imat.GetNeighbors(cord, margin)
	_, tempvals := imat.GetNeighborsAndValues_8(cord, margin)
	nn := misc.IsNumInIntArray(tempvals[0], filter)
	ne := misc.IsNumInIntArray(tempvals[1], filter)
	ee := misc.IsNumInIntArray(tempvals[2], filter)
	se := misc.IsNumInIntArray(tempvals[3], filter)
	ss := misc.IsNumInIntArray(tempvals[4], filter)
	sw := misc.IsNumInIntArray(tempvals[5], filter)
	ww := misc.IsNumInIntArray(tempvals[6], filter)
	nw := misc.IsNumInIntArray(tempvals[7], filter)
	//---------------------------------------------
	// if (!nn) && ((se && !sw) || (!se && sw)) {
	// 	return false
	// }
	// if (!ee) && ((!sw && nw) || (sw && !nw)) {
	// 	return false
	// }
	// if (!ww) && ((!se && ne) || (!ne && se)) {
	// 	return false
	// }
	// if (!ss) && ((!ne && nw) || (ne && !nw)) { //(nn && ee && ww && !ss) && ((!ne && se && sw && nw) || (ne && se && sw && !nw))
	// 	return false
	// }
	if (nn) && ((se && !sw) || (!se && sw) || (se && sw)) {
		return false
	}
	if (ee) && ((!sw && nw) || (sw && !nw) || (sw && nw)) {
		return false
	}
	if (ww) && ((!se && ne) || (!ne && se) || (ne && se)) {
		return false
	}
	if (ss) && ((!ne && nw) || (ne && !nw) || (ne && nw)) { //(nn && ee && ww && !ss) && ((!ne && se && sw && nw) || (ne && se && sw && !nw))
		return false
	}
	return true
}

/*
imported from previous verison;
default filter should be []int{-1, 1, 4}
*/
func (imat *IntegerMatrix2D) PrimMazeGenCell_CheckingRules02(cord coords.CoordInts, filter []int, margin [4]uint) bool {
	// _, tempAr, _ := imat.GetNeighbors(cord, margin)
	_, tempvals := imat.GetNeighborsAndValues_8(cord, margin)
	nn := misc.IsNumInIntArray(tempvals[0], filter)
	ne := misc.IsNumInIntArray(tempvals[1], filter)
	ee := misc.IsNumInIntArray(tempvals[2], filter)
	se := misc.IsNumInIntArray(tempvals[3], filter)
	ss := misc.IsNumInIntArray(tempvals[4], filter)
	sw := misc.IsNumInIntArray(tempvals[5], filter)
	ww := misc.IsNumInIntArray(tempvals[6], filter)
	nw := misc.IsNumInIntArray(tempvals[7], filter)
	//---------------------------------------------

	n0 := (se && sw)
	n1 := (!se && sw)
	n2 := (se && !sw)
	// n3 := ss && n0
	//---------------
	e0 := (sw && nw)
	e1 := (!sw && nw)
	e2 := (sw && !nw)
	// e3 := ee && e0
	//-------------------
	s0 := (ne && nw)
	s1 := (!ne && nw)
	s2 := (ne && !nw)
	// s3 := ss && s0
	//------------------
	w0 := (se && ne)
	w1 := (!se && ne)
	w2 := (se && !ne)
	// w3 := ww && w0

	if (nn) && (n1 || n2 || n0) {
		return false
	}
	if (ee) && (e1 || e2 || e0) {
		return false
	}
	if (ww) && (w1 || w2 || w0) {
		return false
	}
	if (ss) && (s1 || s2 || s0) { //(nn && ee && ww && !ss) && ((!ne && se && sw && nw) || (ne && se && sw && !nw))
		return false
	}
	return true
}
