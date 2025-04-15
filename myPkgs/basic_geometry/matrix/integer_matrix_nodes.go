package matrix

import coords "github.com/KelleyTyler/GridTileEbit04_12/myPkgs/basic_geometry/coords"

type ImatNode struct {
	Position   coords.CoordInts
	Prev       *ImatNode //this is singular
	Next       *ImatNode //make this an array??? (see argument Against such an action)
	Generation int
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

func (node *ImatNode) GetGeneration() int {
	if node.Prev != nil {
		node.Generation = node.Prev.GetGeneration() + 1
		return node.Generation
	} else {
		node.Generation = 0
		return node.Generation
	}
}

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

func (node *ImatNode) SetChildrenOfParents() {
	if node.Prev != nil {
		node.Prev.SetChildrenOfParents()
		node.Prev.Next = node
	}
}

// func (node *ImatNode) FlipOrder() {
// 	if node.Prev != nil {
// 		if node.Next != nil {
// 			temp := node.Next
// 			node.Next = node.Prev
// 			node.Next.FlipOrder()
// 			node.Prev = temp
// 		} else {
// 			node.Next = node.Prev
// 			node.Next.FlipOrder()
// 		}
// 	}else{
// 		f
// 	}
// }

// // func (node *ImatNode) fliporderHelper() {

// // }
// func flipImatOrder(head *ImatNode) {

// }
