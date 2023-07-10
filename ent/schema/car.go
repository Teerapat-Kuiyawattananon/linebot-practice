package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Car holds the schema definition for the Car entity.
type Car struct {
	ent.Schema
}

// Fields of the Car.
func (Car) Fields() []ent.Field {
	return []ent.Field{
		field.String("model"),
		field.Time("registered_at").Default(func() time.Time {
			locaBangkok, _ := time.LoadLocation("Asia/Bangkok")
			return time.Now().In(locaBangkok)
		}),
		field.Int("price").Default(0),
		field.String("image_path").Default("https://digitalfinger.id/wp-content/uploads/2019/12/no-image-available-icon-6.png"),
	}
}

// Edges of the Car.
func (Car) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", LineUser.Type).
		Ref("cars").Unique(),
	}
}
