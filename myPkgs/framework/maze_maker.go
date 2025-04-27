package framework

/*
	This file is intended to provide an interface/wrapper for the various pontial maze generation algorithms that might need to interact with the grid;
*/
import (
	"fmt"
	"image/color"

	"github.com/KelleyTyler/GridTileEbit04_12/myPkgs/basic_geometry/coords"
	mat "github.com/KelleyTyler/GridTileEbit04_12/myPkgs/basic_geometry/matrix"
	ui "github.com/KelleyTyler/GridTileEbit04_12/myPkgs/user_interface"
	"github.com/hajimehoshi/ebiten/v2"
)

type Maze_Tool struct {
	Is_Started, Is_Finished bool
	MazeGen                 mat.MazeMaker
	Imat                    *mat.IntegerMatrix2D
	UI_Backend              *ui.UI_Backend
	Display_Settings_Ptr    *mat.Integer_Matrix_Ebiten_DrawOptions
	DisplaySettings         mat.Integer_Matrix_Ebiten_DrawOptions

	// USER INTERFACE
	// Maze_Selection_Button, Maze_Gen_Button, MazeGenBtn00, MazeGenBtn01 ui.UI_Button
	BarLabel, OUTPUT_LABEL                                                      ui.UI_Label
	Button_SET_POINTS, Button_GENERATE_MAZE, Button_TOGGLE_DIAGONALS, Button_00 ui.UI_Button
	NSelect_MazeGen_00                                                          ui.UI_Num_Select
}

func (mTool *Maze_Tool) Init(intMat *mat.IntegerMatrix2D, backend *ui.UI_Backend, maxPoints int, dsetting *mat.Integer_Matrix_Ebiten_DrawOptions) {
	mTool.Imat = intMat
	mTool.UI_Backend = backend

	mTool.Is_Started = false
	mTool.Is_Finished = false
	mTool.DisplaySettings = *dsetting
	mTool.Display_Settings_Ptr = dsetting

	// tempBoardOps := mat.Integer_Matrix_Ebiten_DrawOptions{
	// 	BoardPosition:     dsetting.BoardPosition,
	// 	BoardMargin:       dsetting.BoardMargin,
	// 	TileSize:          d,
	// 	TileSpacing:       tilespacing,
	// 	ShowTileLines:     []bool{true, true, true},
	// 	TileLineColors:    []color.Color{color.RGBA{150, 150, 150, 255}, color.RGBA{180, 40, 40, 255}, color.RGBA{180, 40, 40, 255}},
	// 	TileLineThickness: []float32{1.0, 1.0, 1.0},
	// }
	mTool.MazeGen.Init(mTool.DisplaySettings, intMat)
}

func (mTool *Maze_Tool) OnValidMouseClickOnGameBoard(pos_X, pos_Y int) (overlay_change, board_change bool) {
	overlay_change = false
	board_change = false
	if mTool.Button_SET_POINTS.GetState() == 2 {
		mTool.MazeGen.CurrentList = append(mTool.MazeGen.CurrentList, coords.CoordInts{X: pos_X, Y: pos_Y})
		mTool.UI_Backend.PlaySound(4)
		overlay_change = true
	}
	return overlay_change, board_change
}

func (mTool *Maze_Tool) Reset(opts mat.Integer_Matrix_Ebiten_DrawOptions) {
	mTool.DisplaySettings.Update_Opts_From_Argument(opts)

	mTool.MazeGen.CurrentList = make(coords.CoordList, 0)
}

func (mTool *Maze_Tool) Redef(opts mat.Integer_Matrix_Ebiten_DrawOptions) {
	mTool.DisplaySettings.Update_Opts_From_Argument(opts)
}

func (mTool *Maze_Tool) Update_Passive() (overlay_change, board_change bool) {
	// if gb.ticker > gb.ticker_max {
	// 	gb.BoardChanges = true
	// 	gb.BoardOverlayChanges = true
	// 	gb.ticker = 0

	// 	gb.MazeTextBox.Text = gb.MazeGen.CurrentList.ToString()
	// 	gb.MazeTextBox.Redraw()
	// 	if gb.MazeTextBox.Parent != nil {
	// 		gb.MazeTextBox.Parent.Redraw()
	// 	}
	// } else {
	// 	gb.ticker++
	// }
	if !mTool.MazeGen.HasFinished { // && !mTool.MazeGen.HasFinished
		mTool.OUTPUT_LABEL.Text = fmt.Sprintf("%s ", mTool.MazeGen.GetString())
		mTool.OUTPUT_LABEL.Redraw()
		// if mTool.OUTPUT_LABEL.Parent != nil {
		// 	mTool.OUTPUT_LABEL.Parent.Redraw()
		// }
	}
	overlay_change = false
	board_change = false
	if mTool.Button_GENERATE_MAZE.GetState() == 2 {
		// mTool.MazeGen.RunPrimlike(32, []int{0}, []int{9}, []int{-1}, [4]uint{1, 1, 1, 1}, mTool.Button_TOGGLE_DIAGONALS.GetState() != 2)
		overlay_change, board_change = mTool.Maze_Gen_Passthrough()
	}
	return overlay_change, board_change
}

func (mTool *Maze_Tool) Maze_Gen_Passthrough() (overlay_change, board_change bool) {
	overlay_change = false
	board_change = false
	if len(mTool.MazeGen.CurrentList) > 0 {
		mTool.MazeGen.HasStarted = true
		mTool.MazeGen.HasFinished = false
		mTool.MazeGen.RunPrimlike(32, []int{0}, []int{9}, []int{-1}, [4]uint{1, 1, 1, 1}, mTool.Button_TOGGLE_DIAGONALS.GetState() != 2)

		if mTool.MazeGen.HasFinished {
			fmt.Printf("done\n")

			mTool.Button_GENERATE_MAZE.DeToggle()
			mTool.Button_SET_POINTS.DeToggle()
			mTool.Button_TOGGLE_DIAGONALS.DeToggle()
			mTool.MazeGen.HasStarted = false
			mTool.OUTPUT_LABEL.Text = fmt.Sprintf("%s ", mTool.MazeGen.GetString())
			mTool.OUTPUT_LABEL.Redraw()
		}
		overlay_change = true
		board_change = true
	}
	return overlay_change, board_change
}

/*
Rather than rely on the native thing in the MazeGen I think I will draw these myself;
*/
func (mTool *Maze_Tool) Draw(screen *ebiten.Image, opts mat.Integer_Matrix_Ebiten_DrawOptions) {
	mTool.DisplaySettings.Update_Opts_From_Argument(opts)
	if len(mTool.MazeGen.CurrentList) > 0 {
		mTool.Imat.DrawCoordListWithLines(screen, mTool.MazeGen.CurrentList, []color.Color{color.RGBA{110, 200, 200, 255}}, mTool.DisplaySettings)
	}

}

/*
This initiates the User Interface
*/
func (mTool *Maze_Tool) UI_Init(parent ui.UI_Object, row_value int) {
	mTool.BarLabel.Init([]string{"maze_gen_label", "MAZE GENERATOR"}, mTool.UI_Backend, nil, coords.CoordInts{X: 0, Y: row_value}, coords.CoordInts{X: 204, Y: 32})
	mTool.BarLabel.TextAlignMode = 10
	mTool.BarLabel.Redraw()
	mTool.BarLabel.Init_Parents(parent)

	row_01_value := row_value + 34
	mTool.Button_SET_POINTS.Init([]string{"maze_gen_btn", "Select\nPoints"}, mTool.UI_Backend, nil, coords.CoordInts{X: 4, Y: row_01_value}, coords.CoordInts{X: 64, Y: 32})
	mTool.Button_SET_POINTS.Btn_Type = 10
	mTool.Button_SET_POINTS.Init_Parents(parent)

	mTool.OUTPUT_LABEL.Init([]string{"maze_gen_label", "------"}, mTool.UI_Backend, nil, coords.CoordInts{X: 70, Y: row_01_value}, coords.CoordInts{X: 130, Y: 66})
	mTool.OUTPUT_LABEL.Init_Parents(parent)

	mTool.Button_GENERATE_MAZE.Init([]string{"maze_gen_btn", "Gen\nPrimlike"}, mTool.UI_Backend, nil, coords.CoordInts{X: 4, Y: row_01_value + 34}, coords.CoordInts{X: 64, Y: 32})
	mTool.Button_GENERATE_MAZE.Btn_Type = 10
	mTool.Button_GENERATE_MAZE.Init_Parents(parent)

	row_03_Value := row_01_value + 68

	mTool.Button_TOGGLE_DIAGONALS.Init([]string{"maze_gen_btn", "Toggle\nDiagonals"}, mTool.UI_Backend, nil, coords.CoordInts{X: 4, Y: row_03_Value}, coords.CoordInts{X: 64, Y: 32})
	mTool.Button_TOGGLE_DIAGONALS.Btn_Type = 10
	mTool.Button_TOGGLE_DIAGONALS.Init_Parents(parent)

	mTool.NSelect_MazeGen_00.Init([]string{"n_map_btn", "mazeGen00"}, mTool.UI_Backend, nil, coords.CoordInts{X: 70, Y: row_03_Value}, coords.CoordInts{X: 64, Y: 32})
	mTool.NSelect_MazeGen_00.Init_Parents(parent)
	mTool.NSelect_MazeGen_00.SetVals(0, 1, 0, 16, 0)

	mTool.Button_00.Init([]string{"maze_gen_btn", "MazeGen00"}, mTool.UI_Backend, nil, coords.CoordInts{X: 136, Y: row_03_Value}, coords.CoordInts{X: 64, Y: 32})
	mTool.Button_00.Btn_Type = 10
	mTool.Button_00.Init_Parents(parent)
}
