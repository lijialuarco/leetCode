package main

// https://leetcode.cn/problems/er-wei-shu-zu-zhong-de-cha-zhao-lcof/
//在一个 n * m 的二维数组中，每一行都按照从左到右 非递减 的顺序排序，每一列都按照从上到下 非递减 //的顺序排序。请完成一个高效的函数，输入这样的一个二维数组和一个整数，判断数组中是否含有该整数。
//示例:
//现有矩阵 matrix 如下：
//
//[
//[1,   4,  7, 11, 15],
//[2,   5,  8, 12, 19],
//[3,   6,  9, 16, 22],
//[10, 13, 14, 17, 24],
//[18, 21, 23, 26, 30]
//]
//给定 target = 5，返回 true。
//
//给定 target = 20，返回 false。

// 这种解法有些
func FindNumberIn2DArray(matrix [][]int, target int) bool {
	width := 0
	height := len(matrix[0]) - 1

	for width < len(matrix)-1 && height >= 0 {
		if matrix[width][height] == target {
			return true
		} else if matrix[width][height] < target {
			width++
		} else {
			height--
		}
	}
	return false
}
