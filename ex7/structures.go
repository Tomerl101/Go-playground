package main

import "fmt"
import "strconv"

type myitem struct {
	id   int
	name string
	cost int
}

// by value
func (m myitem) do1(i int, s string) string {
	return "*" + strconv.FormatInt(int64(i), 10) +
		":" + s + ":" + strconv.FormatInt(int64(m.id), 10) +
		":" + m.name + ":" +
		strconv.FormatInt(int64(m.cost), 10) + "*"
}

// by ref
func (m *myitem) do2() {
	m.id += 1
	m.name = m.name + "2"
	m.cost = int(m.cost * 2)
}

func main() {
	fmt.Println("Hello, World!")
	item1 := myitem{id: 7, name: "a", cost: 77}
	item2 := myitem{8, "b", 99}
	fmt.Println(item1)
	fmt.Println(item2)
	fmt.Println(item1.do1(8, "sss"))
	item1.do2()
	fmt.Println(item1)
	item2.id = 897
	fmt.Println(item2)
}

/*

Go does not have classes
Structures can have functions
Type <name> struct {
<name> <type>
. . .
}
Declaring a variable of type struct
<varname>:=<struct type>{<name:value>…} or no names
Dot notation for accessing elements
Can compress names of fields of same type and use one type at the end


Func (<name> <struct name>) <Functionname> (<params>) <return type> {
<name>.<field> is like this iv C++ and other Lg
} // value receiver will not change the struct
Func (<name> * <struct name>) <Functionname> (<params>) <return type> {
<name>.<field> is like this iv C++ and other Lg
} // pointer receiver will change the struct
*/
