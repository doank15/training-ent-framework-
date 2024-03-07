package schema

import (
	"entgo.io/ent"
)

type Amount float64

// Card holds the schema definition for the Card entity.
type Card struct {
	ent.Schema
}

// Fields of the Card.
// func (Card) Fields() []ent.Field {
// 	return []ent.Field{
// 		field.Float("amount").GoType(Amount(0)),
// 		field.String("name").Optional().GoType(&sql.NullString{}),
// 		field.Float("decimal").
// 		GoType(decimal.Decimal{}).SchemaType(map[string]string{
// 			dialect.MySQL: "decimal(6,2)",
// 			dialect.Postgres: "numeric",
// 		}),
// 	}
// }

// Edges of the Card.
func (Card) Edges() []ent.Edge {
	return nil
}


