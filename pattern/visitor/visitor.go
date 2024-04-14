package main

import "fmt"

type Shape interface {
	accept(Visitor)
}

type Circle struct {
	Radius float64
}

func (c *Circle) accept(v Visitor) {
	v.visitCircle(c)
}

type Square struct {
	SideLength float64
}

func (s *Square) accept(v Visitor) {
	v.visitSquare(s)
}

type Triangle struct {
	Base   float64
	Height float64
}

func (t *Triangle) accept(v Visitor) {
	v.visitTriangle(t)
}

type Visitor interface {
	visitCircle(*Circle)
	visitSquare(*Square)
	visitTriangle(*Triangle)
}

type AreaCalculator struct {
	TotalArea float64
}

func (a *AreaCalculator) visitCircle(c *Circle) {
	area := 3.14 * c.Radius * c.Radius
	fmt.Printf("Circle area: %.2f\n", area)
	a.TotalArea += area
}

func (a *AreaCalculator) visitSquare(s *Square) {
	area := s.SideLength * s.SideLength
	fmt.Printf("Square area: %.2f\n", area)
	a.TotalArea += area
}

func (a *AreaCalculator) visitTriangle(t *Triangle) {
	area := 0.5 * t.Base * t.Height
	fmt.Printf("Triangle area: %.2f\n", area)
	a.TotalArea += area
}

func main() {
	shapes := []Shape{
		&Circle{Radius: 5},
		&Square{SideLength: 4},
		&Triangle{Base: 3, Height: 6},
	}

	areaVisitor := &AreaCalculator{}

	for _, shape := range shapes {
		shape.accept(areaVisitor)
	}

	fmt.Printf("Total area: %.2f\n", areaVisitor.TotalArea)
}

// Паттерн посетитель - это поведенческий паттерн проектирования,
// который позволяет добавлять новые операции к объектам без изменения их классов.
// Он отделяет алгоритмы от структур данных, на которых они оперируют.