package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
	j := 0
	for ; j < 5; j++ {
		fmt.Println(j)
	}
	k := 0
	for k < 5 {
		fmt.Println(k)
		k++
	}
	l := 0
	for aa := 9; l < 5; fmt.Println(aa + l) {
		fmt.Println(l)
		l++
	}
	m := 0
	for {
		m++
		if m == 2 {
			continue
		}
		fmt.Println(m)
		if m > 5 {
			break
		}

	}
	nums := []int{7, 9, 5, 1, 6, 5}
	for i, id := range nums {
		fmt.Printf("%d-%d\n", i, id)
	}
	for _, id := range nums {
		fmt.Printf("-%d\n", id)
	}
	for i := range nums {
		fmt.Printf("%d-\n", i)
	}
	mymap1 := map[string]string{"A": "a", "B": "b"}
	for n, v := range mymap1 {
		fmt.Printf("%s: %s\n", n, v)
	}
	for _, v := range mymap1 {
		fmt.Printf(": %s\n", v)
	}
	for n := range mymap1 {
		fmt.Printf("%s: \n", n)
	}
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
