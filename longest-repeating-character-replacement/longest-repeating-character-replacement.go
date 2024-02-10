package main

import "fmt"

func characterReplacement(s string, k int) int {
	var freq [26]int // Sửa đổi kích thước mảng từ 256 thành 26
	large, start, maxLength := 0, 0, 0

	for end, value := range s {
		freq[value-'A']++ // Sửa đổi từ freq[value-'A'] thành freq[value-'A']
		large = max(large, freq[value-'A'])
		if end-start+1-large > k {
			freq[s[start]-'A']-- // Sửa đổi từ freq[start-'A'] thành freq[s[start]-'A']
			start++
		}
		maxLength = max(maxLength, end-start+1)
	}
	return maxLength
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(characterReplacement("AABABBA", 1))
}
