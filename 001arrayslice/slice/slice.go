package main

import "fmt"

func main() {

	var arr = [4]int{1, 2, 3, 4}

	slc := arr[2:4]

	fmt.Println(slc)
	// output will be [3,4]

	slc[1] = 10
	// slices are reference of array
	// they do not hold data

	fmt.Println(arr)
	// output will be [1,2,3,10]

	// alternative way declare slice
	// make(type,length,capacity)
	// capacity is optional default = length of slice
	slc1 := make([]int, 2, 4)

	fmt.Println(len(slc1)) // output will be 2
	fmt.Println(cap(slc1)) // output will be 4

}
