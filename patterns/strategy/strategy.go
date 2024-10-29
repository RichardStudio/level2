/*
Позволяет выбирать алгоритмы "На ходу".
Плюсы: Легко выбирать алгоритмы; Легко добавлять новые стратегии
Минусы: Большое число классов
Пример: В приложениях для сжатия данных можно использовать паттерн «Стратегия» для выбора алгоритма сжатия
*/
package main

import "fmt"

type Strategy interface {
	Execute(a, b int) int
}

type AddStrategy struct{}

func (s *AddStrategy) Execute(a, b int) int {
	return a + b
}

type SubtractStrategy struct{}

func (s *SubtractStrategy) Execute(a, b int) int {
	return a - b
}

type StratContext struct {
	strategy Strategy
}

func (c *StratContext) SetStrategy(strategy Strategy) {
	c.strategy = strategy
}

func (c *StratContext) ExecuteStrategy(a, b int) int {
	return c.strategy.Execute(a, b)
}

func main() {
	context := &StratContext{}

	addStrategy := &AddStrategy{}
	context.SetStrategy(addStrategy)
	fmt.Println("Result of addition:", context.ExecuteStrategy(10, 5))

	subtractStrategy := &SubtractStrategy{}
	context.SetStrategy(subtractStrategy)
	fmt.Println("Result of subtraction:", context.ExecuteStrategy(10, 5))
}
