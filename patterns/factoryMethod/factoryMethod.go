/*
Паттерн "Фабричный метод" предоставляет интерфейс для создания объектов, но позволяет подклассам изменять тип
создаваемого объекта.
Плюсы: Легко добавить новые типы; Создание объектов отдельным классам
Минусы: Много подклассов, усложняет код
Пример: В системах управления базами данных, где разные базы данных могут требовать разные драйверы,
фабричный метод может создавать соответствующий драйвер в зависимости от выбранной базы данных.
*/
package main

import "fmt"

type Product interface {
	Use() string
}

type ConcreteProductA struct{}

func (p *ConcreteProductA) Use() string {
	return "Using ConcreteProductA"
}

type ConcreteProductB struct{}

func (p *ConcreteProductB) Use() string {
	return "Using ConcreteProductB"
}

type Creator interface {
	FactoryMethod() Product
}

type ConcreteCreatorA struct{}

func (c *ConcreteCreatorA) FactoryMethod() Product {
	return &ConcreteProductA{}
}

type ConcreteCreatorB struct{}

func (c *ConcreteCreatorB) FactoryMethod() Product {
	return &ConcreteProductB{}
}

func main() {
	var creator Creator

	creator = &ConcreteCreatorA{}
	product := creator.FactoryMethod()
	fmt.Println(product.Use())

	creator = &ConcreteCreatorB{}
	product = creator.FactoryMethod()
	fmt.Println(product.Use())
}
