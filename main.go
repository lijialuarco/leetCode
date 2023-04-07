package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(numMovesStones(1, 2, 5))
}

func numMovesStones(a int, b int, c int) []int {
	return numMovesStonesII([]int{a, b, c})
}

func numMovesStonesII(stones []int) []int {
	if len(stones) == 1 {
		return stones
	}
	sort.Ints(stones)
	minGap := stones[len(stones)-1]
	maxStep, minStep, maxGap := 2, 1, 0
	consecutiveCount := 0
	//unConsecutiveCount := 0
	for k := 1; k < len(stones); k++ {
		if stones[k]-stones[k-1] == 1 {
			consecutiveCount++
		}
		if stones[k]-stones[k-1] < minGap {
			minGap = stones[k] - stones[k-1]
		} else if k != 0 && stones[k]-stones[k-1] > maxGap {
			maxGap = stones[k] - stones[k-1]
		}
	}

	if minGap == 1 && consecutiveCount > 1 {
		minStep = 0
		maxStep = 0
	}

	if minGap == 1 && consecutiveCount == 1 {
		minStep = 1

	}

	if maxGap > 1 && consecutiveCount == 1 {
		maxStep = maxGap - 1
	}

	return []int{minStep, maxStep}
}
