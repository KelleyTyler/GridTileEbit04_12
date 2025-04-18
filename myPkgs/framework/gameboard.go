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

	MazeGen mat.MazeMaker

	mapMove       bool
	mapMoveVector coords.CoordInts

	Reset_Map_Btn       ui.UI_Button
	New_Map_Button      ui.UI_Button
	Redraw_Tiles_Button ui.UI_Button
	Load_Map_Button     ui.UI_Button
	Save_Map_Button     ui.UI_Button
	Maze_Gen_Button     ui.UI_Button
	//  ui.UI_Button
	// ui.UI_Button
	NumSelect_TileSize_X    ui.UI_Num_Select
	NumSelect_TileSize_Y    ui.UI_Num_Select
	NumSelect_Tile_Margin_X ui.UI_Num_Select
	NumSelect_Tile_Margin_Y ui.UI_Num_Select
	NumSelect_MapSize_X     ui.UI_Num_Select
	NumSelect_MapSize_Y     ui.UI_Num_Select
}

func (gb *GameBoard) Init(backend *ui.UI_Backend, position, boardMargin, BoardSize, tilesize, tilespacing coords.CoordInts) {
	gb.Position = position
	gb.IMat.Init(BoardSize.Y, BoardSize.X, 10)
	gb.UI_Backend = backend
	gb.BoardOptions = mat.Integer_Matrix_Ebiten_DrawOptions{
		BoardPosition:     boardMargin,
		BoardMargin:       boardMargin,
		TileSize:          tilesize,
		TileSpacing:       tilespacing,
		ShowTileLines:     []bool{false, false, false},
		TileLineColors:    []color.Color{color.Black, color.Black, color.Black},
		TileLineThickness: []float32{1.0, 1.0, 1.0},
	}
	// xx, yy := gb.IMat.GetCursorBounds(gb.BoardOptions)
	// xx += int(float32(gb.BoardOptions.BoardMargin.X) * 2.5)
	// yy += int(float32(gb.BoardOptions.BoardMargin.Y) * 2.5)
	gb.Bounds = coords.CoordInts{X: 590, Y: 590}
	fmt.Printf("INTI: %d %d\n", 590, 590)
	gb.Img = ebiten.NewImage(590, 590)
	gb.Board_Buffer_Img = ebiten.NewImage(590, 590)
	gb.Board_Overlay_Buffer_Img = ebiten.NewImage(590, 590) //644
	gb.ticker = 0
	gb.ticker_max = 6
	gb.DefaultColors = []color.Color{color.RGBA{55, 55, 75, 255}, color.RGBA{125, 125, 150, 255}, color.RGBA{80, 180, 80, 255},
		color.RGBA{0, 150, 150, 255}, color.RGBA{55, 65, 95, 255}, color.RGBA{255, 255, 255, 255}, color.RGBA{75, 75, 75, 255}}
	tempBoardOps := mat.Integer_Matrix_Ebiten_DrawOptions{
		BoardPosition:     boardMargin,
		BoardMargin:       boardMargin,
		TileSize:          tilesize,
		TileSpacing:       tilespacing,
		ShowTileLines:     []bool{true, true, true},
		TileLineColors:    []color.Color{color.RGBA{150, 150, 150, 255}, color.RGBA{180, 40, 40, 255}, color.RGBA{180, 40, 40, 255}},
		TileLineThickness: []float32{1.0, 1.0, 1.0},
	}
	gb.MazeGen.Init(tempBoardOps, &gb.IMat)
	gb.mapMove = false
	gb.mapMoveVector = coords.CoordInts{X: 0, Y: 0}
	gb.UI_INIT()
}

func (gb *GameBoard) Update() {

	if gb.ticker > gb.ticker_max {
		gb.BoardChanges = true
		gb.BoardOverlayChanges = true
		gb.ticker = 0
	} else {
		gb.ticker++
	}
	gb.MouseMove()
	xx, yy := ebiten.CursorPosition()
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButton0) {

		if gb.IsCursorInBounds() {
			tempX, tempY, isOnTile := gb.IMat.GetCoordOfMouseEvent(xx, yy, gb.Position.X, gb.Position.Y, gb.BoardOptions)
			if isOnTile {
				gb.MazeGen.CurrentList = append(gb.MazeGen.CurrentList, coords.CoordInts{X: tempX, Y: tempY})
				gb.UI_Backend.PlaySound(4)
				gb.BoardOverlayChanges = true
				// fmt.Printf("IS ON TILE %d %d\n", tempX, tempY)
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
	gb.UI_UPDATE()
}

func (gb *GameBoard) MazeGenPassthrough() {
	if len(gb.MazeGen.CurrentList) > 0 {
		gb.MazeGen.RunPrimlike(5)

		if gb.MazeGen.HasFinished {
			fmt.Printf("done\n")
			gb.Maze_Gen_Button.State = 0
			gb.Maze_Gen_Button.IsToggled = false
			gb.MazeGen.HasStarted = false
			// gb.Maze_Gen_Button.Redraw()
			// gb.Maze_Gen_Button.Parent.Redraw()
			gb.MazeGen.HasFinished = false
		}
		// gb.BoardChanges = true
		// gb.BoardOverlayChanges = true
	}

}

func (gb *GameBoard) UI_INIT() {

	// Reset_Map_Btn           ui.UI_Button
	// New_Map_Button            ui.UI_Button
	// NumSelect_TileSize_X    ui.UI_Num_Select
	// NumSelect_TileSize_Y    ui.UI_Num_Select
	// NumSelect_Tile_Margin_X ui.UI_Num_Select
	// NumSelect_Tile_Margin_Y ui.UI_Num_Select
	// NumSelect_MapSize_X     ui.UI_Num_Select
	// NumSelect_MapSize_Y     ui.UI_Num_Select

	/*
		game.Board.NumSelect_MapSize_X.SetPosition(coords.CoordInts{X: 4, Y: 72})
		game.Board.NumSelect_MapSize_Y.SetPosition(coords.CoordInts{X: 136, Y: 72})
		game.Board.New_Map_Button.Position = coords.CoordInts{X: 74, Y: 72}
		game.Board.Reset_Map_Btn.Position = coords.CoordInts{X: 74, Y: 112}
		game.Board.NumSelect_TileSize_X.SetPosition(coords.CoordInts{X: 4, Y: 112})
		game.Board.NumSelect_TileSize_Y.SetPosition(coords.CoordInts{X: 136, Y: 112})

		game.Board.NumSelect_Tile_Margin_X.SetPosition(coords.CoordInts{X: 4, Y: 152})
		game.Board.NumSelect_Tile_Margin_Y.SetPosition(coords.CoordInts{X: 136, Y: 152})

	*/
	gb.Load_Map_Button.Init([]string{"r_map_btn", "LOAD\nMAP"}, gb.UI_Backend, nil, coords.CoordInts{X: 70, Y: 36}, coords.CoordInts{X: 64, Y: 32})
	gb.Save_Map_Button.Init([]string{"r_map_btn", "SAVE\nMAP"}, gb.UI_Backend, nil, coords.CoordInts{X: 4, Y: 36}, coords.CoordInts{X: 64, Y: 32})
	gb.Maze_Gen_Button.Init([]string{"maze_gen_btn", "Gen\nPrimlike"}, gb.UI_Backend, nil, coords.CoordInts{X: 4, Y: 192}, coords.CoordInts{X: 64, Y: 32})
	gb.Redraw_Tiles_Button.Init([]string{"r_map_btn", "Redraw\nTiles"}, gb.UI_Backend, nil, coords.CoordInts{X: 70, Y: 152}, coords.CoordInts{X: 64, Y: 32})
	gb.Reset_Map_Btn.Init([]string{"r_map_btn", "ResetMap"}, gb.UI_Backend, nil, coords.CoordInts{X: 70, Y: 112}, coords.CoordInts{X: 64, Y: 32})
	gb.New_Map_Button.Init([]string{"n_map_btn", "New Map"}, gb.UI_Backend, nil, coords.CoordInts{X: 70, Y: 72}, coords.CoordInts{X: 64, Y: 32})

	gb.NumSelect_MapSize_X.Init([]string{"n_map_btn", "Map X"}, gb.UI_Backend, nil, coords.CoordInts{X: 4, Y: 72}, coords.CoordInts{X: 64, Y: 32}) //coords.CoordInts{X: 4, Y: 72}, coords.CoordInts{X: 68, Y: 36}
	gb.NumSelect_MapSize_X.SetVals(len(gb.IMat[0]), 1, 8, 64, 0)
	gb.NumSelect_MapSize_Y.Init([]string{"n_map_btn", "Map Y"}, gb.UI_Backend, nil, coords.CoordInts{X: 136, Y: 72}, coords.CoordInts{X: 64, Y: 32})
	gb.NumSelect_MapSize_Y.SetVals(len(gb.IMat), 1, 8, 64, 0)
	gb.NumSelect_TileSize_X.Init([]string{"n_map_btn", "TileX"}, gb.UI_Backend, nil, coords.CoordInts{X: 4, Y: 112}, coords.CoordInts{X: 64, Y: 32})
	gb.NumSelect_TileSize_X.SetVals(8, 1, 1, 64, 0)

	gb.NumSelect_TileSize_Y.Init([]string{"n_map_btn", "TileY"}, gb.UI_Backend, nil, coords.CoordInts{X: 136, Y: 112}, coords.CoordInts{X: 64, Y: 32})
	gb.NumSelect_TileSize_Y.SetVals(8, 1, 1, 64, 0)

	gb.NumSelect_Tile_Margin_X.Init([]string{"n_map_btn", "TileMX"}, gb.UI_Backend, nil, coords.CoordInts{X: 4, Y: 152}, coords.CoordInts{X: 64, Y: 32})
	gb.NumSelect_Tile_Margin_X.SetVals(2, 1, -1, 16, 0)
	gb.NumSelect_Tile_Margin_Y.Init([]string{"n_map_btn", "TileMY"}, gb.UI_Backend, nil, coords.CoordInts{X: 136, Y: 152}, coords.CoordInts{X: 64, Y: 32})
	gb.NumSelect_Tile_Margin_Y.SetVals(2, 1, -1, 16, 0)

}

func (gb *GameBoard) UI_UPDATE() {
	if gb.Reset_Map_Btn.State == 2 {
		gb.IMat.ClearMatrix_To(0)
	}
	if gb.Redraw_Tiles_Button.State == 2 {
		fmt.Printf("REDRAW TILES BTN PRESSED\n")
		gb.ResetBoard(coords.CoordInts{X: gb.NumSelect_TileSize_X.CurrValue, Y: gb.NumSelect_TileSize_Y.CurrValue}, coords.CoordInts{X: gb.NumSelect_Tile_Margin_X.CurrValue, Y: gb.NumSelect_Tile_Margin_Y.CurrValue})

	}
	if gb.Maze_Gen_Button.GetState() >= 2 {
		gb.MazeGenPassthrough()
	}
}

func (gb *GameBoard) ResetBoard(new_tsize, new_t_spacing coords.CoordInts) {
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
	xx += int(float32(gb.BoardOptions.BoardMargin.X) * 2.5)
	yy += int(float32(gb.BoardOptions.BoardMargin.Y) * 2.5)
	gb.Bounds = coords.CoordInts{X: 582, Y: 582}
	fmt.Printf("INTI: %d %d\n", xx, yy)
	gb.Img = ebiten.NewImage(590, 590)
	gb.Board_Buffer_Img = ebiten.NewImage(xx, yy)
	gb.Board_Overlay_Buffer_Img = ebiten.NewImage(xx, yy) //644

	gb.MazeGen.DisplaySettings.TileSize = new_tsize
	gb.MazeGen.DisplaySettings.TileSpacing = new_t_spacing
	gb.BoardOverlayChanges = true
	gb.BoardChanges = true
}

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
			gb.BoardOverlayChanges = true
		}
	} else {
		if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
			gb.BoardOptions.BoardPosition = gb.BoardOptions.BoardMargin
		}
	}
}
func (gb *GameBoard) Draw(screen *ebiten.Image) {
	if gb.BoardChanges {
		gb.Board_Buffer_Img.Clear()
		gb.IMat.DrawFullGridFromColors(gb.Board_Buffer_Img, gb.DefaultColors, true, &gb.BoardOptions)
		gb.BoardChanges = false
	}
	if gb.BoardOverlayChanges {
		gb.DrawOverlay()
		gb.BoardOverlayChanges = false
	}
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
}
func (gb *GameBoard) DrawOverlay() {
	gb.Board_Overlay_Buffer_Img.Clear()

	if len(gb.MazeGen.CurrentList) > 0 {
		gb.IMat.DrawCoordListWithLines(gb.Board_Overlay_Buffer_Img, gb.MazeGen.CurrentList, []color.Color{color.RGBA{0, 150, 150, 255}}, gb.MazeGen.DisplaySettings)
	}

}
func (gb *GameBoard) ToString() string {
	gb.outMsg = fmt.Sprintf("GAMEBOARD:\n Position: %s\n", gb.BoardOptions.BoardPosition.ToString())
	return gb.outMsg
}

func (gb *GameBoard) IsCursorInBounds() bool {
	xx, yy := ebiten.CursorPosition()
	return (xx > gb.Position.X && xx < gb.Position.X+gb.Bounds.X) && (yy > gb.Position.Y && yy < gb.Position.Y+gb.Bounds.Y)
}
