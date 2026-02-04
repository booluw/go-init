package main

import "fmt"

const (
	B = iota
	KB = iota
	MB = iota * 10
	GB = iota * 10
)

func main() {
	fmt.Println("🚀")
	// fmt.Println(`
	// 	1 There's not a friend like the lowly Jesus.
	// 	No, not one! No, not one!
	// 	None else could heal all our soul's diseases.
	// 	No, not one! No, not one!
	// 	Refrain:
	// 	Jesus knows all about our struggles;
	// 	He will guide till the day is done.
	// 	There's not a friend like the lowly Jesus.
	// 	No, not one! No, not one!
	// 	2 No friend like Him is so high and holy.
	// 	No, not one! No, not one!
	// 	And yet no friend is so meek and lowly.
	// 	No, not one! No, not one! [Refrain]
	// 	3 There's not an hour that He is not near us.
	// 	No, not one! No, not one!
	// 	No night so dark but His love can cheer us.
	// 	No, not one! No, not one! [Refrain]
	// 	4 Did ever saint find this Friend forsake him?
	// 	No, not one! No, not one!
	// 	Or sinner find that He would not take him?
	// 	No, not one! No, not one! [Refrain]
	// 	5 Was e'er a gift like the Savior given?
	// 	No, not one! No, not one!
	// 	Will He refuse us a home in heaven?
	// 	No, not one! No, not one! [Refrain]
	// `)

	// All variables must be used, else throw into the abyss with _
	name, age, _ := "Booluw", 12, 20_02

	fmt.Println(name, age)

	fmt.Println("Printing the HEX and Binary values of different numbers and strings")

	/*
		%s => strings
		%x => HEX
		%d => digits
		%b => Binary
	*/
	fmt.Printf("The HEX for %s is %#x \n", name, name)
	fmt.Printf("The Binary for %d is %b \n", age, age)
	fmt.Println(B, KB, MB, GB)
}