package random

var MinInt = -9007199254740992
var MaxInt = 9007199254740992

func fixIntMinMax(min, max int) (int, int) {
	if min > max {
		return max, min
	}

	return min, max
}

func fixFloat32MinMax(min, max float32) (float32, float32) {
	if min > max {
		return max, min
	}

	return min, max
}

func fixFloat64MinMax(min, max float64) (float64, float64) {
	if min > max {
		return max, min
	}

	return min, max
}

// 映射情况 1: max <= 0 -> max=9007199254740992 (2^53)
// 映射情况 2: min < 0  -> min=0
// 映射情况 3: min < max -> swap(min, max)
func fixIntMinMaxPositive(min, max int) (int, int) {
	if min < 0 {
		min = 0
	}
	if max <= 0 {
		max = MaxInt
	}
	if min == max {
		return min, min
	}
	if min > max {
		return max, min
	}

	return min, max
}
