package main

import (
	"fmt"
)

func mergesort(arr *[]int, l int, m int, r int) {
	n1 := m - l + 1
	n2 := r - m
	k := l

	var a1 = make([]int, n1)
	var a2 = make([]int, n2)
	for i := 0; i < n1; i++ {
		a1[i] = (*arr)[l+i]
	}
	for i := 0; i < n2; i++ {
		a2[i] = (*arr)[m+1+i]
	}

	i := 0
	j := 0

	for i < n1 && j < n2 {
		if a1[i] <= a2[j] {
			(*arr)[k] = a1[i]
			i++
		} else {
			(*arr)[k] = a2[j]
			j++
		}
		k++
	}
	for i < n1 {
		(*arr)[k] = a1[i]
		k++
		i++
	}
	for j < n2 {
		(*arr)[k] = a2[j]
		k++
		j++
	}

}

func merge(arr *[]int, l int, r int) {
	if l >= r {
		return
	}
	m := l + (r-l)/2
	merge(arr, l, m)
	merge(arr, m+1, r)
	mergesort(arr, l, m, r)

}

func main() {
	arr := []int{4, 7, 9, 3, 4, 1}
	merge(&arr, 0, len(arr)-1)
	for i := 0; i < 6; i++ {
		fmt.Println(arr[i])
	}
}
