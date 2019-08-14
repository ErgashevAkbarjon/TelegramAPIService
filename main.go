package main

import (
	"fmt"
	"net/http"
)

func main() {
	handleRequests()
}

func handleRequests() {
	http.HandleFunc("/", mainPageHandler)
	http.HandleFunc("/about", aboutPageHandler)
	http.ListenAndServe(":4343", nil)
}

func mainPageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hellooo world")
	fmt.Println("Endpoint hit: home page")
}

func aboutPageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "About page")
	fmt.Println("Endpoint hit: About page")
}
