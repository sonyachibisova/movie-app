package main

import (
	"log"
	"movie-app/database"
	"movie-app/models"
)

func main() {
	database.InitDB()

	user := models.User{
		Username: "me",
		Email:    "me@example.com",
	}

	result := database.DB.Where("Username = ?", user.Username).FirstOrCreate(&user)

	if result.Error != nil {
		log.Fatal("Failed to create user:", result.Error)
	}

	log.Printf("User created: ID=%d, Username=%s", user.ID, user.Username)
}
