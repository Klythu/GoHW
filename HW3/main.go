package main

import (
	"fmt"
	"math"
)

type Shape struct {
	Name string
}

func (s Shape) GetName() string {
	return s.Name
}

func (s Shape) Area() float64 {
	return 0.0
}

type Circle struct {
	Shape
	radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

type Rectangle struct {
	Shape
	length float64
	height float64
}

func (r Rectangle) Area() float64 {
	return r.height * r.length
}
func main() {
	circle := Circle{Shape{"Круг"}, 5}
	rectangle := Rectangle{Shape{"Прямоугольник"}, 4, 6}

	fmt.Printf("%s: Площадь = %.2f\n", circle.GetName(), circle.Area())
	fmt.Printf("%s: Площадь = %.2f\n", rectangle.GetName(), rectangle.Area())
}
