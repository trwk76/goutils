package spec

type (
	// Specification
	// [Reference]: https://swagger.io/specification/#openapi-object
	Specification struct {
		OpenAPI           string               `json:"openapi" yaml:"openapi"`
		Info              Info                 `json:"info" yaml:"info"`
		JSONSchemaDialect string               `json:"jsonSchemaDialect,omitempty" yaml:"jsonSchemaDialect,omitempty"`
		Servers           Servers              `json:"servers,omitempty" yaml:"servers,omitempty"`
		Paths             Paths                `json:"paths" yaml:"paths"`
		WebHooks          WebHooks             `json:"webhooks,omitempty" yaml:"webhooks,omitempty"`
		Components        Components           `json:"components,omitempty" yaml:"components,omitempty"`
		Security          SecurityRequirements `json:"security,omitempty" yaml:"security,omitempty"`
		Tags              Tags                 `json:"tags,omitempty" yaml:"tags,omitempty"`
		ExternalDocs      *ExternalDoc         `json:"externalDocs,omitempty" yaml:"externalDocs,omitempty"`
	}

	// Info
	// [Reference]: https://swagger.io/specification/#info-object
	Info struct {
		Title          string   `json:"title" yaml:"title"`
		Version        string   `json:"version" yaml:"version"`
		Summary        string   `json:"summary,omitempty" yaml:"summary,omitempty"`
		Description    string   `json:"description,omitempty" yaml:"description,omitempty"`
		TermsOfService string   `json:"termsOfService,omitempty" yaml:"termsOfService,omitempty"`
		Contact        *Contact `json:"contact,omitempty" yaml:"contact,omitempty"`
		License        *License `json:"license,omitempty" yaml:"license,omitempty"`
	}

	// Contact
	// [Reference]: https://swagger.io/specification/#contact-object
	Contact struct {
		Name  string `json:"name,omitempty" yaml:"name,omitempty"`
		URL   string `json:"url,omitempty" yaml:"url,omitempty"`
		Email string `json:"email,omitempty" yaml:"email,omitempty"`
	}

	// License
	// [Reference]: https://swagger.io/specification/#license-object
	License struct {
		Name       string `json:"name" yaml:"name"`
		Identifier string `json:"identifier,omitempty" yaml:"identifier,omitempty"`
		URL        string `json:"url,omitempty" yaml:"url,omitempty"`
	}

	Servers []Server

	// Server
	// [Reference]: https://swagger.io/specification/#server-object
	Server struct {
		URL         string          `json:"url" yaml:"url"`
		Description string          `json:"description,omitempty" yaml:"description,omitempty"`
		Variables   ServerVariables `json:"variables,omitempty" yaml:"variables,omitempty"`
	}

	ServerVariables map[string]ServerVariable

	// ServerVariable
	// [Reference]: https://swagger.io/specification/#server-variable-object
	ServerVariable struct {
		Default     string   `json:"default" yaml:"default"`
		Description string   `json:"description,omitempty" yaml:"description,omitempty"`
		Enum        []string `json:"enum,omitempty" yaml:"enum,omitempty"`
	}

	// Components
	// [Reference]: https://swagger.io/specification/#components-object
	Components struct {
		Schemas         map[string]Schema         `json:"schemas,omitempty" yaml:"schemas,omitempty"`
		Responses       map[string]Response       `json:"responses,omitempty" yaml:"responses,omitempty"`
		Parameters      map[string]Parameter      `json:"parameters,omitempty" yaml:"parameters,omitempty"`
		Examples        map[string]Example        `json:"examples,omitempty" yaml:"examples,omitempty"`
		RequestBodies   map[string]RequestBody    `json:"requestBodies,omitempty" yaml:"requestBodies,omitempty"`
		Headers         map[string]Header         `json:"headers,omitempty" yaml:"headers,omitempty"`
		SecuritySchemes map[string]SecurityScheme `json:"securitySchemes,omitempty" yaml:"securitySchemes,omitempty"`
		Links           map[string]Link           `json:"links,omitempty" yaml:"links,omitempty"`
		Callbacks       map[string]Callback       `json:"callbacks,omitempty" yaml:"callbacks,omitempty"`
		PathItems       map[string]PathItem       `json:"pathItems,omitempty" yaml:"pathItems,omitempty"`
	}
)

const Version string = "3.1.0"
