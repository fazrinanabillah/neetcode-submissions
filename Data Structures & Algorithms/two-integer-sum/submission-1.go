func twoSum(nums []int, target int) []int {
	elemIndex := make(map[int]int)
	for index1, elem1 := range nums {
		elem2 := target - elem1
		if index2, ok := elemIndex[elem2]; ok {
			return []int{index2, index1}
		}
		elemIndex[elem1] = index1
	}
	return []int{}
}
