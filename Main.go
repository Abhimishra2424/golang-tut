package main

import "fmt"

func main() {

	a := [...]int{1, 2, 3, 4, 5, 6, 7, 8}
	b := a[:]
	c := a[2:]
	d := a[:5]
	e := a[2:7]

	fmt.Printf("B: %v\n", a)
	fmt.Printf("A: %v\n", b)
	fmt.Printf("C: %v\n", c)
	fmt.Printf("D: %v\n", d)
	fmt.Printf("E: %v\n", e)

}
