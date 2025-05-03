package framework

import (
	"github.com/KelleyTyler/GridTileEbit04_12/myPkgs/basic_geometry/coords"
	mat "github.com/KelleyTyler/GridTileEbit04_12/myPkgs/basic_geometry/matrix"
	ui "github.com/KelleyTyler/GridTileEbit04_12/myPkgs/user_interface"
	"github.com/hajimehoshi/ebiten/v2"
)

/*
 */
type Mobile_Entity struct {
	Position, Target     coords.CoordInts
	Movement_Speed       int
	Pathfinding_Path     *mat.ImatNode
	Backend              *ui.UI_Backend
	IMat                 *mat.IntegerMatrix2D
	Style                *ui.UI_Object_Style
	Ticker, Ticker_Limit int
}

/*
 */
func (m_e *Mobile_Entity) Init(backend *ui.UI_Backend, position coords.CoordInts, m_Speed, tick_limit int, imat *mat.IntegerMatrix2D) {
	m_e.Position = position
	m_e.Backend = backend
	m_e.Movement_Speed = m_Speed
	m_e.Ticker = 0
	m_e.Ticker_Limit = tick_limit
}

/*
 */
func (m_e *Mobile_Entity) Update() {

}

/*
 */
func (m_e *Mobile_Entity) Draw(screen *ebiten.Image) {

}

/*
 */
type Mobile_Entity_Controller struct {
}

/*
 */
func (mob_ent_cont *Mobile_Entity_Controller) Init() {

}

/*
 */
func (mob_ent_cont *Mobile_Entity_Controller) Update() {

}

/*
 */
func (mob_ent_cont *Mobile_Entity_Controller) Draw(screen *ebiten.Image) {

}
