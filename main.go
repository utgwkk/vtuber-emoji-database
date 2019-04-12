package main

import (
	"fmt"
	"log"
	"net/http"
)

func handleTopPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<h1>It works!</h1>")
}

func main() {
	http.HandleFunc("/", handleTopPage)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
