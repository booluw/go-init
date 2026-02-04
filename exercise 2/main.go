package main

import "fmt"

func main () {

	var ( a uint8; b int8 )

	a = 255
	b = 127

	fmt.Println(a, b)

	// a, b, c, d := 747, 911, 90210, 0220

	// fmt.Printf("%v is \t %b, \t %#x \n", a, a, a)
	// fmt.Printf("%v is \t %b, \t %#x \n", b, b, b)
	// fmt.Printf("%v is %b, %#x \n", c, c, c)
	// fmt.Printf("%v is %b, %#x \n", d, d, d)
}