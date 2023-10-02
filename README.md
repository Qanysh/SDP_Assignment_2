# SDP_Assignment_2
# Coffee Decorator Pattern Implementation in Go

This is a simple implementation of the Decorator design pattern in the Go programming language. The Decorator pattern allows you to add new behavior to objects dynamically without altering their structure. In this example, we have implemented a Coffee interface and two concrete decorators, MilkDecorator and ChocolateDecorator, that enhance the behavior of a simple coffee.

## Overview

This implementation consists of the following key components:

- `Coffee` Interface: Defines the methods that concrete coffee components must implement.
- `SimpleCoffee` Component: A concrete coffee component that implements the `Coffee` interface.
- `CoffeeDecorator` Interface: Defines the methods for coffee decorators.
- `MilkDecorator` and `ChocolateDecorator`: Concrete decorators that implement the `CoffeeDecorator` interface to add functionality to coffee components.


# Factory Pattern Implementation in Go

This is a simple implementation of the Factory design pattern in the Go programming language. The Factory pattern is a creational design pattern that provides an interface for creating objects but allows subclasses to alter the type of objects that will be created. In this example, we have implemented two concrete factories, each of which creates a specific type of product.

## Overview

This implementation consists of the following key components:

- `Product` Interface: Defines the methods that concrete products must implement.
- `ConcreteProduct1` and `ConcreteProduct2`: Concrete product classes that implement the `Product` interface.
- `Factory` Interface: Defines the method for creating products.
- `ConcreteFactory1` and `ConcreteFactory2`: Concrete factory classes that implement the `Factory` interface and create specific products.

## Prerequisites

To run this implementation, you need to have Go (Golang) installed on your system. You can download and install Go from the official [Go website](https://golang.org/dl/).

## Usage

1. Clone this repository to your local machine:

   ```bash
   git clone https://github.com/Qanysh/SDP_Assignment_2

