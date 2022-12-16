package main

import (
	"net/http"

	"internal/infrastructure/handlers/handlers.go"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	router.GET("/hola", handlers.Get)
	router.Run()
}
