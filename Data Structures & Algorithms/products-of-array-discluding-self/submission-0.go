func productExceptSelf(nums []int) []int {
	n := len(nums)
	res := make([]int, n)

	prefixVal := 1
	// loop kiri (prefix)
	for i := 0; i < n; i++ {
		res[i] = prefixVal
		prefixVal *= nums[i]
	}

	suffixVal := 1
	// loop kanan (suffix)
	for i := n - 1; i >= 0; i-- {
		res[i] *= suffixVal
		suffixVal *= nums[i]
	}
	return res
}
