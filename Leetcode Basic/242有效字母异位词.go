package main

import (
	"fmt"
)

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	count := [26]int{}
	for i := 0; i < len(s); i++ {
		count[s[i]-'a']++ //某个字符，s里出现的次数，+1
		count[t[i]-'a']-- //某个字符，t里出现的次数，-1
		//在s和t里出现次数相同那么count就是0
	}
	for _, c := range count {
		if c != 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isAnagram("anagram", "nagaram"))
}
