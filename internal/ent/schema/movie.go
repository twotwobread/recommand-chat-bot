package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Movie holds the schema definition for the Movie entity.
type Movie struct {
	ent.Schema
}

func (Movie) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

// Fields of the Movie.
func (Movie) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").
			NotEmpty().
			StructTag(`json:"title"`),
		field.String("genre").
			NotEmpty().
			StructTag(`json:"genre"`),
		field.String("director").
			NotEmpty().
			StructTag(`json:"director"`),
		field.Strings("actors").
			StructTag(`json:"actors"`),
		field.String("description").
			NotEmpty().
			StructTag(`json:"description"`),
		field.Time("release_date").
			StructTag(`json:"release_date"`),
	}
}

// Edges of the Movie.
func (Movie) Edges() []ent.Edge {
	return nil
}
