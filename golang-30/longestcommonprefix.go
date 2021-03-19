package main

import (
	"fmt"
	"math"
	"sort"
)

func longestCommonPrefix(a []string) string {

	size := len(a)

	// if size is 0, return empty string
	if size == 0 {
		return ""
	}

	if size == 1 {
		return a[0]
	}

	//sort the array of strings
	sort.Strings(a)

	//find the minimum length from
	//first and last string
	end := math.Min(float64(len(a[0])), float64(len(a[size-1])))

	//find the common prefix between
	//the first and last string
	i := 0
	for i < int(end) && a[0][i] == a[size-1][i] {
		i += 1
	}

	pre := a[0][0:i]
	return pre
}

func main() {
	input := []string{"hello", "hey", "hellical", "hell", "helix", "he"}
	fmt.Println(longestCommonPrefix(input))

}
