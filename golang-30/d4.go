package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type person struct {
	age int
}

func (p person) intialise(age int) *person {
	if age < 0 {
		p.age = 0
		fmt.Println("Age is not valid, setting age to 0.")
	} else {
		p.age = age
	}
	return &p
}

func (p person) printage() {
	if p.age < 13 {
		fmt.Println("You are young")
	} else if p.age >= 13 && p.age < 18 {
		fmt.Println("You are teenager")
	} else {
		fmt.Println("You are old")
	}
}

func (p person) yearpasses() *person {
	p.age += 3
	return &p
}

func input() int {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	age, _ := strconv.Atoi(scanner.Text())
	return age
}

func main() {
	tc := input()
	for i := 0; i < tc; i++ {
		age := input()
		per := new(person)
		per = per.intialise(age)
		per.printage()
		per = per.yearpasses()
		per.printage()
		fmt.Println()
	}

}
