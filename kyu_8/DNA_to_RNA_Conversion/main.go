package main

import (
	"fmt"
	"strings"
)

func DNAtoRNA(dna string) string {
	rnaString := strings.ReplaceAll(dna, "T", "U")

	return rnaString
}

func main() {
	fmt.Println(DNAtoRNA("GCAT"))
}
