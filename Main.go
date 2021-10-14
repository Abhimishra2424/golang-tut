package main

import (
	"fmt"
	"strconv"
)

func main() {

	var i = 12
	var j float32 = float32(i)
	var foo string = strconv.Itoa(i)
	fmt.Printf("%v, %T\n", i, i)
	fmt.Printf("%v, %T\n", j, j)
	fmt.Printf("%v, %T\n", foo, foo)
}
