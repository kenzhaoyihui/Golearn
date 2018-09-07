package main

import (
	"fmt"
)

func lengthOfNonRepeat(s string) int {
	lastOccurred := make(map[rune]int)
	start := 0
	maxLength := 0
	for i, ch := range []rune(s) {
		lastI, ok := lastOccurred[ch]
		if ok && lastI >= start {
			start = lastOccurred[ch] + 1
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i
	}
	return maxLength
}

func main() {
	fmt.Println(lengthOfNonRepeat("abcabccc"))
	fmt.Println(lengthOfNonRepeat("bbbbb"))
	fmt.Println(lengthOfNonRepeat("pwwkew"))
	fmt.Println(lengthOfNonRepeat(""))
	fmt.Println(lengthOfNonRepeat("b"))
	fmt.Println(lengthOfNonRepeat("abcdef"))
	fmt.Println(lengthOfNonRepeat("我爱你1我爱你"))
}
