package main

import (
	"strings"
)

//https://leetcode.cn/problems/ti-huan-kong-ge-lcof/

// 请实现一个函数，把字符串 s 中的每个空格替换成"%20"。

// 时间击败100%，空间击败6.7%
func replaceSpace(s string) string {
	var result string
	for _, v := range s {
		if v == ' ' {
			result += "%20"
		} else {
			result += string(v)
		}
	}
	return result
}

// 时间击败7.88%，空间击败90.27%
func replaceSpaceVersion1(s string) string {
	return strings.Replace(s, " ", "%20", -1)
}

func replaceSpaceVersion2(s string) string {
	return strings.Replace(s, " ", "%20", len(s))
}
