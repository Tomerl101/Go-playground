package main

import "fmt"

func main() {
	fmt.Println("fun with slices")
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	b := a[3:6]
	fmt.Println(a, b)
	b[0] = 0
	fmt.Println(a, b)
	c := append(b, 11)
	fmt.Println(a, b, c)
	c[1] = 50
	fmt.Println(a, b, c)
	fmt.Println("***************************")

	fmt.Println("len: ", len(a), " cap: ", cap(a))
	for i, value := range a {
		fmt.Println(i, "-", value)
	}
	fmt.Println("***************************")

	d := make([]int, 3)
	d[0] = 7
	fmt.Println(d)
	d = append(d, 11)
	fmt.Println(d)
	d = append(d, d...)
	fmt.Println(d)
	d = append(d[3:7], d[0:1]...)
	fmt.Println(d)
	d[0]++
	fmt.Println(d)
	fmt.Println("***************************")

	e := []int{}
	fmt.Println(e)
	fmt.Println(e == nil)
	var f []int
	fmt.Println(f)
	fmt.Println(f == nil)
	//f[0] = 1 //# error, it is nil use make
	g := make([]int, 0)
	fmt.Println(g)
	fmt.Println(g == nil)
	fmt.Println("***************************")

	dim := make([][]int, 2)
	fmt.Println(dim)
	fmt.Println(dim[0])
	dim[0] = []int{1, 2, 3}
	fmt.Println(dim[0])
	fmt.Println(dim)
	dim[1] = []int{4, 5, 6, 7}
	fmt.Println(dim)
	dim = append(dim, []int{8, 9})
	fmt.Println(dim)
}

/*

slice = Dynamic arrays (kind of)
<name>:=[]<type>{initial values}
You can create a new slice from slice
<slice> = append(<base slice>, <values>)
Sub slicing using [start:last+1]
https://golang.org/ref/spec#Slice_expressions
https://golang.org/ref/spec#Slice_types

//slices hold pointer
*/
