package main

import "fmt"

func main() {
	s := "this is string"
	s1 := []byte(s)
	fmt.Printf("%v %T", s1, s1)
}
