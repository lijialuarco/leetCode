package main

// https://leetcode.cn/problems/widest-vertical-area-between-two-points-containing-no-points/
// 给你 n 个二维平面上的点 points ，其中 points[i] = [xi, yi] ，请你返回两点之间内部不包含任何点的 最宽垂直区域 的宽度。
// 垂直区域 的定义是固定宽度，而 y 轴上无限延伸的一块区域（也就是高度为无穷大）。 最宽垂直区域 为宽度最大的一个垂直区域。
// 请注意，垂直区域 边上 的点 不在 区域内。
func main() {
	input := [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}, {5, 6}, {6, 7}}
	println(maxWidthOfVerticalArea(input))
}

func maxWidthOfVerticalArea(points [][]int) int {
	max := 0
	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			if points[i][0] > points[j][0] {
				points[i], points[j] = points[j], points[i]
			}
		}
	}

	for i := 0; i < len(points)-1; i++ {
		if points[i+1][0]-points[i][0] > max {
			max = points[i+1][0] - points[i][0]
		}
	}
	return max
}
