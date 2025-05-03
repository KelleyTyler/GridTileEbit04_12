package coords

import "fmt"

/*

 */
type CoordsI_Line_Seg struct {
	Point0, Point1 CoordInts
}

/**/
func Create_CoordsI_LineSeg_00(x0, y0, x1, y1 int) (outLine CoordsI_Line_Seg) {
	outLine = CoordsI_Line_Seg{Point0: CoordInts{X: x0, Y: y0}, Point1: CoordInts{X: x1, Y: y1}}
	return outLine
}

/*
	returns a rounded midpoint as a coordint
*/
func (clis CoordsI_Line_Seg) GetMidpoint_Coordints() (Midpoint CoordInts) {
	x0, y0 := float64(clis.Point0.X), float64(clis.Point0.Y)
	x1, y1 := float64(clis.Point1.X), float64(clis.Point1.Y)
	Midpoint_X := (x1 + x0) / 2.0
	Midpoint_Y := (y1 + y0) / 2.0
	Midpoint = CoordInts{X: int(Midpoint_X), Y: int(Midpoint_Y)}
	return Midpoint
}

/*
 returns the midpoint as two float64 values for the sake of precision in some applications
*/
func (clis CoordsI_Line_Seg) GetMidpoint_2Floats() (Midpoint_X, Midpoint_Y float64) {
	x0, y0 := float64(clis.Point0.X), float64(clis.Point0.Y)
	x1, y1 := float64(clis.Point1.X), float64(clis.Point1.Y)
	Midpoint_X = (x1 + x0) / 2.0
	Midpoint_Y = (y1 + y0) / 2.0

	return Midpoint_X, Midpoint_Y
}

/*
this is to see if another line does in fact intersect;

	totally unsure if this will work at all

*/
func (clis CoordsI_Line_Seg) DoesIntersect(line2 CoordsI_Line_Seg) bool {
	// x0_0, y0_0, x0_1, y0_1 := clis.Get_Vals_As_Floats()
	// x1_0, y1_0, x1_1, y1_1 := line2.Get_Vals_As_Floats()
	dx_0, dy_0 := clis.GetDifferences_Float()
	dx_1, dy_1 := line2.GetDifferences_Float()
	a, b := line2.Point1.GetDifferences_TwoInts(clis.Point0)
	// c, d:=
	p0 := dy_1*(float64(a)) - dx_1*(float64(b))
	a, b = line2.Point1.GetDifferences_TwoInts(clis.Point1)
	p1 := dy_1*(float64(a)) - dx_1*(float64(b))
	a, b = clis.Point1.GetDifferences_TwoInts(line2.Point0)
	p2 := dy_0*(float64(a)) - dx_0*(float64(b))
	a, b = clis.Point1.GetDifferences_TwoInts(line2.Point1)
	p3 := dy_0*(float64(a)) - dx_0*(float64(b))
	return ((p0 * p1) <= 0) && ((p2 * p3) <= 0)
}

/*
returns a string listing the points; this string is about 34-38 characters in length depending on the maximum coordinate point
*/
func (clis CoordsI_Line_Seg) ToString() (outstring string) {
	outstring = fmt.Sprintf("X0: %3d, Y0: %3d; X1: %3d, Y1: %3d", clis.Point0.X, clis.Point0.Y, clis.Point1.X, clis.Point1.Y)
	return outstring
}

/*
	returns both coordinates in the form of four floating point values
*/
func (clis CoordsI_Line_Seg) Get_Vals_As_Floats() (x0, y0, x1, y1 float64) {
	x0 = float64(clis.Point0.X)
	y0 = float64(clis.Point0.Y)
	x1 = float64(clis.Point1.X)
	y1 = float64(clis.Point1.Y)
	return x0, y0, x1, y1
}

/*
	returns both coordinates in the form of four Integer values
*/
func (clis CoordsI_Line_Seg) Get_Vals_As_Ints() (x0, y0, x1, y1 int) {
	x0 = (clis.Point0.X)
	y0 = (clis.Point0.Y)
	x1 = (clis.Point1.X)
	y1 = (clis.Point1.Y)
	return x0, y0, x1, y1
}

/*
returns the differences between the two coordints as
*/
func (clis CoordsI_Line_Seg) GetDifferences_Float() (dx, dy float64) {
	x0, y0, x1, y1 := clis.Get_Vals_As_Floats()
	dx = x1 - x0
	dy = y1 - y0
	return dx, dy
}

/**/
func LineIntersectTest() {
	line0 := Create_CoordsI_LineSeg_00(1, 1, 15, 15)
	line1 := Create_CoordsI_LineSeg_00(3, 3, 5, 25)
	fmt.Printf("Does [%s]and [%s] Intersect: %t\n", line0.ToString(), line1.ToString(), line0.DoesIntersect(line1))
	line0 = Create_CoordsI_LineSeg_00(1, 1, 15, 15)
	line1 = Create_CoordsI_LineSeg_00(16, 0, 16, 25)
	fmt.Printf("Does [%s]and [%s] Intersect: %t\n", line0.ToString(), line1.ToString(), line0.DoesIntersect(line1))
	line0 = Create_CoordsI_LineSeg_00(1, 1, 15, 15)
	line1 = Create_CoordsI_LineSeg_00(15, 0, 15, 25)
	fmt.Printf("Does [%s]and [%s] Intersect: %t\n", line0.ToString(), line1.ToString(), line0.DoesIntersect(line1))
}
