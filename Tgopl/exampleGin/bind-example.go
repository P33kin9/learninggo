package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Login struct {
	User     string `form:"user" json:"user" xml:"user" binding:"required"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
}

func main() {
	router := gin.Default()

	//JSON binding example ({"user": "manu","password": "123"})
	// 	curl -v -X POST \
	//   http://localhost:8080/loginXML \
	//   -H 'content-type: application/json' \
	//   -d '{ "user": "manu", "password": "123" }'
	router.POST("/loginJSON", func(ctx *gin.Context) {
		var json Login
		if err := ctx.ShouldBindJSON(&json); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "unauthorized"})
			return
		}
		if json.User != "manu" || json.Password != "123" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
		}
		ctx.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
	})
	// XML binding example
	//curl -v -X POST "http://localhost:8080/loginXML" \
	// -H "Content-Type: application/xml" \
	// -d '<Login><user>manu</user><password>123</password></Login>'
	router.POST("loginXML", func(ctx *gin.Context) {
		var xml Login
		if err := ctx.ShouldBindXML(&xml); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
	})
	// HTML binding example
	//curl -v -X POST http://localhost:8080/loginForm\?user\=manu\&password\=123
	router.POST("loginForm", func(ctx *gin.Context) {
		var form Login
		if err := ctx.ShouldBind(&form); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"status": "unauthorized"})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
	})
	router.Run(":8080")
}
