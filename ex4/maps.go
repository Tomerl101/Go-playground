package main

import "fmt"

func changemap(m map[int]int) {
	m[0] = 1000
}

func main() {
	var mymap map[int]int
	mymap = make(map[int]int)
	mymap[1] = 1
	fmt.Println(mymap)
	mymap[1] = 2
	fmt.Println(mymap)
	mymap[2] = 10
	fmt.Println(mymap)
	delete(mymap, 1)
	fmt.Println(mymap)
	fmt.Println("**** Map (and slice and channel) is a reference type ******")
	mymap[1] = 0
	fmt.Println(mymap)
	changemap(mymap)
	fmt.Println(mymap)
	fmt.Println("***************************")

	mymap1 := map[string]string{"A": "a", "B": "b"}
	fmt.Println(mymap1)
	fmt.Println("len: ", len(mymap1))
	for key, value := range mymap1 {
		fmt.Println("k:", key, " v:", value)
	}
	val, exist := mymap1["A"]
	fmt.Println(val, exist)
	val, exist = mymap1["C"]
	fmt.Println(val, len(val), exist)
	mymap2 := map[int]int{}
	mymap2[7] = 9
	fmt.Println(mymap2)
	var mymap3 map[int]int
	//# mymap3[7] = 11 //this is an error  (look at the initalize of mymap1 and mymap2)
	fmt.Println(mymap3)

}

/*
must use make otherwise it is nil
Index value pairs
//The index is string the value watever?
Var <name> map[<keys type>]<values type>
<name>:=make(map[<key type>]<value type>,optional initial length)
<name>:=map[<type>]<type>{key:val,...}
<name>[<key>] = <value>
https://golang.org/ref/spec#Map_types
delete(<mapname>, <key>)
len(<mapname>)

*/
