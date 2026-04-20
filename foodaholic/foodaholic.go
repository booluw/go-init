package main

import (
	"bufio"
	"fmt"
	"foodaholic/db"
	"os"
	"text/tabwriter"
)

type Order struct {
	CustomerName string
	Order [][]string
	TotalBill int
	Tip int
}

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
var order = [][]string{}
var bill = Order{}

func main() {
	fmt.Println("Welcome to FoodAholic, where your taste bugs jingle.")
	fmt.Println("What should we call you?")
	fmt.Scanln(&bill.CustomerName)

	fmt.Printf("What would you like to eat %v? \n", bill.CustomerName)

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

func (b Order) showOrder() {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, '.', tabwriter.AlignRight|tabwriter.Debug)
	fmt.Fprintf(w, "Orders Price")
	fmt.Fprintln(w, "a\tb\tc")
	for i := range bill.Order {
		fmt.Println(bill.Order[i][0], bill.Order[i][1])
	}
	fmt.Println("====================== TOTAL: ", b.TotalBill, "=====")
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
		bill.showOrder()
	case "P":
		bill.payForOrder()
	case "M":
		showMenu()
	default:
		for i := range menu {
			if menu[i].abb == input {
				bill.addOrderToBill(menu[i])
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

func (b *Order) payForOrder() {
	fmt.Printf("Any tip?...")
	fmt.Scanln(&b.Tip)

	for i := range b.Order {
		db.WriteToDB(b.Order[i], b.CustomerName)
	}

	db.WriteToDB([]string{fmt.Sprintf("Tip (%v)", b.Tip), fmt.Sprint(b.TotalBill+b.Tip)}, b.CustomerName)
	fmt.Println("Order saved")
}

func (b *Order) addOrderToBill (or Menu) {
	b.Order = append(b.Order, []string{fmt.Sprintf("%v", or.name), fmt.Sprintf("%v", or.price)})
	b.TotalBill += or.price
}
