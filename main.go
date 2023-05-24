package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

func getRoot(w http.ResponseWriter, r *http.Request){
	fmt.Printf("got / request\n")
	io.WriteString(w, "This is my website!\n")
}

func getHello(res http.ResponseWriter, req *http.Request) {
	fmt.Printf("got /hello request\n")
	io.WriteString(res, "Hello, HTTP!\n")
}

func getMew(res http.ResponseWriter, req *http.Request){
	// print route path
	fmt.Printf("got /mew request \n")
	io.WriteString(res,"Hello Babe :D \n")
}

func main(){
	//create route
	http.HandleFunc("/", getRoot)
	http.HandleFunc("/hello", getHello)
	http.HandleFunc("/mew",getMew)

	err := http.ListenAndServe(":3333", nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)}
}