package main

import "fmt"

type Shape interface {
	Area() float64
}

type Rectangle struct {
	width, height float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.radius * c.radius
}

func calculateArea(s Shape) float64 {
	return s.Area()
}

func main () {

	rect := Rectangle{width: 10, height: 20}
	circle := Circle{radius: 5}

	fmt.Println("Area of rectangle is", calculateArea(rect))
	fmt.Println("Area of circle is", calculateArea(circle))

}