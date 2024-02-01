package main

/*
Truong Dinh Kim Luan
Given an integer array nums, find the
subarray
 with the largest sum, and return its sum.
Example 1:

Input: nums = [-2,1,-3,4,-1,2,1,-5,4]
Output: 6
Explanation: The subarray [4,-1,2,1] has the largest sum 6.
Example 2:

Input: nums = [1]
Output: 1
Explanation: The subarray [1] has the largest sum 1.
Example 3:

Input: nums = [5,4,-1,7,8]
Output: 23
Explanation: The subarray [5,4,-1,7,8] has the largest sum 23.
*/
import (
	"fmt"
)

func maxSubArray(nums []int) int {
	max, sum := nums[0], nums[0]
	for _, v := range nums[1:] {
		if sum < 0 {
			sum = v
		} else {
			sum += v
		}
		if max < sum {
			max = sum
		}
	}
	return max
}
func main() {
	a := maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})
	fmt.Println(a) // [-4,-1,2,1]
}
