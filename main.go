package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HandleHello(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello\n vars from server: %v\n", vars)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", HandleHello)
	log.Fatal(http.ListenAndServe(":8081", nil))
}
