package main

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/mandrigin/gin-spa/spa"
)

func main() {
	r := gin.Default()
	r.Use(cors.Default())

	g := r.Group("/api/v1")
  g.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	g.POST("/completion", func(c *gin.Context) {})
	g.POST("/rag", func(c *gin.Context) {})

	r.Use(spa.Middleware("/", "./public"))
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
