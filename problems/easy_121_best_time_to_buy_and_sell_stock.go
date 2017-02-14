package problems

// MaxProfit represents
// Easy 121. Best Time to Buy and Sell Stock
// https://leetcode.com/problems/best-time-to-buy-and-sell-stock/
func MaxProfit(prices []int) int {
	var profit int

	if len(prices) == 0 {
		return profit
	}

	min := prices[0]

	for i := 1; i < len(prices); i++ {
		if prices[i] < min {
			min = prices[i]
		} else if tmp := prices[i] - min; tmp > profit {
			profit = tmp
		}
	}

	return profit
	// var maxCur int
	// var res int
	// for i := 1; i < len(prices); i++ {
	// 	maxCur += prices[i] - prices[i-1]
	// 	maxCur = int(math.Max(0, float64(maxCur)))
	// 	res = int(math.Max(float64(maxCur), float64(res)))
	// }
	// return res
}
