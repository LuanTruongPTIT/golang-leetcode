/*
*
Truong Dinh Kim Luan
Input: prices = [7,1,5,3,6,4]
Output: 5
Explanation: Buy on day 2 (price = 1) and sell on day 5 (price = 6), profit = 6-1 = 5.
Note that buying on day 2 and selling on day 1 is not allowed because you must buy before you sell.
*/
package main

import "fmt"

func max(a int, b int) int {
	if a < b {
		return b
	}
	return a
}
func maxProfit(prices []int) int {
	// var left int = 0
	maxProfit := 0
	if len(prices) == 0 {
		return 0
	}
	var minPrices int = prices[0]
	for i := 1; i < len(prices); i++ {
		if prices[i] < minPrices {
			minPrices = prices[i]
		} else {
			max(maxProfit, prices[i]-minPrices)
		}
	}
	return maxProfit
}

func main() {
	a := maxProfit([]int{7, 1, 5, 6, 10})
	fmt.Println(a) // 5git
}
