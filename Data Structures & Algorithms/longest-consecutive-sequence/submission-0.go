func longestConsecutive(nums []int) int {
	// ensure no duplicate element
	numSet := make(map[int]bool)
	for _, num := range nums {
		numSet[num] = true
	}

	longest := 0

	for num := range numSet {
		if !numSet[num - 1] {
			currentStreak := 1
			currentNum := num

			for numSet[currentNum + 1] {
				currentStreak++
				currentNum++
			}
			if longest < currentStreak {
				longest = currentStreak
			}
		}
	}

	return longest
}
