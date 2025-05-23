package matrix

import (
	"fmt"

	coords "github.com/KelleyTyler/GridTileEbit04_12/myPkgs/basic_geometry/coords"
)

/**/
type MazeMaker struct {
	CurrentList             coords.CoordList
	Fails, maxFails         int
	HasStarted, HasFinished bool

	DisplaySettings Integer_Matrix_Ebiten_DrawOptions
	imat            *IntegerMatrix2D
	//outstrng         string
	Show_CurrentList bool
}

// func (mazeM *MazeMaker) ToggleShowCurrentList() {
// 	mazeM.ShowCurrentList != mazeM.ShowCurrentList
// }
/**/
func (mazeM *MazeMaker) Init(dSettings Integer_Matrix_Ebiten_DrawOptions, intmatrix *IntegerMatrix2D) {
	mazeM.CurrentList = make(coords.CoordList, 0)
	mazeM.Fails = 0
	mazeM.maxFails = 18
	mazeM.HasStarted = false
	mazeM.HasFinished = false
	mazeM.imat = intmatrix
	mazeM.DisplaySettings = dSettings
	mazeM.DisplaySettings.AABody = true
	mazeM.DisplaySettings.AALines = true
	// mazeM.DisplaySettings.TileLineColors = []color.Color{color.RGBA{150, 150, 150, 255}, color.RGBA{180, 40, 40, 255}, color.RGBA{180, 40, 40, 255}}
	// mazeM.DisplaySettings.TileLineThickness = []float32{1.0, 2.0, 2.0}
	// mazeM.DisplaySettings.ShowTileLines = []bool{true, true, true}
}

/**/
func (mazeM *MazeMaker) RunPrimlike(ticks int, floorvals, wallvals, filterFor []int, margin [4]uint, diage bool) {
	for range ticks {
		if len(mazeM.CurrentList) > 0 {
			mazeM.HasStarted = true
			mazeM.CurrentList, mazeM.Fails = mazeM.imat.PrimLike_Maze_Algorithm_Random(mazeM.Fails, mazeM.maxFails, mazeM.CurrentList, floorvals, wallvals, filterFor, margin, diage)
		} else {
			//log.Printf("FINISHED!\n")
			mazeM.HasFinished = true
			mazeM.imat.Convert_All_OldValues_To_NewValue(wallvals[0], 10)
			break
		}
	}
}

/**/
func (mazeM *MazeMaker) GetString() string {
	outstrng := fmt.Sprintf("MAZEGEN:\n HAS IMAT:  %5t\n", mazeM.imat != nil)
	outstrng += fmt.Sprintf("CurrentList: %3d\n", len(mazeM.CurrentList))
	outstrng += fmt.Sprintf("%13s: %5t\n %13s: %5t\n", "HAS_STARTED", mazeM.HasStarted, "HAS_FINISHED", mazeM.HasFinished)
	return outstrng
}

/**/
type Pathfinding_Solver struct {
	Start, Target           coords.CoordInts
	Path                    coords.CoordList
	HasStarted, HasFinished bool
	DisplayOptions          Integer_Matrix_Ebiten_DrawOptions
	Imat                    *IntegerMatrix2D
	ShowPath                bool
}

/**/
func (pfind *Pathfinding_Solver) Init(intmatrix *IntegerMatrix2D, dOps Integer_Matrix_Ebiten_DrawOptions) {
	pfind.DisplayOptions = dOps
	pfind.Imat = intmatrix
}
