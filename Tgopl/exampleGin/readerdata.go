package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)
func main()  {
	router := gin.Default()
	router.GET("/someDataFromReader",func(ctx *gin.Context) {
		response, err := http.Get("https://raw.githubusercontent.com/gin-
gonic/logo/master/color.png")
if err != nil || response. {
	
}
	})
}