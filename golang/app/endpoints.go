package app

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
		var documents = getDocuments(bson.M{})
		c.JSON(http.StatusOK, gin.H{"documents": documents})
	})

	// Return filtered documents
	router.GET("/documents/:name", func(c *gin.Context) {
		name := c.Params.ByName("name")
		var filter = bson.M{
			"name": bson.M{
				"$regex": primitive.Regex{
					Pattern: fmt.Sprintf("%s", name),
					Options: "i",
				},
			},
		}
		var documents = getDocuments(filter)
		c.JSON(http.StatusOK, gin.H{"documents": documents})
	})

	return router
}
