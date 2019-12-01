package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jackyhuang85/gintonic/shorturl"
)

func gethash(c *gin.Context) {
	_url := c.Query("url")
	fmt.Println(_url)
	_res := shorturl.Shorten(_url)
	c.JSON(200, gin.H{
		"message": _res,
	})
}

func main() {
	r := gin.Default()

	// for health test
	r.GET("/hello", func(c *gin.Context) {
		c.Status(200)
	})

	r.GET("/create", gethash)

	// run server
	r.Run()
}
