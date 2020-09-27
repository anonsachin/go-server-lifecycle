package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func dynamic(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprintf(w, "Welcome to ==> dynamic <===  You have arrived at the %s path  \n", vars["path"])
}

func dynamicHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to ==> dynamic <===")
}
