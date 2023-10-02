package main

import (
	"fmt"
)

type Coffee interface {
	GetCost() int
	GetDescription() string
}

type SimpleCoffee struct{}

func (c *SimpleCoffee) GetCost() int {
	return 10
}

func (c *SimpleCoffee) GetDescription() string {
	return "Простой кофе"
}

type CoffeeDecorator interface {
	GetCost() int
	GetDescription() string
}

type MilkDecorator struct {
	Coffee Coffee
}

func (m *MilkDecorator) GetCost() int {
	return m.Coffee.GetCost() + 5
}

func (m *MilkDecorator) GetDescription() string {
	return m.Coffee.GetDescription() + ", с молоком"
}

type ChocolateDecorator struct {
	Coffee Coffee
}

func (c *ChocolateDecorator) GetCost() int {
	return c.Coffee.GetCost() + 7
}

func (c *ChocolateDecorator) GetDescription() string {
	return c.Coffee.GetDescription() + ", с шоколадом"
}

func main() {
	coffee := &SimpleCoffee{}
	fmt.Printf("Описание: %s, Стоимость: %d $\n", coffee.GetDescription(), coffee.GetCost())

	coffeeWithMilk := &MilkDecorator{Coffee: coffee}
	fmt.Printf("Описание: %s, Стоимость: %d $\n", coffeeWithMilk.GetDescription(), coffeeWithMilk.GetCost())

	coffeeWithChocolate := &ChocolateDecorator{Coffee: coffee}
	fmt.Printf("Описание: %s, Стоимость: %d $\n", coffeeWithChocolate.GetDescription(), coffeeWithChocolate.GetCost())

	coffeeWithMilkAndChocolate := &MilkDecorator{Coffee: coffeeWithChocolate}
	fmt.Printf("Описание: %s, Стоимость: %d $\n", coffeeWithMilkAndChocolate.GetDescription(), coffeeWithMilkAndChocolate.GetCost())
}
