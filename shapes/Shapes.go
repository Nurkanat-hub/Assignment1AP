package shapes

import (
	"fmt"
	"math"
)

// Shape
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Rectangle
type Rectangle struct {
	Length float64
	Width  float64
}

func (r Rectangle) Area() float64 {
	return r.Length * r.Width
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Length + r.Width)
}

// Circle
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func PrintShapeDetails(s Shape) {
	fmt.Println("Shape details:")
	fmt.Printf("Area: %.2f\n", s.Area())
	fmt.Printf("Perimeter: %.2f\n", s.Perimeter())
	fmt.Println("-----------------------------")
}

func main() {
	shapes := []Shape{
		Rectangle{Length: 10, Width: 5},
		Circle{Radius: 7},
	}

	for _, shape := range shapes {
		PrintShapeDetails(shape)
	}
}
