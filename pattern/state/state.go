package main

import "fmt"

// State
type TrafficLightState interface {
	TransitionToNextState(light *TrafficLight)
}

// ConcreteState
type RedLight struct{}

func (r *RedLight) TransitionToNextState(light *TrafficLight) {
	fmt.Println("Светофор переключается с красного на желтый")
	light.State = &YellowLight{}
}

// ConcreteState
type YellowLight struct{}

func (y *YellowLight) TransitionToNextState(light *TrafficLight) {
	fmt.Println("Светофор переключается с желтого на зеленый")
	light.State = &GreenLight{}
}

// ConcreteState
type GreenLight struct{}

func (g *GreenLight) TransitionToNextState(light *TrafficLight) {
	fmt.Println("Светофор переключается с зеленого на красный")
	light.State = &RedLight{}
}

// Context
type TrafficLight struct {
	State TrafficLightState
}

func (t *TrafficLight) TransitionToNextState() {
	t.State.TransitionToNextState(t)
}

func main() {
	trafficLight := &TrafficLight{State: &RedLight{}}

	for i := 0; i < 3; i++ {
		trafficLight.TransitionToNextState()
	}
}


// Паттерн cостояние - это поведенческий паттерн проектирования,
// который позволяет объекту изменять свое поведение в зависимости от своего внутреннего состояния. 
// При этом он выглядит так, будто объект меняет свой класс.