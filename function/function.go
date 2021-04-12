package function

import (
	"fmt"
	"reflect"
)

func CallFunction(funcName string, args []interface{}) (interface{}, error) {
	funcs := functions[funcName]
	if len(funcs) == 0 {
		return nil, fmt.Errorf("unknwon function %s", funcName)
	}

	var lastErr error = nil
	for _, f := range funcs {
		inArgs, err := f.Apply(args)
		if err != nil {
			lastErr = fmt.Errorf("args not match")
			continue
		}

		ff := reflect.ValueOf(f)
		result := ff.Call(inArgs)

		if f.MayBeError {
			err := result[1].Elem()
			if !err.IsNil() {
				lastErr = err.Interface().(error)
				continue
			}
		}

		return result[0].Interface(), nil
	}

	//b := strings.Builder{}
	//b.WriteString("[")
	//b.WriteString(funcName)
	//b.WriteString("]{")
	//for i, arg := range args {
	//	if i != 0 {
	//		b.WriteString(", ")
	//	}
	//	b.WriteString(fmt.Sprintf("(%T) %v", arg, arg))
	//}
	//b.WriteString("}")

	return nil, lastErr
}
