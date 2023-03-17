package main

import "fmt"

func main() {
	var arr = [4]int{1, 2, 3}
	// arr[3] = 0 auto-assign

	fmt.Println(arr)
	// output will be [1 2 3]

	arr[2] = 10
	fmt.Println(arr)
	// output will be [1 2 10]

}
