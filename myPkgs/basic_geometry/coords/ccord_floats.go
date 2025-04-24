package coords

import (
	"fmt"
	"math"
)

type CoordFloat2D struct {
	X, Y float64
}

func MakeCoordFloat2D(x, y float64) (outCoord CoordFloat2D) {
	outCoord.X = x
	outCoord.Y = y
	return outCoord
}
func (cf2D CoordFloat2D) ToString() (outstring string) {
	outstring = fmt.Sprintf("X:%7.3f, Y:%7.3f", cf2D.X, cf2D.Y)
	return outstring
}

func (c0 CoordFloat2D) IsEqual(c2 CoordFloat2D) bool {
	return c0.X == c2.X && c0.Y == c2.Y
}

/*
in essence this draws a 'rectangle' around the point and makes one check if the intersection is within it;
*/
func (c0 CoordFloat2D) IsEqual_Range(c2 CoordFloat2D, margin float64) bool {
	x0 := c0.X - margin
	y0 := c0.Y - margin
	x1 := c0.X + margin
	y1 := c0.Y + margin
	isInside0 := (c2.X > x0) && (c2.Y > y0)
	isInside1 := (c2.X < x1) && (c2.Y < y1)
	return isInside0 && isInside1
}

func (c0 CoordFloat2D) Get_Differences_As_CoordFloat2D(c2 CoordFloat2D) (differences CoordFloat2D) {
	differences.X = c2.X - c0.X
	differences.Y = c2.Y - c0.Y

	return differences
}

func (c0 CoordFloat2D) Get_Distance_Hypotenuse(c2 CoordFloat2D) (distance float64) {
	dX := c2.X - c0.X
	dY := c2.Y - c0.Y

	// if dY == 0 || dX == 0 {
	// 	if dX == 0 && dY != 0 {
	// 		return dY
	// 	} else if dX != 0 && dY == 0 {
	// 		return dY
	// 	} else {
	// 		return 0
	// 	}
	// }
	// distance = math.Abs(dY / dX)
	distance = math.Sqrt(math.Pow(dX, 2) + math.Pow(dY, 2))
	return distance
}
func (c0 CoordFloat2D) Get_Slope(c2 CoordFloat2D) (slope float64) {
	dX := c2.X - c0.X
	dY := c2.Y - c0.Y

	if dY == 0 || dX == 0 {
		if dX == 0 && dY != 0 {
			return dY
		} else if dX != 0 && dY == 0 {
			return dY
		} else {
			return 0
		}
	}
	slope = dY / dX
	return slope
}

/*
this takes a coordinate, and produces another when one enters in direction and magnitude
*/
func (c0 CoordFloat2D) Move_According_To_Vector_Radian(angle float64, magnitude float64) (c2 CoordFloat2D) {
	c2 = c0
	//hypotenuse in reverse;
	//so what x2,y2 allows
	xMov := magnitude * math.Cos(angle)
	yMov := magnitude * math.Sin(angle)
	//---do we further normalize?
	c2.X += xMov
	c2.Y += yMov
	return c2
}

/*
this takes a coordinate, and produces another when one enters in direction and magnitude
*/
func (c0 CoordFloat2D) Move_According_To_Vector_Degrees(angle float64, magnitude float64) (c2 CoordFloat2D) {
	c2 = c0
	//hypotenuse in reverse;
	//so what x2,y2 allows
	r_angle := angle * (math.Pi / 180)

	xMov := magnitude * math.Cos(r_angle)
	yMov := magnitude * math.Sin(r_angle)
	//---do we further normalize?
	c2.X += xMov
	c2.Y += yMov
	return c2
}

func CoordFloat2DTest() {
	var cf0 CoordFloat2D = MakeCoordFloat2D(5, 5)
	var cf1 CoordFloat2D = MakeCoordFloat2D(5, 5.2)
	var testval float64 = 0.233
	fmt.Printf("cf0:[%s] cf1:[%s]\tEqual?:%t\t Equal with %5.4f difference?:%t\n", cf0.ToString(), cf1.ToString(), cf0.IsEqual(cf1), testval, cf0.IsEqual_Range(cf1, testval))
	cf0 = MakeCoordFloat2D(5, 3)
	cf1 = MakeCoordFloat2D(5, 3)
	testval = 0.5
	fmt.Printf("cf0:[%s] cf1:[%s]\tEqual?:%t\t Equal with %5.4f difference?:%t\n", cf0.ToString(), cf1.ToString(), cf0.IsEqual(cf1), testval, cf0.IsEqual_Range(cf1, testval))
	fmt.Printf("cf0:[%s] cf1:[%s]\t Differences:[%s]\t Distance: %7.3f \t Slope: %7.3f \n", cf0.ToString(), cf1.ToString(), cf0.Get_Differences_As_CoordFloat2D(cf1).ToString(), cf0.Get_Distance_Hypotenuse(cf1), cf0.Get_Slope(cf1))
	cf0 = MakeCoordFloat2D(1.5, 6)
	cf1 = MakeCoordFloat2D(5.7, 3.21)
	fmt.Printf("cf0:[%s] cf1:[%s]\t Differences:[%s]\t Distance: %7.3f \t Slope: %7.3f \n", cf0.ToString(), cf1.ToString(), cf0.Get_Differences_As_CoordFloat2D(cf1).ToString(), cf0.Get_Distance_Hypotenuse(cf1), cf0.Get_Slope(cf1))

	cf0 = MakeCoordFloat2D(1.5, 6)
	cf1 = MakeCoordFloat2D(5.7, 15.21)
	fmt.Printf("cf0:[%s] cf1:[%s]\t Differences:[%s]\t Distance: %7.3f \t Slope: %7.3f \n", cf0.ToString(), cf1.ToString(), cf0.Get_Differences_As_CoordFloat2D(cf1).ToString(), cf0.Get_Distance_Hypotenuse(cf1), cf0.Get_Slope(cf1))
	cf0 = MakeCoordFloat2D(10, 10)
	cf1 = cf0.Move_According_To_Vector_Radian((math.Pi / 2), 5)
	fmt.Printf("cf0:[%s] cf1:[%s]\t Differences:[%s]\t Distance: %7.3f \t Slope: %7.3f \n", cf0.ToString(), cf1.ToString(), cf0.Get_Differences_As_CoordFloat2D(cf1).ToString(), cf0.Get_Distance_Hypotenuse(cf1), cf0.Get_Slope(cf1))
	cf0 = MakeCoordFloat2D(10, 10)
	cf1 = cf0.Move_According_To_Vector_Radian(-(math.Pi / 2), 5)
	fmt.Printf("cf0:[%s] cf1:[%s]\t Differences:[%s]\t Distance: %7.3f \t Slope: %7.3f \n", cf0.ToString(), cf1.ToString(), cf0.Get_Differences_As_CoordFloat2D(cf1).ToString(), cf0.Get_Distance_Hypotenuse(cf1), cf0.Get_Slope(cf1))
	cf1 = cf0.Move_According_To_Vector_Radian(-(math.Pi / 4), 5)
	fmt.Printf("cf0:[%s] cf1:[%s]\t Differences:[%s]\t Distance: %7.3f \t Slope: %7.3f \n", cf0.ToString(), cf1.ToString(), cf0.Get_Differences_As_CoordFloat2D(cf1).ToString(), cf0.Get_Distance_Hypotenuse(cf1), cf0.Get_Slope(cf1))
	cf1 = cf0.Move_According_To_Vector_Degrees(45, 5)
	fmt.Printf("cf0:[%s] cf1:[%s]\t Differences:[%s]\t Distance: %7.3f \t Slope: %7.3f \n", cf0.ToString(), cf1.ToString(), cf0.Get_Differences_As_CoordFloat2D(cf1).ToString(), cf0.Get_Distance_Hypotenuse(cf1), cf0.Get_Slope(cf1))

}

type CoordFloat3D struct {
	X, Y, Z float64
}

func MakeCoordFloat3D(x, y, z float64) (outCoord CoordFloat3D) {
	outCoord.X = x
	outCoord.Y = y
	outCoord.Z = z
	return outCoord
}

func (cf3D CoordFloat3D) ToString() (outstring string) {
	outstring = fmt.Sprintf("X:%5.2f, Y:%5.2f, Z:%5.2f", cf3D.X, cf3D.Y, cf3D.Z)
	return outstring
}

func (c0 CoordFloat3D) IsEqual(c2 CoordFloat3D) bool {
	return c0.X == c2.X && c0.Y == c2.Y && c0.Z == c2.Z
}
