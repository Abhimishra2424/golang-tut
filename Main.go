package main

import (
	"fmt"
)

func main() {

	shopingCard := map[string]int{
		"keyboard": 100,
		"mouse":    50,
	}
	shopingCard["keyboard"] = 200
	shopingCard["Monitor"] = 600
	_, ok := shopingCard["phone"]
	_, hai := shopingCard["mouse"]

	fmt.Println(shopingCard)
	fmt.Println(ok)
	fmt.Println(hai)
}
