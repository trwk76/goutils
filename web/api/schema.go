package api

import (
	"fmt"
	"reflect"

	"github.com/trwk76/goutils/web/api/spec"
)

type (
	Schemas struct {
		m MediaTypes
		n map[string]Schema
	}

	Schema struct {
		t reflect.Type
		m MediaType
		s spec.Schema
	}
)

func (s Schemas) With(t reflect.Type) {
	if t == nil {
		panic(fmt.Errorf("cannot register nil type"))
	}

	for t.Kind() == reflect.Pointer {
		t = t.Elem()
	}
}

func (s Schemas) Type(t reflect.Type, m MediaType) (string, Schema, bool) {
	for key, item := range s.n {
		if item.t == t && item.m == m {
			return key, item, true
		}
	}

	return "", Schema{}, false
}

func (s Schemas) Reference(t reflect.Type, m MediaType) string {
	key, _, ok := s.Type(t, m)

	if !ok {
		return ""
	}

	return "#/components/schemas/" + key
}

func (s Schema) Type() reflect.Type {
	return s.t
}

func (s Schema) MediaType() MediaType {
	return s.m
}

func (s Schema) Spec() spec.Schema {
	return s.s
}

func newSchemas(m MediaTypes) Schemas {
	return Schemas{
		m: m,
		n: make(map[string]Schema),
	}
}

func (s Schemas) genName(t reflect.Type, m MediaType) string {
	base := typeName(t, m)
	name := base
	i := 1
	_, exists := s.n[name]

	for exists {
		name = fmt.Sprintf("%s%d", base, i)
		i += 1
		_, exists = s.n[name]
	}

	return name
}

func typeName(t reflect.Type, m MediaType) string {
	for t.Kind() == reflect.Pointer {
		t = t.Elem()
	}

	name := t.Name()

	if name == "" {
		switch t.Kind() {
		case reflect.Array, reflect.Slice:
			name = typeName(t.Elem(), m) + "Array"
		case reflect.Map:
			name = typeName(t.Elem(), m) + "Map"
		default:
			name = "Anonymous"
		}
	}

	if m != nil {
		name = m.Prefix() + name
	}

	return name
}
