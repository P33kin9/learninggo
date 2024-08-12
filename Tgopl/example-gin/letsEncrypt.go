package main

import (
	"log"

	"github.com/gin-gonic/autotls"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/acme/autocert"
)

// func main() {
// 	r := gin.Default()
// 	r.GET("/ping", func(ctx *gin.Context) {
// 		ctx.String(200, "pong")
// 	})
// 	log.Fatal(autotls.Run(r, "0.0.0.0", "example2.com"))
// }

func main() {
	r := gin.Default()
	r.GET("/ping", func(ctx *gin.Context) {
		ctx.String(200, "Pong")
	})
	m := autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		HostPolicy: autocert.HostWhitelist("example2.com", "example1.com"),
		Cache:      autocert.DirCache("/var/www/.cache"),
	}
	log.Fatal(autotls.RunWithManager(r, &m))
}
