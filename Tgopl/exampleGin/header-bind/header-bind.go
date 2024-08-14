package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type testHeader struct {
	Rate   int    `header:"Rate"`
	Domain string `header:"Domain"`
}

func main() {
	r := gin.Default()
	r.GET("/", func(ctx *gin.Context) {
		h := testHeader{}

		if err := ctx.ShouldBindHeader(&h); err != nil {
			ctx.JSON(200, err)
		}

		fmt.Printf("%#v\n", h)
		ctx.JSON(200, gin.H{"Rate": h.Rate, "Domain": h.Domain})
	})
	r.Run()
}

// client
// curl -H "rate:300" -H "domain:music" 127.0.0.1:8080/
// output
// {"Domain":"music","Rate":300}
