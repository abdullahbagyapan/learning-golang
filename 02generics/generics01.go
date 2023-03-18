package main

import (
	"fmt"
	"reflect"
)

func main() {

	x := convertAnyToString(1232.2)
	fmt.Println(x)

	typeOfX := reflect.ValueOf(x).Kind() // type of x
	fmt.Println(typeOfX)
	// output will be string

}

func convertAnyToString(x any) string {
	y := fmt.Sprint(x)
	return y
}
