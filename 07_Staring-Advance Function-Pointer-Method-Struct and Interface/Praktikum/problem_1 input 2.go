package main

import (
	"fmt"
	"strings"
)

func main() {
	a := "AKA"
	b := "AKASHI"
	fmt.Println(Compare(b, a))

}

func Compare(b, a string) string {
	shortest, longest := b, a
	if len(b) > len(a) {
		longest = b
		shortest = a
	}
	if strings.Contains(longest, shortest) {
		return shortest
	}
	return ""
}
