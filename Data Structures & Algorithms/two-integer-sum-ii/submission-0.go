func twoSum(numbers []int, target int) []int {
	left := 0
	right := len(numbers) - 1

	for left < right {
		val := numbers[left] + numbers[right]
		if val == target {
			return []int{left + 1, right + 1}
		} else if val > target {
			right--
		} else {
			left++
		}

	}
	return []int{}
}
