package app

import (
	"jerrybook/database"
	"jerrybook/service"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func Start() {
	database.Connect()

	app := fiber.New()

	app.Use(logger.New())

	app.Use(cors.New(cors.Config{
		AllowOrigins: "https://book.jerrykang.com, https://host.jerrykang.com",
		AllowHeaders: "Origin, Content-Type, Accept",
		AllowMethods: "GET,POST,PUT,DELETE",
	}))

	bh := BookHandler{service.NewBookService(database.Client)}
	app.Get("/book", bh.getBookList)
	app.Post("/book", bh.createBook)
	app.Put("/book/:uuid", bh.updateBook)
	app.Delete("/book/:uuid", bh.deleteBook)


	app.Listen(":8001")
}
