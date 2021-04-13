package random

import "strings"

// 生成随机字符串
// pool - 池
// min - 最小生成数量
// max - 最大生成数量
func String(pool string, min, max int) string {
	b := strings.Builder{}

	l := Natural(min, max)

	mappedPool := MapCharacterPool(pool)

	for i := 0; i < l; i++ {
		b.WriteRune(CharacterR(mappedPool))
	}

	return b.String()
}

var Str = String
