package main

import "fmt"

func main() {
	arrSatu := []string{"kazuya", "jin", "lee"}
	arrDua := []string{"kazuya", "feng"}
	fmt.Println(ArrayMerge(arrSatu, arrDua))
}

func ArrayMerge(arrSatu, arrDua []string) (hasilGabung []string) {
	hmap := make(map[string]bool)
	for _, valueSatu := range arrSatu {
		hmap[valueSatu] = true
		hasilGabung = append(hasilGabung, valueSatu)
	}
	for _, valueDua := range arrDua {
		_, isFound := hmap[valueDua]
		if isFound == false {
			hmap[valueDua] = true
			hasilGabung = append(hasilGabung, valueDua)
		}
	}
	return hasilGabung
}
