package main

import (
	"log"
	"your-module/controllers"
	"your-module/routes"
	"your-module/services"
)

func main() {
	// 1. Initialize dependencies (DB, Services, Controllers)
	menuService := services.NewMenuService()
	menuController := controllers.NewMenuController(menuService)

	// 2. Setup Router
	r := routes.SetupRouter(menuController)

	// 3. Start Server
	log.Println("Starting server on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
