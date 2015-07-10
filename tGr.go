package main

import (
	"fmt"
	"runtime"
)


type Test struct {
	msgCh chan string
}

func(t *Test) Crt(n int) {
	for i:=0; i<n; i++ {
		t.msgCh <- fmt.Sprintf("hello %d", i)
	}
}

func(t *Test) Prt(s string) {
	for {
		select {
		case msg, ok := <-t.msgCh:
			if ok {
				fmt.Println(s, msg)
			}
			// runtime.Gosched()
		}
	}
}



func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	var input string
	t := new(Test)
	t.msgCh = make(chan string, 2)
	go t.Prt("1:")
	go t.Prt("2:")
	go t.Prt("3:")
	go t.Prt("4:")
	go t.Crt(500)

	fmt.Scanln(&input)
}

