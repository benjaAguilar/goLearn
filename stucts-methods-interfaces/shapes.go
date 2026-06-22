package stuctsmethodsinterfaces

import "math"

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radio float64
}

type Triangle struct {
	Base   float64
	Height float64
}

type Shape interface {
	Area() float64
}

func Perimeter(rect Rectangle) float64 {
	return (rect.Width + rect.Height) * 2
}

func (rect Rectangle) Area() float64 {
	return rect.Width + rect.Height
}

func (circle Circle) Area() float64 {
	return math.Pi * (circle.Radio * circle.Radio)
}

func (t Triangle) Area() float64 {
	return 0.5 * (t.Base * t.Height)
}
