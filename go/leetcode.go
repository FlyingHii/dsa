package _go

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

     4
xxxx xxxx
2345 1345

*/
func productExceptSelf(nums []int) []int {
	lenght := len(nums)
	convertedNums := []int{}
	for i := 0; i < lenght; i++ {
		convertedNums = append(convertedNums, nums[:i]...)   // elements before i
		convertedNums = append(convertedNums, nums[i+1:]...) // elements after i
	}

	nums = []int{}
	multiply := 1
	for i := 0; i < len(convertedNums); i++ {
		if i%lenght-1 != 0 {
			multiply = multiply * convertedNums[i]
			continue
		}

		nums = append(nums, multiply)
		multiply = 1
	}

	return nums
}
func main() {
	productExceptSelf([]int{1, 2, 3, 4})
}
