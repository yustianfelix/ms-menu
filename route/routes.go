package routes

import (
	"your-module/controllers"

	"github.com/gin-gonic/gin"
)

// SetupRouter encapsulates all routing logic
func SetupRouter(menuController *controllers.MenuController) *gin.Engine {
	router := gin.Default()

	// Grouping routes
	api := router.Group("/api/v1")
	{
		menuRoutes := api.Group("/menus")
		{
			menuRoutes.GET("/", menuController.GetAllMenus)
			menuRoutes.POST("/", menuController.CreateMenu)
		}
	}

	return router
}
