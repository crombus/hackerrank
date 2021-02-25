package main

import (
	"fmt"
)

func value(char string) int {
	val := map[string]int{
		"M": 1000,
		"D": 500,
		"C": 100,
		"L": 50,
		"X": 10,
		"V": 5,
		"I": 1,
	}
	return val[char]
}

func answer(val string) int {
	ans := 0
	prev := 0
	for s, i := range val {
		if value(fmt.Sprintf("%c", i)) >= prev {
			ans += value(fmt.Sprintf("%c", i))
		} else {
			ans -= value(fmt.Sprintf("%c", i))
		}
		prev = value(fmt.Sprintf("%c", i))
		fmt.Printf("%3d  %05d\n", s, i)
	}
	return ans
}

func main() {
	val := "MCMIV"

	fmt.Println(answer(val))
}
