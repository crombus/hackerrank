package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func input() int {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	number, _ := strconv.Atoi(scanner.Text())
	return number
}

func conditional(number int) {
	if number == 0 || number == 1 || number%2 != 0 {
		fmt.Println("Weird")
	}
	if number%2 == 0 {
		if number >= 6 && number < 21 {
			fmt.Println("Weird")
		} else {
			fmt.Println("Not Weird")
		}
	}
}

func main() {
	number := input()
	conditional(number)
}
