package main

import "fmt"

func main() {

	var identityMatrix [3][3]int = [3][3]int{
		{1, 0, 0},
		{0, 1, 1},
		{0, 0, 1},
	}
	identityMatrix[1][2] = 7
	fmt.Println(identityMatrix)
}
