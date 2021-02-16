package main

import (
	"fmt"
)

func input() int {
	var nu int
	fmt.Scan(&nu)
	return nu

}

func main() {
	nu := input()
	for i := 1; i <= 10; i++ {
		sum := nu * i
		fmt.Printf("%d X %d = %d\n", nu, i, sum)
	}
}
