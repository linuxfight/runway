package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

type {{ .Title }} struct {
	ent.Schema
}

func ({{ .Title }}) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			NotEmpty(),

		field.Time("created_at").
			Default(time.Now),
	}
}
