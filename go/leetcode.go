package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

func mergeAlternately(word1 string, word2 string) string {
	var retByte []byte
	for i := 0; i < 100; i++ {
		if i > len(word1) && i > len(word2) {
			break
		}
		if i < len(word1) {
			retByte = append(retByte, word1[i])
		}
		if i < len(word2) {
			retByte = append(retByte, word2[i])
		}
	}
	return string(retByte)
}

// 345
/**
xx xx xx xx
*/
func reverseVowels(s string) string {
	vowels := map[byte]bool{
		'a': true, 'e': true, 'i': true, 'o': true, 'u': true,
		'A': true, 'E': true, 'I': true, 'O': true, 'U': true,
	} // use map (for create set) for quick lookup
	left := 0
	right := len(s) - 1
	runes := []byte(s) // use byte instead of rune
	for left < right {
		if !vowels[runes[right]] {
			right--
			continue
		}

		if !vowels[runes[left]] {
			left++
			continue
		}
		runes[right], runes[left] = runes[left], runes[right]
		right--
		left++
	}
	return string(runes)
}

// 151
func reverseWords(s string) string {
	var words []byte
	var word []byte
	for i := 0; i < len(s); i++ {
		if (s)[i] != ' ' {
			word = append(word, (s)[i])
		}

		if ((s)[i] == ' ' || i == len(s)-1) && len(word) != 0 {
			if len(words) > 0 {
				words = append([]byte{' '}, words...)
			}
			words = append(word, words...)
			word = []byte{}
		}
	}

	return string(words)
}

// 238
/*
xxxxx
12345

1    5    ... 20
xxxx xxxx
2345 1345

{
	1:total
	2:total
	3:total
}
*/
// TODO
/*
			1st loop 	 {
					1:total
					2:total
					3:total
					4:total
				}
	2nd loop left sum
		3nd loop right sum
*/

func productExceptSelf(nums []int) []int {
	mapNums := map[int]int{}
	uniqueNums := []int{}
	for i := 0; i < len(nums); i++ {
		_, ok := mapNums[nums[i]]
		if !ok {
			mapNums[nums[i]] = 1
			uniqueNums = append(uniqueNums, nums[i])
			continue
		}
		mapNums[nums[i]]++
	}

	rightSumNums := make(map[int]int, len(uniqueNums))
	rightestUniqueNums := uniqueNums[len(uniqueNums)-1]
	rightSumNums[rightestUniqueNums] = 1
	for i := len(uniqueNums) - 2; i >= 0; i-- {
		// si = ni-1 * ni-1-1 *...
		// si = ni-1 * ni-1-1 *...
		// si = ni-1 * si-1
		curVal := uniqueNums[i]
		lastVal := uniqueNums[i+1]
		lastCnt := mapNums[lastVal]
		lastSum := rightSumNums[lastVal]
		rightSumNums[curVal] = int(math.Pow(float64(lastVal), float64(lastCnt))) * lastSum
	}
	leftSumNums := make(map[int]int, len(uniqueNums))
	leftestUniqueNums := uniqueNums[0]
	leftSumNums[leftestUniqueNums] = 1
	for i := 1; i < len(uniqueNums); i++ {
		// si = ni-1 * ni-1-1 *...
		// si = ni-1 * ni-1-1 *...
		// si = ni-1 * si-1
		curVal := uniqueNums[i]
		lastVal := uniqueNums[i-1]
		lastCnt := mapNums[lastVal]
		lastSum := leftSumNums[lastVal]
		leftSumNums[curVal] = lastSum * int(math.Pow(float64(lastVal), float64(lastCnt)))
	}

	for i := 0; i < len(nums); i++ {
		value := nums[i]
		count := mapNums[value]
		nums[i] = int(math.Pow(float64(value), float64(count-1))) * leftSumNums[value] * rightSumNums[value]
	}

	return nums
}
func productExceptSelf_AI(nums []int) []int {
	length := len(nums)
	result := make([]int, length)

	// Initialize result array with 1
	for i := range result {
		result[i] = 1
	}

	// Calculate left products
	leftProd := 1
	for i := 0; i < length; i++ {
		result[i] = leftProd
		leftProd *= nums[i]
	}

	// Calculate right products and finalize result
	rightProd := 1
	for i := length - 1; i >= 0; i-- {
		result[i] *= rightProd
		rightProd *= nums[i]
	}

	return result
}

// https://github.com/doocs/leetcode/blob/main/solution/0200-0299/0238.Product%20of%20Array%20Except%20Self/README_EN.md
func productExceptSelf_online(nums []int) []int {
	n := len(nums)
	ans := make([]int, n)
	left, right := 1, 1
	for i, x := range nums {
		ans[i] = left
		left *= x
	}
	for i := n - 1; i >= 0; i-- {
		ans[i] *= right
		right *= nums[i]
	}
	return ans
}
func main() {
	memBefore := getMemStats()
	start := time.Now()

	a := productExceptSelf([]int{5, 9, 2, -9, -9, -7, -8, 7, -9, 10})

	elapsed := time.Since(start)
	// Capture memory stats after the function call
	memAfter := getMemStats()

	fmt.Println("------ Algo Statistic ----------")
	fmt.Printf("Function took %s\n", elapsed)
	fmt.Printf("Memory used by function: %v bytes\n", memAfter.Alloc-memBefore.Alloc)
	fmt.Println("------ End Algo Statistic ----------")
	b := []int{-51438240, -28576800, -128595600, 28576800, 28576800, 36741600, 32148900, -36741600, 28576800, -25719120}
	fmt.Println(a)
	fmt.Println(slicesEqual(a, b))
}
func slicesEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
func getMemStats() runtime.MemStats {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	return m
}
