package main

import (
	"net/url"
	"reflect"
	"strconv"
)

// Includes in the payloads.
type Includes struct {

	// StemCellVersion flag.
	StemCellVersion bool `qs:"scv"`
}

// NewIncludes will create a new includes object.
func NewIncludes(queryString url.Values) (includes *Includes) {
	includes = &Includes{}

	if len(queryString) > 0 {
		t := reflect.TypeOf(*includes)
		elem := reflect.ValueOf(includes).Elem()

		for i := 0; i < t.NumField(); i++ {
			structField := t.Field(i)
			field := elem.Field(i)

			if field.CanSet() {
				qsName := structField.Tag.Get("qs")
				strVal := queryString.Get(qsName)

				if val, err := strconv.ParseBool(strVal); err == nil {
					field.SetBool(val)
				}
			}
		}
	}

	return
}
