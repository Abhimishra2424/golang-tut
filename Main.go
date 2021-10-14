package main

import "fmt"

type Processor struct {
	processorName string
	cores         int
}

type Memory struct {
	memoryCapacity int
	memoryName     string
}

type Computer struct {
	Processor
	Memory
	price int
}

func main() {

	c := Computer{}
	c.price = 5000
	c.processorName = "intel i5"
	c.cores = 5
	c.memoryCapacity = 8
	c.memoryName = "DDR4"

	fmt.Println(c)

}
