package schema

import (
	"errors"
	"net/mail"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Validate(func(s string) error {
			if strings.TrimSpace(s) == "" {
				return errors.New("name required")
			}
			return nil
		}),
		field.String("email").Unique().Validate(func(s string) error {
			_, err := mail.ParseAddress(s)
			return err
		}),
		field.Enum("role").Values("admin", "employee"),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}

func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}
