package main

import (
	"fmt"
)

func swap(n1 *int, n2 *int) {
	x := *n1
	*n1 = *n2
	*n2 = x
}

func insertionsort(arr *[]int, n int) {
	for i := 1; i < n; i++ {
		key := (*arr)[i]
		j := i - 1
		for j >= 0 && (*arr)[j] > key {
			(*arr)[j+1] = (*arr)[j]
			j--
		}
		(*arr)[j+1] = key
	}

}

func selectionsort(arr *[]int, n int) {
	for i := 0; i < n-1; i++ {
		minid := i
		for j := 1 + i; j < n; j++ {
			if (*arr)[j] < (*arr)[minid] {
				minid = j
			}
			swap(&(*arr)[i], &(*arr)[minid])
		}
	}
}

func main() {
	arr := []int{4, 7, 9, 3, 4, 1}
	insertionsort(&arr, len(arr))
	for i := 0; i < 6; i++ {
		fmt.Println(arr[i])
	}
	selectionsort(&arr, len(arr))
	for i := 0; i < 6; i++ {
		fmt.Println(arr[i])
	}
}
