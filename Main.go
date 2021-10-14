package main

import (
	"fmt"
)

func main() {

	shopingCard := map[string]int{
		"keyboard": 100,
		"mouse":    50,
	}

	delete(shopingCard, "mouse")
	fmt.Println(shopingCard)

}
