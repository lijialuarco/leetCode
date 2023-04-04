package main

import "fmt"

func main() {
	fmt.Println(prevPermOpt1([]int{1, 9, 4, 6, 7}))
}

func prevPermOpt1(arr []int) []int {
	maxIndex := 0
	minIndex := 0
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				if arr[i] > arr[maxIndex] {
					maxIndex = i
				}

				if arr[j] < arr[minIndex] {
					minIndex = j
				}

			}
		}
	}
	if maxIndex != 0 && minIndex != 0 {
		arr[maxIndex], arr[minIndex] = arr[minIndex], arr[maxIndex]
	}
	return arr
}
