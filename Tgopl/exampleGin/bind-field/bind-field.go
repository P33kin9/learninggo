package main

import "github.com/gin-gonic/gin"

type structA struct {
	FieldA string `form:"field_a"`
}

type structB struct {
	NestedStruct structA
	FieldB       string `form:"field_b"`
}

type structC struct {
	NestedStructPointer *structA
	FieldC              string `form:"field_c"`
}

type structD struct {
	NestedAnyStruct struct {
		FieldX string `form:"field_x"`
	}
	FieldD string `form:"field_d"`
}

func GetDataB(c *gin.Context) {
	var b structB
	c.Bind(&b)
	c.JSON(200, gin.H{
		"a": b.NestedStruct,
		"b": b.FieldB,
	})
}

func GetDataC(c *gin.Context) {
	var b structC
	c.Bind(&b)
	c.JSON(200, gin.H{
		"a": b.NestedStructPointer,
		"c": b.FieldC,
	})
}

func GetDataD(c *gin.Context) {
	var b structD
	c.Bind(&b)
	c.JSON(200, gin.H{
		"x": b.NestedAnyStruct,
		"d": b.FieldD,
	})
}

func main() {
	r := gin.Default()
	r.GET("/getb", GetDataB)
	r.GET("/getc", GetDataC)
	r.GET("/getd", GetDataD)

	r.Run()
}
