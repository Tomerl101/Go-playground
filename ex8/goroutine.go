package main

import (
	"fmt"
	"strconv"
	"time"
)

// DoWork this is DoWork
func DoWork(id string) {
	for i := 0; i < 5; i++ {
		fmt.Printf("#" + id + "\n")
		time.Sleep(time.Millisecond * 500) // play with the value

	}
}

func main() {
	fmt.Println("Hello World")
	for i := 0; i < 3; i++ {
		go DoWork(strconv.Itoa(i))
	}
	fmt.Println("Press ENTER to terminate")
	fmt.Scanln() //this line will make sure main func will not exit before workers
	fmt.Println("Bye World")
}
