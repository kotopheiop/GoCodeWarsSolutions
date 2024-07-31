package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func Crack(hash string) string {
	var pin string
	for i := 0; i <= 99999; i++ {
		pin = fmt.Sprintf("%05d", i)
		if getMD5Hash(pin) == hash {
			return pin
		}
	}

	return ""
}

func getMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))

	return hex.EncodeToString(hash[:])
}

func main() {
	fmt.Println(Crack("86aa400b65433b608a9db30070ec60cd")) // 00078
	fmt.Println(Crack("827ccb0eea8a706c4c34a16891f84e7b")) // 12345
	fmt.Println(Crack("dcddb75469b4b4875094e14561e573d8")) // 00000
	fmt.Println(Crack("d3eb9a9233e52948740d7eb8c3062d14")) // 99999
}
