package main

import (
	"ms-menu/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Routes point to Controller functions
	router.GET("/menu", controllers.GetMenu)
	router.POST("/menu", controllers.AddMenuItem)

	router.Run(":8080")
}
