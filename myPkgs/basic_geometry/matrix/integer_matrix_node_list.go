package matrix

import (
	"slices"

	"github.com/KelleyTyler/GridTileEbit04_12/myPkgs/basic_geometry/coords"
)

//type IntMatNodeList []ImatNode
//alternatively lets try some slices;

func NodeList_ContainsPoint(point coords.CoordInts, S []*ImatNode) bool {
	return slices.ContainsFunc(S, func(E *ImatNode) bool {
		return E.Position.IsEqual(point)
	})
}

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
