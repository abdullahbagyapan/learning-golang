package main

import (
	"fmt"
)

func main() {

	var str string = "A" // initialize string

	// string
	fmt.Printf("%s \n", str)
	// output = A

	// string with quotes
	fmt.Printf("%q \n", str)
	// output = "A"

	// hexadecimal format of string in ASCII
	fmt.Printf("%x \n", str)
	// output = 41

	var isOkay bool = true

	// boolean format
	fmt.Printf("%t \n", isOkay)
	// output = true

	// address of variable (pointer)
	fmt.Printf("%p \n", &isOkay)
	// output = 0x001...
}
