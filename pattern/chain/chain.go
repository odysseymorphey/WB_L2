package main

import "fmt"

// Handler
type Approver interface {
	SetNext(Approver)
	ProcessRequest(int)
}

// ConcreteHandler
type DepartmentHead struct {
	next Approver
}

func (d *DepartmentHead) SetNext(next Approver) {
	d.next = next
}

func (d *DepartmentHead) ProcessRequest(days int) {
	if days <= 5 {
		fmt.Println("Заявка одобрена отделом")
	} else if d.next != nil {
		fmt.Println("Заявка передана руководителю подразделения")
		d.next.ProcessRequest(days)
	} else {
		fmt.Println("Заявка отклонена")
	}
}

// ConcreteHandler
type DivisionHead struct {
	next Approver
}

func (d *DivisionHead) SetNext(next Approver) {
	d.next = next
}

func (d *DivisionHead) ProcessRequest(days int) {
	if days <= 10 {
		fmt.Println("Заявка одобрена подразделением")
	} else if d.next != nil {
		fmt.Println("Заявка передана директору")
		d.next.ProcessRequest(days)
	} else {
		fmt.Println("Заявка отклонена")
	}
}

// ConcreteHandler
type Director struct{}

func (d *Director) SetNext(next Approver) {
	// Director is the last in the chain, so no next handler to set
}

func (d *Director) ProcessRequest(days int) {
	if days <= 15 {
		fmt.Println("Заявка одобрена директором")
	} else {
		fmt.Println("Заявка отклонена")
	}
}

func main() {
	director := &Director{}
	divisionHead := &DivisionHead{}
	departmentHead := &DepartmentHead{}

	departmentHead.SetNext(divisionHead)
	divisionHead.SetNext(director)

	// Подаем заявку с разным количеством дней
	departmentHead.ProcessRequest(3)
	departmentHead.ProcessRequest(7)
	departmentHead.ProcessRequest(12)
	departmentHead.ProcessRequest(20)
}


// Паттерн цепочка обязанностей - это поведенческий паттерн проектирования,
// который позволяет передавать запросы последовательно по цепочке обработчиков.
// Запрос проходит через каждый обработчик, пока один из них не обработает запрос