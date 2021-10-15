package main

import (
	"fmt"
)

func main() {
	// msg := "hello here"
	// wirteMessage(msg)
	// fmt.Println(msg)
	value := divide(3, 2)
	fmt.Print(value)

	fun := func() {
		fmt.Println("this is anonymons function")
	}
	fun()

}

// func wirteMessage(msg string) {
// 	fmt.Println(msg)
// }

func divide(a float64, b float64) float64 {
	return a / b
}
