package matrix

import coords "github.com/KelleyTyler/GridTileEbit04_12/myPkgs/basic_geometry/coords"

type ImatNode struct {
	Position, Start, Target coords.CoordInts
	Prev                    *ImatNode //this is singular
	Next                    *ImatNode //make this an array??? (see argument Against such an action)
	Iteration               int
	ValueOnGrid             int
}

/*ARGUEMENTS AGAINST ARRAY FOR "ImatNode.NEXT:"

->it's unncessary
->it promotes an architecuture/design pattern that would be very bad and lead to massive headaches;


*/

// func (node *ImatNode) getGenerationHelper(val int) int {
// 	if val == -1 {
// 		if node.Prev != nil {
// 			node.Prev.getGenerationHelper(-1)
// 		}else{

// 		}
// 	}
// }

/*
 */
func GetNode(start, point, target coords.CoordInts, imat IntegerMatrix2D, parent *ImatNode) ImatNode {
	temp := ImatNode{
		Start:    start,
		Position: point,
		Target:   target,
	}
	if imat.IsValidCoords(point) {
		temp.ValueOnGrid = imat.GetValueOnCoord(point)
	}
	if parent != nil {
		temp.Prev = parent
		temp.Iteration = temp.GetInteration()
	}
	return temp
}

/*
 */
func (node *ImatNode) GetInteration() int {
	if node.Prev != nil {
		node.Iteration = node.Prev.GetInteration() + 1
		return node.Iteration
	} else {
		node.Iteration = 0
		return node.Iteration
	}
}

/*
 */
func (node *ImatNode) GetDistances(starting, ending coords.CoordInts, straightLine bool) (toStart, toEnd int) {
	if straightLine {
		//get the hypotenuse
		toStart = node.Position.GetHypotenuseDistance_Int(starting)
		toEnd = node.Position.GetHypotenuseDistance_Int(ending)
	} else {
		toStart = node.Position.GetManhattanDistance_Int(starting)
		toEnd = node.Position.GetManhattanDistance_Int(ending)

	}
	return toStart, toEnd
}

/*
 */
func (node *ImatNode) SetChildrenOfParents_Recursive() {
	if node.Prev != nil {
		node.Prev.SetChildrenOfParents_Recursive()
		node.Prev.Next = node
	}

}

/*
 */
func (node *ImatNode) SetChildrenOfParents() {
	if node.Prev != nil {
		node.Prev.Next = node
	}
}

/*
 */
func (node *ImatNode) GetFValue() (fVal float64) {
	fVal = node.GetGValue() + node.GetHValue()
	return fVal
}

/*
	this is the cost-so-far so the movement distance to start;
*/
func (node *ImatNode) GetGValue() (gVal float64) {
	gVal = node.GetMovementDistanceToStart()
	return gVal
}

/*
 */
func (node *ImatNode) GetHValue() (hVal float64) {
	hVal = node.Target.GetHypotenuseDistance_Float(node.Target)
	return hVal
}

/*
 */
func (node *ImatNode) GetMovementDistanceToStart() (m_dist_to_start float64) {
	m_dist_to_start = float64(node.Iteration)
	return
}

/*
 */
func (node *ImatNode) Swap(node2 *ImatNode) {
	tempChild := node.Next
	tempParent := node.Prev
	node.Next = node2.Next
	node.Prev = node2.Prev
	node2.Next = tempChild
	node2.Prev = tempParent
}

/*
this only works if all the Prev pointers are lined up properly; otherwise you're kind of screwed;
*/
func (node *ImatNode) GetHead() *ImatNode {
	if node.Prev != nil {
		return node.Prev.GetHead()
	} else {
		return node
	}
}

/*
this only works if all the Next pointers are lined up properly; otherwise you're kind of screwed;
*/
func (node *ImatNode) GetTail() *ImatNode {
	if node.Next != nil {
		return node.Next.GetTail()
	} else {
		return node
	}
}

// func (node *ImatNode) PopHead() (oldhead *ImatNode, newhead *ImatNode) {

// 	return
// }

// /*
// 	This only works if the Head is selected; nothing else
// */
// func (node *ImatNode) pop_from_head_helper() (oldhead *ImatNode, newhead *ImatNode, err error) {
// 	oldhead = node
// 	newhead = node.Next
// 	err = nil
// 	if node.Next.Prev != nil {

// 	}

// 	return
// }

/*
This sets;


*/
func (node *ImatNode) Set_Heads_Tails_On_Up() {
	if node.Prev != nil {
		node.Prev.Next = node
		node.Prev.Set_Heads_Tails_On_Up()
	}
}

/*

 */
func (node *ImatNode) FlipOrder() {
	// node.fliporderHelper(false)
	tempHead := node.GetHead()
	tempHead.fliporderHelper(true)
}

/*

 */
func (node *ImatNode) fliporderHelper(FromHead bool) {
	if FromHead {
		if node.Next != nil {
			temp := node.Prev
			node.Prev = node.Next
			node.Next = temp
			node.Prev.fliporderHelper(true)
		}
	} else if node.Prev != nil {
		node.Prev.fliporderHelper(false)

	} else {
		if node.Next != nil {
			node.Prev = node.Next
			node.Prev.fliporderHelper(true)
		}
	}
}

/*

 */
func (node *ImatNode) PrintFromHead() {

}
func (node ImatNode) GetNode() ImatNode {
	return node
}

// type NodeComparisonType int

// const (
// 	TotallyEqual NodeComparisonType = iota
// 	SamePostion
// 	SameParent
// 	Same
// )

// func flipImatOrder(head *ImatNode) {

// }
