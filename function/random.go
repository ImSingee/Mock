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

	MustRegisterFunction("number", func(a int) int {
		return a
	})

}
