package main

import (
	"fmt"
	"math"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func solve(nums []int, n int, dp []int) int {

	if n >= len(nums)-1 {
		return 0
	}

	if dp[n] != -1 {
		return dp[n]
	}
	var m = math.MaxInt
	for i := n + 1; i <= min(len(nums)-1, n+nums[n]); i++ {
		val := solve(nums, i, dp)
		if val < m {
			m = val
		}
	}
	if m == math.MaxInt {
		dp[n] = m
		return dp[n]
	}
	dp[n] = m + 1
	return dp[n]
}

func jump(nums []int) int {
	n := len(nums)
	var dp = make([]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = -1
	}
	val := solve(nums, 0, dp)
	if val == math.MaxInt {
		return 0
	}
	return val
}

func main() {
	nums := []int{2, 3, 1, 1, 4}
	fmt.Println(jump(nums))
}
