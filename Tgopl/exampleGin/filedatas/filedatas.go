package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/local/file", func(ctx *gin.Context) {
		ctx.File("local/file.go")
	})
	var fs http.FileSystem = http.Dir("./moredir")
	router.GET("/fs/:file", func(ctx *gin.Context) {
		filename := ctx.Param("file")
		ctx.FileFromFS(filename, fs)

	})
	router.Run(":8080")
}
