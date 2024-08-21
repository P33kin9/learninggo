package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
		log.Printf("endpoint %v %v %v %v", httpMethod, absolutePath, handlerName, nuHandlers)
	}
	r.POST("/foo", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "foo")
	})

	r.GET("/bar", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "bar")
	})

	r.GET("/status", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "ok")
	})
	r.Run()
}
