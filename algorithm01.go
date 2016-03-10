package main

import (
	"fmt"
)

// 二分查找
func binary_search(array []int, l, value int) int {
	left, right := 0, l-1
	middle := -1
	for left <= right {
		middle = left + (right-left)>>1
		if array[middle] > value {
			right = middle - 1
		} else if array[middle] < value {
			left = middle + 1
		} else {
			return middle
		}
	}
	return -1
}

// 快速排序
func quick_sort(array []int, left, right int) {
	if left >= right {
		return
	}

	m := left + (right-left)>>1
	array[left], array[m] = array[m], array[left]
	i, j, x := left, right, array[left]
	for i < j {
		for i < j && array[j] > x {
			j--
		}
		if i < j {
			array[i] = array[j]
			i++
		}

		for i < j && array[i] < x {
			i++
		}
		if i < j {
			array[j] = array[i]
			j--
		}
		array[i] = x
		quick_sort(array, left, i-1)
		quick_sort(array, i+1, right)
	}
}

// 插入排序
func insert_sort(a []int) {
	l := len(a)
	for i := 1; i < l; i++ {
		k := a[i]
		j := 0
		for j = i - 1; j >= 0 && a[j] > k; j-- {
			a[j+1] = a[j]
		}
		a[j+1] = k
	}
}

// 冒泡排序
func bubble_sort(a []int) {
	var i, j int
	f := true
	for i = 0; i < len(a) && f; i++ {
		f = false
		for j = len(a) - 1; j > i; j-- {
			if a[j-1] > a[j] {
				a[j-1], a[j] = a[j], a[j-1]
				f = true
			}
		}
	}
}

func main() {
	test := []int{5, 3, 2, 4, 1, 6, 7, 15, 8, 10, 30, 9, 18, 23, 32, 21, 24, 27, 29}
	quick_sort(test, 0, 18)
	fmt.Println(test)

	idx := binary_search(test, len(test), 3)
	fmt.Println(idx)
}
