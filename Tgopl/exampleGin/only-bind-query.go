package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

type Person2 struct {
	Name    string `form:"name"`
	Address string `form:"address"`
}

func main() {
	route := gin.Default()
	route.Any("/testing", startPage1)
	route.Run(":8085")
}

func startPage1(c *gin.Context) {
	var person Person2
	if c.ShouldBindQuery(&person) == nil {
		log.Println("====Only By Query String=====")
		log.Println(person.Name)
		log.Println(person.Address)
	}
	c.String(200, "Success")
}
