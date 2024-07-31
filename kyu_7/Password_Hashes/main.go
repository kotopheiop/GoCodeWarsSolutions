package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func PassHash(str string) string {
	hash := md5.Sum([]byte(str))

	return hex.EncodeToString(hash[:])
}

func main() {
	fmt.Println(PassHash("password")) // 5f4dcc3b5aa765d61d8327deb882cf99
	fmt.Println(PassHash("abc123"))   //e99a18c428cb38d5f260853678922e03
}
