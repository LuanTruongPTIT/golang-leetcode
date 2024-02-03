package main

import "fmt"

/*
 * s = "cbaebabacd"
 * p = "abc"
 * s1 = "abab"
 * p1= "ab"
 */
func findAnagrams(s string, p string) []int {
	right, left, output, count := 0, 0, []int{}, len(p)
	var freq [256]int
	for _, value := range p {
		freq[value-'a']++
	}
	for right < len(s) {
		if freq[s[right]-'a'] >= 1 {
			count--
		}
		freq[s[right]-'a']--
		right++

		if count == 0 {
			output = append(output, left)
		}
		if right-left == len(p) {
			if freq[s[left]-'a'] >= 0 {
				count++
			}
			freq[s[left]-'a']++
			left++
		}
	}
	return output
}
func main() {
	result := findAnagrams("cbaebabacd", "abc")
	fmt.Println(result)
}
