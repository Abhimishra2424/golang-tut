package main

import (
	"fmt"
)

func main() {

	// control flow statements

	// if i := 2; i == 2 {
	// 	fmt.Println("this is simple if statement")
	// }

	// shopingCart := make(map[string]int)
	// shopingCart = map[string]int{
	// 	"keyboard": 100,
	// 	"Mouse":    89,
	// }

	// shopingCart["Laptop"] = 1500

	// if _, ok := shopingCart["Laptop"]; ok {
	// 	fmt.Println("item exist in the shopCart")
	// }

	//////////////  if else if statement   ////////////////////
	// if i := 2; i == 4 {
	// 	fmt.Println("this is simple if statement")
	// } else if i == 2 {
	// 	fmt.Println("this is simple else if statement")
	// } else {
	// 	fmt.Println("this is simple else block")
	// }

	// And operator dono ko check karega
	// i := 10
	// j := 10
	// if i > 0 && j > 0 {
	// 	fmt.Println("i and j greater than 0")
	// }

	// oR operator dono main se koi ek hoga to
	// i := 10
	// j := 10
	// if i > 0 || j > 0 {
	// 	fmt.Println("i and j greater than 0")
	// }

	// based on tag and expresion
	// switch 1 {
	// case 1, 3, 5, 7, 9:
	// 	fmt.Println("this is odd")
	// case 2, 4, 6, 8:
	// 	fmt.Println("this is even")
	// default:
	// 	fmt.Println("this is Default")
	// }

	// i := 2 + 3
	// switch {
	// case i > 0:
	// 	fmt.Println("this is odd")
	// case i < 5:
	// 	fmt.Println("this is even")
	// default:
	// 	fmt.Println("this is Default")
	// }

	// only one loop for loop in go lang

	// for i := 0; i < 5; i++ {
	// 	fmt.Println(i)
	// }

	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			fmt.Println(i * j)
		}
	}
}
