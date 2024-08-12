package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, world!")
}
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/hello", hello).Methods("GET")
	r.HandleFunc("/hello", hello).Methods("POST")
	r.HandleFunc("/hello/{name}", hello).Methods("GET")
	r.HandleFunc("/hello/{name}", hello).Methods("POST")
	r.HandleFunc("/hello/{name}/{age}", hello).Methods("GET")
	r.HandleFunc("/hello/{name}/{age}", hello).Methods("POST")
	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}
