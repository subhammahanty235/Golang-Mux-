package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Hello")
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler).Methods("GET")

	log.Fatal(http.ListenAndServe(":5000", r))
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>hello server </h1>"))
}
