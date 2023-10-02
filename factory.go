package main

import (
	"fmt"
)

type Product interface {
	GetName() string
}

type ConcreteProduct1 struct{}

func (p *ConcreteProduct1) GetName() string {
	return "Product 1"
}

type ConcreteProduct2 struct{}

func (p *ConcreteProduct2) GetName() string {
	return "Product 2"
}

type Factory interface {
	CreateProduct() Product
}

type ConcreteFactory1 struct{}

func (f *ConcreteFactory1) CreateProduct() Product {
	return &ConcreteProduct1{}
}

type ConcreteFactory2 struct{}

func (f *ConcreteFactory2) CreateProduct() Product {
	return &ConcreteProduct2{}
}

func main() {
	factory1 := &ConcreteFactory1{}
	product1 := factory1.CreateProduct()
	fmt.Printf("Created %s\n", product1.GetName())

	factory2 := &ConcreteFactory2{}
	product2 := factory2.CreateProduct()
	fmt.Printf("Created %s\n", product2.GetName())
}
