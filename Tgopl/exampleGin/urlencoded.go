package main

import (
	"mime/multipart"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProfileForm struct {
	Name   string                `form:"name" binding:"required"`
	Avatar *multipart.FileHeader `form:"avatar" binding:"required"`
}

func main() {
	router := gin.Default()
	router.POST("/profile", func(ctx *gin.Context) {
		var form ProfileForm
		if err := ctx.ShouldBind(&form); err != nil {
			ctx.String(http.StatusBadRequest, "bad request")
			return
		}

		err := ctx.SaveUploadedFile(form.Avatar, form.Avatar.Filename)
		if err != nil {
			ctx.String(http.StatusInternalServerError, "unknown error")
			return
		}
		ctx.String(http.StatusOK, "ok")
	})
	router.Run(":8080")
}
