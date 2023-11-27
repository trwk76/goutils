package api

import (
	"fmt"
	"reflect"

	"github.com/trwk76/goutils/check"
	"github.com/trwk76/goutils/web/api/spec"
)

type (
	Schemas struct {
		m MediaTypes
		n map[string]*Schema
	}

	Schema struct {
		t reflect.Type
		m MediaType
		s spec.Schema
	}

	SimpleSchema interface {
		Schema() spec.Schema
	}
)

func (s Schemas) Add(t reflect.Type, m MediaTypes) {
	var cnst check.Constraint

	if t == nil {
		panic(fmt.Errorf("cannot register nil type"))
	}

	for t.Kind() == reflect.Pointer {
		t = t.Elem()
	}

	val := reflect.New(t)

	if st, ok := getInterface[SimpleSchema](val); ok {
		name := s.genName(t, nil)
		item := &Schema{t: t, m: nil, s: st.Schema()}

		s.n[name] = item
		return
	}

	if checked, ok := getInterface(val, check.Constrained); ok {
		cnst = checked.Constraint()
	}

	switch t.Kind() {
		
	}
}

func (s Schemas) Type(t reflect.Type, m MediaType) (string, *Schema, bool) {
	for key, item := range s.n {
		if item.t == t && item.m == m {
			return key, item, true
		}
	}

	return "", nil, false
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
		n: make(map[string]*Schema),
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

func getInterface[I any](val reflect.Value) (T, bool) {
	if res, ok := val.Interface().(T); ok {
		return res, true
	}

	val = reflect.Indirect(val)
	return val.Interface().(T)
}
