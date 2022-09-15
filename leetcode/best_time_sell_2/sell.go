package best_time_sell_2

func MaxProfit(prices []int) int {
	if len(prices) == 1 {
		return 0
	}
	var max int
	for i := 0; i < len(prices)-1; i++ {
		if prices[i] < prices[i+1] {
			max += prices[i+1] - prices[i]
		}
	}
	return max
}
