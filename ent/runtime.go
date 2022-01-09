// Code generated by entc, DO NOT EDIT.

package ent

import (
	"jerrybook/ent/book"
	"jerrybook/ent/schema"
	"time"

	"github.com/google/uuid"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	bookFields := schema.Book{}.Fields()
	_ = bookFields
	// bookDescUUID is the schema descriptor for uuid field.
	bookDescUUID := bookFields[0].Descriptor()
	// book.DefaultUUID holds the default value on creation for the uuid field.
	book.DefaultUUID = bookDescUUID.Default.(func() uuid.UUID)
	// bookDescCreatedAt is the schema descriptor for created_at field.
	bookDescCreatedAt := bookFields[1].Descriptor()
	// book.DefaultCreatedAt holds the default value on creation for the created_at field.
	book.DefaultCreatedAt = bookDescCreatedAt.Default.(func() time.Time)
	// bookDescUpdatedAt is the schema descriptor for updated_at field.
	bookDescUpdatedAt := bookFields[2].Descriptor()
	// book.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	book.DefaultUpdatedAt = bookDescUpdatedAt.Default.(func() time.Time)
	// book.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	book.UpdateDefaultUpdatedAt = bookDescUpdatedAt.UpdateDefault.(func() time.Time)
	// bookDescBackgroundColor is the schema descriptor for background_color field.
	bookDescBackgroundColor := bookFields[4].Descriptor()
	// book.DefaultBackgroundColor holds the default value on creation for the background_color field.
	book.DefaultBackgroundColor = bookDescBackgroundColor.Default.(string)
}
