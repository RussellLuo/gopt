package gopt

import (
	"fmt"
	"reflect"
)

// ReflectSet assigns the field named name to value for the struct s.
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
		panic(fmt.Errorf("field %T.%s not found", s, name))
	}

	f := vv.FieldByName(name)
	if !f.CanSet() {
		panic(fmt.Errorf("field %T.%s cannot be set", s, name))
	}

	if !reflect.TypeOf(value).AssignableTo(f.Type()) {
		panic(fmt.Errorf("value %#v (of type %T) is not assignable to field %T.%s (of type %T)", value, value, s, name, f.Interface()))
	}

	f.Set(reflect.ValueOf(value))
}
