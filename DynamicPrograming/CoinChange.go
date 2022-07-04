package main

import "fmt"

/*
Find the number of ways to change the coin by given coins are 2,3,5
*/

func changeCoin(n int, x, y, z int, dp []int) int {
	if n < 0 {
		return 0
	}
	if n == 0 {
		return 1
	}
	if dp[n] != -1 {
		return dp[n]
	}
	dp[n] = changeCoin(n-x, x, y, z, dp) + changeCoin(n-y, x, y, z, dp) + changeCoin(n-z, x, y, z, dp)
	return dp[n]
}

func main() {
	n := 7
	dp := make([]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = -1
	}
	fmt.Println(changeCoin(n, 2, 3, 5, dp))
}
