package main

import (
	"fmt"
	"log"

	"github.com/devanand100/eCommerce/apis"
	"github.com/devanand100/eCommerce/apis/user"
	"github.com/devanand100/eCommerce/database"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	app := fiber.New(fiber.Config{
		Immutable: true,
	})
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database.DbConnect()

	database.DB.AutoMigrate(user.User{})
	fmt.Println("ased",&user.User{})
	apis.Setup(app)
	app.Listen(":1234")

}
