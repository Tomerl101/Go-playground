package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// Slice the slice to give it zero length.
	s = s[:0]
	printSlice(s)

	// Extend its length.
	s = s[:4]
	printSlice(s)

	// Drop its first two values.
	s = s[2:]
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

/*

Dynamic arrays (kind of)
<name>:=[]<type>{initial values}
You can create a new slice from slice
<slice> = append(<base slice>, <values>)
Sub slicing using [start:last+1]
https://golang.org/ref/spec#Slice_expressions
https://golang.org/ref/spec#Slice_types


*/
