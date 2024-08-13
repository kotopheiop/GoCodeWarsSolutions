package main

import (
	"fmt"
	"net"
)

func Is_valid_ip(ip string) bool {
	ipAddress := net.ParseIP(ip)
	if ipAddress.DefaultMask() != nil {
		return true
	}

	return false
}

func main() {
	fmt.Println(Is_valid_ip("255.255.255.255"))                         // true
	fmt.Println(Is_valid_ip("255.255.255"))                             // false
	fmt.Println(Is_valid_ip("255.255.255.785"))                         // false
	fmt.Println(Is_valid_ip("123.45.67.89"))                            // true
	fmt.Println(Is_valid_ip("123.045.067.089"))                         // false
	fmt.Println(Is_valid_ip("2001:0000:130F:0000:0000:09C0:876A:130B")) // false
}
