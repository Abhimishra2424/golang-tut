package main

import (
	"fmt"
)

func main() {

	// shopingCard := make(map[strings]int)

	shopingCard := map[string]int{
		"keyboard": 100,
		"mouse":    50,
	}
	shopingCard["keyboard"] = 200
	shopingCard["Monitor"] = 600

	fmt.Println(shopingCard)
}
