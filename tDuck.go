package main

import (
	"fmt"
)

type Persion struct {
	name string
	age  int
}

func (p *Persion) Names() string {
	fmt.Println("names:", p.name)
	return p.name
}

func (p *Persion) Ages() int {
	return p.age
}

type PP interface {
	Names() string
	Ages() int
}

func main() {
	var p PP
	per := &Persion{"liwq", 33}
	p = per
	p.Names()
	// fmt.Println(p.age)  //这样直接调用是有问题的。
	fmt.Println(p.(*Persion).age)
	fmt.Println(per.name)
}
