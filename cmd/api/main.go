package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mod/internal/infrastructure/handlers"
)

func main() {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	router.GET("/hola", handlers.Get)
	router.POST("v1/items", handlers.PostItem)
	router.Run()
}
