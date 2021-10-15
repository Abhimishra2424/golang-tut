package main

import (
	"fmt"
)

func main() {

	// if i := 2; i == 2 {
	// 	fmt.Println("this is simple if statement")
	// }

	shopingCart := make(map[string]int)
	shopingCart = map[string]int{
		"keyboard": 100,
		"Mouse":    89,
	}

	shopingCart["Laptop"] = 1500

	if _, ok := shopingCart["Laptop"]; ok {
		fmt.Println("item exist in the shopCart")
	}

}
