package main

import (
	"go-fiber-rest-api/database"
	"go-fiber-rest-api/models"
)

func main() {
	database.DB.AutoMigrate(&models.Book{})
}
