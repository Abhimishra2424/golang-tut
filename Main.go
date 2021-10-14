package main

import "fmt"

func main() {

	a := [...]int{1, 2, 3, 4, 5, 6, 7, 8}
	b := a[:]

	fmt.Printf("B: %v\n", a)
	fmt.Printf("A: %v\n", b)

}
