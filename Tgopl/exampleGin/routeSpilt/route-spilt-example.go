package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// MyBenchLogger 中间件
func MyBenchLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 这里可以添加日志逻辑
		// fmt.Println("MyBenchLogger 中间件已被执行")
		c.Next()
	}
}

// AuthRequired 中间件
func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 这里可以添加授权逻辑
		// 假设所有请求都被授权
		c.Next()
	}
}

// 处理函数示例
func benchEndpoint(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Benchmark endpoint",
	})
}

func loginEndpoint(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Login endpoint",
	})
}

func submitEndpoint(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Submit endpoint",
	})
}

func readEndpoint(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Read endpoint",
	})
}

func analysisEndpoint(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Analytics endpoint",
	})
}

func main() {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/benchmark", MyBenchLogger(), benchEndpoint)

	authorized := r.Group("/")
	authorized.Use(AuthRequired())
	{
		authorized.POST("/login", loginEndpoint)
		authorized.POST("/submit", submitEndpoint)
		authorized.POST("/read", readEndpoint)

		testing := authorized.Group("testing")
		testing.GET("/analytics", analysisEndpoint)
	}

	r.Run(":8080") // 启动服务器，监听8080端口
}
