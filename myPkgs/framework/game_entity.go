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
	UI_Backend           *ui.UI_Backend
	IMat                 *mat.IntegerMatrix2D
	Style                *ui.UI_Object_Style
	Ticker, Ticker_Limit int
}

/*
Init([]string{"gameboard_panel_label", "GAMEBOARD"}, gb.UI_Backend, nil, coords.CoordInts{X: 0, Y: 0}, coords.CoordInts{X: 204, Y: 16})
*/
func (m_e *Mobile_Entity) Init(backend *ui.UI_Backend, position coords.CoordInts, m_Speed, tick_limit int, imat *mat.IntegerMatrix2D) {
	m_e.Position = position
	m_e.UI_Backend = backend
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
	IMat            *mat.IntegerMatrix2D
	Style           *ui.UI_Object_Style
	UI_Backend      *ui.UI_Backend
	GridDrawOptions mat.Integer_Matrix_Ebiten_DrawOptions
	//------------------------------------------------------
	Panel_Label         ui.UI_Label
	UI_Panel            ui.UI_Object_Primitive
	Btn_Place, Btn_Rest ui.UI_Button
}

/*
 */
func (mob_ent_cont *Mobile_Entity_Controller) Init(backend *ui.UI_Backend, style *ui.UI_Object_Style, options *mat.Integer_Matrix_Ebiten_DrawOptions, Pos coords.CoordInts, m_Speed, tick_limit int, imat *mat.IntegerMatrix2D) {
	mob_ent_cont.UI_Backend = backend
	if style != nil {
		mob_ent_cont.Style = style
	} else {
		mob_ent_cont.Style = &mob_ent_cont.UI_Backend.Style
	}

}

/*
 */
func (mob_ent_cont *Mobile_Entity_Controller) Init_User_Interface(parent ui.UI_Object, y_position int) {
	mob_ent_cont.UI_Panel.Init([]string{"mob_ent_cont_panel", "PRIMITIVE"}, mob_ent_cont.UI_Backend, mob_ent_cont.Style, coords.CoordInts{X: 0, Y: 0}, coords.CoordInts{X: 204, Y: 138})
	mob_ent_cont.Panel_Label.Init([]string{"mob_ent_cont_panel_label", "Mobile Entity Controller"}, mob_ent_cont.UI_Backend, mob_ent_cont.Style, coords.CoordInts{X: 0, Y: y_position}, coords.CoordInts{X: 204, Y: 16})
	mob_ent_cont.Panel_Label.TextAlignMode = 10
	mob_ent_cont.Btn_Place.Init([]string{"mob_ent_cont_btn", "Entity"}, mob_ent_cont.UI_Backend, mob_ent_cont.Style, coords.CoordInts{X: 4, Y: 34}, coords.CoordInts{X: 64, Y: 32})

	mob_ent_cont.Btn_Rest.Init([]string{"mob_ent_cont_btn", "Entity"}, mob_ent_cont.UI_Backend, mob_ent_cont.Style, coords.CoordInts{X: 70, Y: 34}, coords.CoordInts{X: 64, Y: 32})
}

/*
 */
func (mob_ent_cont *Mobile_Entity_Controller) Update() {

}

/*
 */
func (mob_ent_cont *Mobile_Entity_Controller) Draw(screen *ebiten.Image) {

}
