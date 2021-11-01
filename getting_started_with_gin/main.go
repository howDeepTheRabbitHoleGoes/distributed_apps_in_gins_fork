package main

import (
	"encoding/xml"

	"github.com/gin-gonic/gin"
)

type Person struct {
	XMLName   xml.Name `xml:"person"`
	FirstName string   `xml:"firstName"`
	LastName  string   `xml:lastName`
}

func IndexHandler(c *gin.Context) {
	// name := c.Params.ByName("name")

	// c.JSON(200, gin.H{
	// 	"message": "hello " + name,
	// })
	c.XML(200,
		Person{
			FirstName: "Vince",
			LastName:  "Carter",
		})
}

func main() {

	router := gin.Default()
	//router.GET("/:name", IndexHandler)
	router.GET("/", IndexHandler)
	router.Run()
}
