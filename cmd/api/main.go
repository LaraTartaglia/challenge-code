package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"go.mod/internal/core/services"
	"go.mod/internal/infrastructure/handlers"
	"go.mod/internal/infrastructure/repositories/mongodb"
)

func main() {
	router := gin.Default()

	connection, err := mongodb.GetConnection()
	if err != nil {
		panic("error connecting to db: " + err.Error())
	}

	itemRepository, err := mongodb.NewItemRepository(connection)
	if err != nil {
		panic("error creating item repository: " + err.Error())
	}

	itemService, err := services.NewItemService(itemRepository)
	if err != nil {
		panic("error creating item service: " + err.Error())
	}

	itemHandler, err := handlers.NewItemHandler(itemService)
	if err != nil {
		panic("error connecting to item handler: " + err.Error())
	}

	router.GET("v1/item/:id", itemHandler.GetItems)
	router.POST("v1/items", itemHandler.PostItem)

	router.Run()
	fmt.Println("running")
}
