package main

import "fmt"

/*
rob cannot rob at adjasent homes and we have to return the max he can rob
*/

// Max function
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Recursion + Memoization
func solveMem(nums []int, n int, dp []int) int {
	if n < 0 {
		return 0
	}
	if dp[n] != -1 {
		return dp[n]
	}
	// Recursive relation
	// Rober can rob (n-2)th if he rob from current house or (n-1)th house if he has not rob current house
	dp[n] = max(solveMem(nums, n-1, dp), solveMem(nums, n-2, dp)+nums[n])
	return dp[n]
}

// Iterative + Tabulation method
func solveTab(nums []int) int {
	n := len(nums)
	prev2 := 0
	prev1 := nums[0]

	for i := 1; i < n; i++ {
		ans := max(prev1, prev2+nums[i])
		prev2 = prev1
		prev1 = ans
	}
	return prev1
}

func rob(nums []int) int {
	// n := len(nums)
	// var dp = make([]int, n)
	// for i:=0;i<n;i++{
	//     dp[i] = -1
	// }
	// return solveMem(nums,n-1,dp)

	return solveTab(nums)
}

func main() {
	nums := []int{1, 2, 3, 1}
	fmt.Println(rob(nums))
}
