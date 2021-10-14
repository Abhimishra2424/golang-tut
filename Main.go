package main

import "fmt"

// User Constant
const (
	User    string = "admin"
	Product string = "Product"
)

func main() {

	const i int = 12
	const j float32 = 3.14
	const k string = "abhi"
	const l bool = true
	fmt.Println(i)
	fmt.Println(j)
	fmt.Println(k)
	fmt.Println(l)

	var a int = 13
	fmt.Println(i + a)

}
