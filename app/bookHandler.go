package app

import (
	"jerrybook/models"
	"jerrybook/service"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type BookHandler struct {
	service service.DefaultBookServie
}

func (bh BookHandler) getBookList(c *fiber.Ctx) error {
	books, err := bh.service.GetAllBook()

	if err != nil {
		return err
	}

	return c.JSON(books)
}

func (bh BookHandler) createBook(c *fiber.Ctx) error {
	payload := models.Book{}

	if err := c.BodyParser(&payload); err != nil {
		return err
	}

	book, err := bh.service.CreateBook(payload)

	if err != nil {
		return err;
	}

	return c.JSON(book)
}

func (bh BookHandler) deleteBook(c *fiber.Ctx) error {

	uuid, err := uuid.Parse(c.Params("uuid"))

	if err != nil {
		return err
	}

	err = bh.service.DeleteBook(uuid)

	if err != nil {
		return err
	}

	return c.SendString("ok")
}