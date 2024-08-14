package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Static("/assets", "./assets")
	r.StaticFS("/more_static", http.Dir("moredir"))
	r.StaticFile("/favicon.ico", "./resoures/favicon.ico")

	r.Run(":8080")
}
