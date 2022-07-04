package main

import "fmt"

/*
	Kadane's Algo to find the largest sum of sub array
	Kadane's Algo is DP solution
*/

// 1,-2,3,-1,5

//             5
//         -1  4
//       3  2  7
//   -2  1  0  5
// 1 -1  2  1  6

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func findSum(nums []int, n int) int {
	if n == 1 {
		return nums[0]
	}
	val := findSum(nums, n-1)
	return max(val+nums[n-1], nums[n-1])
}

func findSum4(nums []int, n int) (int, []int) {
	if n == 1 {
		return nums[0], []int{nums[0]}
	}
	val, arr := findSum4(nums, n-1)
	if val+nums[n-1] > nums[n-1] {
		arr = append(arr, nums[n-1])
		return val + nums[n-1], arr
	}
	return nums[n-1], []int{nums[n-1]}
}

func findSum2(nums []int, n int) int {
	if n == 1 {
		return nums[0]
	}
	return max(findSum2(nums, n-1)+nums[n-1], nums[n-1])
}

//Kadane's algo
func findSum3(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	current := 0
	maxVal := 0

	for i := 0; i < n; i++ {
		current = max(current+nums[i], nums[i])
		maxVal = max(maxVal, current)
	}
	return maxVal
}

func main() {
	nums := []int{1, -2, 3, -1, 5}
	// fmt.Println(findSum2(nums, 5))

	//fmt.Println(findSum3(nums))

	val, arr := findSum4(nums, 5)
	fmt.Println("val : ", val)
	fmt.Println("arr : ", arr)
}
