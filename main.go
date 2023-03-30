package main

func main() {
	println(countVowelStrings(2))
}

func countVowelStrings(n int) int {
	vowel := []int{
		1, 1, 1, 1, 1,
	}
	for i := 1; i < n; i++ {
		for j := 1; j < 5; j++ {
			vowel[j] += vowel[j-1]
		}
	}
	ret := 0
	for _, v := range vowel {
		ret += v
	}
	return ret

}
