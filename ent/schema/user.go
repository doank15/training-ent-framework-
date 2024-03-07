package schema

import (
	"net/url"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Unique(),
		field.String("password").Sensitive(),
		field.Int("age").Positive(),
		field.Bool("active").Default(false),
		field.Time("created_at").Default(time.Now()),
		field.JSON("url", &url.URL{}).Optional(),
		field.Enum("state").Values("on", "off").Optional(),
		field.UUID("id", uuid.UUID{}).Default(uuid.New).StorageKey("oid"),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("cars", Car.Type),
		edge.From("groups", Group.Type).Ref("users"),
	}
}
