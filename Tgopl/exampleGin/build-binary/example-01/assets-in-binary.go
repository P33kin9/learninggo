package main

import (
	"html/template"
	"io"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	t, err := loadTemplate()
	if err != nil {
		panic(err)
	}
	r.SetHTMLTemplate(t)

	r.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "/html/index.tmpl", gin.H{"Foo": "World"})
	})

	r.GET("/bar", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "/html/bar.tmpl", gin.H{"Bar": "World"})
	})
	r.Run(":8088")
}

func loadTemplate() (*template.Template, error) {
	t := template.New("")
	for name, file := range Assets.Files {
		if file.IsDir() || !strings.HasSuffix(name, ".tmpl") {
			continue
		}
		h, err := io.ReadAll(file)
		if err != nil {
			return nil, err
		}
		t, err = t.New(name).Parse(string(h))
		if err != nil {
			return nil, err
		}
	}
	return t, nil
}