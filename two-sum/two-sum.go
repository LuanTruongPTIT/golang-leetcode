package main

import (
	"fmt"
)

func twosum(nums []int, target int) []int {
	m := make(map[int]int)
	for k, v := range nums {
		if idx, ok := m[target-v]; ok {
			return []int{idx, k}
		}
		m[v] = k
	}
	return nil
}

func main() {
	fmt.Println("main")
	m := twosum([]int{1, 3, 4}, 7)
	fmt.Println(m)
}
