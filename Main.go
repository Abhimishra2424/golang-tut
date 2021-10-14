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
		name:   "abhi",
		rollNo: 24,
		subjects: []string{
			"React",
			"Next",
			"go",
		},
	}
	Student1.rollNo = 55
	fmt.Println(Student1)

}
