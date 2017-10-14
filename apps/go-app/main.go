package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

var DB = make(map[string]string)

func main() {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"response": "Hello World", "by": "go-app"})
	})

	// Listen and Server in 0.0.0.0:80
	port := os.Getenv("PORT")
	if port == "" {
		port = "80"
	}

	r.Run(fmt.Sprintf(":%s", port))
}
