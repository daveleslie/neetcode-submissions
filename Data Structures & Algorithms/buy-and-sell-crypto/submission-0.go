func maxProfit(prices []int) int {
	var maxProfit int
	buy, sell := 0, 1
	for sell < len(prices) {
		profit := prices[sell] - prices[buy]
		if profit > 0 {
			maxProfit = max(maxProfit, profit)
		} else if profit < 0 {
			buy = sell
		}
		sell++
	}
	return maxProfit
}
