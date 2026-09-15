func topKFrequent(nums []int, k int) []int {
	counts := map[int]int{}

	for _, num := range nums {
			counts[num]++
	}

	bucket := make([][]int, len(nums)+1);

	for num, count := range counts {
		bucket[count] = append(bucket[count], num)
	}

	res := make([]int, 0, k);
	for i := len(bucket) - 1; i >= 0; i-- {
		for _, b := range bucket[i] {
			res = append(res, b)
			if len(res) == k {
				return res
			}
		}
	}
	return res
}
