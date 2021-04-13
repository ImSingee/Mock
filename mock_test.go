package mock

import (
	"fmt"
	"github.com/ImSingee/dt"
	"github.com/ImSingee/mock/random"
	"github.com/ImSingee/tt"
	"strings"
	"testing"
)

func TestMockNone(t *testing.T) {
	mocks := map[string]string{
		"ABCD":    "ABCD",
		"@some":   "@some",
		"@@":      "@",
		"@@abc()": "@abc()",
	}

	for a, b := range mocks {
		t.Run("test for "+a, func(t *testing.T) {
			mock, err := Mock(a)
			tt.AssertIsNil(t, err)
			fmt.Println(mock)
			tt.AssertEqual(t, b, mock)
		})
	}

	//fmt.Println(Mock(`@string("aeiou", 3, 5)`))
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

func TestMockFloat(t *testing.T) {
	t.Run("no param", func(t *testing.T) {
		for i := 0; i < 100; i++ {
			mock, err := Mock("@float()")
			tt.AssertIsNil(t, err)
			fmt.Println(mock)
			num, ok := dt.NumberFromString(mock)
			tt.AssertTrue(t, ok)
			tt.AssertTrue(t, num.Float())
		}
	})

	t.Run("one param (min)", func(t *testing.T) {
		for i := 0; i < 100; i++ {
			r := random.Float64(100, 1000)
			mock, err := Mock(fmt.Sprintf("@float(%.3f)", r))
			fmt.Println(mock)
			tt.AssertIsNil(t, err)
			num, ok := dt.NumberFromString(mock)
			tt.AssertTrue(t, ok)
			tt.AssertTrue(t, num.Float())
			tt.AssertTrue(t, num.Float64() >= float64(int64(r)))
		}
	})

	t.Run("two param (min, max)", func(t *testing.T) {
		for i := 0; i < 100; i++ {
			r := random.Float64(100, 1000)
			q := random.Float64(10000, 100000)
			mock, err := Mock(fmt.Sprintf("@float(%.3f, %.3f)", r, q))
			fmt.Println(mock)
			tt.AssertIsNil(t, err)
			num, ok := dt.NumberFromString(mock)
			tt.AssertTrue(t, ok)
			tt.AssertTrue(t, num.Float())
			tt.AssertTrue(t, num.Float64() >= float64(int64(r)))
			tt.AssertTrue(t, num.Float64() <= float64(int64(q)+1))
		}
	})

	t.Run("two param (min, max, d)", func(t *testing.T) {
		for i := 0; i < 100; i++ {
			r := random.Float64(100, 1000)
			q := random.Float64(10000, 100000)
			d := random.Integer(1, 5)
			mock, err := Mock(fmt.Sprintf("@float(%.3f, %.3f, %d)", r, q, d))
			fmt.Println(mock)
			tt.AssertIsNil(t, err)
			num, ok := dt.NumberFromString(mock)
			tt.AssertTrue(t, ok)
			tt.AssertTrue(t, num.Float())
			tt.AssertTrue(t, num.Float64() >= float64(int64(r)))
			tt.AssertTrue(t, num.Float64() <= float64(int64(q)+1))
			tt.AssertTrue(t, strings.Contains(mock, "."))
			tt.AssertEqual(t, d, len(mock)-strings.Index(mock, ".")-1)
		}
	})
}
