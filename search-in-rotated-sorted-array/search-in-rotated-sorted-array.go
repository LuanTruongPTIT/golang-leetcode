// [4,5,6,7,0,1,2] target = 0
// high = 6 low = 0
package main

import (
	"fmt"
)

func search(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		mid := (low + high) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[low] <= nums[mid] {
			if nums[low] <= target && nums[mid] > target {
				high = mid - 1
			} else {
				low = mid + 1
			}
		} else {
			if nums[mid] <= target && target > nums[high] {
				low = low + 1
			} else {
				high = mid - 1
			}
		}
	}
	return -1
}
func main() {
	fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 0)) // 4
}
