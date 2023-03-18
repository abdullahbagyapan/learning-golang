package main

import "fmt"

func main() {

	x := sumInts(10, 20)
	fmt.Println(x)

	y := sumFloats(11.2, 13.3)
	fmt.Println(y)

	// [type](value1,value2)
	z := sumIntsOrFloats[float32](12.4, 51.4)
	fmt.Println(z)

}

// non generic
// for integer
func sumInts(x int, y int) int {
	return x + y
}

// non generic
// for float
func sumFloats(x float32, y float32) float32 {
	return x + y
}

// take "T" type and return "T" type
// "T" can be int or float
// inputs have to be same type
func sumIntsOrFloats[T int | float32](x T, y T) T {
	return x + y
}
