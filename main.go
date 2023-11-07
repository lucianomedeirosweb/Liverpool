package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/createPerson", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "frsfsr!")
	})

	http.HandleFunc("/getPerson", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "getPerson!")
	})

	http.HandleFunc("/updatePerson", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "updatePerson!")
	})

	http.HandleFunc("/deletePerson", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "deleteperson!")
	})

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
