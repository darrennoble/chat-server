package config

import (
	"fmt"
	"reflect"
)

func copyNonzero(src, dest interface{}) error {
	srcType := reflect.TypeOf(src)
	if srcType.Kind() != reflect.Struct {
		return fmt.Errorf("src is of type %T, which is not a struct", src)
	}

	destType := reflect.TypeOf(dest)
	if destType.Kind() != reflect.Ptr {
		return fmt.Errorf("dest is of type %T, which is not a pointer", dest)
	}

	if destType.Elem() != srcType {
		return fmt.Errorf("type mismatch: src is '%T' and dest is '%T'", src, dest)
	}

	srcV := reflect.ValueOf(src)
	destV := reflect.ValueOf(dest).Elem()

	for i := 0; i < srcV.NumField(); i++ {
		srcF := srcV.Field(i)
		destF := destV.Field(i)

		if !destF.CanSet() {
			continue
		}

		if isZero(destF.Interface()) {
			destF.Set(srcF)
		} else {
			if srcF.Type().Kind() == reflect.Struct {
				err := copyNonzero(srcF.Interface(), destF.Addr().Interface())
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func isZero(i interface{}) bool {
	v := reflect.ValueOf(i)
	switch v.Kind() {
	case reflect.Func, reflect.Map, reflect.Slice:
		return v.IsNil()
	case reflect.Array:
		for i := 0; i < v.Len(); i++ {
			if !isZero(v.Index(i).Interface()) {
				return false
			}
		}
		return true
	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			if !isZero(v.Field(i).Interface()) {
				return false
			}
		}
		return true
	case reflect.Ptr:
		return isZero(v.Elem().Interface())
	}

	z := reflect.Zero(v.Type())
	return z.Interface() == v.Interface()
}
