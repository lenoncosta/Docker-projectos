package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/mdc", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Mentoria Devops Cloud!")
	})

	log.Println("Server started on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

