package main

import "fmt"

type House struct {
	Floor string
	Wall  string
	Roof  string
}

type HouseBuilder interface {
	BuildFloor()
	BuildWalls()
	BuildRoof()
}

type CivilBuilder struct {
	house House
}

func NewCivilBuilder() *CivilBuilder {
	return &CivilBuilder{
		house: House{},
	}
}

func (c *CivilBuilder) BuildFloor() {
	fmt.Println("Floor builded")
}

func (c *CivilBuilder) BuildWalls() {
	fmt.Println("Walls builded")
}

func (c *CivilBuilder) BuildRoof() {
	fmt.Println("Roof builded")
}

type Director struct {
	builder HouseBuilder
}

func NewDirector(bldr HouseBuilder) *Director {
	return &Director{
		builder: bldr,
	}
}

func (d *Director) BuildHouse() {
	d.builder.BuildFloor()
	d.builder.BuildWalls()
	d.builder.BuildRoof()
}

func main() {
	cb := NewCivilBuilder()
	d := NewDirector(cb)
	d.BuildHouse()
}