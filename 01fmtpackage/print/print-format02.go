package main

import (
	"fmt"
)

type student struct {
	name string
	age  int
}

func main() {

	var std = student{
		name: "Jack",
		age:  18,
	}

	// print struct instances without key
	fmt.Printf("%v \n", std)
	// output = {Jack 18}

	// with key
	fmt.Printf("%+v \n", std)
	// output = {name:Jack age:18}
}
