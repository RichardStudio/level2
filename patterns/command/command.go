/*
Паттерн «Команда» инкапсулирует запрос как объект, позволяя параметризировать объекты с различными запросами и
поддерживать операции, такие как отмена, очередь и логирование.
Плюсы: Разделение обязанностей; Легко добавить новые команды, можно управлять очередью команд.
Минусы: Усложняет структуру проекта, много классов
Примеры: Система дистанционного управления для электронных устройств в умном доме.
*/

package main

import "fmt"

type Command interface {
	Execute()
}

type ConcreteCommandA struct {
	receiver *Receiver
}

func (c *ConcreteCommandA) Execute() {
	c.receiver.ActionA()
}

type ConcreteCommandB struct {
	receiver *Receiver
}

func (c *ConcreteCommandB) Execute() {
	c.receiver.ActionB()
}

type Receiver struct{}

func (r *Receiver) ActionA() {
	fmt.Println("Receiver: Action A")
}

func (r *Receiver) ActionB() {
	fmt.Println("Receiver: Action B")
}

type Invoker struct {
	commands []Command
}

func (i *Invoker) AddCommand(cmd Command) {
	i.commands = append(i.commands, cmd)
}

func (i *Invoker) ExecuteCommands() {
	for _, cmd := range i.commands {
		cmd.Execute()
	}
}

func main() {
	receiver := &Receiver{}

	commandA := &ConcreteCommandA{receiver}
	commandB := &ConcreteCommandB{receiver}

	invoker := &Invoker{}
	invoker.AddCommand(commandA)
	invoker.AddCommand(commandB)

	invoker.ExecuteCommands()
}
