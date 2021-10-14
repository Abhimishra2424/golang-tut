package main

import (
	"fmt"
)

func main() {

	var i = 12
	var j float32 = float32(i)
	fmt.Printf("%v, %T\n", i, i)
	fmt.Printf("%v, %T\n", j, j)
}
