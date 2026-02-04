package main

import (
	"firstmodule/reverse"
	"fmt"
	"github.com/google/go-cmp/cmp"
)

func main() {
	str := "Booluw"

	fmt.Println(str, &str)

	fmt.Println(reverse.ReverseString("Boluwatife"))
	fmt.Println(cmp.Diff("Booluw", "Booluwatife"))
}
