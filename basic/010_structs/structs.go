package main

import "math"

type Shape interface {
	Area() float64
}

func PerimeterVal(width, height float64) float64 {
	return 2 * (width + height)
}

func AreaVal(width, height float64) float64 {
	return width * height
}

type Rectangle struct {
	Width  float64
	Height float64
}


func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}


type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

type Triangle struct {
	Base   float64
	Height float64
}

func (t Triangle) Area() float64 {
	return 0.5 * t.Base * t.Height
}
