package main

import (
	"html/template"
	"log"

	"github.com/gin-gonic/gin"
)

var html = template.Must(template.New("https").Parse(`
<html>
<head>
<title>Https TEST</title>
<script src="/assets/app.js"></script>
</head>
<body>
<h1 style="color:red;">Welcome,Ginner!</h1>
</body>
</html>
`))

func main() {
	r := gin.Default()
	r.Static("/assets", "./assets")
	r.SetHTMLTemplate(html)

	r.GET("/", func(ctx *gin.Context) {
		if pusher := ctx.Writer.Pusher(); pusher != nil {
			if err := pusher.Push("/assets/app.js", nil); err != nil {
				log.Printf("Failed to push: %v", err)
			}
		}
		ctx.HTML(200, "https", gin.H{
			"status": "success",
		})
	})
	r.RunTLS(":8080", "./testdata/server.pem", "./testdata/server.key")
}
