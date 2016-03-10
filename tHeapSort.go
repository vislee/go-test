package main

import (
	"fmt"
)

func max_heap_fixdown(a []int, i, n int) {
	tmp := a[i]
	j := i<<1 + 1
	for j < n {
		if j+1 < n && a[j+1] > a[j] {
			j++
		}

		if a[j] < tmp {
			break
		}

		a[i] = a[j]
		i = j
		j = i<<1 + 1
	}

	a[i] = tmp
}

func make_max_heap(a []int) {
	for i := len(a) >> 1; i >= 0; i-- {
		max_heap_fixdown(a, i, len(a))
	}
}

func heap_sort(a []int) {
	make_max_heap(a)
	for i := len(a) - 1; i >= 0; i-- {
		a[i], a[0] = a[0], a[i]
		max_heap_fixdown(a, 0, i)
	}
}

func main() {
	a := []int{9, 1, 2, 8, 7, 3, 4, 6, 5, 0}
	fmt.Println(a)
	heap_sort(a)
	fmt.Println(a)
}
