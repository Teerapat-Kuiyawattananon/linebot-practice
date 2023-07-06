package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// LineLog holds the schema definition for the LineLog entity.
type LineLog struct {
	ent.Schema
}

// Fields of the LineLog.
func (LineLog) Fields() []ent.Field {
	return []ent.Field{
		field.String("action").Default("unknown"),
		field.String("message").Default("unknown"),
		field.Time("created_at").Default(time.Now()),
	}
}

// Edges of the LineLog.
func (LineLog) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", LineUser.Type).
			Ref("linelogs").Unique(),
	}
}
