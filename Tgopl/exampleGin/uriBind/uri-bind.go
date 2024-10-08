package main

import (
	"github.com/gin-gonic/gin"
)

// curl -v localhost:8088/thinkerou/987fbc97-4bed-5078-9f07-9141ba07c9f3
// curl -v localhost:8088/thinkerou/not-uuid

type Person1 struct {
	ID   string `uri:"id" binding:"required,uuid"`
	Name string `uri:"name" binding:"required"`
}

func main() {
	route := gin.Default()
	route.GET("/:name/:id", func(ctx *gin.Context) {
		var person Person1
		if err := ctx.ShouldBindUri(&person); err != nil {
			ctx.JSON(400, gin.H{"msg": err})
			return
		}
		ctx.JSON(200, gin.H{"name": person.Name, "uuid": person.ID})
	})
	route.Run(":8088")
}
