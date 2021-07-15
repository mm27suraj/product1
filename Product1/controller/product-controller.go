package controller

import (
	"Product1/model"
	"Product1/service"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func FetchProductsFromInventory(c *gin.Context) {
	var Inventorys []model.Inventory
	err := service.FetchProductsFromInventory(&Inventorys)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, Inventorys)
	}
}

func AddProductToInventory(c *gin.Context) {
	var Inventory model.Inventory

	c.BindJSON(&Inventory)
	log.Printf(" Prod to be Created %v", Inventory)
	err := service.AddProductToInventory(&Inventory)
	if err != nil {
		log.Println("Error", err)
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		log.Printf("Inventory Created %v", Inventory)
		c.JSON(http.StatusOK, Inventory)

	}
}

func GetInventory(c *gin.Context) {
	id := c.Params.ByName("id")
	var Inventory model.Inventory
	err := service.GetInventory(&Inventory, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, Inventory)
	}
}

func UpdateInventory(c *gin.Context) {
	var Inventory model.Inventory
	id := c.Params.ByName("id")
	err := service.GetInventory(&Inventory, id)
	if err != nil {
		c.JSON(http.StatusNotFound, Inventory)
	}
	c.BindJSON(&Inventory)
	err = service.UpdateInventory(&Inventory, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, Inventory)
	}
}

func DeleteInventory(c *gin.Context) {
	var Inventory model.Inventory
	id := c.Params.ByName("id")
	err := service.DeleteInventory(&Inventory, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id:" + id: "deleted"})
	}
}

func GetProductDetailsByName(c *gin.Context) {
	name := c.Params.ByName("prodname")
	var inv []model.Inventory
	err := service.GetProductDetailsByName(&inv, name)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, inv)
	}
}

func GetProductDetailsById(c *gin.Context) {
	id := c.Params.ByName("prodid")
	var inv []model.Inventory
	err := service.GetProductDetailsById(&inv, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, inv)
	}
}
