package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "THIS IS A TEST TEXT"
	fmt.Println(strings.Contains(str, "TeST"))
	fmt.Println(strings.Count(str, "TEST"))
	fmt.Println(str[:1])

	str1 := "AABAACAADAABAABA"
	pat := "AABA"

	h := true
	for h {
		in := strings.Index(str1, pat)
		str1 = str1[in+len(pat):]
		fmt.Println(in)
		if !strings.Contains(str1, pat) {
			h = false
		}
	}
}

func IsLetter(s string) bool {
	for _, r := range s {
		if (r < 'a' || r > 'z') && (r < 'A' || r > 'Z') {
			return false
		}
	}
	return true
}
