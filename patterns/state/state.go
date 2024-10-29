/*
Позволяет управлять поведением объекта в зависимости от его состояния
Плюсы: Меняет поведение объекта "на ходу"; Позволяет изолировать состояния; Легко добавить новые состояния
Минусы: Много классов
Пример: В автоматах по продаже товаров, где состояние автомата
*/
package main

import "fmt"

type Context struct {
	state State
}

func (c *Context) SetState(state State) {
	c.state = state
}

func (c *Context) Request() {
	c.state.Handle(c)
}

type State interface {
	Handle(context *Context)
}

type ConcreteStateA struct{}

func (s *ConcreteStateA) Handle(context *Context) {
	fmt.Println("ConcreteStateA handling request and transitioning to ConcreteStateB")
	context.SetState(&ConcreteStateB{})
}

type ConcreteStateB struct{}

func (s *ConcreteStateB) Handle(context *Context) {
	fmt.Println("ConcreteStateB handling request and transitioning to ConcreteStateA")
	context.SetState(&ConcreteStateA{})
}

func main() {
	context := &Context{state: &ConcreteStateA{}}

	context.Request()
	context.Request()
	context.Request()
}
