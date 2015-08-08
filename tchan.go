package main

import (
	"fmt"
	"sync"
	"os"
	"os/signal"
	"syscall"
	"time"
)


func main() {
	var w sync.WaitGroup
	var ch1 = make(chan int, 0)
	var ch2 = make(chan int, 0)
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		w.Add(1)
		defer w.Done()
		fmt.Println("go 2")
		ch2 <- 2
		fmt.Println("go 2 done")
	}()

	go func() {
		w.Add(1)
		defer w.Done()
		fmt.Println("go 1")
		ch1 <- 3
		fmt.Println("go 1 done")
	}()


	c := time.Tick(3 * time.Second)
	for i := 0; i < 5; i++ {
		select {
		case <- ch1:
			fmt.Println("ch1 ok")
		case <- ch2:
			fmt.Println("ch2 ok")
		case <- c:
			fmt.Println("time out")
		// default:
		// 	fmt.Println("error")
		}
	}

	fmt.Println("done")
	fmt.Println(<-ch)
	fmt.Println("done2")
	w.Wait()
	fmt.Println("done3")
}
