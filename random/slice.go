package random

import (
	"github.com/ImSingee/mock/iter"
	"sort"
)

// 从 [min, max) 区间范围内随机返回 count 数量的不同的数字，按照升序排列
func Unique(min, max, count int) []int {
	if count <= 0 || max <= min {
		return nil
	}
	if count >= max-min {
		return iter.List(min, max)
	}

	m := make(map[int]struct{}, count)

	for len(m) != count {
		m[Int(min, max-1)] = struct{}{}
	}

	result := make([]int, 0, count)
	for k := range m {
		result = append(result, k)
	}
	sort.Ints(result)

	return result
}

func SelectInts(slice []int, count int) []int {
	if count <= 0 {
		return nil
	}
	if count >= len(slice) {
		return slice
	}

	indices := Unique(0, len(slice), count)

	result := make([]int, count)
	for i, v := range indices {
		result[i] = slice[v]
	}

	return result
}
