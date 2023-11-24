package api

type (
	MediaTypes []MediaType

	MediaType interface {
		ContentType() string
	}
)
