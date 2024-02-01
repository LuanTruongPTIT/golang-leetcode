package main

import "fmt"

func isAnagram(s string, t string) bool {
	hash := map[rune]int{}
	for _, value := range s {
		hash[value]++
	}
	for idx, _ := range hash {
		fmt.Println(idx)
	}
	for _, value := range t {
		hash[value]--
	}
	for _, value := range hash {
		if value != 0 {
			return false
		}
	}
	return true
}
func main() {
	result := isAnagram("listen", "slient")
	fmt.Println(result)
}
