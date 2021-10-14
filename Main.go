package main

import "fmt"

//Student Struct
type Student struct {
	name     string
	rollNo   int
	subjects []string
}

func main() {

	Student1 := Student{
		"abhi",
		24,
		[]string{
			"React",
			"Next",
			"go",
		},
	}
	Student1.rollNo = 55
	fmt.Println(Student1)

}
