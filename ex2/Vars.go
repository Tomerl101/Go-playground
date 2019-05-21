package main

import "fmt"

const n1 = iota
const n2 = iota * 10

//iota incremet the number by 1 every time it been calld
//it only work in cont (var1, var2,...)
const (
	_  = iota // _ is ignore saving the return value so we cant use it
	n3 = iota
	n4 = iota
	n5 = iota * 10
)

func main() {
	c := 6
	fmt.Printf("%T\n", c)
	fmt.Println("==== seperate consts with iota ====")
	fmt.Println(n1)
	fmt.Println(n2)
	fmt.Println("==== chaning consts with iota ====")
	fmt.Println(n3)
	fmt.Println(n4)
	fmt.Println(n5)
}
