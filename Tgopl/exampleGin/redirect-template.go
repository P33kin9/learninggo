package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/foo", func(ctx *gin.Context) {
		ctx.Redirect(http.StatusMovedPermanently, "https://www.google.com/")
	})

	router.POST("/test", func(ctx *gin.Context) {
		ctx.Redirect(http.StatusFound, "/foo")
	})

	router.Run(":8080")
}
