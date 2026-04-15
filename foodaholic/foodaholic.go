package main

import (
	"bufio"
	"fmt"
	"foodaholic/db"
	"os"
)

type Menu struct {
	name  string
	price int
	abb   string
}

var scanner = bufio.NewScanner(os.Stdin)
var menu = []Menu{
	{name: "Salad", price: 20, abb: "Salad"},
	{name: "Jollof Rice", price: 50, abb: "JR"},
	{name: "White Rice", price: 30, abb: "WR"},
	{name: "Coke", price: 10, abb: "Coke"},
	{name: "Fanta", price: 8, abb: "Fanta"},
}
var order = []string{}

func main() {
	fmt.Println("Welcome to FoodAholic, where your taste bugs jingle")
	fmt.Println("What would you like to eat?")

	showMenu()
}

func showMenu() {
	fmt.Println("==================")
	for i := range menu {
		fmt.Printf("> %v(%v) ---- $%v \n", menu[i].name, menu[i].abb, menu[i].price)
	}

	input := acceptInput()
	processInput(input)
}

func showOrder() {
	fmt.Println("=================")
	fmt.Println("ORDER")
	fmt.Println("=================")
	fmt.Println()
	for i := range order {
		fmt.Println(order[i])
	}
	fmt.Println("=================")
	showOptions()

	input := acceptInput()
	processInput(input)
}

func showOptions() {
	fmt.Println("Esc: End \t M: Menu \t O: Show Order \t P: Pay")
}

func acceptInput() string {
	var input string

	fmt.Scanln(&input)
	return input
}

func processInput(input string) {
	switch input {
	case "Esc":
		return
	case "O":
		showOrder()
	case "P":
		payForOrder()
	case "M":
		showMenu()
	default:
		for i := range menu {
			if menu[i].abb == input {
				order = append(order, fmt.Sprintf("%v, %v", menu[i].name, menu[i].price))
				showMenu()
				return
			}
		}

		fmt.Println("Wrong Order")
		showOptions()

		val := acceptInput()
		processInput(val)
	}
}

func payForOrder() {
	db.WriteToDB(order)

	fmt.Println("Order saved")
}
