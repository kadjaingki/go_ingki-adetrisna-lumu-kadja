package main

import "fmt"

func main() {
	s := "76523752"
	AngkaMunculSekali(s)
}

func AngkaMunculSekali(s string) []int {

	for i := 0; i < len(s); i++ {
		fmt.Println(string(s[i]))
		val, isFound := hmap[string(s[i])]
		if isFound {
			hmap[string(s)[i]] = val + 1
		} else {
			hmap[string(s[i])] = 1
		}

	}
	return []int{}
}
