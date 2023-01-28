package main

import (
	"fmt"
	"log"
	"net/http"
)

func HandleHello(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprint(w, "Hello \n")
}

func main() {
	http.HandleFunc("/", HandleHello)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
