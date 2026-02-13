package controllers

import (
	"ms-menu/models"
	"ms-menu/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetMenu(c *gin.Context) {
	items := services.GetAllItems()
	c.JSON(http.StatusOK, items)
}

func AddMenuItem(c *gin.Context) {
	var newItem models.MenuItem
	if err := c.ShouldBindJSON(&newItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := services.CreateItem(newItem)
	c.JSON(http.StatusCreated, result)
}
