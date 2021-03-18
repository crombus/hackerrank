package main

import (
	"fmt"
)

func main() {

	set := make(map[int]struct{})
	set[1] = struct{}{}
	set[2] = struct{}{}
	set[1] = struct{}{}
	// ...

	for key := range set {
		fmt.Println(key)
	}
	// each value will be printed only once, in no particular order

	// you can use the ,ok idiom to check for existing keys
	if _, ok := set[1]; ok {
		fmt.Println("element found")
	} else {
		fmt.Println("element not found")
	}
}
