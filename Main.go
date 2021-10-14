package main

import "fmt"

func main() {
	s := "this is string"
	s1 := "this is string"
	fmt.Printf("%v %T", s+s1, s+s1)
}
