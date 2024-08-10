package main

import (
	"fmt"
	"strconv"
	"strings"
)

func IpsBetween(start, end string) int {
	startIP := ipToInt(start)
	endIP := ipToInt(end)

	return endIP - startIP
}

func ipToInt(ip string) int {
	var result int

	parts := strings.Split(ip, ".")
	for _, part := range parts {
		num, _ := strconv.Atoi(part)
		result = result<<8 + num
	}

	return result
}

func main() {
	fmt.Println(IpsBetween("10.0.0.0", "10.0.0.50")) // 50
	fmt.Println(IpsBetween("20.0.0.10", "20.0.1.0")) // 246
}
