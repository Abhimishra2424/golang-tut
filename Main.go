package main

import "fmt"

func main() {

	var amts [3]int = [3]int{10, 20, 30}
	amt := [...]int{30, 40, 40, 0, 3, 2, 4, 52, 1}
	fmt.Printf("amounts:%v", amts)
	fmt.Printf("amounts:%v", amt)
	fmt.Printf("%v\n", len(amt))

}
