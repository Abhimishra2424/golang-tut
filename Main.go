package main

import "fmt"

func main() {

	var amts [3]int = [3]int{10, 20, 30}
	a := amts
	fmt.Printf("Amouts: %v\n", amts)
	fmt.Printf("A: %v\n", a)

}
