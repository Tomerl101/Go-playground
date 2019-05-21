package main

// printf is a package so we import fmt
//use go build <file_name> to create runnable file

//go run ex1/*.go both ex2 ex1 need to run because ex1 use ex2 inside of him
import "fmt"

func main() {
	fmt.Printf("hello, Tomer\n")
	// can use ex2 because it is in the same package (main)
	ex2()
}
