package main

import (
	"fmt"
	"strconv"
)

type IProduct interface {
	setStock(stock int)
	getStock() int
	setName(name string)
	getName() string
}

type Computer struct {
	name  string
	stock int
}

func (c *Computer) setStock(stock int) {
	c.stock = stock
}

func (c *Computer) getStock() int {
	return c.stock
}

func (c *Computer) setName(name string) {
	c.name = name
}

func (c *Computer) getName() string {
	return c.name
}

type Laptop struct {
	Computer
}

type DeskTop struct {
	Computer
}

func newLaptop() IProduct {
	return &Laptop{
		Computer: Computer{
			name:  "Laptop computer",
			stock: 25,
		},
	}
}

func newDesktop() IProduct {
	return &Laptop{
		Computer: Computer{
			name:  "Desktop computer",
			stock: 30,
		},
	}
}

func getComputerFactory(typeComputer string) (IProduct, error) {
	if typeComputer == "laptop" {
		return newLaptop(), nil
	}

	if typeComputer == "desktop" {
		return newDesktop(), nil
	}

	return nil, fmt.Errorf("Computer type incorrect")
}

func printNameAndStock(p IProduct) {
	fmt.Println("Product name: " + p.getName() + ", with stock " + strconv.Itoa(p.getStock()) + "")
}

func main() {

	laptop, _ := getComputerFactory("laptop")
	desktop, _ := getComputerFactory("desktop")

	printNameAndStock(laptop)
	printNameAndStock(desktop)

}
