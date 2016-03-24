package main

import (
	"fmt"
)

type node struct {
	data int
	flag bool
}

// 螺旋打印二维数组
func lxprt(a [4][4]int) {
	i, j := 0, 0
	m, n := 3, 3
	for i <= m && j <= n {
		for k := j; k <= n; k++ {
			fmt.Println(a[i][k])
		}
		i += 1

		for k := i; k <= m; k++ {
			fmt.Println(a[k][n])
		}
		n -= 1

		for k := n; k >= j; k-- {
			fmt.Println(a[m][k])
		}
		m -= 1

		for k := m; k >= i; k-- {
			fmt.Println(a[k][j])
		}
		j += 1
	}
}

// 蛇形打印二维数组
func sxprt(a [4][4]int) {
	for i := 0; i <= 3; i++ {
		if i%2 == 0 {
			for k := 0; k <= 3; k++ {
				fmt.Println(a[i][k])
			}
		} else {
			for k := 3; k >= 0; k-- {
				fmt.Println(a[i][k])
			}
		}
	}
}

// 从二维数组中查找一个连续的串，斜对角的值不算连续，不能循环
func lookupXY(a [4][4]node, x, y, m, n int, b []int, k int) bool {
	if k+1 >= len(b) {
		return true
	}
	fmt.Printf("%d,%d:%d\n", x, y, b[k+1])

	if x < m && y < n {
		if x > 0 && a[x-1][y].data == b[k+1] && !a[x-1][y].flag {
			fmt.Println("1==>", x-1, y, b[k+1])
			a[x-1][y].flag = true
			bl := lookupXY(a, x-1, y, m, n, b, k+1)
			a[x-1][y].flag = false
			if bl {
				return bl
			}
		}

		if y > 0 && a[x][y-1].data == b[k+1] && !a[x][y-1].flag {
			fmt.Println("2==>", x, y-1, b[k+1])
			a[x][y-1].flag = true
			bl := lookupXY(a, x, y-1, m, n, b, k+1)
			a[x][y-1].flag = false
			if bl {
				return bl
			}
		}

		if x+1 < m && a[x+1][y].data == b[k+1] && !a[x+1][y].flag {
			fmt.Println("3==>", x+1, y, b[k+1])
			a[x+1][y].flag = true
			bl := lookupXY(a, x+1, y, m, n, b, k+1)
			a[x+1][y].flag = false
			if bl {
				return bl
			}
		}

		if y+1 < n && a[x][y+1].data == b[k+1] && !a[x][y+1].flag {
			fmt.Println("4==>", x, y+1, b[k+1])
			a[x][y+1].flag = true
			bl := lookupXY(a, x, y+1, m, n, b, k+1)
			a[x][y+1].flag = false
			if bl {
				return bl
			}
		}

	}

	return false
}

func searchLL(a [4][4]node, b []int) bool {
	m, n := 4, 4
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			k := 0
			if a[i][j].data == b[k] {
				a[i][j].flag = true
				if lookupXY(a, i, j, m, n, b, k) {
					return true
				}
				a[i][j].flag = false
			}
		}
	}
	return false
}

func main() {
	a := [4][4]node{[4]node{node{data: 0}, node{data: 7}, node{data: 1}, node{data: 2}}, [4]node{node{data: 4}, node{data: 5}, node{data: 20}, node{data: 8}}, [4]node{node{data: 9}, node{data: 4}, node{data: 3}, node{data: 0}}, [4]node{node{data: 5}, node{data: 2}, node{data: 1}, node{data: 3}}}
	var b []int = []int{0, 7, 1, 2, 8, 0, 3, 1, 2, 5, 9, 4, 3, 20, 5, 4}

	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			fmt.Printf("%d,%d:[%2d]  ", i, j, a[i][j].data)
		}
		fmt.Println("")
	}
	fmt.Println("-----------")
	fmt.Println(b)

	fmt.Println("================")
	fmt.Printf("\n-------\n[ %v ]\n", searchLL(a, b))
	fmt.Println("-------")

}
