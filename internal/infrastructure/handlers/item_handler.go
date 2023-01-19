package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mod/internal/core/ports"
)

type ItemHandler interface {
	GetById(c *gin.Context)
	GetItems(c *gin.Context)
	PostItem(c *gin.Context)
}

type itemHandler struct {
	itemService ports.ItemService
}

func NewItemHandler(itemService ports.ItemService) (ItemHandler, error) {
	if itemService == nil {
		return nil, fmt.Errorf("service cannot be nil")
	}

	return &itemHandler{
		itemService: itemService,
	}, nil
}

func (i *itemHandler) GetById(c *gin.Context) {
	id := c.Param("id")

	err := i.itemService.GetById(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id existente": id})
}

func (i *itemHandler) GetItems(c *gin.Context) {}

func (i *itemHandler) PostItem(c *gin.Context) {
	reqBody := new(Item)

	err := c.ShouldBindJSON(&reqBody)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	itemToSave := reqBody.toDomain()

	err = i.itemService.PostItem(c.Request.Context(), &itemToSave)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, itemToSave)
}
