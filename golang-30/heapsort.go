package main

import (
	"fmt"
)

func swap(n1 *int, n2 *int) {
	x := *n1
	*n1 = *n2
	*n2 = x
}

func heapify(arr *[]int, n int, nu int) {
	largest := nu
	l := 2*nu + 1
	r := 2*nu + 2

	if l < n && (*arr)[l] > (*arr)[largest] {
		largest = l
	}
	if r < n && (*arr)[r] > (*arr)[largest] {
		largest = r
	}

	if largest != nu {
		swap(&(*arr)[largest], &(*arr)[nu])

		heapify(arr, n, largest)
	}
}

func heapsort(arr *[]int, n int, k int) {
	for i := n/2 - 1; i >= 0; i-- {
		heapify(arr, n, i)
	}

	for i := n - 1; i > 0; i-- {
		swap(&(*arr)[0], &(*arr)[i])

		heapify(arr, i, 0)
	}
}

func printsort(arr []int, s int) {
	for i := 0; i < s; i++ {
		fmt.Print(" ", arr[i])
	}
	fmt.Println()
}

func main() {
	var arr = []int{12, 11, 13, 5, 6, 7}
	printsort(arr, len(arr))
	heapsort(&arr, len(arr), 4)
	printsort(arr, 4)
}
