package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"sync"
)

func main() {
	db := databaseOne{map[string]dollars{"shoes": 50, "socks": 5}, sync.Mutex{}}
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	http.HandleFunc("/update", db.update)
	http.HandleFunc("/delete", db.delete)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type databaseOne struct {
	R     map[string]dollars
	mutex sync.Mutex
}

func (db databaseOne) list(w http.ResponseWriter, req *http.Request) {
	var shopList = template.Must(template.New("shopList").Parse(`
	<h1>shopList</h1>
	<table>
	<tr style='text-align: left'>
	<th>item</th>
	<th>	</th>
	<th>price</th>
	</tr>
	</table>
	`))
	shopList.Execute(w, nil)

	const templ = `<p>{{.A}}	{{.B}}</p>`
	type data struct {
		A string
		B dollars
	}

	t := template.Must(template.New("escape").Parse(templ))

	for item, price := range db.R {
		var dat = data{item, price}

		err := t.Execute(w, dat)
		if err != nil {
			log.Fatal(err)
		}
	}

}

func (db databaseOne) update(w http.ResponseWriter, req *http.Request) {
	db.mutex.Lock()
	defer db.mutex.Unlock()

	item := req.URL.Query().Get("item")
	price := req.URL.Query().Get("price")

	p, err := strconv.ParseFloat(price, 64)
	if err != nil {
		fmt.Fprint(w, "invalid price")
	} else {
		db.R[item] = dollars(p)
		fmt.Fprint(w, "Update price success")
	}
}

func (db databaseOne) delete(w http.ResponseWriter, req *http.Request) {
	db.mutex.Lock()
	defer db.mutex.Unlock()

	item := req.URL.Query().Get("item")
	delete(db.R, item)
	fmt.Fprint(w, "Success")
}

func (db databaseOne) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if price, ok := db.R[item]; ok {
		fmt.Fprintf(w, "%s\n", price)
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item: %q\n", item)
	}
}
