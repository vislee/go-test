package main

import (
	"fmt"
	"sort"
	"time"
)

type AkRequest struct {
	Ak    string
	Times int64
	Req   int
}

type TopReq []*AkRequest

func (tr TopReq) Len() int {
	return len(tr)
}

func (tr TopReq) Less(i, j int) bool {
	tt := time.Now().Unix()
	if tt-tr[i].Times > 300 && tt-tr[j].Times < 300 {
		return true
	}
	if tt-tr[i].Times < 300 && tt-tr[j].Times > 300 {
		return false
	}
	return tr[i].Req < tr[j].Req
}

func (tr TopReq) Swap(i, j int) {
	tr[i], tr[j] = tr[j], tr[i]
}

type TimeTopN struct {
	num    int
	topReq TopReq
}

func NewTimeTopN(n int) *TimeTopN {
	return &TimeTopN{n, make([]*AkRequest, 0, n+1)}
}

func (ttn *TimeTopN) Insert(ar *AkRequest) {
	if ttn.topReq.Len() >= ttn.num {
		ttn.topReq = append(ttn.topReq[:ttn.num], ar)
	} else {
		ttn.topReq = append(ttn.topReq, ar)
	}
	sort.Sort(sort.Reverse(ttn.topReq))
}

func (ttn *TimeTopN) String() string {
	var s string
	for k, v := range ttn.topReq[:ttn.num] {
		s += fmt.Sprintf("seq:%d ak:%s times:%d request:%d. \n", k, v.Ak, v.Times, v.Req)
	}
	return s
}

func main() {
	trq := NewTimeTopN(10)
	nt := time.Now().Unix()
	lnt := time.Now().Add(-10 * time.Minute).Unix()
	var trq11 []*AkRequest
	trq11 = append(trq11, &AkRequest{"tt1", lnt, 3200})
	trq11 = append(trq11, &AkRequest{"tt2", lnt, 3100})
	trq11 = append(trq11, &AkRequest{"tt3", nt, 120})
	trq11 = append(trq11, &AkRequest{"tt4", nt, 130})
	trq11 = append(trq11, &AkRequest{"tt5", nt, 140})
	trq11 = append(trq11, &AkRequest{"tt6", nt, 150})
	trq11 = append(trq11, &AkRequest{"tt7", nt, 160})
	trq11 = append(trq11, &AkRequest{"tt8", nt, 150})
	trq11 = append(trq11, &AkRequest{"tt9", nt, 140})
	trq11 = append(trq11, &AkRequest{"tt10", nt, 130})
	trq11 = append(trq11, &AkRequest{"tt11", nt, 120})

	for _, v := range trq11 {
		trq.Insert(v)
	}

	fmt.Println(trq.String())
}
