package main

import (
	"reflect"
	"fmt"
)

type item struct {
	Age int
	Year float64
}

type persion struct {
	name string
	Itt  item
}

func (p *persion) Test() *item {
	fmt.Println("test test")
	// fmt.Println(s)
	return &item{22, 333.00}
}


func (p *persion) Hello(s []string, ss ... string) int {
	fmt.Println("data:", *p)
	fmt.Println("hello:", s)
	return 333
}


func (p *persion) Run(s string) {
	// sz := []string{"zhou", "liwq"}
	// params := make([]reflect.Value, 2)
	v := reflect.ValueOf(p)
	// params[0] = reflect.ValueOf(sz)
	// params[1] = reflect.ValueOf("mumu")
	rt := v.MethodByName(s).Call(nil)
	ss := rt[0].Interface()
	fmt.Println(ss.(*item))
	// fmt.Println(ss.Type())
	// switch ss.(type) {
	// case item:
	// 	it, ok := ss.(item)
	// 	fmt.Println("item", it, ok)
	// case *item:
	// 	fmt.Println("*item")
	// 	it, ok := ss.(*item)
	// 	fmt.Println(*it, ok)

	// default:
	// 	fmt.Println("error")
	// }
}


func main() {
	p := persion{name:"liwq",}
	p.Itt.Age = 30
	p.Itt.Year = 22.0

	// s := reflect.ValueOf(&p).Elem()
	// val := reflect.ValueOf(item{333, 0.0})
	// s.FieldByName("Itt").Set(val)
	// fmt.Println(n.Float())
	// val.SetFloat(38.0)

	fmt.Println(p)

	p.Run("Test")

}
