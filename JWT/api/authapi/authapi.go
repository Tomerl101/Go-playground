package authapi

import (
	"fmt"
	"net/http"
)

func getToken(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(response, "API 1 func ")
}

func useToken(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(response, "API 2 func ")
}
