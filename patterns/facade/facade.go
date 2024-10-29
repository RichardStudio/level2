/*
Паттерн «Фасад» упрощает взаимодействие со сложной подсистемой, предоставляя единый интерфейс для работы с ней.
Плюсы: упрощает работу с системой; упрощает понимание системы; дает ровно ту функциональность, которая требуется,
скрывая ненужное; дает возможность модифицировать систему не влияя на пользователей фасада.
Минусы: Дополнительный уровень абстракции, ограниченный функционал.
Примеры: GUI приложения
*/

package main

import (
	"fmt"
)

type SubsystemA struct{}

func (s *SubsystemA) OperationA() string {
	return "SubsystemA: Operation A"
}

type SubsystemB struct{}

func (s *SubsystemB) OperationB() string {
	return "SubsystemB: Operation B"
}

type SubsystemC struct{}

func (s *SubsystemC) OperationC() string {
	return "SubsystemC: Operation C"
}

type Facade struct {
	subsystemA *SubsystemA
	subsystemB *SubsystemB
	subsystemC *SubsystemC
}

func NewFacade() *Facade {
	return &Facade{
		subsystemA: &SubsystemA{},
		subsystemB: &SubsystemB{},
		subsystemC: &SubsystemC{},
	}
}

func (f *Facade) Operation() string {
	return fmt.Sprintf("%s\n%s\n%s",
		f.subsystemA.OperationA(),
		f.subsystemB.OperationB(),
		f.subsystemC.OperationC())
}

func main() {
	facade := NewFacade()
	fmt.Println(facade.Operation())
}
