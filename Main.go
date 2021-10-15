package main

import (
	"fmt"
)

func main() {
	rect := rectangle{
		width: 20, height: 10,
	}
	area := rect.area()
	fmt.Println("area of rectangle:", area)
}

type rectangle struct {
	width, height int
}

func (r rectangle) area() int {
	return r.width * r.height
}
