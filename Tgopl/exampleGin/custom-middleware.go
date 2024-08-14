package main

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		t := time.Now()
		ctx.Set("example", "12345")
		ctx.Next()
		latency := time.Since(t)
		log.Println(latency)
	}
}

func main() {
	router := gin.New()
	router.Use(Logger())
	router.GET("/test", func(ctx *gin.Context) {
		example := ctx.MustGet("example").(string)
		log.Println(example)
	})

	router.Run(":8080")
}
