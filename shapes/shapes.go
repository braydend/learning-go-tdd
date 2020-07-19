package shapes

import "math"

type Shape interface {
	Area() float64
	Perimiter() float64
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

func (r Rectangle) Perimiter() float64 {
	return 2 * (r.Width + r.Height)
}

func (c Circle) Perimiter() float64 {
	return (2 * math.Pi) * c.Radius
}

func (t Triangle) Perimiter() float64 {
	a := t.Height
	b := (t.Base * 0.5)
	c := math.Sqrt(math.Pow(a, 2) + math.Pow(b, 2))

	return t.Base + c + c
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (t Triangle) Area() float64 {
	return (t.Base * t.Height) * 0.5
}
