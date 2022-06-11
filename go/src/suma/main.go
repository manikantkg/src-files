// suma project main.go
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World!")
	maxProfit(10, []int{1, 2, 3, 4, 5})

}
func maxProfit(k int, prices []int) {
	// Corner case
	if prices.length == 0 {
		return 0
	}

	// Corner case: reduce to 122 when k > prices / 2 (can buy as many as we want)
	if k > prices.length/2 {
		return quickSolve(prices)
	}

	// Init DP table
	var DP *[5]int = new([5]int)(prices.length)

	for n := 1; n < DP.length; n++ {
		// While we scan the prices, keep the min so far
		min := prices[0]

		for i := 1; i < DP[0].length; i++ {
			min = Math.min(min, prices[i]-DP[n-1][i-1])

			// DP[n][i] = max{DP[n][i - 1], prices[i] - prices[j] + DP[n - 1][j - 1]} for j in [0, i - 1]
			DP[n][i] = Math.max(DP[n][i-1], prices[i]-min)
		}
	}

	return DP[k][prices.length-1]
}
