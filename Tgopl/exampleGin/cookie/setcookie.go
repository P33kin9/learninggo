package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/cookie", func(ctx *gin.Context) {
		cookie, err := ctx.Cookie("gin_cookies")
		if err != nil {
			cookie = "NotSet"
			ctx.SetCookie("gin_cookies", "test", 3600, "/", "localhost", false, true)
		}
		fmt.Printf("Cookie value: %s \n", cookie)
	})
	r.Run()
}
