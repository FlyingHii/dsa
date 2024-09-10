package _go

func kidsWithCandies(candies []int, extraCandies int) (ret []bool) {
	// Step 1: Find the current maximum number of candies any kid has
	maxCandies := 0
	for _, candy := range candies {
		if candy > maxCandies {
			maxCandies = candy
		}
	}

	// Step 2: Create a result slice to store the boolean values
	result := make([]bool, len(candies))

	// Step 3: Check if giving each kid the extra candies makes them have the most candies
	for i, candy := range candies {
		result[i] = candy+extraCandies >= maxCandies
	}

	return result
}
