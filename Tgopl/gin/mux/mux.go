package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func hello1(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, world!")
}
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/hello", hello1).Methods("GET")
	r.HandleFunc("/hello", hello1).Methods("POST")
	r.HandleFunc("/hello/{name}", hello1).Methods("GET")
	r.HandleFunc("/hello/{name}", hello1).Methods("POST")
	r.HandleFunc("/hello/{name}/{age}", hello1).Methods("GET")
	r.HandleFunc("/hello/{name}/{age}", hello1).Methods("POST")
	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}
