package gopt

import (
	"fmt"
	"reflect"
)

func ErrNotFound(s Setter, name string) error {
	return fmt.Errorf("field %T.%s not found", s, name)
}

func ErrBadType(s Setter, name string, target, value any) error {
	return fmt.Errorf("value %#v (of type %T) is not assignable to field %T.%s (of type %T)", value, value, s, name, target)
}

// ReflectSet sets the field named name to value for the struct s.
// It panics if it fails to assign the value, since any error is considered
// as a programming error.
//
// Note that this is implemented by using reflection, so only exported
// struct fields can be assigned.
func ReflectSet(s Setter, name string, value any) {
	v := reflect.ValueOf(s)
	if v.Kind() != reflect.Ptr {
		panic(fmt.Errorf("want a struct pointer but got %T", s))
	}

	vv := v.Elem()
	if vv.Kind() != reflect.Struct {
		panic(fmt.Errorf("want a struct pointer but got %T", s))
	}

	if _, ok := vv.Type().FieldByName(name); !ok {
		panic(ErrNotFound(s, name))
	}

	f := vv.FieldByName(name)
	if !f.CanSet() {
		panic(fmt.Errorf("field %T.%s cannot be set", s, name))
	}

	if !reflect.TypeOf(value).AssignableTo(f.Type()) {
		panic(ErrBadType(s, name, f.Interface(), value))
	}

	f.Set(reflect.ValueOf(value))
}

// Field represents a struct field to set.
type Field[T any] struct {
	s         Setter
	name      string
	targetPtr *T
}

// F creates a new Field with the given properties.
func F[T any](s Setter, name string, targetPtr *T) Field[T] {
	return Field[T]{s: s, name: name, targetPtr: targetPtr}
}

// Set assigns value to f's target.
// It panics if value is not assignable to the type of f's target.
func (f Field[T]) Set(value any) {
	v, ok := value.(T)
	if !ok {
		panic(ErrBadType(f.s, f.name, *f.targetPtr, value))
	}
	*f.targetPtr = v
}
