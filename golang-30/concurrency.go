package main

import (
	"fmt"
)

func goro() int {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	return sum
}

func main() {
	//sum := make(chan int)
	var sum1 int
	go func() {
		var x int
		for i := 0; i < 10; i++ {
			x += i
			sum1 += i
		}
		//sum <- x
	}()
	fmt.Println(sum1)
}
