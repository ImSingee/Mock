package function

import (
	"fmt"
	"github.com/ImSingee/mock/random"
	"math"
	"strconv"
)

func init() {
	// @bool() @boolean()
	MustRegisterFunction("bool", "boolean", random.Boolean)

	// @natural()
	// @natural(min)
	// @natural(min, max)
	MustRegisterFunction("natural",
		func() int {
			return random.Natural(0, math.MaxInt32)
		},
		func(min int) int {
			return random.Natural(min, math.MaxInt32)
		},
		random.Natural,
	)

	// @integer()          @int()
	// @integer(min)       @int(min)
	// @integer(min, max)  @int(min, max)
	MustRegisterFunction("int", "integer",
		func() int {
			return random.Integer(math.MinInt32, math.MaxInt32)
		},
		func(min int) int {
			return random.Integer(min, math.MaxInt32)
		},
		random.Int,
	)

	// @float()
	// @float(min)
	// @float(min, max)
	// @float(min, max, d)
	MustRegisterFunction("float",
		func() float64 {
			return random.Float64(float64(math.MinInt32), float64(math.MaxInt32))
		},
		func(min float64) float64 {
			return random.Float64(min, float64(math.MaxInt32))
		},
		random.Float,
		func(min, max float64, d int) string {
			return fmt.Sprintf("%."+strconv.Itoa(d)+"f", random.Float64(min, max))
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
			return random.Character("numletter")
		},
		func(pool string) rune {
			return random.Character(pool)
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
			return random.String("numletter", 1, 20)
		},
		func(length int) string {
			return random.String("numletter", length, length)
		},
		func(min, max int) string {
			return random.String("numletter", min, max)
		},
		func(pool string) string {
			return random.String(pool, 1, 20)
		},
		func(pool string, length int) string {
			return random.String(pool, length, length)
		},
		func(pool string, min, max int) string {
			return random.String(pool, min, max)
		},
	)

	// @sentence
	MustRegisterFunction("sentence", random.Sentence)

	// @increment()
	// @increment(step)
	// @increment(step, delta)
	// @increment(namespace)
	// @increment(namespace, step)
	// @increment(namespace, step, delta)
	MustRegisterFunction("increment",
		func() int64 {
			return random.Increment()
		},
		func(step int64) int64 {
			return random.IncrementWithStep(step)
		},
		func(step, delta int64) int64 {
			return random.IncrementWithStepAndDelta(step, delta)
		},
		func(namespace string) int64 {
			return random.IncrementN(namespace)
		},
		func(namespace string, step int64) int64 {
			return random.IncrementNWithStep(namespace, step)
		},
		func(namespace string, step, delta int64) int64 {
			return random.IncrementNWithStepAndDelta(namespace, step, delta)
		},
	)
}
