package main

import (
	"fmt"
	"strconv"
)

func MaxRot(n int64) int64 {
	str := strconv.FormatInt(n, 10)
	maxStr := str
	for i := 0; i < len(str); i++ {
		str = str[:i] + str[i+1:] + string(str[i])
		if str > maxStr {
			maxStr = str
		}
	}
	maxN, _ := strconv.ParseInt(maxStr, 10, 64)
	return maxN
}

func main() {
	fmt.Println(MaxRot(38458215))  // 85821534
	fmt.Println(MaxRot(195881031)) // 988103115
	fmt.Println(MaxRot(896219342)) // 962193428
	fmt.Println(MaxRot(69418307))  // 94183076
}
