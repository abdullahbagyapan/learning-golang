package main

import (
	"fmt"
	"time"
)

func main() {

	go func() { // if we do not wait goroutine it behave nondeterministic
		fmt.Println("Hello from func")
	}()

	fmt.Println("Hello from main")

	time.Sleep(time.Second) // we wait for goroutine

	// output will be :
	// 1) Hello from main
	// 2) Hello from func

}
