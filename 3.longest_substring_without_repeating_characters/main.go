package main

import (
	"fmt"
)

func main() {
	fmt.Println(lengthOfLongestSubstring("dfghjkfghjgggggjjjgvjcccr"))
}

func lengthOfLongestSubstring(s string) int {
	m := make(map[rune]int)
	start, max := -1, 0

	for k, v := range s {
		//如果发现重复字符
		if last, ok := m[v]; ok && last > start {
			start = last
		}
		m[v] = k
		//保存最大值
		if k-start > max {
			max = k - start
		}
	}
	return max
}
