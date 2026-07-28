package services

import "ms-menu/models"

// In-Memory store
var menu = []models.MenuItem{
	{ID: "1", Name: "Coffee", Price: 3.50},
}

func GetAllItems() []models.MenuItem {
	return menu
}

func CreateItem(item models.MenuItem) models.MenuItem {
	menu = append(menu, item)
	return item
}
