package main

import (
	"fmt"
	"net/http"

	"./api/accountapi"
)

func main() {
	http.HandleFunc("/api/account/generatetoken", accountapi.GenerateToken)
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		fmt.Println(err)
	}
}

//go run JWT/main.go

//use postman to send POST request to localhost:3000/api/account/generatetoken with body {"username":"test1","password":"aabbcc"} to get token
