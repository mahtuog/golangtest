package structs

import "math"

type Shape interface {
	Area() float64
}

type Rectangle struct {
	width float64
	height float64
}

type Triangle struct {
	base float64
	height float64
}

type Circle struct {
	radius float64
}

func (r Rectangle) Area() (area float64){
	area = r.width * r.height
	return
}

func (r Rectangle) Perimeter() (perimeter float64){
	perimeter = 2 * (r.width + r.height)
	return
}

func (c Circle) Area() (area float64){
	area = math.Pi * (c.radius * c.radius)
	return
}

func (c Circle) Perimeter() (perimeter float64){
	perimeter = 2 * math.Pi * c.radius
	return
}

func (t Triangle) Area() (area float64){
	return 0.5 * t.base * t.height
}