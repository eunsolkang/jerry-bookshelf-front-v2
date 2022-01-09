package handlers

import (
	"jerrybook/database"
	"jerrybook/models"

	"github.com/gofiber/fiber/v2"
)

func GetBookList(c *fiber.Ctx) error{
	books, err := database.Client.Book.Query().All(c.Context())
	if err != nil {
		return err
	}
	return c.JSON(books)
}

func CreateBook(c *fiber.Ctx) error {
	payload := models.Book{}

	if err := c.BodyParser(&payload); err != nil {
		return err
	}
	
	book := database.Client.Book.
		Create().
		SetName(payload.Name).
		SetAuthor(payload.Author).
		SetImageURL(payload.ImageUrl).
		SetReport(payload.Report).
		SetBackgroundColor(payload.BackgroundColor).
		SetRating(payload.Rating).SaveX(c.Context())

	return c.JSON(book)
}