/*
Паттерн «Посетитель» позволяет добавлять новые операции к объектам без изменения их классов, внедряя внешний класс, который реализует эти операции.
Плюсы: Разделение операций и данных; Легко добавить новые операции.
Минусы: Сложная структура, сложно добавить новые элементы.
Примеры: В компиляторах часто используется паттерн «Посетитель» для обхода синтаксического дерева и выполнения разных
действий на каждом узле
*/

package main

import "fmt"

type Element interface {
	Accept(visitor Visitor)
	GetVisitCount() int
}

type ConcreteElementA struct {
	visitCount int
}

func (e *ConcreteElementA) Accept(visitor Visitor) {
	e.visitCount++
	visitor.VisitConcreteElementA(e)
}

func (e *ConcreteElementA) GetVisitCount() int {
	return e.visitCount
}

type ConcreteElementB struct {
	visitCount int
}

func (e *ConcreteElementB) Accept(visitor Visitor) {
	e.visitCount++
	visitor.VisitConcreteElementB(e)
}

func (e *ConcreteElementB) GetVisitCount() int {
	return e.visitCount
}

type Visitor interface {
	VisitConcreteElementA(*ConcreteElementA)
	VisitConcreteElementB(*ConcreteElementB)
}

type ConcreteVisitor struct{}

func (v *ConcreteVisitor) VisitConcreteElementA(e *ConcreteElementA) {
	fmt.Println("Visiting ConcreteElementA")
}

func (v *ConcreteVisitor) VisitConcreteElementB(e *ConcreteElementB) {
	fmt.Println("Visiting ConcreteElementB")
}

func main() {
	elements := []Element{
		&ConcreteElementA{},
		&ConcreteElementB{},
	}

	visitor := &ConcreteVisitor{}

	for _, element := range elements {
		element.Accept(visitor)
		fmt.Printf("Visit count: %d\n", element.GetVisitCount())
	}
}
