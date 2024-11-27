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
	fmt.Printf("Listening on : %d\n", 8070)

	err := http.ListenAndServe(":8070", nil)
	if err != nil {
		fmt.Printf("Error starting server: %s\n", err)

	}
}
