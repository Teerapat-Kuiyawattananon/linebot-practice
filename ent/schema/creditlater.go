package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// CreditLater holds the schema definition for the CreditLater entity.
type CreditLater struct {
	ent.Schema
}

// Fields of the CreditLater.
func (CreditLater) Fields() []ent.Field {
	return []ent.Field{
		field.String("transaction_ref").NotEmpty(),
		field.String("date").Default(time.Now().Format("02/01/2006 15:04")),
		field.String("branch").Default("Center"),
		field.Int("amount").Default(0),
		field.Int("installment").Default(0),
		field.String("detail").Default("-"),
	}
}

// Edges of the CreditLater.
func (CreditLater) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", LineUser.Type).
			Ref("creditlaters").Unique(),
	}
}
