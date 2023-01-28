package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HandleHello(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte("Hello\n"))
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", HandleHello)
	log.Fatal(http.ListenAndServe(":8081", nil))
}
