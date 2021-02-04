package main

import (
	"fmt"
	"log"
	"net/http"
)

var countView int

func homeWelcome(w http.ResponseWriter, r *http.Request) {
	countView++

	fmt.Fprintf(w, "Welcome to DB API on Go \nPage Views: %d", countView)
	fmt.Printf("Endpoint Hit: Welcome Count %d", countView)
}

func handleRequests() {
	http.HandleFunc("/", homeWelcome)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	handleRequests()
}
