package main

import "fmt"

func aaa() {
	fmt.Println("entering A")
	defer fmt.Println("in A first defer")
	defer fmt.Println("in A second defer")
	fmt.Println("Exiting A")
	return
}

func main() {
	fmt.Println("Hello, World!")
	defer fmt.Println("in main first defer")
	fmt.Println("In main calling A")
	defer fmt.Println("in main second defer")
	aaa()
	defer fmt.Println("in main third defer")
	fmt.Println("In main after A")
	defer fmt.Println("in main forth defer")
	fmt.Println("In main exiting main")
}

/*

For <name>:=<init value>; <stay condition>;<<name> manipulation> {
}
For <stay condition> {
}
For index,value:=range <iterate-able variable>  {
}
If you don’t need tha index in range use _

*/
