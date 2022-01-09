package main

import (
	"jerrybook/database"
	"jerrybook/handlers"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
)

func main() {
	database.Connect()

	app := fiber.New()

	app.Post("/book",)
	app.Get("/book", handlers.GetBookList)

	app.Listen(":8001")
}
