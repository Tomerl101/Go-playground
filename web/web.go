package main

import (
	"fmt"
	"net/http"
)

//go run web/*.go
// then in the browser go to localhost:3000

func index(w http.ResponseWriter, r *http.Request) {
	res := sendData()
	fmt.Fprintf(w, "%s", res)
	fmt.Fprintf(w, "hello world from web")
}
func index1(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello about world from web")
}
func main() {
	fmt.Println("Hello World")
	http.HandleFunc("/", index)
	http.HandleFunc("/about", index1)
	http.ListenAndServe(":3000", nil)
}
