package main

import (
	"fmt"
	"strings"
)

func main() {
	a := "KANGGOORO"
	b := "KANG"
	fmt.Println(Compare(a, b))

}

func Compare(a, b string) string {
	shortest, longest := a, b
	if len(a) > len(b) {
		longest = a
		shortest = b
	}
	if strings.Contains(longest, shortest) {
		return shortest
	}
	return ""
}
