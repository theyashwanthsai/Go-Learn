package structs

import "math"



type Shape interface{
	Area() float64
}



type Rectangle struct{
	Width float64
	Height float64
}

func (rectangle Rectangle) Perimeter() float64{
	return 2*(rectangle.Width+rectangle.Height)
}

func (rectangle Rectangle) Area() float64{
	return rectangle.Height*rectangle.Width
}

type Circle struct{
	Radius float64
}

func (circle Circle) Area() float64{
	return math.Pi*circle.Radius*circle.Radius
}

type Triangle struct{
	Base float64
	Height float64
}

func (t Triangle) Area() float64{
	return (0.5)*t.Base*t.Height
}