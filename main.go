package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type MenuItem struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

var menu = []MenuItem{
	{ID: "1", Name: "Coffee", Price: 35000},
	{ID: "2", Name: "Tea", Price: 15000},
}

func main() {
	// 1. Initialize the Gin router
	router := gin.Default()

	// 2. Define Routes
	router.GET("/menu", func(c *gin.Context) {
		c.JSON(http.StatusOK, menu)
	})

	router.POST("/menu", func(c *gin.Context) {
		var newItem MenuItem
		// Bind JSON request body to MenuItem struct
		if err := c.ShouldBindJSON(&newItem); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		menu = append(menu, newItem)
		c.JSON(http.StatusCreated, newItem)
	})

	// 3. Start the server (Defaults to :8080)
	router.Run()
}
