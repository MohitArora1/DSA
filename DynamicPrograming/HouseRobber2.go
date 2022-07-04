package main

import "fmt"

/* this problem is rob cannot rob at adjacent home and first and last home is connected
befor reading this read HouseRobber first
*/

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
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

func rob2(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	if n == 1 {
		return nums[0]
	}

	return max(solveTab(nums[1:]), solveTab(nums[:n-1]))
}

func main() {
	nums := []int{1, 2, 3, 1, 5}
	fmt.Println(rob2(nums))
}
