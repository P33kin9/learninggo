package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/JSONP", func(ctx *gin.Context) {
		data := gin.H{
			"foo": "bar",
		}
		ctx.JSONP(http.StatusOK, data)
	})
	r.Run(":8080")
}
