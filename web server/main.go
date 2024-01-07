package main

import (
	"fmt"
	"log"
	"net/http"
)


func helloHandler(w http.ResponseWriter, r *http.Request){
	// first take care of errors then main code
	
	
	
	fmt.Fprintf(w, "hello!!")
	return
}

func formHandler(w http.ResponseWriter, r *http.Request){
	if err := r.ParseForm(); err != nil{
		fmt.Fprintf(w, "Error parsing form %s", err)
		return
	}

	fmt.Fprintf(w, "Post Request Successful\n")
	name := r.FormValue("Name")
	fmt.Fprintf(w, "Name: %s", name) 
	// return
}

func main(){
	// static folder, serve file
	// create routes for /, /hello, /form

	fileserver := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileserver)
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/form", formHandler)
	fmt.Println("Server Running on Port 8080:")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}