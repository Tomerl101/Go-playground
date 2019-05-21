package main

import "fmt"

// Hello function
func Hello(name string) string {
	return "Hello " + name
}

// DoSum function
func DoSum(num1, num2 int) int {
	return num1 + num2
}
func aoSum(num1, num2 int) int {
	return num1 + num2
}
func main() {
	fmt.Println("Hello, World!")
	fmt.Println(Hello("kuku"))
	fmt.Println(DoSum(5, 6))
	fmt.Println(aoSum(5, 6))
}

/*

Func <name>(<name> <type>,…) <return type> {
Return <whatever>
}
The same shorting of x,y type is OK here too
Can have multiple return values
If the return value(s) are named you can assign them along the function and just return “naked” without any value. It is not recommended in long functions because it diminish readability.
The second return value is usually am error object
Go does not have exceptions !!!

*/
