package main

import (
	"fmt"
	"io"
	"net/http"

	"./api/authapi"
	"./middleware/authmiddleware"
)

func handlePrivate(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	io.WriteString(response, `{"ok":"valid_key", "msg":"hello user :)"}`)
	//we can also get the token and validate here if we like to...
}

func main() {
	http.HandleFunc("/auth/generatetoken", authapi.GenerateToken)
	http.HandleFunc("/auth/getjwtclaims", authapi.GetJwtClaims)

	//wrap in a HandlerFunc
	handlePrivate := http.HandlerFunc(handlePrivate)
	http.Handle("/private", authmiddleware.AuthMiddleware(handlePrivate))

	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		fmt.Println(err)
	}
}

//go run JWT/main.go

//use postman to send POST request to localhost:3000/api/account/generatetoken with body {"username":"test1","password":"aabbcc"} to get token

//use postman to GET response from server with the jwt user claims , need to pass header token:"tokenString"
