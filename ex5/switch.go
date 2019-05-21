package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
	var x int = 8
	switch x {
	case 10 - 2:
		fmt.Println("x>5")
	case 9:
		fmt.Println("x==8")
	default:
		fmt.Println("none")
	}
	var y bool = true
	switch y {
	case x > 5:
		fmt.Println("x>5")
	case false:
		fmt.Println("x==8")
		fallthrough // the oposit of C break
	default:
		fmt.Println("none")
	}
	var st string = "ggd"
	switch st {
	case "a", "b":
		fmt.Println("by")
		fallthrough
	default:
		fmt.Println("hi")
	}
	switch {
	case 2 == 2:
		fmt.Println("a")
	case 4 > 7:
		fmt.Println("by")
	}
}

/*

switch <expression> {
case <value:
     . . .
case . . .
     . . .
default:
     . . .
}
  No break needed as in C

*/
