/*
Паттерн «Цепочка вызовов» позволяет передавать запросы по цепочке обработчиков, пока один из них не обработает
запрос.
Плюсы: Убирает зависимости между отправителем и получателем запроса; Легко добавлять новые обработчики или изменять
цепочку; Легко добавлять новые типы обработчиков без изменения существующего кода.
Минусы: Запрос может быть не обработан; Сложность отладки; Длинна цепочки влияет на производительность.
Примеры: В веб-серверах, где запросы HTTP проходят через несколько уровней фильтров и обработчиков
*/
package main

import "fmt"

type Handler interface {
	SetNext(handler Handler)
	Handle(request string)
}

type BaseHandler struct {
	next Handler
}

func (h *BaseHandler) SetNext(handler Handler) {
	h.next = handler
}

func (h *BaseHandler) Handle(request string) {
	if h.next != nil {
		h.next.Handle(request)
	}
}

type ConcreteHandlerA struct {
	BaseHandler
}

func (h *ConcreteHandlerA) Handle(request string) {
	if request == "A" {
		fmt.Println("ConcreteHandlerA handled request:", request)
	} else {
		h.BaseHandler.Handle(request)
	}
}

type ConcreteHandlerB struct {
	BaseHandler
}

func (h *ConcreteHandlerB) Handle(request string) {
	if request == "B" {
		fmt.Println("ConcreteHandlerB handled request:", request)
	} else {
		h.BaseHandler.Handle(request)
	}
}

type ConcreteHandlerC struct {
	BaseHandler
}

func (h *ConcreteHandlerC) Handle(request string) {
	if request == "C" {
		fmt.Println("ConcreteHandlerC handled request:", request)
	} else {
		h.BaseHandler.Handle(request)
	}
}

func main() {
	handlerA := &ConcreteHandlerA{}
	handlerB := &ConcreteHandlerB{}
	handlerC := &ConcreteHandlerC{}

	handlerA.SetNext(handlerB)
	handlerB.SetNext(handlerC)

	requests := []string{"A", "B", "C", "D"}

	for _, req := range requests {
		fmt.Println("Sending request:", req)
		handlerA.Handle(req)
	}
}
