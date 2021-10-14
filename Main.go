package main

import "fmt"

func main() {

	var amts [3]int = [3]int{10, 20, 30}
	amt := [3]int{30, 40, 40}
	fmt.Printf("amounts:%v", amts)
	fmt.Printf("amounts:%v", amt)
	fmt.Printf("%v\n", len(amt))

}
