package main

import (
	"fmt"
	"sync"
	"time"
)

// DoWork this is DoWork
func DoWork2(id int) {
	for i := 0; i < 5; i++ {
		fmt.Printf("#%v\n", id)
		time.Sleep(time.Millisecond * 500)
	}
	fmt.Printf("\n")
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 2; i++ { // try with 2
		wg.Add(1)
		go func(id int) {
			DoWork2(id)
			wg.Done()
		}(i)
	}
	fmt.Println("Waiting.....")
	wg.Wait()
	fmt.Println("End Waiting.....")
}

// wg (WaitGroup) will wait till the WaitGroup counter
// will reach to zero
