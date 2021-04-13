package mock

import (
	"fmt"
	"github.com/ImSingee/dt"
	"github.com/ImSingee/mock/random"
	"github.com/ImSingee/tt"
	"testing"
)

func TestMockNone(t *testing.T) {
	mock, err := Mock("ABCD")
	tt.AssertIsNil(t, err)
	fmt.Println(mock)
	tt.AssertEqual(t, "ABCD", mock)

	mock, err = Mock("@some")
	tt.AssertIsNil(t, err)
	fmt.Println(mock)
	tt.AssertEqual(t, "@some", mock)

	//fmt.Println(Mock("@number"))
	//fmt.Println(Mock("@number()"))
	//fmt.Println(Mock("@number(3)"))
	//fmt.Println(Mock("@number(3,5)"))
	//fmt.Println(Mock("@number(3, 5)"))
	//fmt.Println(Mock(`@string("aeiou", 3, 5)`))
	//fmt.Println(Mock(`@bool(true)`))
	//fmt.Println(Mock(`@test(true, 'a')`))
	//fmt.Println(Mock(`@test(nil)`))
	//fmt.Println(Mock(`@something(1.2, 3.4, )`))
}

func TestMockBoolean(t *testing.T) {
	trueCount := 0
	falseCount := 0
	for i := 0; i < 100; i++ {
		mock, err := Mock("@bool()")
		tt.AssertIsNil(t, err)
		fmt.Println(mock)
		switch mock {
		case "true":
			trueCount += 1
		case "false":
			falseCount += 1
		default:
			t.FailNow()
		}
	}
	tt.AssertNotEqual(t, 0, trueCount)
	tt.AssertNotEqual(t, 0, falseCount)
}

func TestMockInteger(t *testing.T) {
	mockName := []string{"int", "integer"}
	for _, name := range mockName {
		t.Run("mock use @"+name, func(t *testing.T) {
			t.Run("no param", func(t *testing.T) {
				for i := 0; i < 100; i++ {
					mock, err := Mock("@" + name + "()")
					tt.AssertIsNil(t, err)
					num, ok := dt.NumberFromString(mock)
					tt.AssertTrue(t, ok)
					tt.AssertTrue(t, num.IsInt64())
				}
			})

			t.Run("one param (min)", func(t *testing.T) {
				for i := 0; i < 100; i++ {
					r := int64(random.Integer(100, 1000))
					mock, err := Mock(fmt.Sprintf("@%s(%d)", name, r))
					tt.AssertIsNil(t, err)
					num, ok := dt.NumberFromString(mock)
					tt.AssertTrue(t, ok)
					tt.AssertTrue(t, num.IsInt64())
					tt.AssertTrue(t, num.Int64() >= r)
				}
			})

			t.Run("two param (min, max)", func(t *testing.T) {
				for i := 0; i < 100; i++ {
					r := int64(random.Integer(100, 1000))
					q := int64(random.Integer(10000, 100000))
					mock, err := Mock(fmt.Sprintf("@%s(%d, %d)", name, r, q))
					tt.AssertIsNil(t, err)
					num, ok := dt.NumberFromString(mock)
					tt.AssertTrue(t, ok)
					tt.AssertTrue(t, num.IsInt64())
					tt.AssertTrue(t, num.Int64() >= r)
					tt.AssertTrue(t, num.Int64() <= q)
				}
			})
		})
	}
}
