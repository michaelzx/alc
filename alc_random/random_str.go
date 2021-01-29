package alc_random

import (
	"math/rand"
	"time"
)

type RandomType uint

const (
	RandomTypeNum   RandomType = 0 // 纯数字
	RandomTypeLower RandomType = 1 // 小写字母
	RandomTypeUpper RandomType = 2 // 大写字母
	RandomTypeAll   RandomType = 3 // 数字、大小写字母
)

func RandomNumStr(size int) string {
	return string(randomStr(size, RandomTypeNum))
}
func RandomAllStr(size int) string {
	return string(randomStr(size, RandomTypeAll))
}
func RandomStr(size int, randomType RandomType) string {
	return string(randomStr(size, randomType))
}

// 随机字符串
func randomStr(size int, randomType RandomType) []byte {
	kind, kinds, result := randomType, [][]int{{10, 48}, {26, 65}, {26, 97}}, make([]byte, size)
	isAll := randomType > 2 || randomType < 0
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		if isAll {
			// random kind
			kind = RandomType(rand.Intn(3))
		}
		scope, base := kinds[kind][0], kinds[kind][1]
		result[i] = uint8(base + rand.Intn(scope))
	}
	return result
}
