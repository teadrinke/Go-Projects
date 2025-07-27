package main

import (
	"fmt"
	"log"
	"net/http"

)

func formHandler(w http.ResponseWriter,r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
		//This line checks if the form submission could not be parsed. If it could not, it logs the error and returns.
	}
	fmt.Fprintf(w, "Post request successful\n")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name: %s\n", name)
	fmt.Fprintf(w, "Address: %s\n", address)
	//This line retrieves the form values (name and address) from the request and prints them to the console.
}
func helloHandler(w http.ResponseWriter,r *http.Request) {
	if r.URL.Path != "/hello"{
		http.Error(w, "404 not found", http.StatusNotFound)
		return
		//This line checks if the requested URL path is not "/hello". If it is not, it returns a 404 Not Found error.
		//This is necessary as /hello/hi or in general /hello/... would also be able to call hello handler but since this check is done in the helloHandler function, it would not be executed only strictly the /hello will only be executed.
	}
	if r.Method != "GET" {
		http.Error(w, "Method is not supported", http.StatusNotFound)
		return
		//This line checks if the HTTP method used for the request is not GET. If it is not, it returns a status Not Found error.
	}

	fmt.Fprintf(w, "Hello, World!")
}
func main() {
	//A file server is a centralized system that stores and shares files over a network.
	fileServer := http.FileServer(http.Dir("./static"))
	//It creates a file server that serves static files (like HTML, CSS, JS, images) from the ./static folder on your computer.
	http.Handle("/", fileServer)
	//The http.Handle function registers the file server to handle requests to the root URL ("/").
	http.HandleFunc("/form", formHandler)
	//This line registers a handler function for the "/form" URL path, which will handle form submissions.
	http.HandleFunc("/hello", helloHandler)
	//This line registers a handler function for the "/hello" URL path, which will display a simple "Hello, World!" message.

	fmt.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080",nil); err !=nil{
		log.Fatal(err)
		//This line starts the HTTP server on port 8080 and logs any errors that occur while starting the server.
	}
}