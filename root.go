package main

import (
	"fmt"
	"net/http"
)

func base(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome, You have arrived at the ROOT path  -- %s --\n", r.URL.Path)
}
