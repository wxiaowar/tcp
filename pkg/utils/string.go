package utils

import (
	"math/rand"
	"time"
)

/*
字符串截取函数

参数：
	str:带截取字符串
	begin:开始截取位置
	length:截取长度
*/
func SubString(str string, begin, length int) (substr string) {
	// 将字符串的转换成[]rune
	rs := []rune(str)
	lth := len(rs)
	// 简单的越界判断
	if begin < 0 {
		begin = 0
	}
	if begin >= lth {
		begin = lth
	}
	end := begin + length
	if end > lth {
		end = lth
	}
	// 返回子串
	return string(rs[begin:end])
}

// 生成一个随机int64
func RandInt(i int64) int64 {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Int63n(i)
}