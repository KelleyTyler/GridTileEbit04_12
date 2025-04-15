package framework

import (
	"fmt"
	"image/color"
	"math"

	coords "github.com/KelleyTyler/GridTileEbit04_12/myPkgs/basic_geometry/coords"
	mat "github.com/KelleyTyler/GridTileEbit04_12/myPkgs/basic_geometry/matrix"
	ui "github.com/KelleyTyler/GridTileEbit04_12/myPkgs/userinterface"
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
	xx, yy := gb.IMat.GetCursorBounds(gb.BoardOptions)
	xx += int(float32(gb.BoardOptions.BoardMargin.X) * 2.5)
	yy += int(float32(gb.BoardOptions.BoardMargin.Y) * 2.5)
	gb.Bounds = coords.CoordInts{X: xx, Y: yy}
	fmt.Printf("INTI: %d %d\n", xx, yy)
	gb.Img = ebiten.NewImage(xx, yy)
	gb.Board_Buffer_Img = ebiten.NewImage(xx, yy)
	gb.Board_Overlay_Buffer_Img = ebiten.NewImage(xx, yy) //644
	gb.ticker = 0
	gb.ticker_max = 16
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
				//fmt.Printf("IS ON TILE %d %d\n", tempX, tempY)
			}
		} else {
			//fmt.Printf("cursor out of bounds\n")
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

}

func (gb *GameBoard) MazeGenPassthrough() {
	gb.MazeGen.RunPrimlike(5)
	gb.BoardChanges = true
	gb.BoardOverlayChanges = true

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

	if len(gb.MazeGen.CurrentList) > 1 {
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
