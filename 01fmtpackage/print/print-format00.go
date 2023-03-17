package main

import "fmt"

func main() {

	var x int = 10

	// decimal -> base 10
	fmt.Printf("%d \n", x)
	// output = 10

	// binary -> base 2
	fmt.Printf("%b \n", x)
	// output = 1010

	// octal -> base 8
	fmt.Printf("%o \n", x)
	// output = 12

	// hexadecimal -> base 16
	fmt.Printf("%x \n", x)
	// output = a

	var y float32 = 10.5

	// decimal point
	fmt.Printf("%f \n", y)
	// output = 10.500000

	// same as f but without 0 after point
	fmt.Printf("%g \n", y)
	// output = 10.5
}
