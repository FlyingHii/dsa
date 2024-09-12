package main

import "fmt"

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
				}
	2nd loop left sum
		3nd loop right sum
*/

func productExceptSelf(nums []int) []int {
	mapNums := map[int]int{}
	uniqueNums := []int{}
	for i := 1; i < len(nums); i++ {
		_, ok := mapNums[i]
		if !ok {
			mapNums[nums[i]] = 1
			uniqueNums = append(uniqueNums, nums[i])
			continue
		}
		mapNums[i]++
	}
	rightSumNums := make([]int, len(uniqueNums))
	rightSumNums[len(uniqueNums)-1] = 1
	for i := len(uniqueNums) - 2; i >= 0; i-- {
		value := uniqueNums[i+1]
		count := mapNums[uniqueNums[i+1]]
		rightSumNums[i] = rightSumNums[i+1] * count * value
	}
	leftSumNums := make([]int, len(uniqueNums))
	leftSumNums[0] = 1
	for i := 1; i < len(uniqueNums); i++ {
		leftSumNums[i] = leftSumNums[i-1] * mapNums[uniqueNums[i-1]] * uniqueNums[i-1]
	}

	for i := 0; i < len(nums); i++ {
		value := nums[i]
		count := mapNums[nums[i]]
		nums[i] = value * (count - 1) * leftSumNums[i] * rightSumNums[i]
	}

	return nums
}
func main() {
	a := productExceptSelf([]int{1, 2, 3, 4})
	fmt.Print(a)
}
