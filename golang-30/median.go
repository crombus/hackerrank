package main

import (
	"fmt"
)

func maximum(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func minimum(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func findMedianSortedArrays(a []int, n int, b []int, m int) int {

	median := 0
	i := 0
	j := 0
	min_index := 0
	max_index := n

	fmt.Println("in funct")
	for min_index <= max_index {

		i = int((min_index + max_index) / 2)
		j = int(((n + m + 1) / 2) - i)
		fmt.Println("in loop")
		fmt.Println(min_index, max_index, median, i, j)

		if i < n && j > 0 && b[j-1] > a[i] {
			min_index = i + 1
		} else if i > 0 && j < m && b[j] < a[i-1] {
			max_index = i - 1
		} else {

			if i == 0 {
				median = b[j-1]
			} else if j == 0 {
				median = a[i-1]
			} else {
				median = maximum(a[i-1], b[j-1])
			}
			break
		}
	}

	if (n+m)%2 == 1 {
		return median
	}
	if i == n {
		return ((median + b[j]) / 2)
	}
	if j == m {
		return ((median + a[i]) / 2)
	}
	return ((median + minimum(a[i], b[j])) / 2)
}

func main() {
	a := []int{900}
	b := []int{10, 13, 14}
	n := len(a)
	m := len(b)

	if n < m {
		fmt.Println(findMedianSortedArrays(a, n, b, m))
	} else {
		fmt.Println(findMedianSortedArrays(b, m, a, n))
	}
}
