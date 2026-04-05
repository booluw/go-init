package main

import "fmt"

func main () {
	ice_creams := map[string]int{"Gelato": 200, "Vanilla": 400, "Strawberry": 300, "Dark Chocolate": 300 }

	fmt.Println("Ice cream Menu")

	for key, value := range ice_creams {
		fmt.Printf("Flavour: %v | Price: %v \n", key, value)
	}
}