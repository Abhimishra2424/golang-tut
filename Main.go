package main

import "fmt"

func main() {
	i := 3.14
	j := 1.7e12
	k := 2.3e12
	fmt.Printf("%v %T", i, i)
	fmt.Printf("%v %T", j, j)
	fmt.Printf("%v %T", k, k)
}
