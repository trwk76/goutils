package spec

type (
	// ExternalDoc
	// [Reference]: https://swagger.io/specification/#external-documentation-object
	ExternalDoc struct {
		URL         string `json:"url" yaml:"url"`
		Description string `json:"description,omitempty" yaml:"description,omitempty"`
	}

	Examples map[string]ExampleOrRef

	ExampleOrRef struct {
		Reference
		Example
	}

	// Example
	// [Reference]: https://swagger.io/specification/#example-object
	Example struct {
		Summary       string `json:"summary,omitempty" yaml:"summary,omitempty"`
		Description   string `json:"description,omitempty" yaml:"description,omitempty"`
		Value         any    `json:"value,omitempty" yaml:"value,omitempty"`
		ExternalValue string `json:"externalValue,omitempty" yaml:"externalValue,omitempty"`
	}

	// Reference
	// [Reference]: https://swagger.io/specification/#reference-object
	Reference struct {
		Ref string `json:"$ref,omitempty" yaml:"$ref,omitempty"`
	}
)
