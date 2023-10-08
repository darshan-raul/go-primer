package structmethod

import (
	"math"
)

type Shape interface{

	Area() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64

}
type Triangle struct {
	Base   float64
	Height float64
}


func (r Rectangle) Area() float64{
	area := r.Height*r.Width
	return area
}

func (c Circle) Area() float64{

	area := math.Pi * c.Radius * c.Radius
	return area
}


func (t Triangle) Area() float64 {
	area := (t.Base*t.Height)*0.5
	return area
}

