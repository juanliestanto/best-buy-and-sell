package main

import "fmt"

func maxProfit(k int, prices []int) (int, int) {
	n := len(prices)
	transactions := 0

	if n <= 1 {
		return 0, transactions
	}

	if k >= n/2 {
		maxProfit := 0
		for i := 1; i < n; i++ {
			if prices[i] > prices[i-1] {
				maxProfit += prices[i] - prices[i-1]
				transactions++
			}
		}
		return maxProfit, transactions
	}

	dayPrices := make([][]int, k+1)
	for i := range dayPrices {
		dayPrices[i] = make([]int, n)
	}

	for i := 1; i <= k; i++ {
		maxPrev := -prices[0]
		for j := 1; j < n; j++ {
			dayPrices[i][j] = max(dayPrices[i][j-1], maxPrev+prices[j])
			maxPrev = max(maxPrev, dayPrices[i-1][j]-prices[j])
		}
	}

	return dayPrices[k][n-1], k
}

func main() {
	prices := []int{6, 1, 3, 2, 4, 7}
	transactions := 10

	profit, numTransactions := maxProfit(transactions, prices)
	fmt.Println("Maximum Profit : ", profit)
	fmt.Println("Number Of Transactions : ", numTransactions)
}
