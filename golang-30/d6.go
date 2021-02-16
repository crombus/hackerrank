package main

import (
	"bufio"
	"fmt"
	"os"
)

func input() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

func main() {
	var tc int
	fmt.Scan(&tc)
	for j := 0; j < tc; j++ {
		str := input()
		var st string
		var st1 string
		//st := string(str[0])
		//st1 := string(str[1])
		for i := 0; i < len(str); i += 2 {
			st = st + string(str[i])
		}
		for i := 1; i < len(str); i += 2 {
			st1 = st1 + string(str[i])
		}
		fmt.Printf("%s %s\n", st, st1)

	}
}

/*
accepted by hackerrank

package main

import (
	"bufio"
	"fmt"
	"os"
)


func main() {
	var tc int
	fmt.Scan(&tc)
    scanner := bufio.NewScanner(os.Stdin)
	for j := 0; j < tc; j++ {
        scanner.Scan()
		str := scanner.Text()
		st := ""
        st1 := ""

		//var st1 string
		//st := string(str[0])
		//st1 := string(str[1])
		for i := 0; i < len(str); i += 2 {
			st = st + string(str[i])
		}
		for i := 1; i < len(str); i += 2 {
			st1 = st1 + string(str[i])
		}

		fmt.Printf("%s %s", st, st1)
        fmt.Println()

	}
}
*/
