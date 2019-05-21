package main

import "fmt"

func inc(num *int) *int {
	*num++
	return num
}

func main() {
	myNum := 10
	inc(&myNum)
	fmt.Println(myNum)
}
