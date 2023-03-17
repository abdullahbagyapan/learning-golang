package main

import "fmt"

func main() {

	var x int

	// if no error isOkay will be 1
	// if is not , isOkay will be 0
	// err will report why
	isOkay, err := fmt.Scan(&x)

	fmt.Printf("%d \n", x)
	fmt.Printf("%d \n", isOkay)
	fmt.Printf("%d \n", err)

	var str string

	// read until space
	isOkay, err = fmt.Scan(&str)

	fmt.Printf("%s \n", str)
	fmt.Printf("%d \n", isOkay)
	fmt.Printf("%d \n", err)

}
