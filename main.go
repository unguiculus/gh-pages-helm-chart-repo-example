package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Starting example app...")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello Helm!")
	})

	fmt.Println("Serving on port 8080...")
	http.ListenAndServe(":8080", nil)
}
