package main

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

// curl -X GET "http://localhost:8085/testing?\
// name=John&address=123+Main+St&birthday=2000-01-01\
// &createTime=1630385117000000000&unixTime=1630385117"

// curl -X POST "http://localhost:8085/testing" \
// -F "name=John" -F "address=123 Main St" \
// -F "birthday=2000-01-01" -F "createTime=1630385117000000000" \
// -F "unixTime=1630385117"
type Person3 struct {
	Name       string    `form:"name"`
	Address    string    `form:"address"`
	Birthday   time.Time `form:"birthday" time_format:"2006-01-02" time_utc:"1"`
	CreateTime time.Time `form:"createTime" time_format:"unixNano"`
	unixtime   time.Time `form:"unixTime" time_format:"unix"`
}

func main() {
	route := gin.Default()
	route.GET("/testing", startPage)
	route.POST("/testing", startPage)
	route.Run(":8085")
}

func startPage(c *gin.Context) {
	var person Person3

	if c.ShouldBind(&person) == nil {
		log.Println(person.Name)
		log.Println(person.Address)
		log.Println(person.Birthday)
		log.Println(person.CreateTime)
		log.Println(person.unixtime)

	}
	c.String(200, "Success GET or POST")
}
