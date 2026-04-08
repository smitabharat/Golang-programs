package main

import "fmt"

//interface example

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	Width, Height float64
}

type Circle struct {
	Radius float64
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (c Circle) Perimeter() float64 {
	return 2 * 3.14 * c.Radius
}

func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

func main() {
	Rect := Rectangle{Height: 10, Width: 10}
	Circ := Circle{Radius: 5}

	var shape Shape

	shape = Rect
	fmt.Printf("Rectangle Area: %.2f, Perimeter: %.2f\n",shape.Area(),shape.Perimeter())

	 shape = Circ
	fmt.Printf("Circle Area: %.2f, Perimeter: %.2f\n",shape.Area(),shape.Perimeter())
}
