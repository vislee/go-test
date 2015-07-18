package main

import (
	"fmt"
	"sync"
)

var mmbuf = sync.Pool{
	New: func() interface{} {
		b := make([]byte, 16)
		return &b
	},
}

func main() {
	k := mmbuf.Get().(*[]byte)
	copy(*k, "hello liwq")
	fmt.Println(1, string(*k))
	mmbuf.Put(k)

	m := mmbuf.Get().(*[]byte)
	copy(*m, "hello zhou")
	fmt.Println(6, string(*m))
	mmbuf.Put(m)
}
