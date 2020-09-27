package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", base).Methods("GET")

	dyn := r.PathPrefix("/dynamic").Subrouter()

	dyn.HandleFunc("/", dynamicHello).Methods("GET")

	dyn.HandleFunc("/{path}", dynamic).Methods("GET")

	fmt.Println("Starting ...")

	http.ListenAndServe(":3000", r)
}
