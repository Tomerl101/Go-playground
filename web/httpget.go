package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

//files with the same packge name cannot have the same func name
func main2() {
	fmt.Println("Hello World")
	fmt.Println("*********************************************")
	res, _ := http.Get("https://jsonplaceholder.typicode.com/users/1")
	page, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()
	fmt.Printf("%s", page)
	fmt.Println("*********************************************")
	res, err := http.Get("https://jsonplaceholder.typicode.com/users/1")
	if err != nil {
		fmt.Println("error on get")
		return
	}
	page, err = ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		fmt.Println("error on readall")
		return
	}
	fmt.Printf("%s", page)
}

func sendData() []byte {
	res, _ := http.Get("https://jsonplaceholder.typicode.com/users/1")
	page, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()
	return page
}
