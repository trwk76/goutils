package spec

type (
	// Paths
	// [Reference]: https://swagger.io/specification/#paths-object
	Paths map[string]PathItem

	// PathItem
	// [Reference]: https://swagger.io/specification/#path-item-object
	PathItem struct {
		Reference
		GET        *Operation `json:"get,omitempty" yaml:"get,omitempty"`
		PUT        *Operation `json:"put,omitempty" yaml:"put,omitempty"`
		POST       *Operation `json:"post,omitempty" yaml:"post,omitempty"`
		DELETE     *Operation `json:"delete,omitempty" yaml:"delete,omitempty"`
		OPTIONS    *Operation `json:"options,omitempty" yaml:"options,omitempty"`
		HEAD       *Operation `json:"head,omitempty" yaml:"head,omitempty"`
		PATCH      *Operation `json:"patch,omitempty" yaml:"patch,omitempty"`
		TRACE      *Operation `json:"trace,omitempty" yaml:"trace,omitempty"`
		Parameters Parameters `json:"parameters,omitempty" yaml:"parameters,omitempty"`
		Servers    Servers    `json:"servers,omitempty" yaml:"servers,omitempty"`
	}

	// Operation
	// [Reference]: https://swagger.io/specification/#operation-object
	Operation struct {
		OperationID  string                   `json:"operationId,omitempty" yaml:"operationId,omitempty"`
		Summary      string                   `json:"summary,omitempty" yaml:"summary,omitempty"`
		Description  string                   `json:"description,omitempty" yaml:"description,omitempty"`
		Deprecated   bool                     `json:"deprecated,omitempty" yaml:"deprecated,omitempty"`
		ExternalDocs *ExternalDoc             `json:"externalDocs,omitempty" yaml:"externalDocs,omitempty"`
		Tags         []string                 `json:"tags,omitempty" yaml:"tags,omitempty"`
		Parameters   Parameters               `json:"parameters,omitempty" yaml:"parameters,omitempty"`
		RequestBody  *RequestBodyOrRef        `json:"requestBody,omitempty" yaml:"requestBody,omitempty"`
		Responses    Responses                `json:"responses,omitempty" yaml:"responses,omitempty"`
		Security     SecurityRequirements     `json:"security,omitempty" yaml:"security,omitempty"`
		Callbacks    map[string]CallbackOrRef `json:"callbacks,omitempty" yaml:"callbacks,omitempty"`
		Servers      Servers                  `json:"servers,omitempty" yaml:"servers,omitempty"`
	}

	Parameters []ParameterOrRef

	ParameterOrRef struct {
		Reference
		Parameter
	}

	// Parameter
	// [Reference]: https://swagger.io/specification/#parameter-object
	Parameter struct {
		Name            string      `json:"name" yaml:"name"`
		In              ParameterIn `json:"in" yaml:"in"`
		Summary         string      `json:"summary,omitempty" yaml:"summary,omitempty"`
		Description     string      `json:"description,omitempty" yaml:"description,omitempty"`
		Required        bool        `json:"required,omitempty" yaml:"required,omitempty"`
		Deprecated      bool        `json:"deprecated,omitempty" yaml:"deprecated,omitempty"`
		AllowEmptyValue bool        `json:"allowEmptyValue,omitempty" yaml:"allowEmptyValue,omitempty"`
		Style           string      `json:"style,omitempty" yaml:"style,omitempty"`                 // Simple
		Explode         bool        `json:"explode,omitempty" yaml:"explode,omitempty"`             // Simple
		AllowReserved   bool        `json:"allowReserved,omitempty" yaml:"allowReserved,omitempty"` // Simple
		Schema          SchemaOrRef `json:"schema,omitempty" yaml:"schema,omitempty"`               // Simple
		Example         any         `json:"example,omitempty" yaml:"example,omitempty"`             // Simple
		Examples        Examples    `json:"examples,omitempty" yaml:"examples,omitempty"`           // Simple
		Content         MediaTypes  `json:"content,omitempty" yaml:"content,omitempty"`             // Complex
	}

	ParameterIn string

	RequestBodyOrRef struct {
		Reference
		RequestBody
	}

	// RequestBody
	// [Reference]: https://swagger.io/specification/#request-body-object
	RequestBody struct {
		Content     MediaTypes `json:"content" yaml:"content"`
		Description string     `json:"description,omitempty" yaml:"description,omitempty"`
		Required    bool       `json:"required,omitempty" yaml:"required,omitempty"`
	}

	// Responses
	// [Reference]: https://swagger.io/specification/#responses-object
	Responses map[string]ResponseOrRef

	ResponseOrRef struct {
		Reference
		Response
	}

	// Response
	// [Reference]: https://swagger.io/specification/#response-object
	Response struct {
		Description string     `json:"description" yaml:"description"`
		Headers     Headers    `json:"headers,omitempty" yaml:"headers,omitempty"`
		Content     MediaTypes `json:"content,omitempty" yaml:"content,omitempty"`
		Links       Links      `json:"links,omitempty" yaml:"links,omitempty"`
	}

	Headers map[string]HeaderOrRef

	HeaderOrRef struct {
		Reference
		Header
	}

	// Header
	// [Reference]: https://swagger.io/specification/#header-object
	Header struct {
		Description     string      `json:"description,omitempty" yaml:"description,omitempty"`
		Required        bool        `json:"required,omitempty" yaml:"required,omitempty"`
		Deprecated      bool        `json:"deprecated,omitempty" yaml:"deprecated,omitempty"`
		AllowEmptyValue bool        `json:"allowEmptyValue,omitempty" yaml:"allowEmptyValue,omitempty"`
		Style           string      `json:"style,omitempty" yaml:"style,omitempty"`                 // Simple
		Explode         bool        `json:"explode,omitempty" yaml:"explode,omitempty"`             // Simple
		AllowReserved   bool        `json:"allowReserved,omitempty" yaml:"allowReserved,omitempty"` // Simple
		Schema          SchemaOrRef `json:"schema,omitempty" yaml:"schema,omitempty"`               // Simple
		Example         any         `json:"example,omitempty" yaml:"example,omitempty"`             // Simple
		Examples        Examples    `json:"examples,omitempty" yaml:"examples,omitempty"`           // Simple
		Content         MediaTypes  `json:"content,omitempty" yaml:"content,omitempty"`             // Complex
	}

	MediaTypes map[string]MediaType

	// MediaType
	// [Reference]: https://swagger.io/specification/#media-type-object
	MediaType struct {
		Schema   SchemaOrRef `json:"schema,omitempty" yaml:"schema,omitempty"`
		Example  any         `json:"example,omitempty" yaml:"example,omitempty"`
		Examples Examples    `json:"examples,omitempty" yaml:"examples,omitempty"`
		Encoding Encodings   `json:"encodings,omitempty" yaml:"encodings,omitempty"`
	}

	Encodings map[string]Encoding

	// Encoding
	// [Reference]: https://swagger.io/specification/#encoding-object
	Encoding struct {
		ContentType   string  `json:"contentType,omitempty" yaml:"contentType,omitempty"`
		Headers       Headers `json:"headers,omitempty" yaml:"headers,omitempty"`
		Style         string  `json:"style,omitempty" yaml:"style,omitempty"`
		Explode       bool    `json:"explode,omitempty" yaml:"explode,omitempty"`
		AllowReserved bool    `json:"allowReserved,omitempty" yaml:"allowReserved,omitempty"`
	}
)

const (
	ParameterInCookie ParameterIn = "cookie"
	ParameterInHeader ParameterIn = "header"
	ParameterInPath   ParameterIn = "path"
	ParameterInQuery  ParameterIn = "query"
)
