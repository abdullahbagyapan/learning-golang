package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		i := i

		wg.Add(1) // new wait group
		go func() {
			fmt.Println(i)
			wg.Done() // tell wait group -> I am done
		}()

		// we block runtime
		wg.Wait() // we wait for wait group to be done
	}

	// deterministic output -> in order from 1 to 100

	time.Sleep(time.Second * 1)
}
