package main

//
// 给定一个正整数数组,元素重复个数只可能是两个结果,且两个结果互质,其中一个结果只出现一次
// 要求空间复杂度为O(1) 找到这个元素
func main() {
	input := []int{1, 2, 3, 4, 5, 6, 7, 7, 6, 5, 4, 3, 2, 1, 5, 5}
	println(findSingleCount(input, 4))
}

func findSingleCount(nums []int, a int) int {
	mask := 1 // 用来找到只出现一次的元素在二进制表示中的位置
	for (a & mask) == 0 {
		mask <<= 1
	}

	result := 0
	for i := 0; i < len(nums); i++ {
		if (nums[i] & mask) != 0 {
			result ^= nums[i] // 异或操作，相同的数异或结果为0
		}
	}

	return result
}
