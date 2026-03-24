package main

import (
	"movie-app/database"

	"github.com/gin-gonic/gin"
)

func main() {
	database.InitDB()

	r := gin.Default()

	r.LoadHTMLGlob("templates/*")

	r.Static("/static", "./static")

	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})

	r.Run(":8080")
}
