package service

import (
	"context"
	"jerrybook/database"
	"jerrybook/ent"
	"jerrybook/ent/book"
	"jerrybook/models"
	"github.com/google/uuid"
)

type BookServie interface {
	GetAllBook()
	CreateBook()
	DeleteBook()
}

type DefaultBookServie struct {
	client *ent.Client
}

func (d DefaultBookServie) GetAllBook () ([]*ent.Book, error) {
	books, err := d.client.Book.Query().All(context.Background())

	if err != nil {
		return nil, err
	}

	return books, nil
}

func (d DefaultBookServie) CreateBook(payload models.Book) (*ent.Book, error) {

	book := database.Client.Book.
		Create().
		SetName(payload.Name).
		SetAuthor(payload.Author).
		SetImageURL(payload.ImageUrl).
		SetReport(payload.Report).
		SetBackgroundColor(payload.BackgroundColor).
		SetRating(payload.Rating).SaveX(context.Background())
	
	return book, nil;
}

func (d DefaultBookServie) DeleteBook (uuid uuid.UUID) (error) {

	b, err := database.Client.Book.Query().Where(book.UUID(uuid)).Only(context.Background())
	if err != nil {
		return err
	}

	err = database.Client.Book.DeleteOne(b).Exec(context.Background())
	if err != nil {
		return err
	}

	return nil
}

func NewBookService(client *ent.Client) DefaultBookServie{
	return DefaultBookServie{client}
}