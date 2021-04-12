package random

import "github.com/ImSingee/mock/function"

func Bool() bool {
	return r.Intn(2) == 1
}

var Boolean = Bool

func init() {
	function.MustRegisterFunction(Boolean, "bool", "boolean")
}
