package main

import (
	"fmt"
	_ "net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	// router.GET("/user/:name", func(c *gin.Context) {
	// 	name := c.Param("name")
	// 	c.String(http.StatusOK, "Hello %s", name)
	// })

	// router.GET("/user/:name/*action", func(c *gin.Context) {
	// 	name := c.Param("name")
	// 	action := c.Param("action")
	// 	message := name + " is " + action
	// 	c.String(http.StatusOK, message)
	// })

	// router.POST("/user/:name/*action", func(c *gin.Context) {
	// 	c.FullPath()
	// })

	// router.POST("/form_post", func(c *gin.Context) {
	// 	message := c.PostForm("message")
	// 	nick := c.DefaultPostForm("nick", "anonymous")
	// 	c.JSON(200, gin.H{
	// 		"status":  "posted",
	// 		"message": message,
	// 		"nick":    nick,
	// 	})
	// })

	// router.POST("/post", func(c *gin.Context) {
	// 	id := c.Query("id")
	// 	page := c.DefaultQuery("page", "0")
	// 	name := c.PostForm("name")
	// 	message := c.PostForm("message")
	// 	fmt.Printf("id: %s; page: %s; name: %s; message: %s", id, page, name,
	// 		message)
	// })

	// router.POST("/post", func(c *gin.Context) {
	// 	ids := c.QueryMap("ids")
	// 	names := c.PostFormMap("names")
	// 	fmt.Printf("ids: %v; names: %v", ids, names)
	// })

	router.POST("/post", func(c *gin.Context) {
		ids := c.QueryMap("ids")
		names := c.PostFormMap("names")
		fmt.Printf("ids: %v; names: %v", ids, names)
	})
	router.Run(":8080")
}
