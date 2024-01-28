/*
Truong Dinh Kim Luan
Given an array of positive integers nums and a positive integer target, return the minimal length of a
subarray whose sum is greater than or equal to target. If there is no such subarray, return 0 instead.

Example 1:

Input: target = 7, nums = [2,3,1,2,4,3]
Output: 2
Explanation: The subarray [4,3] has the minimal length under the problem constraint.
Example 2:

Input: target = 4, nums = [1,4,4]
Output: 1
Example 3:

Input: target = 11, nums = [1,1,1,1,1,1,1,1]
Output: 0
*/
package main

import (
	"fmt"
)

func minSubArrayLen(target int, nums []int) int {
	left, sum, res := 0, 0, len(nums)+1
	for idx, value := range nums {
		sum += value
		for sum >= target {
			sum -= nums[left]
			res = min(res, idx-left+1)
			left++
		}
	}
	if res == len(nums)+1 {
		return 0
	}
	return res
}
func min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}
func main() {
	//case1
	a := minSubArrayLen(7, []int{2, 3, 1, 2, 4, 3})
	//case2
	b := minSubArrayLen(4, []int{1, 4, 4})
	//case3
	c := minSubArrayLen(11, []int{1, 1, 1, 1, 1, 1, 1, 1})
	fmt.Println(a, b, c) // 2,1,0
}
