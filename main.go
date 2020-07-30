package main

import (
	"fmt"
)

type Shape interface {
	CalcScale() float64
}

type Rect struct {
	width  float64
	height float64
}

type Circle struct {
	radius float64
}

type Triangle struct {
	height float64
	bottom float64
}

type Trapezoid struct {
	 height float64
	 top    float64
	 bottom float64
}

func (r Rect) CalcScale() float64 {
	return r.width * r.height
}

func (c Circle) CalcScale() float64 {
	return 3.14 * c.radius * c.radius
}

func (t Triangle) CalcScale() float64 {
	return t.height * t.bottom / 2.0
}

func (t Trapezoid) CalcScale() float64 {
	return t.height * (t.bottom＋t.top)
}
func main() {
	shape := []Shape{
		Rect{width: 100, height: 20},
		Rect{width: 200, height: 40},
		Rect{width: 300, height: 80},
		Circle{radius: 10},
		Circle{radius: 20},
		Circle{radius: 30},
		Triangle{height: 10, bottom: 20},
		Triangle{height: 10, bottom: 40},
		Triangle{height: 10, bottom: 50},
		Trapezoid{height: 10, bottom: 10,top: 10},
		Trapezoid{}
	}

	for _, s := range shape {
		scale := s.CalcScale()
		fmt.Printf("%f\n", scale)
	}
}	