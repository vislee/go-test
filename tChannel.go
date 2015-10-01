package main

import (
	"fmt"
	"sync"
	"time"
)

type GoSrv struct {
	sync.WaitGroup
}

//Wrap go the func
func (s *GoSrv) Wrap(cb func()) {
	s.Add(1)
	go func() {
		cb()
		s.Done()
	}()
}
func (s *GoSrv) Wrap1(cb func(ch chan interface{}), ch12 chan interface{}) {
	s.Add(1)
	go func() {
		cb(ch12)
		s.Done()
	}()
}

func (s *GoSrv) Wrap11(cb func(interface{}), ch12 interface{}) {
	s.Add(1)
	go func() {
		cb(ch12)
		s.Done()
	}()
}

func create(i interface{}) {
	ch, ok := i.(chan int)
	if !ok {
		fmt.Println("type error")
		return
	}
	var t = 0
	for {
		t += 1
		if t > 10 {
			goto end
		}
		ch <- t
		time.Sleep(100 * time.Millisecond)
	}
end:
	fmt.Println("go done")
}

func main() {
	var max = 5
	var srv GoSrv
	var ch []chan int
	var qch chan int
	qch = make(chan int, 0)

	ch = make([]chan int, max)
	for i := 0; i < max; i += 1 {
		ch[i] = make(chan int, 5)
	}
	for i := 0; i < max; i += 1 {
		srv.Wrap11(create, ch[i])
	}

	go func() {
		for {
			select {
			case <-qch:
				fmt.Println("end end done")
				return
			default:
			}
			for i := 0; i < max; i += 1 {
				select {
				case k, ok := <-ch[i]:
					if !ok {
						fmt.Println(i, "break")
						break
					}
					fmt.Println(i, k)
				default:
					fmt.Println("nil")
					time.Sleep(1 * time.Second)
					break
				}
			}
		}

	}()

	srv.Wait()
	time.Sleep(1 * time.Second)
	for i := 0; i < max; i += 1 {
		close(ch[i])
	}
	qch <- 1
	fmt.Println("done")
}
