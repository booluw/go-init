package main

import "fmt"

/*
	Variadic functions, methods, interfaces
*/

type Person struct {
	firstName string
	lastName string
	words string
}

type Animal struct {
	group string
	sound string
}

type Mamals interface {
	speak()
}

func main () {
	sliceOfInt :=[]int{1, 10, 99, -10}

	total := sum(sliceOfInt...)
	fmt.Println(total)

	dog := Animal {
		group: "dog",
		sound: "Woof! Woof!!",
	}

	booluw := Person{
		firstName: "Booluw",
		lastName: "Johnson",
		words: "Hello, my name is ",
	}

	// dog.speak()
	// booluw.speak()

	shout(dog)
	shout(booluw)
}

func (p Person) speak() {
	fmt.Printf("HUMAN: %v %v \n", p.words, p.firstName)
}

func (a Animal) speak() {
	fmt.Printf("ANIMAL: I am a %v, so I %v\n", a.group, a.sound)
}

func shout(m Mamals) {
	m.speak()
}

func sum(i ...int) int {
	total := 0
	
	for _, v := range i {
		total += v
	}

	return total
}