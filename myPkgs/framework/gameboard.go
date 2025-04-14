package framework

import (
	"fmt"
	"image/color"

	coords "github.com/KelleyTyler/GridTileEbit04_12/myPkgs/basic_geometry/coords"
	mat "github.com/KelleyTyler/GridTileEbit04_12/myPkgs/basic_geometry/matrix"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type GameBoard struct {
	Board_Buffer_Img         *ebiten.Image
	Board_Overlay_Buffer_Img *ebiten.Image
	Img                      *ebiten.Image
	Position                 coords.CoordInts
	outMsg                   string
	IMat                     mat.IntegerMatrix2D
	DefaultColors            []color.Color
	BoardOptions             mat.Integer_Matrix_Ebiten_DrawOptions
	BoardChanges             bool
	BoardOverlayChanges      bool
	ticker                   int
	ticker_max               int

	MazeGen mat.MazeMaker
}

func (gb *GameBoard) Init(position, bposition, boardMargin, BoardSize, tilesize, tilespacing coords.CoordInts) {
	gb.Position = position
	gb.IMat.Init(BoardSize.Y, BoardSize.X, 10)
	gb.BoardOptions = mat.Integer_Matrix_Ebiten_DrawOptions{
		BoardPosition:     bposition,
		BoardMargin:       boardMargin,
		TileSize:          tilesize,
		TileSpacing:       tilespacing,
		ShowTileLines:     []bool{false, false, false},
		TileLineColors:    []color.Color{color.Black, color.Black, color.Black},
		TileLineThickness: []float32{1.0, 1.0, 1.0},
	}
	xx, yy := gb.IMat.GetCursorBounds(gb.BoardOptions)
	gb.Img = ebiten.NewImage(644, 644)
	gb.Board_Buffer_Img = ebiten.NewImage(xx, yy)
	gb.Board_Overlay_Buffer_Img = ebiten.NewImage(xx, yy)
	gb.ticker = 0
	gb.ticker_max = 16
	gb.DefaultColors = []color.Color{color.RGBA{55, 55, 75, 255}, color.RGBA{125, 125, 150, 255}, color.RGBA{80, 180, 80, 255},
		color.RGBA{0, 150, 150, 255}, color.RGBA{55, 65, 95, 255}, color.RGBA{255, 255, 255, 255}, color.RGBA{75, 75, 75, 255}}
	tempBoardOps := mat.Integer_Matrix_Ebiten_DrawOptions{
		BoardPosition:     bposition,
		BoardMargin:       boardMargin,
		TileSize:          tilesize,
		TileSpacing:       tilespacing,
		ShowTileLines:     []bool{true, true, true},
		TileLineColors:    []color.Color{color.RGBA{150, 150, 150, 255}, color.RGBA{180, 40, 40, 255}, color.RGBA{180, 40, 40, 255}},
		TileLineThickness: []float32{1.0, 1.0, 1.0},
	}
	gb.MazeGen.Init(tempBoardOps, &gb.IMat)
}

func (gb *GameBoard) Update() {
	if gb.ticker > gb.ticker_max {
		gb.BoardChanges = true
		gb.BoardOverlayChanges = true
		gb.ticker = 0
	} else {
		gb.ticker++
	}
	//onclick()
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButton0) {
		xx, yy := ebiten.CursorPosition()
		if gb.IsCursorInBounds() {
			tempX, tempY, isOnTile := gb.IMat.GetCoordOfMouseEvent(xx, yy, gb.Position.X, gb.Position.Y, gb.BoardOptions)
			if isOnTile {
				gb.MazeGen.CurrentList = append(gb.MazeGen.CurrentList, coords.CoordInts{X: tempX, Y: tempY})
				fmt.Printf("IS ON TILE %d %d\n", tempX, tempY)
			}
		}
	}
}

func (gb *GameBoard) MazeGenPassthrough() {
	gb.MazeGen.RunPrimlike(5)
	gb.BoardChanges = true
	gb.BoardOverlayChanges = true

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
	return (xx > gb.Position.X && xx < gb.Position.X+644) && (yy > gb.Position.Y && yy < gb.Position.Y+644)
}
