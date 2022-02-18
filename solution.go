package square

import "math"

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)
//подобрал формулу  для треугольника) там сторона в степень 2 не возводится. Итоговая формула - (сторона * корень(3))/4
//a*a*math.Sqrt(3)/4
//короч с окружностью просто на 2 разделить
//а квадрат это а*a

type myType int

const SidesTriangle myType = 3
const SidesSquare myType = 4
const SidesCircle myType = 0

func CalcSquare(sideLen float64, sidesNum myType) (result float64) {
	switch sidesNum {
	case 3:
		result = (math.Pow(sideLen, 2.0) * math.Sqrt(3.0)) / 4
	case 4:
		result = math.Pow(sideLen, 2.0)
	case 0:
		result = math.Pow(sideLen, 2.0) * math.Pi
	}
	return
}
