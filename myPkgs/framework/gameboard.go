package framework

import (
	"fmt"
	"image/color"
	"math"

	coords "github.com/KelleyTyler/GridTileEbit04_12/myPkgs/basic_geometry/coords"
	mat "github.com/KelleyTyler/GridTileEbit04_12/myPkgs/basic_geometry/matrix"
	ui "github.com/KelleyTyler/GridTileEbit04_12/myPkgs/user_interface"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

/**/
type GameBoard struct {
	Board_Buffer_Img         *ebiten.Image
	Board_Overlay_Buffer_Img *ebiten.Image
	Img                      *ebiten.Image
	UI_Backend               *ui.UI_Backend
	Position, Bounds         coords.CoordInts
	outMsg                   string
	IMat                     mat.IntegerMatrix2D
	DefaultColors            []color.Color
	BoardOptions             mat.Integer_Matrix_Ebiten_DrawOptions
	BoardChanges             bool
	BoardOverlayChanges      bool
	ticker                   int
	ticker_max               int

	mapMove       bool
	mapMoveVector coords.CoordInts

	SavePath           string
	GameBoard_UI_STATE uint8 // 10 is the normal 30 is 'loading window up' 40 is 'Save Window Up'

	Reset_Map_Btn         ui.UI_Button
	New_Map_Button        ui.UI_Button
	Redraw_Tiles_Button   ui.UI_Button
	Load_Map_Button       ui.UI_Button
	Save_Map_Button       ui.UI_Button
	NumSelect_ResetNumber ui.UI_Num_Select
	//------

	//  ui.UI_Button
	// ui.UI_Button
	NumSelect_TileSize_X    ui.UI_Num_Select
	NumSelect_TileSize_Y    ui.UI_Num_Select
	NumSelect_Tile_Margin_X ui.UI_Num_Select
	NumSelect_Tile_Margin_Y ui.UI_Num_Select
	NumSelect_MapSize_X     ui.UI_Num_Select
	NumSelect_MapSize_Y     ui.UI_Num_Select

	Window_Save, Window_Load ui.UI_Window
	//========================
	// GridmapPallet01 ui.UI_Button
	drawTool  Drawing_Tool
	mazeTool  Maze_Tool
	pfindTest Pathfind_Tester
}

/*
	Init

this initializes the gameboard

	[]color.Color{
			color.RGBA{125, 125, 125, 255},
			color.RGBA{115, 115, 115, 255},
			color.RGBA{80, 180, 80, 255}, //color.RGBA{80, 180, 80, 255},
			color.RGBA{0, 150, 150, 255}, //color.RGBA{0, 150, 150, 255},
			color.RGBA{55, 55, 75, 255},  //color.RGBA{55, 55, 75, 255},
			color.RGBA{55, 65, 95, 255},  //color.RGBA{55, 65, 95, 255},
			color.RGBA{65, 65, 65, 255},
			color.RGBA{55, 55, 55, 255},
			color.RGBA{45, 45, 45, 255},
			color.RGBA{35, 35, 35, 255},
			color.RGBA{25, 25, 25, 255},
			// color.RGBA{15, 15, 15, 255},
		}
*/
func (gb *GameBoard) Init(backend *ui.UI_Backend, UI_Panel_Parent ui.UI_Object, position, boardMargin, BoardSize, tilesize, tilespacing coords.CoordInts) {
	gb.Position = position
	gb.IMat.Init(BoardSize.Y, BoardSize.X, 10)
	gb.UI_Backend = backend
	gb.SavePath = "bin/Output"
	gb.GameBoard_UI_STATE = 10
	gb.BoardOptions = mat.Integer_Matrix_Ebiten_DrawOptions{
		BoardPosition:     boardMargin,
		BoardMargin:       boardMargin,
		TileSize:          tilesize,
		TileSpacing:       tilespacing,
		ShowTileLines:     []bool{false, false, false},
		TileLineColors:    []color.Color{color.Black, color.Black, color.Black},
		TileLineThickness: []float32{1.0, 1.0, 1.0},
	}
	gb.Redraw_Board_New_Params(tilesize, tilespacing)
	/*
		 []color.Color{
		 color.RGBA{55, 55, 75, 255},
		 color.RGBA{125, 125, 150, 255},
		 color.RGBA{80, 180, 80, 255},
			color.RGBA{0, 150, 150, 255},
			color.RGBA{55, 65, 95, 255},
			color.RGBA{255, 255, 255, 255},
			color.RGBA{75, 75, 75, 255},
			}


		[]color.Color{color.RGBA{255, 0, 0, 255}, color.RGBA{255, 255, 0, 255}, color.RGBA{0, 255, 0, 255},
		color.RGBA{0, 255, 255, 255}, color.RGBA{0, 0, 255, 255}, color.RGBA{255, 0, 255, 255}, color.RGBA{75, 75, 75, 255}}


		 []color.Color{
		color.RGBA{255, 0, 0, 255},
		color.RGBA{255, 150, 0, 255}, //1
		color.RGBA{255, 255, 0, 255},
		color.RGBA{150, 255, 0, 255}, //3
		color.RGBA{0, 255, 0, 255},
		color.RGBA{0, 255, 150, 255}, //5
		color.RGBA{0, 255, 255, 255},
		color.RGBA{0, 150, 255, 255}, //7
		color.RGBA{0, 0, 255, 255},
		color.RGBA{150, 0, 255, 255}, //9
		color.RGBA{255, 0, 255, 255},
		color.RGBA{255, 0, 150, 255}, //11
		}
	*/

	// xx, yy := gb.IMat.GetCursorBounds(gb.BoardOptions)
	// xx += int(float32(gb.BoardOptions.BoardMargin.X) * 2.5)
	// yy += int(float32(gb.BoardOptions.BoardMargin.Y) * 2.5)
	// gb.Bounds = coords.CoordInts{X: 5, Y: 590}
	// fmt.Printf("INTI: %d %d\n", 590, 590)
	// gb.Img = ebiten.NewImage(590, 590)
	// gb.Board_Buffer_Img = ebiten.NewImage(590, 590)
	// gb.Board_Overlay_Buffer_Img = ebiten.NewImage(590, 590) //644
	gb.ticker = 0
	gb.ticker_max = 64
	gb.DefaultColors = []color.Color{
		color.RGBA{125, 125, 125, 255},
		color.RGBA{255, 0, 0, 255},
		color.RGBA{255, 150, 0, 255},
		color.RGBA{255, 255, 0, 255}, //color.RGBA{80, 180, 80, 255},
		color.RGBA{150, 255, 0, 255}, //color.RGBA{0, 150, 150, 255},
		color.RGBA{0, 255, 0, 255},   //color.RGBA{55, 55, 75, 255},
		color.RGBA{0, 255, 150, 255}, //color.RGBA{55, 65, 95, 255},
		color.RGBA{0, 255, 255, 255},
		color.RGBA{0, 150, 255, 255},
		color.RGBA{65, 65, 70, 255},
		color.RGBA{55, 55, 75, 255},
		// color.RGBA{15, 15, 15, 255},
	}
	tempBoardOps := mat.Integer_Matrix_Ebiten_DrawOptions{
		BoardPosition:     boardMargin,
		BoardMargin:       boardMargin,
		TileSize:          tilesize,
		TileSpacing:       tilespacing,
		ShowTileLines:     []bool{true, true, true},
		TileLineColors:    []color.Color{color.RGBA{150, 150, 150, 255}, color.RGBA{180, 40, 40, 255}, color.RGBA{180, 40, 40, 255}},
		TileLineThickness: []float32{1.0, 1.0, 1.0},
	}

	gb.mapMove = false
	gb.mapMoveVector = coords.CoordInts{X: 0, Y: 0}
	gb.UI_INIT()
	gb.mazeTool.Init(&gb.IMat, backend, 4, &gb.BoardOptions)
	gb.drawTool.Init(&gb.IMat, backend, 2, tempBoardOps)
	gb.pfindTest.Init(&gb.IMat, gb.UI_Backend, &gb.drawTool.DisplaySettings)
	gb.SetParents(UI_Panel_Parent)

}

/**/
func (gb *GameBoard) Update() {

	if gb.GameBoard_UI_STATE == 10 {
		gb.MouseMove()
		xx, yy := ebiten.CursorPosition()
		if inpututil.IsMouseButtonJustPressed(ebiten.MouseButton0) {

			if gb.IsCursorInBounds() {
				tempX, tempY, isOnTile := gb.IMat.GetCoordOfMouseEvent(xx, yy, gb.Position.X, gb.Position.Y, gb.BoardOptions)
				if isOnTile {

					// gb.BoardOverlayChanges, gb.BoardChanges =
					if b, a := gb.drawTool.OnValidMouseClickOnGameBoard(tempX, tempY); a || b {
						if !gb.BoardChanges && a {
							gb.BoardChanges = true
						}
						if !gb.BoardOverlayChanges && b {
							gb.BoardOverlayChanges = true
						}
					}
					if b, a := gb.pfindTest.OnValidMouseClickOnGameBoard(tempX, tempY); a || b {
						if !gb.BoardChanges && a {
							gb.BoardChanges = true
						}
						if !gb.BoardOverlayChanges && b {
							gb.BoardOverlayChanges = true
						}
					}
					// fmt.Printf("IS ON TILE %d %d\n", tempX, tempY)
					if b, a := gb.mazeTool.OnValidMouseClickOnGameBoard(tempX, tempY); a || b {
						if !gb.BoardChanges && a {
							gb.BoardChanges = true
						}
						if !gb.BoardOverlayChanges && b {
							gb.BoardOverlayChanges = true
						}
					}
				}
			} else {
				// fmt.Printf("cursor out of bounds\n")
			}
		}
		if inpututil.IsMouseButtonJustPressed(ebiten.MouseButton2) && gb.IsCursorInBounds() { //or if touchinput??
			gb.mapMove = true
			gb.mapMoveVector = coords.CoordInts{X: xx, Y: yy}
		}
		if inpututil.IsMouseButtonJustReleased(ebiten.MouseButton2) && gb.mapMove { //or if touchinput??

			gb.mapMove = false
			gb.mapMoveVector = coords.CoordInts{X: xx, Y: yy}
		}
	} else {
		gb.Save_Load_Update()
	}
	gb.UI_UPDATE()

}

/*
 */
func (gb *GameBoard) UI_INIT() {

	gb.Load_Map_Button.Init([]string{"r_map_btn", "LOAD\nMAP"}, gb.UI_Backend, nil, coords.CoordInts{X: 136, Y: 34}, coords.CoordInts{X: 64, Y: 32})
	gb.Save_Map_Button.Init([]string{"r_map_btn", "SAVE\nMAP"}, gb.UI_Backend, nil, coords.CoordInts{X: 4, Y: 34}, coords.CoordInts{X: 64, Y: 32})
	gb.NumSelect_ResetNumber.Init([]string{"n_map_btn", "mazeGen00"}, gb.UI_Backend, nil, coords.CoordInts{X: 70, Y: 34}, coords.CoordInts{X: 64, Y: 32})
	gb.NumSelect_ResetNumber.SetVals(10, 1, 0, 10, 0)
	gb.New_Map_Button.Init([]string{"n_map_btn", "New Map"}, gb.UI_Backend, nil, coords.CoordInts{X: 70, Y: 68}, coords.CoordInts{X: 64, Y: 32})
	gb.NumSelect_MapSize_X.Init([]string{"n_map_btn", "Map X"}, gb.UI_Backend, nil, coords.CoordInts{X: 4, Y: 68}, coords.CoordInts{X: 64, Y: 32}) //coords.CoordInts{X: 4, Y: 72}, coords.CoordInts{X: 68, Y: 36}
	gb.NumSelect_MapSize_X.SetVals(len(gb.IMat[0]), 1, 8, gb.UI_Backend.Settings.GameBoardXMax, 0)
	gb.NumSelect_MapSize_Y.Init([]string{"n_map_btn", "Map Y"}, gb.UI_Backend, nil, coords.CoordInts{X: 136, Y: 68}, coords.CoordInts{X: 64, Y: 32})
	gb.NumSelect_MapSize_Y.SetVals(len(gb.IMat), 1, 8, gb.UI_Backend.Settings.GameBoardYMax, 0)
	gb.NumSelect_TileSize_X.Init([]string{"n_map_btn", "TileX"}, gb.UI_Backend, nil, coords.CoordInts{X: 4, Y: 68 + 34}, coords.CoordInts{X: 64, Y: 32})
	gb.NumSelect_TileSize_X.SetVals(8, 1, 1, 64, 0)
	gb.Reset_Map_Btn.Init([]string{"r_map_btn", "ResetMap"}, gb.UI_Backend, nil, coords.CoordInts{X: 70, Y: 68 + 34}, coords.CoordInts{X: 64, Y: 32})

	gb.NumSelect_TileSize_Y.Init([]string{"n_map_btn", "TileY"}, gb.UI_Backend, nil, coords.CoordInts{X: 136, Y: 68 + 34}, coords.CoordInts{X: 64, Y: 32})
	gb.NumSelect_TileSize_Y.SetVals(8, 1, 1, 64, 0)

	gb.NumSelect_Tile_Margin_X.Init([]string{"n_map_btn", "TileMX"}, gb.UI_Backend, nil, coords.CoordInts{X: 4, Y: 68 + 68}, coords.CoordInts{X: 64, Y: 32})
	gb.NumSelect_Tile_Margin_X.SetVals(0, 1, 0, 16, 0)
	gb.Redraw_Tiles_Button.Init([]string{"r_map_btn", "Redraw\nTiles"}, gb.UI_Backend, nil, coords.CoordInts{X: 70, Y: 68 + 68}, coords.CoordInts{X: 64, Y: 32})
	gb.NumSelect_Tile_Margin_Y.Init([]string{"n_map_btn", "TileMY"}, gb.UI_Backend, nil, coords.CoordInts{X: 136, Y: 68 + 68}, coords.CoordInts{X: 64, Y: 32})
	gb.NumSelect_Tile_Margin_Y.SetVals(0, 1, 0, 16, 0)

	gb.Window_Save.Init([]string{"window_save", "SAVE WINDOW"}, gb.UI_Backend, nil, coords.CoordInts{X: 55, Y: 150}, coords.CoordInts{X: 256, Y: 128})
	gb.Window_Save.Redraw()
	gb.Window_Load.Init([]string{"window_load", "LOAD WINDOW"}, gb.UI_Backend, nil, coords.CoordInts{X: 55, Y: 150}, coords.CoordInts{X: 256, Y: 128})
	gb.Window_Load.Redraw()
}

/**/
func (gb *GameBoard) UI_UPDATE() {
	// strng :=
	// g.MazeTextBox.
	if gb.Reset_Map_Btn.GetState() == 2 {
		gb.IMat.ClearMatrix_To(gb.NumSelect_ResetNumber.CurrValue)
		gb.mazeTool.Reset(gb.BoardOptions)
		gb.drawTool.Clear()
		gb.Redraw_Board_New_Params(coords.CoordInts{X: gb.NumSelect_TileSize_X.CurrValue, Y: gb.NumSelect_TileSize_Y.CurrValue}, coords.CoordInts{X: gb.NumSelect_Tile_Margin_X.CurrValue, Y: gb.NumSelect_Tile_Margin_Y.CurrValue})

		gb.BoardChanges = true

	}
	if gb.Redraw_Tiles_Button.GetState() == 2 {
		fmt.Printf("REDRAW TILES BTN PRESSED\n")
		gb.Redraw_Board_New_Params(coords.CoordInts{X: gb.NumSelect_TileSize_X.CurrValue, Y: gb.NumSelect_TileSize_Y.CurrValue}, coords.CoordInts{X: gb.NumSelect_Tile_Margin_X.CurrValue, Y: gb.NumSelect_Tile_Margin_Y.CurrValue})
		gb.BoardChanges = true

		gb.Redraw_Board()

	}
	if gb.New_Map_Button.GetState() == 2 {
		//gb.NumSelect_ResetNumber
		gb.NewBoard(coords.CoordInts{X: gb.NumSelect_MapSize_X.CurrValue, Y: gb.NumSelect_MapSize_Y.CurrValue})
		gb.Redraw_Board_New_Params(coords.CoordInts{X: gb.NumSelect_TileSize_X.CurrValue, Y: gb.NumSelect_TileSize_Y.CurrValue}, coords.CoordInts{X: gb.NumSelect_Tile_Margin_X.CurrValue, Y: gb.NumSelect_Tile_Margin_Y.CurrValue})
		gb.BoardChanges = true

	}

	if b, a := gb.mazeTool.Update_Passive(); a || b {
		if !gb.BoardChanges && a {
			gb.BoardChanges = true
		}
		if !gb.BoardOverlayChanges && b {
			gb.BoardOverlayChanges = true
		}
	}

	if gb.Save_Map_Button.GetState() == 2 {
		gb.Save_Button_Pressed()
	}
	if gb.Load_Map_Button.GetState() == 2 {
		gb.Load_Button_Pressed()
	}

	if gb.pfindTest.Update_Passive() {
		gb.BoardOverlayChanges = true
		gb.BoardChanges = true

	}

}

/**/
func (gb *GameBoard) SetParents(parent ui.UI_Object) {
	gb.Load_Map_Button.Init_Parents(parent)
	gb.Save_Map_Button.Init_Parents(parent)
	gb.NumSelect_ResetNumber.Init_Parents(parent)
	//-----------------------
	gb.New_Map_Button.Init_Parents(parent)
	gb.Reset_Map_Btn.Init_Parents(parent) //MazeGenBtn00
	//------
	gb.NumSelect_MapSize_X.Init_Parents(parent)
	gb.NumSelect_MapSize_Y.Init_Parents(parent)
	gb.NumSelect_TileSize_Y.Init_Parents(parent)
	gb.NumSelect_TileSize_X.Init_Parents(parent)
	gb.NumSelect_Tile_Margin_X.Init_Parents(parent)
	gb.NumSelect_Tile_Margin_Y.Init_Parents(parent)
	gb.Redraw_Tiles_Button.Init_Parents(parent)

	gb.mazeTool.UI_Init(parent, 170)
	gb.drawTool.InitUI(parent, 306)
	gb.pfindTest.UI_Init(parent, 444)

	//--------------------------------
	// gb.Window_Save.Init_Parents(parent)
}

/**/
func (gb *GameBoard) NewBoard(new_Bsize coords.CoordInts) {
	gb.IMat.Init(new_Bsize.X, new_Bsize.Y, gb.NumSelect_ResetNumber.CurrValue)
}

/**/
func (gb *GameBoard) Redraw_Board_New_Params(new_tsize, new_t_spacing coords.CoordInts) {
	gb.BoardOptions.BoardPosition = gb.BoardOptions.BoardMargin
	gb.BoardOptions = mat.Integer_Matrix_Ebiten_DrawOptions{
		BoardPosition:     gb.BoardOptions.BoardPosition,
		BoardMargin:       gb.BoardOptions.BoardMargin,
		TileSize:          new_tsize,
		TileSpacing:       new_t_spacing,
		ShowTileLines:     []bool{false, false, false},
		TileLineColors:    []color.Color{color.Black, color.Black, color.Black},
		TileLineThickness: []float32{1.0, 1.0, 1.0},
	}
	xx, yy := gb.IMat.GetCursorBounds(gb.BoardOptions)
	xx += int(float32(gb.BoardOptions.BoardMargin.X) * 2.0)
	yy += int(float32(gb.BoardOptions.BoardMargin.Y) * 2.0)
	gb.Bounds = coords.CoordInts{X: 590, Y: 590}
	fmt.Printf("INTI: %d %d\n", xx, yy)
	gb.Img = ebiten.NewImage(590, 590)
	gb.Board_Buffer_Img = ebiten.NewImage(xx, yy)
	gb.Board_Overlay_Buffer_Img = ebiten.NewImage(xx, yy) //644
	gb.mazeTool.Redef(gb.BoardOptions)
	gb.drawTool.DisplaySettings.TileSize = new_tsize
	gb.drawTool.DisplaySettings.TileSpacing = new_t_spacing
	gb.BoardOverlayChanges = true
	gb.BoardChanges = true
}

/**/
func (gb *GameBoard) Redraw_Board() {
	gb.BoardOptions.BoardPosition = gb.BoardOptions.BoardMargin

	xx, yy := gb.IMat.GetCursorBounds(gb.BoardOptions)
	xx += int(float32(gb.BoardOptions.BoardMargin.X) * 2.0)
	yy += int(float32(gb.BoardOptions.BoardMargin.Y) * 2.0)
	gb.Bounds = coords.CoordInts{X: 590, Y: 590}
	fmt.Printf("INTI: %d %d\n", xx, yy)
	gb.Img = ebiten.NewImage(590, 590)
	gb.Board_Buffer_Img = ebiten.NewImage(xx, yy)
	gb.Board_Overlay_Buffer_Img = ebiten.NewImage(xx, yy) //644
	gb.mazeTool.Redef(gb.BoardOptions)
	gb.BoardOverlayChanges = true
	gb.BoardChanges = true
}

/**/
func (gb *GameBoard) MouseMove() {
	if gb.mapMove {
		//temp := coords.CoordInts{X: xx, Y: yy}
		x0, y0 := ebiten.CursorPosition()
		x1 := gb.mapMoveVector.X
		y1 := gb.mapMoveVector.Y
		dx, dy := (gb.mapMoveVector.X - x0), (gb.mapMoveVector.Y - y0)
		t0 := gb.mapMoveVector.X == x0 && gb.mapMoveVector.Y == y0
		t1 := int(math.Abs(float64(dx))) < 2 && int(math.Abs(float64(dy))) < 2
		if t0 || t1 {
			gb.mapMoveVector = coords.CoordInts{X: x0, Y: y0}
		} else {
			x2, y2 := (x1-x0)/4.0, (y1-y0)/4.0
			//fmt.Printf("----- %d %d \n", x2, y2)
			gb.BoardOptions.BoardPosition.Y -= y2
			gb.BoardOptions.BoardPosition.X -= x2
			gb.mapMoveVector = coords.CoordInts{X: x0, Y: y0}
			gb.BoardChanges = true
			// gb.BoardOverlayChanges = true
		}
	} else {
		if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
			gb.BoardOptions.BoardPosition = gb.BoardOptions.BoardMargin
			// gb.BoardOverlayChanges = true

		}
	}
}

/**/
func (gb *GameBoard) Draw(screen *ebiten.Image) {
	gb.Redraw()
	ops := ebiten.DrawImageOptions{}
	ops.GeoM.Reset()
	ops.GeoM.Translate(float64(gb.BoardOptions.BoardPosition.X), float64(gb.BoardOptions.BoardPosition.Y))
	ops.GeoM.Scale(1.0, 1.0)
	gb.Img.Fill(color.RGBA{20, 20, 20, 255})
	gb.Img.DrawImage(gb.Board_Buffer_Img, &ops)
	gb.Img.DrawImage(gb.Board_Overlay_Buffer_Img, &ops)
	ops.GeoM.Reset()
	ops.GeoM.Translate(float64(gb.Position.X-gb.BoardOptions.BoardMargin.X), float64(gb.Position.Y-gb.BoardOptions.BoardMargin.Y))
	screen.DrawImage(gb.Img, &ops)

	///--------------------
	ops.GeoM.Reset()
	gb.Window_Save.Draw(screen)
	gb.Window_Load.Draw(screen)

	// if gb.GameBoard_UI_STATE == 40 {
	// 	gb.Window_Save.Draw(screen)
	// 	// screen.DrawImage(gb.Window_Save.Image, &ops)
	// }
	// if gb.GameBoard_UI_STATE == 30 {
	// 	gb.Window_Load.Draw(screen)

	// 	// screen.DrawImage(gb.Window_Load.Image, &ops)
	// }
}

/**/
func (gb *GameBoard) Redraw() {
	if gb.BoardChanges {
		gb.Board_Buffer_Img.Clear()
		gb.IMat.DrawFullGridFromColors(gb.Board_Buffer_Img, gb.DefaultColors, true, &gb.BoardOptions)
		gb.BoardChanges = false
	}
	if gb.BoardOverlayChanges {
		gb.DrawOverlay()
		gb.BoardOverlayChanges = false
	}
}

/**/
func (gb *GameBoard) DrawOverlay() {
	gb.Board_Overlay_Buffer_Img.Clear()

	if len(gb.drawTool.Points) > 0 {
		gb.IMat.DrawCoordListWithLines(gb.Board_Overlay_Buffer_Img, gb.drawTool.Points, []color.Color{color.RGBA{0, 150, 150, 255}}, gb.drawTool.DisplaySettings)
	}
	gb.pfindTest.Draw(gb.Board_Overlay_Buffer_Img, &gb.drawTool.DisplaySettings)
	gb.mazeTool.Draw(gb.Board_Overlay_Buffer_Img, gb.BoardOptions)
}

/**/
func (gb *GameBoard) ToString() string {
	gb.outMsg = fmt.Sprintf("GAMEBOARD:\n Position: %s\n", gb.BoardOptions.BoardPosition.ToString())
	return gb.outMsg
}

/**/
func (gb *GameBoard) IsCursorInBounds() bool {
	xx, yy := ebiten.CursorPosition()
	return (xx > gb.Position.X && xx < gb.Position.X+gb.Bounds.X) && (yy > gb.Position.Y && yy < gb.Position.Y+gb.Bounds.Y)
}
