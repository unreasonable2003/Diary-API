package main

import (
	"diary_api/database"
	"diary_api/model"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	loadEnv()
	loadDatabase()
}

func loadDatabase() {
	// Connect to the database
	database.Connect()

	// Perform database migrations
	err := database.Database.AutoMigrate(&model.User{}, &model.Entry{})
	if err != nil {
		log.Fatal("Error performing database migrations:", err)
	}

}

func loadEnv() {
	err := godotenv.Load(".env.local")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
