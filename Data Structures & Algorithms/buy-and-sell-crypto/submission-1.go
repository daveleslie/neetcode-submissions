func maxProfit(prices []int) int {
	var res int
	buy, sell := 0, 1
	for sell < len(prices) {
		if prices[sell] < prices[buy] {
			buy = sell
		}
		res = max(res, prices[sell] - prices[buy])
		sell++
	}
	return res
}
