package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize Gin router
	r := gin.Default()

	// Define a simple GET route
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, Gin!",
		})
	})

	// Run the server on port 8080
	r.Run(":8080")
}
