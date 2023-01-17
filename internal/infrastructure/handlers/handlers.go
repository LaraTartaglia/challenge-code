package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mod/internal/core/domain"
)

type Handlers interface {
	Get(c *gin.Context)
	PostItem(c *gin.Context)
}

type handlers struct {
}

func Get(c *gin.Context) error {
	c.JSON(http.StatusOK, gin.H{
		"message": "hola",
	})
}

func PostItem(c *gin.Context) error {
	c.JSON(http.StatusOK, gin.H{
		"message": "hola",
	})

	var reqBody domain.Item
	err := c.BindJSON(&reqBody)
	if err != nil {
		return err
	}

}
