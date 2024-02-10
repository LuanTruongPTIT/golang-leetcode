package main

import "fmt"

func findMins(nums []int) int {
	low, high := 0, len(nums)-1
	for low < high {
		mid := low + (high-low)/2
		fmt.Println(mid, high, low)
		if nums[mid] < nums[high] {
			high = mid
		} else if nums[mid] > nums[high] {
			low = mid + 1
		}
	}
	return nums[low]
}
func main() {
	fmt.Println(findMins([]int{3, 4, 5, 1, 2})) // 1
}
