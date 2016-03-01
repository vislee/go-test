package main

import (
	"fmt"
	"net"
)

func GetLocalIpv4(dev string) (string, error) {
	itf, err := net.InterfaceByName(dev)
	if err != nil {
		return "", err
	}

	ads, err := itf.Addrs()
	if err != nil {
		return "", err
	}

	for _, addr := range ads {
		if in, ok := addr.(*net.IPNet); ok && !in.IP.IsLoopback() {
			if in.IP.To4() != nil {
				return in.IP.String(), nil
			}
		}
	}

	return "", fmt.Errorf("interface %s no ipv4", dev)
}

func main() {
	ipv4, err := GetLocalIpv4("en0")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(ipv4)
}
