package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, Worldss!")
}

func main() {
	http.HandleFunc("/", helloHandler)
	fmt.Println("Server is running on port 9090...")
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println("Error:", err)
	}
}

