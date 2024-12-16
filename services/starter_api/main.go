package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Routes
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// Start the server
	r.Run(":8080") // Runs on http://localhost:8080
}
