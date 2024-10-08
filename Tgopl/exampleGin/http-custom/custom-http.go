package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// func main() {
// 	router := gin.Default()
// 	router.GET("/", func(ctx *gin.Context) {
// 		ctx.JSON(http.StatusOK, "hello, world!")
// 	})
// 	http.ListenAndServe(":8080", router)
// }

func main() {
	r := gin.Default()
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "hello world!!!!")
	})
	s := &http.Server{
		Addr:           ":8080",
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
