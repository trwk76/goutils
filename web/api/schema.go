package api

import (
	"reflect"

	"github.com/trwk76/goutils/web/api/spec"
)

type (
	Schemas map[string]Schema

	Schema struct {
		t reflect.Type
		m MediaType
		s spec.Schema
	}
)
