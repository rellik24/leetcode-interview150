package top150

// fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
func MaxProfit(prices []int) int {
	var minValue int = prices[0]
	var profit int = 0
	for _, v := range prices {
		if v < minValue {
			minValue = v
		} else if v-minValue > profit {
			profit = v - minValue
		}
	}
	return profit
}
