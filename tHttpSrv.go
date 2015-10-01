package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
)

type HandlerSrv struct {
}

func (hl *HandlerSrv) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Printf("remote ip: %s, url: %s\n", req.RemoteAddr, req.URL.Path)
	hl.router(w, req)
}

func (hl *HandlerSrv) router(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/ping":
		hl.ping(w, req)
	default:
		hl.noPage(w, req)
	}
}

func (hl *HandlerSrv) ping(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Length", "2")
	io.WriteString(w, "OK")
}

func (hl *HandlerSrv) noPage(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Length", "7")
	io.WriteString(w, "NO PAGE")
}

func main() {
	ls, err := net.Listen("tcp", "127.0.0.1:6378")
	if err != nil {
		log.Fatalf("listen error. error: %s", err.Error())
	}
	defer ls.Close()

	var hl HandlerSrv
	hr := &http.Server{
		Handler: &hl,
	}
	err = hr.Serve(ls)
	if err != nil {
		log.Fatalf("serve error. error: %s", err.Error())
	}
}
