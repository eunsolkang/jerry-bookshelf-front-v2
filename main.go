package main

import (
	"jerrybook/database"
	"jerrybook/handlers"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	database.Connect()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://book.jerrykang.com",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	app.Post("/book", handlers.CreateBook)
	app.Delete("/book/:uuid", handlers.DeleteBook)
	app.Get("/book", handlers.GetBookList)

	app.Listen(":8001")
}
