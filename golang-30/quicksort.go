package main

import (
	"fmt"
)

func swap(n1 *int, n2 *int) {
	x := *n1
	*n1 = *n2
	*n2 = x
}

func partition(arr *[]int, low int, high int) int {
	pivot := (*arr)[high]
	i := (low - 1)

	for j := low; j <= high-1; j++ {
		if (*arr)[j] < pivot {
			i++
			swap(&(*arr)[i], &(*arr)[j])
		}
	}
	swap(&(*arr)[i+1], &(*arr)[high])
	return (i + 1)
}

func quickSort(arr *[]int, low int, high int) {
	if low < high {

		pi := partition(arr, low, high)

		quickSort(arr, low, pi-1)
		quickSort(arr, pi+1, high)
	}
}

func main() {
	arr := []int{4, 7, 9, 3, 4, 1}
	quickSort(&arr, 0, len(arr)-1)
	for i := 0; i < 6; i++ {
		fmt.Println(arr[i])
	}
}
