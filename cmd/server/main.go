package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// create simple server
	r := gin.Default()

	// define route
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})

	// run server
	r.Run(":80")
}
