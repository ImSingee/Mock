package random

var CharacterPool = map[string][]rune{
	"lower":     []rune("abcdefghijklmnopqrstuvwxyz"),
	"upper":     []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ"),
	"letter":    []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"),
	"number":    []rune("0123456789"),
	"symbol":    []rune("!@#$%^&*()[]"),
	"numletter": []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"),
	"any":       []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()[]"),

	// 为了和 mockjs 保持一致
	"alpha": []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"),
	"":      []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()[]"),
}

// pool 映射
func MapCharacterPool(pool string) []rune {
	if p := CharacterPool[pool]; len(p) != 0 {
		return p
	} else {
		return []rune(pool)
	}
}

// 利用 pool 生成随机一个随机字符
func Character(pool string) rune {
	return CharacterR(MapCharacterPool(pool))
}

// 利用 pool 生成随机一个随机字节
// 与 Character 的区别：不会进行 pool 映射
func CharacterR(pool []rune) rune {
	return pool[r.Intn(len(pool))]
}

var Char = Character
