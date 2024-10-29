/*
Паттерн «Строитель» позволяет пошагово создавать сложные объекты, разделяя процесс создания на несколько этапов.
Плюсы: Позволяет создавать различные вариации объекта; Контроль над процессом создания; Разделение на шаги
Минусы: Сложность кода из-за большого кол-ва объектов и методов.
Примеры: Система конфигурации автомобиля с разными компонентами, генерация http запросов
*/

package main

import "fmt"

type Product struct {
	partA string
	partB string
	partC string
}

type Builder struct {
	product *Product
}

func NewBuilder() *Builder {
	return &Builder{&Product{}}
}

func (b *Builder) BuildPartA(partA string) {
	b.product.partA = partA
}

func (b *Builder) BuildPartB(partB string) {
	b.product.partB = partB
}

func (b *Builder) BuildPartC(partC string) {
	b.product.partC = partC
}

func (b *Builder) GetProduct() *Product {
	return b.product
}

type Director struct {
	builder *Builder
}

func NewDirector(builder *Builder) *Director {
	return &Director{builder}
}

func (d *Director) Construct() {
	d.builder.BuildPartA("Custom PartA")
	d.builder.BuildPartB("Custom PartB")
	d.builder.BuildPartC("Custom PartC")
}

func main() {
	builder := NewBuilder()
	director := NewDirector(builder)
	director.Construct()
	product := builder.GetProduct()
	fmt.Printf("Product parts: %s, %s, %s\n", product.partA, product.partB, product.partC)
}
