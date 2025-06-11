package main

import (
	"fmt"
	"net/http"
	"time"
)

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World! %s", time.Now())
}

func newMethod(w http.ResponseWriter, r *http.Request) {
	fmt.Println("This is a new method.")
}

func main() {
	http.HandleFunc("/", greet)
	http.HandleFunc("/new_method", newMethod) 
	http.ListenAndServe(":8080", nil)
}
