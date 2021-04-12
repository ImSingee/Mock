package mock

import (
	"fmt"
	"testing"
)

func TestMock(t *testing.T) {
	fmt.Println(Mock("ABC"))

	fmt.Println(Mock("@number"))
	fmt.Println(Mock("@number()"))
	fmt.Println(Mock("@number(3)"))
	fmt.Println(Mock("@number(3,5)"))
	fmt.Println(Mock("@number(3, 5)"))
	fmt.Println(Mock(`@string("aeiou", 3, 5)`))
	fmt.Println(Mock(`@bool(true)`))
	fmt.Println(Mock(`@test(true, 'a')`))
	fmt.Println(Mock(`@test(nil)`))
	fmt.Println(Mock(`@something(1.2, 3.4, )`))
}
