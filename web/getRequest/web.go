package main

import (
	"fmt"
	"net/http"
)

//go run web/getRequest/*.go
//we run httpget.go and web.go
// then in the browser go to localhost:3000

func indexHandle(w http.ResponseWriter, r *http.Request) {
	res := getUsers()
	fmt.Fprintf(w, "%s", res)
	fmt.Fprintf(w, "hello world from web")
}
func aboutHandle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "this is about page")
}
func main() {
	http.HandleFunc("/", indexHandle)
	http.HandleFunc("/about", aboutHandle)
	http.ListenAndServe(":3000", nil)
}
