package schema

import (
	"time"

	"entgo.io/ent"
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
		field.String("name").Default("noname"),
		field.Enum("signin_with").Values("email", "google", "twitter"),
		field.String("email").Unique(),
		field.String("password").Sensitive(),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now),
		field.Time("logged_out_at").Default(time.Now),
		field.UUID("uuid", uuid.UUID{}).
			Default(uuid.New),
	}
}
