func maxProfit(prices []int) int {
    buy := 0
	maxProfit:= 0

	for sell := 1; sell < len(prices); sell++ {
		if prices[buy] < prices[sell] {
			currentProfit := prices[sell] - prices[buy]
			if maxProfit < currentProfit {
				maxProfit = currentProfit
			}
		} else {
			buy = sell
		}
	}

	return maxProfit

}
