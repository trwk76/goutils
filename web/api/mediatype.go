package api

type (
	MediaTypes []MediaType

	MediaType interface {
		Prefix() string
		ContentType() string
	}
)
