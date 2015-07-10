package main

import (
	"fmt"
	"runtime"
)


func Crt (n int, ch chan string) {
	for i:=0; i<n; i++ {
		ch<- fmt.Sprintf("hello %d", i)
	}
}

func Prt (ch chan string, s string) {
	for {
		select {
		case m, ok := <-ch:
			if ok {
				fmt.Println(s, m)
			}
		}
	}
}


func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	mc := make(chan string, 3)
	go Prt(mc, "1:")
	go Prt(mc, "2:")
	go Prt(mc, "3:")
	go Prt(mc, "4:")
	go Crt(500, mc)

	var e string
	fmt.Scanln(&e)
}
