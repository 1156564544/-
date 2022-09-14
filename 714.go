package Question_by_day

func maxProfit(prices []int, fee int) int {
	meiyou := 0
	you := -prices[0]
	n := len(prices)
	for i := 1; i < n; i++ {
		_meiyou, _you := meiyou, you
		meiyou = max(_meiyou, _you+prices[i]-fee)
		you = max(_you, _meiyou-prices[i])
	}
	return meiyou
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
