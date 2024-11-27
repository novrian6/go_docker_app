package main

import (
	"fmt"
	"net/http"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}
func main() {
	http.HandleFunc("/", helloWorld)

	err := http.ListenAndServe(":8070", nil)
	if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}
