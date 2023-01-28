package main

import (
	"fmt"
	"log"
	"net/http"
)

func HandleHello(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello\n")
}

func main() {
	http.HandleFunc("/", HandleHello)
	log.Fatal(http.ListenAndServe(":8081", nil))
}
