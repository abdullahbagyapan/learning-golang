package main

import "fmt"

// we can declare generic types as an interface
type number interface {
	int | float64
}

func main() {

	result := sumTwoValue(10.5, 20.3)
	fmt.Println(result)

}

func sumTwoValue[T number](x T, y T) T {
	return x + y
}
