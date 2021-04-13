package random

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
	k := r.Float64()
	if min < 0 {
		return (k*max + min) - k*min
	} else {
		return k*(max-min) + min
	}
}

func Float32(min, max float32) float32 {
	min, max = fixFloat32MinMax(min, max)
	return min + r.Float32()*(max-min)
}

var Float = Float64
