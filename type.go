package util

import (
	"reflect"
	"strings"
)

// GetTypeName returns the name of the type T.
// If T is a pointer type, it returns the name with a leading asterisk (e.g., "*MyStruct").
func GetTypeName[T any]() string {
	_, name := typeInfoAndName[T]()
	return name
}

// GetTypeNameFromValue returns the name of the type of the given value.
// If the value is a pointer, it returns the name with a leading asterisk (e.g., "*MyStruct").
// It panics if value is nil.
func GetTypeNameFromValue(value any) string {
	_, name := typeInfoAndNameFromValue(value)
	return name
}

// GetTypeString returns the string representation of the type T.
// This includes the package path if the type is from another package.
func GetTypeString[T any]() string {
	typeOf, _ := typeInfoAndName[T]()
	return typeOf.String()
}

// Deprecated: Use GetTypeNameFromValue instead, as GetTypeStringFromValue might return unexpected results for anonymous types or interfaces.
func GetTypeStringFromValue(value any) string {
	typeOf, _ := typeInfoAndNameFromValue(value)
	return typeOf.String()
}

// typeInfoAndName returns the reflect.Type and the name of the type T.
// It handles pointer types by returning the element type's name with a leading asterisk.
//
// It returns the same reflect.Type instance for the same type.
func typeInfoAndName[T any]() (reflect.Type, string) {
	typeOf := reflect.TypeOf((*T)(nil)).Elem()
	if typeOf.Kind() == reflect.Ptr {
		return typeOf, "*" + typeOf.Elem().Name()
	}
	return typeOf, typeOf.Name()
}

// typeInfoAndNameFromValue returns the reflect.Type and the name of the type of the given value.
// It handles multiple levels of pointers and anonymous types.
// It panics if value is nil.
func typeInfoAndNameFromValue(value any) (reflect.Type, string) {
	if value == nil {
		panic("value cannot be nil")
	}
	typeOf := reflect.TypeOf(value)
	deep := 0
	current := typeOf
	for current.Name() == "" && current.Kind() == reflect.Ptr {
		deep++
		current = current.Elem()
	}
	if current.Name() == "" {
		v := value
		cur := reflect.TypeOf(v)
		for cur.Name() == "" && cur.Kind() == reflect.Ptr {
			v = *(v.(*any))
			cur = reflect.TypeOf(v)
		}
		if cur.Name() != "" {
			return typeOf, strings.Repeat("*", deep) + cur.Name()
		}
		return typeOf, strings.Repeat("*", deep) + cur.String()
	}
	return typeOf, strings.Repeat("*", deep) + current.Name()
}