package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
)

func main() {
	db := database{map[string]dollars1{"shoes": 50, "socks": 5}, sync.Mutex{}}
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	http.HandleFunc("/update", db.update)
	http.HandleFunc("/delete", db.delete)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

type dollars1 float32

func (d dollars1) String() string { return fmt.Sprintf("$%.2f", d) }

type database struct {
	R     map[string]dollars1
	mutex sync.Mutex
}

func (db *database) list(w http.ResponseWriter, req *http.Request) {
	for item, price := range db.R {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func (db *database) update(w http.ResponseWriter, req *http.Request) {
	db.mutex.Lock()
	defer db.mutex.Unlock()

	item := req.URL.Query().Get("item")
	price := req.URL.Query().Get("price")

	p, err := strconv.ParseFloat(price, 64)
	if err != nil {
		fmt.Fprint(w, "invalid price")
	} else {
		db.R[item] = dollars1(p)
		fmt.Fprint(w, "Update price success")
	}
}

func (db *database) delete(w http.ResponseWriter, req *http.Request) {
	db.mutex.Lock()
	defer db.mutex.Unlock()

	item := req.URL.Query().Get("item")
	delete(db.R, item)
	fmt.Fprint(w, "Success")
}

func (db *database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if price, ok := db.R[item]; ok {
		fmt.Fprintf(w, "%s\n", price)
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item: %q\n", item)
	}
}
