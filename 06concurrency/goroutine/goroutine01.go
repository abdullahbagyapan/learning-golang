package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 100; i++ {
		a := i // we copy i to a for each lap
		// new integer variable for each lap

		go func() { // nondeterministic output -> random output from 0 to 100
			fmt.Println(a)
			time.Sleep(time.Second * 5)
		}()
	}

	time.Sleep(time.Second * 1)

}
