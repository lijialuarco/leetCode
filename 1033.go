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
	c := stones[2]
	b := stones[1]
	a := stones[0]
	max := c - a - 2
	min := 2
	if a+1 == b || b+1 == c || a+2 == b || b+2 == c {
		min = 1
	}
	if a+1 == b && b+1 == c {
		min = 0
	}
	return []int{min, max}

}
