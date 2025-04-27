package framework

import (
	"fmt"
	"image/color"

	"github.com/KelleyTyler/GridTileEbit04_12/myPkgs/basic_geometry/coords"
	mat "github.com/KelleyTyler/GridTileEbit04_12/myPkgs/basic_geometry/matrix"
	ui "github.com/KelleyTyler/GridTileEbit04_12/myPkgs/user_interface"
	"github.com/hajimehoshi/ebiten/v2"
)

type Pathfind_Tester struct {
	Start, Target                                                coords.CoordInts
	ClosedList, OpenList, BlockedList                            []*mat.ImatNode
	IsReady, HasStarted, IsFinished, IsStartInput, IsTargetInput bool

	EndNode *mat.ImatNode

	ShowOnScreen        bool
	ShowOnScreenOptions uint8
	GridOptions         mat.Integer_Matrix_Ebiten_DrawOptions
	IMat                *mat.IntegerMatrix2D

	UI_Backend *ui.UI_Backend

	max_fails, curr_fails int
	//------- User Interface and User Interface Buttons
	Button_Panel_Label                                                                      ui.UI_Label
	Button_Select_Points, Button_Pathfind_Tick, Button_Reset                                ui.UI_Button
	Button_Pathfind_Auto                                                                    ui.UI_Button
	Button_SHOW_OPENLIST, Button_SHOW_CLOSEDLIST, Button_SHOW_BLOCKEDLIST, Button_SHOW_PATH ui.UI_Button
	tickCount                                                                               int
}

func (pfind *Pathfind_Tester) Init(imat *mat.IntegerMatrix2D, backend *ui.UI_Backend, options *mat.Integer_Matrix_Ebiten_DrawOptions) {
	pfind.IMat = imat
	pfind.Reset()
	pfind.GridOptions = options.GetOptsByValue()
	pfind.GridOptions.TileLineColors = []color.Color{color.Black, color.Black, color.Black}
	pfind.UI_Backend = backend
}
func (pfind *Pathfind_Tester) Reset() {
	pfind.IsTargetInput = false
	pfind.IsStartInput = false
	pfind.IsFinished = false
	pfind.IsReady = false
	pfind.HasStarted = false
	pfind.ClosedList = make([]*mat.ImatNode, 0)
	pfind.OpenList = make([]*mat.ImatNode, 0)
	pfind.BlockedList = make([]*mat.ImatNode, 0)
	pfind.EndNode = nil
	pfind.max_fails = 100
	pfind.curr_fails = 0
	fmt.Printf("RESET PFIND TEST \n")
	pfind.tickCount = 0
}

func (pfind *Pathfind_Tester) UI_Init(parent ui.UI_Object, pfindBtnRow int) {
	pfind.Button_Panel_Label.Init([]string{"pfind_b_panel_label", "Pathfind"}, pfind.UI_Backend, nil, coords.CoordInts{X: 0, Y: pfindBtnRow}, coords.CoordInts{X: 204, Y: 32})
	pfind.Button_Panel_Label.TextAlignMode = 10
	pfind.Button_Panel_Label.Redraw()
	pfind.Button_Panel_Label.Init_Parents(parent)
	pfind.Button_Select_Points.Init([]string{"pfind_set_points", "SET\nPOINTS"}, pfind.UI_Backend, nil, coords.CoordInts{X: 4, Y: pfindBtnRow + 34}, coords.CoordInts{X: 64, Y: 32})
	pfind.Button_Select_Points.Btn_Type = 10
	pfind.Button_Select_Points.Init_Parents(parent)
	pfind.Button_Pathfind_Tick.Init([]string{"pfind_Tick", "TICK"}, pfind.UI_Backend, nil, coords.CoordInts{X: 70, Y: pfindBtnRow + 34}, coords.CoordInts{X: 64, Y: 32})
	pfind.Button_Pathfind_Tick.Init_Parents(parent)
	pfind.Button_Reset.Init([]string{"pfind_reset_points", "RESET"}, pfind.UI_Backend, nil, coords.CoordInts{X: 136, Y: pfindBtnRow + 34}, coords.CoordInts{X: 64, Y: 32})
	pfind.Button_Reset.Init_Parents(parent)

	pfind.Button_Pathfind_Auto.Init([]string{"pfind_Auto", "AUTO\nPathfind"}, pfind.UI_Backend, nil, coords.CoordInts{X: 4, Y: pfindBtnRow + 68}, coords.CoordInts{X: 64, Y: 32})
	pfind.Button_Pathfind_Auto.Btn_Type = 10
	pfind.Button_Pathfind_Auto.Init_Parents(parent)

	pfind.Button_SHOW_OPENLIST.Init([]string{"pfind_Auto", "SHOW\nO_LIST"}, pfind.UI_Backend, nil, coords.CoordInts{X: 4, Y: pfindBtnRow + 102}, coords.CoordInts{X: 48, Y: 32})
	pfind.Button_SHOW_OPENLIST.Btn_Type = 10
	pfind.Button_SHOW_OPENLIST.Init_Parents(parent)

	pfind.Button_SHOW_CLOSEDLIST.Init([]string{"pfind_Auto", "SHOW\nC_LIST"}, pfind.UI_Backend, nil, coords.CoordInts{X: 53, Y: pfindBtnRow + 102}, coords.CoordInts{X: 48, Y: 32})
	pfind.Button_SHOW_CLOSEDLIST.Btn_Type = 10
	pfind.Button_SHOW_CLOSEDLIST.Init_Parents(parent)

	pfind.Button_SHOW_BLOCKEDLIST.Init([]string{"pfind_Auto", "SHOW\nB_LIST"}, pfind.UI_Backend, nil, coords.CoordInts{X: 103, Y: pfindBtnRow + 102}, coords.CoordInts{X: 48, Y: 32})
	pfind.Button_SHOW_BLOCKEDLIST.Btn_Type = 10
	pfind.Button_SHOW_BLOCKEDLIST.Init_Parents(parent)

	pfind.Button_SHOW_PATH.Init([]string{"pfind_Auto", "SHOW\nPATH"}, pfind.UI_Backend, nil, coords.CoordInts{X: 152, Y: pfindBtnRow + 102}, coords.CoordInts{X: 48, Y: 32})
	pfind.Button_SHOW_PATH.Btn_Type = 10
	pfind.Button_SHOW_PATH.Init_Parents(parent)
}

func (pfind *Pathfind_Tester) OnValidMouseClickOnGameBoard(pos_X, pos_Y int) (overlay_change, board_change bool) {
	overlay_change = false
	board_change = false

	if pfind.Button_Select_Points.GetState() == 2 {
		pfind.InputCoords(coords.CoordInts{X: pos_X, Y: pos_Y})
		overlay_change = true
	}

	return overlay_change, board_change
}

func (pfind *Pathfind_Tester) Update_Passive() (overlay_change bool) {
	overlay_change = false
	if pfind.Button_Pathfind_Tick.GetState() == 2 {
		// fmt.Printf("BUTTON TICK\n")
		// pfind.Process(1)
		go pfind.Process02()
		overlay_change = true
	}
	if pfind.Button_Reset.GetState() == 2 {
		fmt.Printf("RESET BUTTON\n")

		pfind.Reset()
		overlay_change = true
	}
	if pfind.Button_Pathfind_Auto.GetState() == 2 {
		// fmt.Printf("AUTO TICK\n")

		pfind.Process(110)
		overlay_change = true
	}
	return overlay_change
}

func (pfind *Pathfind_Tester) InputCoords(coord coords.CoordInts) {
	if !pfind.IsStartInput {
		pfind.Start = coord

		fmt.Printf("START INPUT\n")
		pfind.IsStartInput = true
		pfind.UI_Backend.PlaySound(4)

	} else if !pfind.IsTargetInput {
		if !coord.IsEqual(pfind.Start) {
			pfind.Target = coord
			pfind.IsTargetInput = true
			pfind.IsReady = true
			pfind.IsFinished = false
			// pfind.EndNode = nil
			fmt.Printf("END INPUT\n")
			pfind.UI_Backend.PlaySound(4)

		}
	}
}

// -----replicating the process we've seen before;
func (pfind *Pathfind_Tester) Process(ticks int) error {
	var err error

	// fmt.Printf("ENDNODE %t %t %t %t\n\n", pfind.EndNode != nil, pfind.HasStarted, pfind.IsFinished, pfind.IsReady)

	if pfind.IsReady {
		for range ticks {
			if !pfind.HasStarted {
				fmt.Printf("STARTING!\n")
				startnode := mat.GetNode(pfind.Start, pfind.Start, pfind.Target, *pfind.IMat, nil)
				pfind.ClosedList = append(pfind.ClosedList, &startnode)
				temp := mat.NodeList_GetNeighbors_4_Filtered_Hypentenuse(&startnode, pfind.Start, pfind.Target, pfind.IMat, []int{0, 1}, []int{9, 10}, [4]uint{1, 1, 1, 1})
				mat.NodeList_SortByFValue_DESC(temp, pfind.Start, pfind.Target)
				pfind.OpenList = append(pfind.OpenList, temp...)
				fmt.Printf("STARTUP! %d %d \n", len(pfind.ClosedList), len(pfind.OpenList))

				// pfind.OpenList, pfind.ClosedList, pfind.BlockedList, pfind.IsFinished, pfind.curr_fails, err = mat.Pathfind_Phase1_Tick(pfind.Start, pfind.Target, pfind.OpenList, pfind.ClosedList, pfind.BlockedList, pfind.IsFinished, pfind.curr_fails, pfind.max_fails, *pfind.IMat, []int{0, 1}, []int{9, 10}, [4]uint{1, 1, 1, 1})
				pfind.HasStarted = true
			} else {
				if !pfind.IsFinished {
					if len(pfind.OpenList) < 1 {
						// mat.NodeList_SortByHValue_Ascending(pfind.ClosedList)
						// mat.NodeList_SortByFValue_Ascending(pfind.ClosedList, pfind.Start, pfind.Target)

						pfind.ClosedList = mat.NodeList_SortByFValue_Desc_toReturn(pfind.ClosedList, pfind.Start, pfind.Target)

						// mat.NodeList_SortByFValue_DESC(pfind.ClosedList, pfind.Start, pfind.Target)

						// slices.Reverse(pfind.ClosedList)
						temp := mat.NodeList_GetNeighbors_4_Filtered_Hypentenuse(pfind.ClosedList[0], pfind.Start, pfind.Target, pfind.IMat, []int{0, 1}, []int{9, 10}, [4]uint{1, 1, 1, 1})
						temp = mat.NodeList_FILTER_LIST(temp, pfind.OpenList)
						temp = mat.NodeList_FILTER_LIST(temp, pfind.ClosedList)
						pfind.OpenList = append(pfind.OpenList, temp...)
						//fmt.Printf("TICK!----\n")

						// mat.NodeList_SortByFValue_Ascending(pfind.OpenList, pfind.Start, pfind.Target)
						// slices.Reverse(pfind.OpenList)
					}

					pfind.OpenList, pfind.ClosedList, pfind.BlockedList, pfind.IsFinished, pfind.EndNode, pfind.curr_fails, err = mat.Pathfind_Phase1_Tick(pfind.Start, pfind.Target, pfind.OpenList, pfind.ClosedList, pfind.BlockedList, pfind.IsFinished, pfind.curr_fails, pfind.max_fails, *pfind.IMat, []int{0, 1}, []int{9, 10}, [4]uint{1, 1, 1, 1})
					//fmt.Printf("RUNNING! %d %d \n", len(pfind.ClosedList), len(pfind.OpenList))
					pfind.tickCount++
				} else {
					fmt.Printf("FOUND! %t---- TICKS: %d", pfind.EndNode != nil, pfind.tickCount)
					// fmt.Printf("FOUND! %t---- TICKS: %d\n\n", pfind.EndNode != nil, pfind.tickCount)

					if pfind.EndNode != nil {
						fmt.Printf(" length: %d", pfind.EndNode.GetInteration())
						pfind.EndNode.Set_Heads_Tails_On_Up()
					}
					fmt.Printf("\n")
					// pfind.EndNode.Set_Heads_Tails_On_Up() //<---- for some reason the whole thing fucking freezes if this is put into any kind of storage;
					if pfind.Button_Pathfind_Auto.GetState() == 2 {
						pfind.Button_Pathfind_Auto.DeToggle()
					}
					break
				}
			}
		}
	}
	return err
}

func (pfind *Pathfind_Tester) Process02() error {
	var err error
	// isDone := false
	// fmt.Printf("ENDNODE %t %t %t %t\n\n", pfind.EndNode != nil, pfind.HasStarted, pfind.IsFinished, pfind.IsReady)

	if pfind.IsReady {
		pfind.IsFinished, pfind.EndNode = mat.Pathfind_Phase1A(pfind.Start, pfind.Target, *pfind.IMat, []int{0, 1}, []int{9, 10}, [4]uint{1, 1, 1, 1})
		if pfind.IsFinished {
			fmt.Printf("DONE %d\n\n", pfind.EndNode.GetInteration())
		}
	}
	return err
}

func (pfind *Pathfind_Tester) Draw(screen *ebiten.Image, drawOpts *mat.Integer_Matrix_Ebiten_DrawOptions) {
	pfind.GridOptions = *drawOpts
	pfind.GridOptions.TileLineColors = []color.Color{color.Black, color.Black, color.Black}

	if pfind.IsStartInput {
		pfind.IMat.DrawAGridTile_With_Lines(screen, pfind.Start, color.RGBA{128, 255, 128, 255}, &pfind.GridOptions)

	}
	if pfind.IsTargetInput {
		pfind.IMat.DrawAGridTile_With_Lines(screen, pfind.Target, color.RGBA{255, 128, 0, 255}, &pfind.GridOptions)
	}

	if len(pfind.ClosedList) > 1 && pfind.Button_SHOW_CLOSEDLIST.GetState() != 2 {
		pfind.IMat.NodeList_DrawList(screen, pfind.ClosedList, []color.Color{color.RGBA{255, 24, 125, 255}, color.RGBA{200, 24, 100, 255}, color.RGBA{180, 24, 70, 255}}, &pfind.GridOptions)

		// for _, node := range pfind.ClosedList {
		// 	pfind.IMat.DrawAGridTile_With_Lines(screen, node.Position, color.RGBA{255, 24, 125, 255}, &pfind.GridOptions)

		// }
	}
	if len(pfind.OpenList) > 1 && pfind.Button_SHOW_OPENLIST.GetState() != 2 {
		// for _, node := range pfind.OpenList {
		// 	pfind.IMat.DrawAGridTile_With_Lines(screen, node.Position, color.RGBA{125, 128, 255, 255}, &pfind.GridOptions)

		// }
		pfind.IMat.NodeList_DrawList(screen, pfind.OpenList, []color.Color{color.RGBA{125, 128, 255, 255}, color.RGBA{100, 108, 222, 255}, color.RGBA{100, 108, 200, 255}}, &pfind.GridOptions)

	}
	if len(pfind.BlockedList) > 1 && pfind.Button_SHOW_BLOCKEDLIST.GetState() != 2 {
		pfind.IMat.NodeList_DrawList(screen, pfind.BlockedList, []color.Color{color.RGBA{75, 75, 75, 255}, color.RGBA{50, 50, 50, 255}, color.RGBA{125, 125, 125, 255}}, &pfind.GridOptions)
	}
	if pfind.EndNode != nil && pfind.Button_SHOW_PATH.GetState() != 2 { //Button_SHOW_PATH
		// pfind.IMat.DrawAGridTile_With_Lines(screen, pfind.EndNode.Position, color.RGBA{255, 255, 255, 255}, &pfind.GridOptions)
		pfind.IMat.ImatNode_Draw(screen, pfind.EndNode, []color.Color{color.RGBA{180, 255, 125, 255}, color.RGBA{120, 155, 155, 255}, color.RGBA{120, 120, 75, 255}}, &pfind.GridOptions)
	}
}
