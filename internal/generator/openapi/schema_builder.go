package openapi

import (
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
)

type SchemaBuilder struct {
	schemas map[string]*openapi3.SchemaRef
}

func NewSchemaBuilder() *SchemaBuilder {
	return &SchemaBuilder{
		schemas: map[string]*openapi3.SchemaRef{},
	}
}

func (b *SchemaBuilder) Components() openapi3.Schemas {
	return b.schemas
}

func (b *SchemaBuilder) Ref(name string) *openapi3.SchemaRef {
	return &openapi3.SchemaRef{
		Ref: "#/components/schemas/" + name,
	}
}

func (b *SchemaBuilder) BuildFromType(t *TypeMeta) {
	if t == nil {
		return
	}

	if _, exists := b.schemas[t.Name]; exists {
		return
	}

	schema := &openapi3.Schema{
		Type:       &openapi3.Types{"object"},
		Properties: openapi3.Schemas{},
	}

	requiredFields := []string{}

	for _, f := range t.Fields {
		propName := f.JSON
		if propName == "" {
			propName = strings.ToLower(f.Name)
		}

		prop := b.fieldSchema(&f)

		if prop.Value != nil {
			if applyValidation(prop.Value, &f) {
				requiredFields = append(requiredFields, propName)
			}
		}

		schema.Properties[propName] = prop
	}

	if len(requiredFields) > 0 {
		schema.Required = requiredFields
	}

	b.schemas[t.Name] = &openapi3.SchemaRef{
		Value: schema,
	}
}

func (b *SchemaBuilder) fieldSchema(f *FieldMeta) *openapi3.SchemaRef {

	// ARRAY
	if isSlice(f.Type) {

		elemType := sliceElem(f.Type)

		// nested object
		if f.Nested != nil {
			b.BuildFromType(f.Nested)

			return &openapi3.SchemaRef{
				Value: &openapi3.Schema{
					Type:  &openapi3.Types{"array"},
					Items: b.Ref(f.Nested.Name),
				},
			}
		}

		// primitive array
		return &openapi3.SchemaRef{
			Value: &openapi3.Schema{
				Type: &openapi3.Types{"array"},
				Items: &openapi3.SchemaRef{
					Value: &openapi3.Schema{
						Type: &openapi3.Types{mapPrimitive(elemType)},
					},
				},
			},
		}
	}

	// OBJECT
	if f.Nested != nil {
		b.BuildFromType(f.Nested)
		return b.Ref(f.Nested.Name)
	}

	// PRIMITIVE
	s := &openapi3.Schema{
		Type: &openapi3.Types{mapPrimitive(f.Type)},
	}

	return &openapi3.SchemaRef{
		Value: s,
	}
}

func mapPrimitive(t string) string {
	switch {
	case strings.Contains(t, "string"):
		return "string"
	case strings.Contains(t, "int"):
		return "integer"
	case strings.Contains(t, "float"):
		return "number"
	case strings.Contains(t, "bool"):
		return "boolean"
	default:
		return "string"
	}
}

func applyValidation(s *openapi3.Schema, f *FieldMeta) (required bool) {
	if f.Validate == "" {
		return false
	}

	rules := strings.Split(f.Validate, ",")

	for _, r := range rules {
		switch {
		case r == "required":
			required = true

		case strings.HasPrefix(r, "min="):
			if isString(s) {
				s.MinLength = parseUint(r[4:])
			}

		case strings.HasPrefix(r, "max="):
			if isString(s) {
				s.MaxLength = uint64Ptr(parseUint(r[4:]))
			}

		case strings.HasPrefix(r, "gte="):
			s.Min = float64Ptr(parseFloat(r[4:]))

		case strings.HasPrefix(r, "lte="):
			s.Max = float64Ptr(parseFloat(r[4:]))

		case r == "email":
			s.Format = "email"
		}
	}

	return
}

func isSlice(t string) bool {
	return strings.HasPrefix(t, "[]")
}

func sliceElem(t string) string {
	return strings.TrimPrefix(t, "[]")
}
