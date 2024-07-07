package main

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/mandrigin/gin-spa/spa"
)

func main() {
	// Create a router
	r := gin.Default()
	// Use CORS
	r.Use(cors.Default())

	// Create a route group
	g := r.Group("/api/v1")
	
	// Greate a GET PING endpoint
	g.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	// Create a completion POST endpoint
	g.POST("/completion", func(c *gin.Context) {})
	// Create a rag POST endpoint
	g.POST("/rag", func(c *gin.Context) {})

	// Enable static files
	r.Use(spa.Middleware("/", "./public"))
	
	// Serve
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
