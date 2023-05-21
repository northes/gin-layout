package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"time"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id"),
		field.Time("create_at").Default(time.Now),
		field.Time("update_at").Default(time.Now).UpdateDefault(time.Now),
		field.Int64("age").Positive(),
		field.String("name").Default("unknown"),
		field.Int64("phone").Unique(),
		field.String("password").Sensitive(),
	}
}

func (User) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("phone"),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
