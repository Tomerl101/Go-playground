package main

import "fmt"

func hello(name string) string {
	return "Hello " + name
}

func doSum(num1, num2 int) int {
	return num1 + num2
}

// for multipale return value we need to wrap the return type in brackets!
func returnMultiVal() (int, int) {
	a := 10
	b := 20
	return a, b
}

func main() {
	fmt.Println(hello("Tomer"))
	fmt.Println(doSum(5, 5))
	fmt.Println(returnMultiVal())

	//Anonymous function
	func() {
		fmt.Println("i run my self! and i dont have name!")
	}()
}
