package main

import "fmt"

// Product
type Transport interface {
	Deliver()
}

// ConcreteProduct
type Truck struct{}

func (t *Truck) Deliver() {
	fmt.Println("Доставка груза грузовиком")
}

// ConcreteProduct
type Ship struct{}

func (s *Ship) Deliver() {
	fmt.Println("Доставка груза кораблем")
}

// ConcreteProduct
type Airplane struct{}

func (a *Airplane) Deliver() {
	fmt.Println("Доставка груза самолетом")
}

// Creator
type Logistics interface {
	CreateTransport() Transport
}

// ConcreteCreator
type RoadLogistics struct{}

func (r *RoadLogistics) CreateTransport() Transport {
	return &Truck{}
}

// ConcreteCreator
type SeaLogistics struct{}

func (s *SeaLogistics) CreateTransport() Transport {
	return &Ship{}
}

// ConcreteCreator
type AirLogistics struct{}

func (a *AirLogistics) CreateTransport() Transport {
	return &Airplane{}
}

func main() {
	roadLogistics := &RoadLogistics{}
	truck := roadLogistics.CreateTransport()
	truck.Deliver()

	seaLogistics := &SeaLogistics{}
	ship := seaLogistics.CreateTransport()
	ship.Deliver()

	airLogistics := &AirLogistics{}
	airplane := airLogistics.CreateTransport()
	airplane.Deliver()
}

// Паттерн фабричный метод - это порождающий паттерн проектирования,
// который определяет интерфейс для создания объектов,
// но оставляет выбор конкретного класса создаваемого объекта производным классам. 
// Таким образом, он делегирует создание экземпляров подклассам.