package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// how to handle GET or POST request
func getOrPosrRequest(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	switch r.Method {
	case "GET":
		http.ServeFile(w, r, "form.html")
	case "POST":
		// Call ParseForm() to parse the raw query and update r.PostForm and r.Form.
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}
		fmt.Fprintf(w, "Post from website! r.PostFrom = %v\n", r.PostForm)
		name := r.FormValue("name")
		address := r.FormValue("address")
		fmt.Fprintf(w, "Name = %s\n", name)
		fmt.Fprintf(w, "Address = %s\n", address)
	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
}

func getUsers() []byte {
	res, _ := http.Get("https://jsonplaceholder.typicode.com/users/1")
	page, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()
	return page
}

func getUsersAndHandleErrors() {
	res, err1 := http.Get("https://jsonplaceholder.typicode.com/users/1")
	if err1 != nil {
		fmt.Println("error on get")
		return
	}
	page, err2 := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err2 != nil {
		fmt.Println("error on readall")
		return
	}
	fmt.Printf("%s", page)
}
