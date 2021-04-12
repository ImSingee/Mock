package function

import (
	"fmt"
	"reflect"
)

var functions = make(map[string][]Function, 32)

type Function struct {
	F          interface{}
	Args       []Argument
	MayBeError bool // 返回值是 error
}

type Argument struct {
	Name string

	InType  reflect.Kind // 用户字面书写的转换后参数类型，支持 string, bool, *GenericNumber
	OutType reflect.Kind // 函数原型参数类型，支持 string, bool, int[X], uint[X]
	// 暂不支持可变参数
}

func MustRegisterFunction(f interface{}, names ...string) {
	err := RegisterFunction(f, names...)
	if err != nil {
		panic(err)
	}
}

var ErrIsNotFunction = fmt.Errorf("is not function")
var errorInterface = reflect.TypeOf((*error)(nil)).Elem()

func RegisterFunction(f interface{}, names ...string) error {
	for _, name := range names {
		err := registerFunction(name, f)
		if err != nil {
			return err
		}
	}

	return nil
}

func registerFunction(name string, f interface{}) error {
	ff := reflect.TypeOf(f)
	if ff.Kind() != reflect.Func {
		return ErrIsNotFunction
	}

	function := Function{F: f}

	numOut := ff.NumOut()
	switch numOut {
	case 0:
		return fmt.Errorf("function return at least one value")
	case 1:
		function.MayBeError = false
	case 2:
		out := ff.Out(1)
		if out.Kind() != reflect.Interface {
			return fmt.Errorf("function's second return value must be error interface")
		}

		if !out.Implements(errorInterface) {
			return fmt.Errorf("function's second return value must be error")
		}

		function.MayBeError = true
	default:
		return fmt.Errorf("function return too much values")
	}

	argsCount := ff.NumIn()
	args := make([]Argument, argsCount)
	for i := range args {
		arg := ff.In(i)

		args[i].Name = arg.Name()
		args[i].OutType = arg.Kind()
		// TODO inType (在 dt 中完成）
	}

	// TODO 重载冲突检查
	functions[name] = append(functions[name], function)

	return nil
}
