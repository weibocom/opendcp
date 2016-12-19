package util

import (
	"math/rand"
	"runtime"
	"strings"
	"time"
)

const (
	RAND_KIND_NUM   = 0 // 纯数字
	RAND_KIND_LOWER = 1 // 小写字母
	RAND_KIND_UPPER = 2 // 大写字母
	RAND_KIND_ALL   = 3 // 数字、大小写字母
)

func RandomStr(size int, kind int) string {
	ikind, kinds, result := kind, [][]int{[]int{10, 48}, []int{26, 97}, []int{26, 65}}, make([]byte, size)
	is_all := kind > 2 || kind < 0
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		if is_all {
			ikind = rand.Intn(3)
		}
		scope, base := kinds[ikind][0], kinds[ikind][1]
		result[i] = uint8(base + rand.Intn(scope))
	}
	return string(result)
}

// Get function name by skip
// which means the differs between Caller and Callers
func GetFuncName(skip int) string {
	pc, _, _, _ := runtime.Caller(skip)
	name := runtime.FuncForPC(pc).Name()
	i := strings.LastIndex(name, ".")
	if i >= 0 {
		name = name[i+1:]
	}
	return name
}
