package random

import (
	"fmt"
	"math"
	"strconv"
)

func MustRegister(MustRegisterFunction func(params ...interface{})) {
	// @bool() @boolean()
	MustRegisterFunction("bool", "boolean", Boolean)

	// @natural()
	// @natural(min)
	// @natural(min, max)
	MustRegisterFunction("natural",
		func() int {
			return Natural(0, math.MaxInt32)
		},
		func(min int) int {
			return Natural(min, math.MaxInt32)
		},
		Natural,
	)

	// @integer()          @int()
	// @integer(min)       @int(min)
	// @integer(min, max)  @int(min, max)
	MustRegisterFunction("int", "integer",
		func() int {
			return Integer(math.MinInt32, math.MaxInt32)
		},
		func(min int) int {
			return Integer(min, math.MaxInt32)
		},
		Int,
	)

	// @float()
	// @float(min)
	// @float(min, max)
	// @float(min, max, d)
	MustRegisterFunction("float",
		func() float64 {
			return Float64(float64(math.MinInt32), float64(math.MaxInt32))
		},
		func(min float64) float64 {
			return Float64(min, float64(math.MaxInt32))
		},
		Float,
		func(min, max float64, d int) string {
			return fmt.Sprintf("%."+strconv.Itoa(d)+"f", Float64(min, max))
		},
	)

	// @number(a) -> 永远返回 a
	MustRegisterFunction("number", func(a int) int {
		return a
	})

	// @character()       @char()
	// @character(pool)   @char(pool)
	MustRegisterFunction("char", "character",
		func() rune {
			return Character("numletter")
		},
		func(pool string) rune {
			return Character(pool)
		},
	)

	// @string()
	// @string(length)
	// @string(min, max)
	// @string(pool)
	// @string(pool, length)
	// @string(pool, min, max)
	MustRegisterFunction("string",
		func() string {
			return String("numletter", 1, 20)
		},
		func(length int) string {
			return String("numletter", length, length)
		},
		func(min, max int) string {
			return String("numletter", min, max)
		},
		func(pool string) string {
			return String(pool, 1, 20)
		},
		func(pool string, length int) string {
			return String(pool, length, length)
		},
		func(pool string, min, max int) string {
			return String(pool, min, max)
		},
	)

	// @sentence
	MustRegisterFunction("sentence", Sentence)

	// @increment()
	// @increment(step)
	// @increment(step, delta)
	// @increment(namespace)
	// @increment(namespace, step)
	// @increment(namespace, step, delta)
	MustRegisterFunction("increment",
		func() int64 {
			return Increment()
		},
		func(step int64) int64 {
			return IncrementWithStep(step)
		},
		func(step, delta int64) int64 {
			return IncrementWithStepAndDelta(step, delta)
		},
		func(namespace string) int64 {
			return IncrementN(namespace)
		},
		func(namespace string, step int64) int64 {
			return IncrementNWithStep(namespace, step)
		},
		func(namespace string, step, delta int64) int64 {
			return IncrementNWithStepAndDelta(namespace, step, delta)
		},
	)
}
