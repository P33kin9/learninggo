package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	r.GET("/json", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"html": "<b>Hello, world!/<br>",
		})
	})
	r.GET("purejson", func(ctx *gin.Context) {
		ctx.PureJSON(200, gin.H{
			"html": "<br>HELLO, WORLD!!!/<br>",
		})
	})
	r.Run(":8080")
}
