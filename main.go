package main

import (
    "net/http"
    "github.com/gorilla/mux"
)

func main() {
    r := mux.NewRouter()

	r.HandleFunc("/",base).Methods("GET")
	
	dyn := r.PathPrefix("/dynamic").Subrouter()

	dyn.HandleFunc("/",dynamicHello).Methods("GET")

	dyn.HandleFunc("/{path}",dynamic).Methods("GET")

    http.ListenAndServe(":3000", r)
}