package check

import (
	"reflect"
)

type (
	Constraint interface {
		Check(ctx Context, val reflect.Value) Error
	}

	Constrained interface {
		Constraint() Constraint
	}
)
