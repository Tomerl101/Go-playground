package main

// printf is a package so we import fmt
//use go build <file_name> to create runnable file

//go run ex1/ex1.go
import "fmt"

func main() {
	fmt.Printf("hello, Tomer\n")
	// can use ex2 because it is in the same package (main)
	ex2()
}
