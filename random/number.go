package random

import (
	"github.com/ImSingee/mock/function"
	"math"
)

// 返回一个随机的自然数（大于等于 0 的整数）
// - min 最小值
// - max 最大值
func Natural(min, max int) int {
	min, max = fixIntMinMaxPositive(min, max)

	return r.Intn(max-min+1) + min
}

// 返回一个随机的整数
// - min 最小值
// - max 最大值
//
// 映射情况 1: min < max -> swap(min, max)
func Integer(min, max int) int {
	min, max = fixIntMinMax(min, max)

	return r.Intn(max-min) + min
}

var Int = Integer

func Float64(min, max float64) float64 {
	min, max = fixFloat64MinMax(min, max)
	return min + r.Float64()*(max-min)
}

func Float32(min, max float32) float32 {
	min, max = fixFloat32MinMax(min, max)
	return min + r.Float32()*(max-min)
}

var Float = Float64

func init() {
	function.MustRegisterFunction("int", "integer",
		func() int { // @integer() @int()
			return Integer(math.MinInt32, math.MaxInt32)
		},
		func(min int) int { // @integer(10000)  @int(10000)
			return Integer(min, math.MaxInt32)
		},
		Int, // @integer(60, 100) @int(60, 100)
	)
}
