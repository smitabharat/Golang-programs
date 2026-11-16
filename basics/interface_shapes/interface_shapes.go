// Program to show how an interface lets different types share behaviour.
package main

import (
	"fmt"
	"math"
)

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

func (r Rectangle) Area() float64      { return r.Width * r.Height }
func (r Rectangle) Perimeter() float64 { return 2 * (r.Width + r.Height) }

func (c Circle) Area() float64      { return math.Pi * c.Radius * c.Radius }
func (c Circle) Perimeter() float64 { return 2 * math.Pi * c.Radius }

func main() {
	shapes := map[string]Shape{
		"Rectangle": Rectangle{Width: 10, Height: 10},
		"Circle":    Circle{Radius: 5},
	}

	for _, name := range []string{"Rectangle", "Circle"} {
		s := shapes[name]
		fmt.Printf("%s Area: %.2f, Perimeter: %.2f\n", name, s.Area(), s.Perimeter())
	}
}

// Output:
// Rectangle Area: 100.00, Perimeter: 40.00
// Circle Area: 78.54, Perimeter: 31.42
