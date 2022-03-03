//Разработать программу нахождения расстояния между двумя точками,
//которые представлены в виде структуры Point с инкапсулированными параметрами x,y и конструктором
package main

import (
	"fmt"
	"math"
)

type point struct {
	x float64
	y float64
}

func newPoint(x float64, y float64) *point {
	return &point{
		x: x,
		y: y,
	}
}

func distance(point1 *point, point2 *point) float64 {
	return math.Sqrt(math.Pow((point1.x-point2.x), 2) + math.Pow((point1.y-point2.y), 2))
}

func main() {
	a := newPoint(0, 0)
	b := newPoint(10, 10)
	fmt.Println(distance(a, b))
}
