package random

var CharacterPool = map[string]string{
	"lower":     "abcdefghijklmnopqrstuvwxyz",
	"upper":     "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
	"letter":    "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ",
	"number":    "0123456789",
	"symbol":    "!@#$%^&*()[]",
	"numletter": "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789",
	"any":       "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()[]",

	// 为了和 mockjs 保持一致
	"alpha": "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ",
	"":      "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()[]",
}

// pool 映射
func MapCharacterPool(pool string) string {
	if p := CharacterPool[pool]; p != "" {
		return p
	} else {
		return pool
	}
}

// 利用 pool 生成随机一个随机字节
func Character(pool string) byte {
	return CharacterR(MapCharacterPool(pool))
}

// 利用 pool 生成随机一个随机字节
// 与 Character 的区别：不会进行 pool 映射
func CharacterR(pool string) byte {
	return pool[r.Intn(len(pool))]
}

var Byte = Character
var Char = Character
