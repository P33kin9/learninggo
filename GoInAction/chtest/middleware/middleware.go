package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/urfave/negroni"
	//	"github.com/urfave/negroni"
)

func main() {
	n := negroni.New()
	n.UseFunc(printAuthorInof)

	router := http.NewServeMux()
	router.Handle("/", handler())

	n.UseHandler(router)

	n.Run(":1234")
}

func printAuthorInof(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	fmt.Println("Blog:www.flysnow.org")
	fmt.Println("Wechat:flysnow_org")
	next(rw, r)
}

func handler() http.Handler {
	return http.HandlerFunc(myHandler)
}

func myHandler(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "text/plain")
	io.WriteString(rw, "Hello World")
}
