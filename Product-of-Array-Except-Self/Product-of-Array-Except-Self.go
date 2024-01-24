/*
Truong Dinh Kim Luan
Given an integer array nums, return an array answer such that answer[i] is equal to the product of all the elements of nums except nums[i].

The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer.

You must write an algorithm that runs in O(n) time and without using the division operation.

Example 1:

Input: nums = [1,2,3,4]
Output: [24,12,8,6]
Example 2:

Input: nums = [-1,1,0,-3,3]
Output: [0,0,9,0,0]
*/
package main

import (
	"fmt"
)

func productExceptSelf(nums []int) []int {
	size := len(nums)
	res := make([]int, size)
	var leftMult int = 1
	var rightMult int = 1
	for i := size - 1; i >= 0; i-- {
		res[i] = rightMult
		rightMult *= nums[i]
	}
	for j := 0; j < size; j++ {
		res[j] *= leftMult
		leftMult *= nums[j]
	}
	return res
}
func main() {
	arr := productExceptSelf([]int{1, 2, 3, 4})
	fmt.Println(arr)
}
