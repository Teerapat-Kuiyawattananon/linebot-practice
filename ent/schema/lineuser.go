package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// LineUser holds the schema definition for the LineUser entity.
type LineUser struct {
	ent.Schema
}

// Fields of the LineUser.
func (LineUser) Fields() []ent.Field {
	return []ent.Field{
		field.String("userId").Unique().NotEmpty(),
		field.String("displyaName"),
		field.Time("registered_at").Default(time.Now()),
		field.Bool("active").Default(false),
	}
}

// Edges of the LineUser.
func (LineUser) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("linelogs", LineLog.Type).
		Annotations(entsql.OnDelete(entsql.Cascade)),

		// Buy Pay Later
		edge.To("creditlaters", CreditLater.Type).
		Annotations(entsql.OnDelete(entsql.Cascade)).
			Unique(),
	}
}
