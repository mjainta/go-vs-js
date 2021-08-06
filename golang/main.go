package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)


var db = make(map[string]string)

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	// Return documents
	r.GET("/documents", func(c *gin.Context) {
		var documents = getAllDocuments()
		c.JSON(http.StatusOK, gin.H{"documents": documents})
	})

	// Return filtered documents
	r.GET("/documents/:name", func(c *gin.Context) {
		name := c.Params.ByName("name")
		var documents = getFilteredDocuments(name)
		c.JSON(http.StatusOK, gin.H{"documents": documents})
	})

	return r
}

func main() {
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
