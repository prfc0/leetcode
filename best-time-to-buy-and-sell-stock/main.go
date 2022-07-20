package main

/*
https://leetcode.com/problems/best-time-to-buy-and-sell-stock/
You are given an array prices where prices[i] is the price of a given stock on the ith day.
You want to maximize your profit by choosing a single day to buy one stock and choosing a different day in the future to sell that stock.
Return the maximum profit you can achieve from this transaction. If you cannot achieve any profit, return 0.
Example 1:
Input: prices = [7,1,5,3,6,4]
Output: 5
Explanation: Buy on day 2 (price = 1) and sell on day 5 (price = 6), profit = 6-1 = 5.
Note that buying on day 2 and selling on day 1 is not allowed because you must buy before you sell.
*/

import "fmt"

func maxProfit(prices []int) int {
	N := len(prices)
	maxPrices := make([]int, N)
	max := -1
	// maxPrices[N-1] = max
	for i := N - 1; i >= 0; i-- {
		if prices[i] < max {
			maxPrices[i] = max
		} else {
			maxPrices[i] = prices[i]
		}
		if prices[i] > max {
			max = prices[i]
		}
	}
	ans := 0
	for i := 0; i < N; i++ {
		profit := maxPrices[i] - prices[i]
		if profit > ans {
			ans = profit
		}
	}
	// fmt.Println(maxPrices)
	return ans
}

func main() {
	// prices := []int{7, 1, 5, 3, 6, 4}
	prices := []int{1, 2, 3, 4, 5, 6, 7, 6, 5, 4, 3, 2, 1}
	// prices := []int{1, 2}
	fmt.Println(maxProfit(prices))
}
