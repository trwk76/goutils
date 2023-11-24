package spec

type (
	Schema struct {
		Reference
		Type SchemaType `json:"type,omitempty" yaml:"type,omitempty"`
	}

	SchemaType string
)

const (
	SchemaTypeBoolean SchemaType = "boolean"
	SchemaTypeInteger SchemaType = "integer"
	SchemaTypeNumber  SchemaType = "number"
	SchemaTypeString  SchemaType = "string"
	SchemaTypeArray   SchemaType = "array"
	SchemaTypeObject  SchemaType = "object"
)
