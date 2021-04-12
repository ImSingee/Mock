package mock

import (
	"fmt"
	"github.com/ImSingee/dt"
	"github.com/ImSingee/mock/function"
	"go/ast"
	"go/parser"
	"regexp"
	"strings"
)

var re = regexp.MustCompile(`@(?P<call>(?P<function>[a-zA-Z0-9\-_]+)(?P<args>\(.*?\)))`)

func mapArgs(args []ast.Expr) ([]interface{}, error) {
	mapped := make([]interface{}, len(args))
	var err error

	for i := range args {
		//fmt.Printf("param %d: %#+v\n", i, args[i])

		switch arg := args[i].(type) {
		case *ast.BasicLit:
			mapped[i], err = mapArg(arg.Value)
		case *ast.Ident:
			mapped[i], err = mapArg(arg.String())
		}

		if err != nil {
			return nil, fmt.Errorf("invalid param (E03)\n")
		}
	}

	return mapped, nil
}

func mapArg(arg string) (interface{}, error) {
	// 关键字（*ast.Ident）   <- bool, nil
	switch arg {
	case "true":
		return true, nil
	case "false":
		return false, nil
	case "nil":
		return nil, nil
	}

	// 字面值量 (*ast.BasicLit) <- number, string

	if strings.HasPrefix(arg, `'`) || strings.HasPrefix(arg, `"`) {
		if !strings.HasPrefix(arg, `'`) && !strings.HasPrefix(arg, `"`) {
			return nil, fmt.Errorf("invalid literal")
		}
		return arg[1 : len(arg)-1], nil
	}

	if num, ok := dt.NumberFromString(arg); ok {
		return num, nil
	}

	return nil, fmt.Errorf("invalid param (E5) V=%s", arg)
}

func Mock(source string) (s string, err error) {
	defer func() {
		e := recover()
		if ee, ok := e.(error); ok {
			err = ee
			return
		} else if e != nil {
			err = fmt.Errorf("unknown Error (%v)", e)
		}
	}()

	return re.ReplaceAllStringFunc(source, func(s string) string {
		f := source[1:]

		tree, err := parser.ParseExpr(f)

		if err != nil {
			panic(fmt.Errorf("parse error (%s) (%w)", f, err))
			return ""
		}
		//fmt.Printf("%s: %#+v\n", f, tree)

		call, ok := tree.(*ast.CallExpr)
		if !ok {
			panic(fmt.Errorf("parse as Call error (%s) (%T)", f, tree))
			return ""
		}

		funcName := call.Fun.(*ast.Ident).String()
		args, err := mapArgs(call.Args)

		if err != nil {
			panic(err)
		}

		result, err := function.CallFunction(funcName, args)
		if err != nil {
			panic(err)
		}

		return dt.ToString(result)
	}), nil
}
