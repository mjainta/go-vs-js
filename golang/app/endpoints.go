package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
)


func SetupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	router := gin.Default()

	// Ping test
	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	// Return documents
	router.GET("/documents", func(c *gin.Context) {
		var documents = getAllDocuments()
		c.JSON(http.StatusOK, gin.H{"documents": documents})
	})

	// Return filtered documents
	router.GET("/documents/:name", func(c *gin.Context) {
		name := c.Params.ByName("name")
		var documents = getFilteredDocuments(name)
		c.JSON(http.StatusOK, gin.H{"documents": documents})
	})

	return router
}
