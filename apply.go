package mock

import (
	"reflect"
)

func Apply(source interface{}) (err error) {
	return apply(reflect.ValueOf(source))
}

func apply(source reflect.Value) (err error) {
	if source.Kind() != reflect.Ptr || source.IsNil() {
		return ErrInvalidType
	}

	switch source.Elem().Kind() {
	case reflect.String:
		return applyOnString(source.Interface().(*string))
	case reflect.Struct:
		return applyOnStruct(source)
	default:
		return ErrInvalidType
	}
}

func applyOnString(source *string) (err error) {
	result, err := Mock(*source)
	if err != nil {
		return err
	}
	*source = result
	return nil
}

func applyOnStruct(source reflect.Value) (err error) {
	e := source.Elem()
	numField := e.NumField()
	for i := 0; i < numField; i++ {
		ev := e.Field(i)
		err = nil

		if ev.Kind() == reflect.Ptr {
			if !ev.IsNil() {
				err = apply(ev)
			}
		} else {
			if ev.CanAddr() {
				err = apply(ev.Addr())
			}
		}

		if err != nil && err != ErrInvalidType {
			return err
		}
	}

	return nil
}
