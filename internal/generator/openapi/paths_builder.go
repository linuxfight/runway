package openapi

import (
	"strconv"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
)

func addRoute(doc *openapi3.T, r RouteMeta, sb *SchemaBuilder) {
	path := convertPath(r.Path)

	item := doc.Paths.Find(path)
	if item == nil {
		item = &openapi3.PathItem{}
		doc.Paths.Set(path, item)
	}

	op := &openapi3.Operation{
		OperationID: r.ID,
		Summary:     r.Summary,
		Description: r.Description,
		Tags:        r.Tags,
		Parameters:  buildPathParams(r),
		Responses:   buildResponses(r, sb),
	}

	op.Parameters = append(op.Parameters, buildExtraParams(r)...)

	// request body
	if r.Request != nil {
		sb.BuildFromType(r.Request)

		op.RequestBody = &openapi3.RequestBodyRef{
			Value: &openapi3.RequestBody{
				Required: true,
				Content: openapi3.Content{
					"application/json": &openapi3.MediaType{
						Schema: sb.Ref(r.Request.Name),
					},
				},
			},
		}
	}

	switch strings.ToUpper(r.Method) {
	case "GET":
		item.Get = op
	case "POST":
		item.Post = op
	case "PUT":
		item.Put = op
	case "PATCH":
		item.Patch = op
	case "DELETE":
		item.Delete = op
	}
}

func buildResponses(r RouteMeta, sb *SchemaBuilder) *openapi3.Responses {
	responses := openapi3.NewResponses()

	if r.Response == nil {
		responses.Set("204", &openapi3.ResponseRef{
			Value: &openapi3.Response{
				Description: strPtr("No Content"),
			},
		})
		return responses
	}

	sb.BuildFromType(r.Response)

	code := r.StatusCode
	if code == 0 {
		code = 200
	}

	responses.Set(
		strconv.Itoa(code),
		&openapi3.ResponseRef{
			Value: &openapi3.Response{
				Description: strPtr("OK"),
				Content: openapi3.Content{
					"application/json": &openapi3.MediaType{
						Schema: sb.Ref(r.Response.Name),
					},
				},
			},
		},
	)

	return responses
}

func buildPathParams(r RouteMeta) openapi3.Parameters {
	var params openapi3.Parameters

	for _, p := range extractPathParams(r.Path) {
		params = append(params, &openapi3.ParameterRef{
			Value: &openapi3.Parameter{
				Name:     p,
				In:       "path",
				Required: true,
				Schema: &openapi3.SchemaRef{
					Value: &openapi3.Schema{
						Type: &openapi3.Types{"string"},
					},
				},
			},
		})
	}

	return params
}

func buildExtraParams(r RouteMeta) openapi3.Parameters {
	var params openapi3.Parameters

	if r.Request == nil {
		return params
	}

	for _, f := range r.Request.Fields {
		name := f.JSON

		var in string

		switch {
		case f.Query != "":
			in = "query"
			name = f.Query

		case f.Header != "":
			in = "header"
			name = f.Header

		case f.Param != "":
			in = "path"
			name = f.Param

		default:
			continue
		}

		schema := &openapi3.Schema{
			Type: &openapi3.Types{mapPrimitive(f.Type)},
		}

		applyValidation(schema, &f)

		params = append(params, &openapi3.ParameterRef{
			Value: &openapi3.Parameter{
				Name:     strings.ToLower(name),
				In:       in,
				Required: strings.Contains(f.Validate, "required"),
				Schema:   &openapi3.SchemaRef{Value: schema},
			},
		})
	}

	return params
}
