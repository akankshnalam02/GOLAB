package main
import (
	"fmt"
	"math"
)
type Shape interface {
	Area() float64
	Perimeter() float64
}
type Circle struct {
	Radius float64
}
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}
type Rectangle struct {
	Width, Height float64
}
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}
func main() {
	var s Shape
	c := Circle{Radius: 5}
	r := Rectangle{Width: 3, Height: 4}
	s = c
	fmt.Printf("Area of circle: %.2f\n", s.Area())
	fmt.Printf("Perimeter of circle: %.2f\n\n", s.Perimeter())
	s = r
	fmt.Printf("Area of rectangle: %.2f\n", s.Area())
	fmt.Printf("Perimeter of rectangle: %.2f\n\n", s.Perimeter())
}




