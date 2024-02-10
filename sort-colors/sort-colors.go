package main

func sortColors(nums []int) []int {
	temp, size := 0, len(nums)-1
	for idx, _ := range nums {
		for i := 0; i < size-idx; i++ {
			if nums[i] > nums[i+1] {
				temp = nums[i]
				nums[i] = nums[i+1]
				nums[i+1] = temp
			}
		}
	}
	return nums
}

func main() {
	sortColors([]int{2, 0, 2, 1, 1, 0})
}
