package main

import (
	"fmt"
	"math"
	"runtime"
	"sort"
	"strings"
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

// 334
func increasingTriplet_naive(nums []int) bool {

	for i := len(nums) - 1; i >= 0; i-- {
		for i2 := len(nums) - 1; i2 >= i; i2-- {
			for i3 := len(nums) - 1; i3 >= i; i3-- {
				/*
					if nums[i] < nums[i2] < nums[i3] {
						return true
					}
				*/
			}
		}
	}

	return false
}

// failed - solution
func increasingTripletFailed(nums []int) bool {
	if len(nums) < 3 {
		return false
	}
	lenNums := len(nums)
	maxRightITh := nums[lenNums-1] // max(nums[i2-1:])
	maxRightISe := nums[lenNums-2] // max(nums[i-1:])
	iMaxRightITh := lenNums - 1
	iMaxRightISe := lenNums - 2
	// iMaxRightISe < iMaxRightITh
	for i := lenNums - 3; i >= 0; i-- {
		if nums[i] < maxRightISe && maxRightISe < maxRightITh && iMaxRightISe < iMaxRightITh {
			return true
		}

		switch {
		case maxRightISe > maxRightITh:
			maxRightITh = maxRightISe
			iMaxRightITh = iMaxRightISe
			maxRightISe = nums[i]
			iMaxRightISe = i
			continue
		case nums[i] > maxRightISe:
			maxRightISe = nums[i]
			iMaxRightISe = i
			continue
		case iMaxRightISe >= iMaxRightITh, maxRightISe == maxRightITh, nums[i] == maxRightISe:
			continue
		}
	}

	return false
}

// greedy approach:
// why: no-backtracking, local optimal (first & second)
/*
mi and mid serve as placeholders to store the smallest and middle elements of the potential increasing triplet subsequence. By initializing them to infinity, it ensures that any number in the array will be smaller, allowing for proper updating of these variables.
The main algorithm unfolds within a single pass through the array of numbers (nums):

A for loop iterates through the nums array. For each num in nums, there is a series of checks and updates:
	if num > mid: This is the condition that tells us we have found a valid triplet. If the current num is greater than our mid value, then we already have a mi which is less than mid, and hence, we have found a sequence where mi < mid < num. We return True immediately.
	if num <= mi: If the current num is smaller than or equal to the current smallest value mi, it means that we can potentially start a new triplet sequence with this num as the smallest element, thus we update mi with the value of num.
	else: If the current num is greater than mi and not greater than mid, it fits between the mi and mid, so we update the mid to be num since it could potentially be the middle element of a valid triplet.
It's important to note that the code uses a greedy approach to always maintain the smallest possible values for mi and mid as it iterates over nums.
By consistently updating these with the smallest possible values at each step, it optimizes the chance of finding a valid triplet later in the array.
*/
func increasingTriplet_AI(nums []int) bool {
	first, second := int(^uint(0)>>1), int(^uint(0)>>1) // max int value

	for _, num := range nums {
		if num <= first {
			first = num // update first if num is smaller
		} else if num <= second {
			second = num // update second if num is between first and second
		} else {
			return true // found a number greater than both first and second
		}
	}

	return false // no valid triplet found
}

// main

// 49
func getRuneCountKey(s string) string {
	// Create a map to count the frequency of each rune
	runeCount := make(map[rune]int)
	for _, r := range s {
		runeCount[r]++
	}

	// To ensure the order doesn't matter, create a sorted string key
	var sortedRuneCount []string
	for r, count := range runeCount {
		sortedRuneCount = append(sortedRuneCount, fmt.Sprintf("%c:%d", r, count))
	}

	sort.Strings(sortedRuneCount) // Sort the key for consistency
	return strings.Join(sortedRuneCount, ",")
}

func groupAnagrams(strs []string) [][]string {
	showed := make(map[string][]string) // map to store anagram groups

	for _, value := range strs {
		// Get the "runes & count" key for the current string
		// TODO: no need to use "runes&count" as key => "sortedRunes" as key instead (a is anagrams of b => sortedA == sortedB)
		key := getRuneCountKey(value)

		// Add the string to the group corresponding to its "rune & count" key
		showed[key] = append(showed[key], value)
	}

	// Collect all the anagram groups
	var group [][]string
	for _, anagrams := range showed {
		group = append(group, anagrams)
	}

	return group
}

/*
36. Valid Sudoku
Medium
Topics
Companies
Determine if a 9 x 9 Sudoku board is valid. Only the filled cells need to be validated according to the following rules:

Each row must contain the digits 1-9 without repetition.
Each column must contain the digits 1-9 without repetition.
Each of the nine 3 x 3 sub-boxes of the grid must contain the digits 1-9 without repetition.
Note:

A Sudoku board (partially filled) could be valid but is not necessarily solvable.
Only the filled cells need to be validated according to the mentioned rules.

Example 1:

Input: board =
[["5","3",".",".","7",".",".",".","."]
,["6",".",".","1","9","5",".",".","."]
,[".","9","8",".",".",".",".","6","."]
,["8",".",".",".","6",".",".",".","3"]
,["4",".",".","8",".","3",".",".","1"]
,["7",".",".",".","2",".",".",".","6"]
,[".","6",".",".",".",".","2","8","."]
,[".",".",".","4","1","9",".",".","5"]
,[".",".",".",".","8",".",".","7","9"]]
Output: true
Example 2:

Input: board =
[["8","3",".",".","7",".",".",".","."]
,["6",".",".","1","9","5",".",".","."]
,[".","9","8",".",".",".",".","6","."]
,["8",".",".",".","6",".",".",".","3"]
,["4",".",".","8",".","3",".",".","1"]
,["7",".",".",".","2",".",".",".","6"]
,[".","6",".",".",".",".","2","8","."]
,[".",".",".","4","1","9",".",".","5"]
,[".",".",".",".","8",".",".","7","9"]]
Output: false
Explanation: Same as Example 1, except with the 5 in the top left corner being modified to 8. Since there are two 8's in the top left 3x3 sub-box, it is invalid.

Constraints:

board.length == 9
board[i].length == 9
board[i][j] is a digit 1-9 or '.'.
----------
*/
//['1','2',...]
func check19_without_repetition(cellValues [9]byte) bool {
	cellValuesSet := map[byte]bool{}
	for _, cellValue := range cellValues {
		if cellValue == '.' || cellValue == 0 {
			continue
		}
		if cellValuesSet[cellValue] == true {
			return false
		}
		cellValuesSet[cellValue] = true
	}
	return true
}

func getX33sIJ(location [2]int) [][2]int {
	ret := make([][2]int, 0, 9)
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			ret = append(ret, [2]int{location[0] + i, location[1] + j})
		}
	}
	return ret
}
func isValidSudoku(board [][]byte) bool {
	x33sStartLoc := [9][2]int{{0, 0}, {3, 0}, {6, 0}, {0, 3}, {3, 3}, {6, 3}, {0, 6}, {3, 6}, {6, 6}}
	for i := 0; i < 9; i++ {
		var row [9]byte // Array of 9 bytes
		copy(row[:], board[i])
		if check19_without_repetition(row) == false {
			return false
		}

		col := [9]byte{}
		for colIdx := 0; colIdx < 9; colIdx++ {
			col[colIdx] = board[colIdx][i]
		}
		if check19_without_repetition(col) == false {
			return false
		}

		x33 := [9]byte{}
		for in, location := range getX33sIJ(x33sStartLoc[i]) {
			x33[in] = board[location[0]][location[1]]
		}
		if check19_without_repetition(x33) == false {
			return false
		}
	}

	return true
}

/*
128. Longest Consecutive Sequence
Medium
Topics
Companies
Given an unsorted array of integers nums, return the length of the longest consecutive elements sequence.

You must write an algorithm that runs in O(n) time.

Example 1:

Input: nums = [100,4,200,1,3,2]
Output: 4
Explanation: The longest consecutive elements sequence is [1, 2, 3, 4]. Therefore its length is 4.
Example 2:

Input: nums = [0,3,7,2,5,8,4,6,0,1]
Output: 9

Constraints:

0 <= nums.length <= 105
-109 <= nums[i] <= 109

------------
*/
func longestConsecutive(nums []int) int {

	//-----------0
	// []consecutive int array
	//rets is [][]int

	//ret is longest among rets
	//&& for retVal = ret -> if retVal[i] !== retVal[i-1]+1 => isconsecutive = false

	//-----------0.1
	// convert nums to set map[num]
	minNum := int(math.Pow(10, 9))
	numsSet := map[int]bool{}
	for _, num := range nums {
		numsSet[num] = true
		if num < minNum {
			minNum = num
		}
	}

	//build endStartMap with consecutive
	//-----------0.1
	// map[end]start
	//endStartMap:=map[int]int
	// map[start]end
	//startEndMap:=map[int]int
	//for num,_ = range numsSet {//i 0 num 8
	//-----------0.1
	//if endStartMap[num] not exist {
	//  if endStartMap[num-1] not exist {
	//    endStartMap[num]=num //num or nums[0]??
	//  } else {
	//    //increase end
	//    endStartMap[num]=endStartMap[num-1]
	//    delete(endStartMap[num-1])
	//  }
	//
	//} else {
	//  //endStartMap[num]=start
	//}
	//}

	//-----------0.2
	maxStreak := 0
	streak := 1
	// stop condition??
	// 1 2 3 5 6
	start := minNum
	end := minNum
	for i := 1; len(numsSet) != 0; i++ {
		//end of consequentive, update new max, reset streak, choose new init item
		if numsSet[end+i] == false && numsSet[start-i] == false {
			//set new init item
			for key := range numsSet {
				start = key
				end = key
				break
			}
			delete(numsSet, start)
			streak = 1
			i = 0
		}

		//expand right
		if numsSet[end+i] == true {
			streak++
			delete(numsSet, end+i)
		}
		//expand left
		if numsSet[start-i] == true {
			streak++
			delete(numsSet, start-i)
		}
		if streak > maxStreak {
			maxStreak = streak
		}
	}

	// find ret is longest among endStartMap
	//maxLength:=0
	//for end,start = range endStartMap {
	//	if end-start > maxLength {
	//		maxLength=length
	//	}
	//}
	return maxStreak
}
func main() {
	memBefore := getMemStats()
	start := time.Now()

	board := [][]byte{
		{'.', '.', '.', '.', '5', '.', '.', '1', '.'},
		{'.', '4', '.', '3', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '3', '.', '.', '1'},
		{'8', '.', '.', '.', '.', '.', '.', '2', '.'},
		{'.', '.', '2', '.', '7', '.', '.', '.', '.'},
		{'.', '1', '5', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '2', '.', '.', '.'},
		{'.', '2', '.', '9', '.', '.', '.', '.', '.'},
		{'.', '.', '4', '.', '.', '.', '.', '.', '.'},
	}
	a := isValidSudoku(board)

	println(a)

	elapsed := time.Since(start)
	// Capture memory stats after the function call
	memAfter := getMemStats()

	fmt.Println("------ Algo Statistic ----------")
	fmt.Printf("Function took %s\n", elapsed)
	fmt.Printf("Memory used by function: %v bytes\n", memAfter.Alloc-memBefore.Alloc)
	fmt.Println("------ End Algo Statistic ----------")
}
func getMemStats() runtime.MemStats {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	return m
}
