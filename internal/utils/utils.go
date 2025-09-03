package utils

import (
	"reflect"
	"strings"
)

func ParseFieldNames[T any](s T) map[string]struct{} {
	res := make(map[string]struct{})

	t := reflect.TypeOf(s)

	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		fieldName := LowercaseFirst(f.Name)
		res[fieldName] = struct{}{}
	}

	return res
}

func LowercaseFirst(s string) string {
	if len(s) == 0 {
		return s
	}
	return strings.ToLower(s[:1]) + s[1:]
}
