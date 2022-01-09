package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Book holds the schema definition for the Book entity.
type Book struct {
	ent.Schema
}

// Fields of the Book.
func (Book) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("uuid", uuid.UUID{}).Default(uuid.New),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
        field.String("name"),
		field.String("background_color").Default("#fff"),
		field.String("author"),
		field.String("image_url").Optional(),
		field.Text("report"),
		field.Float("rating"),
    }
}

// Edges of the Book.
func (Book) Edges() []ent.Edge {
	return nil
}
