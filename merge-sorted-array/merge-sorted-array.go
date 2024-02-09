package main

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int) []int {
	i, j, k := m-1, n-1, m+n-1
	for j >= 0 {
		if i >= 0 && nums1[i] > nums2[j] {
			nums1[k] = nums1[i]
			i--
		} else {
			nums1[k] = nums2[j]
			j--
		}
		k--
	}
	return nums1
}

func main() {
	nums := merge([]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3)
	fmt.Println(nums)
}
