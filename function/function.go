package function

import (
	"fmt"
	"strings"
)

func CallFunction(funcName string, args []interface{}) (interface{}, error) {
	b := strings.Builder{}
	b.WriteString("[")
	b.WriteString(funcName)
	b.WriteString("]{")
	for i, arg := range args {
		if i != 0 {
			b.WriteString(", ")
		}
		b.WriteString(fmt.Sprintf("(%T) %v", arg, arg))
	}
	b.WriteString("}")

	// TODO
	return b.String(), nil
}
