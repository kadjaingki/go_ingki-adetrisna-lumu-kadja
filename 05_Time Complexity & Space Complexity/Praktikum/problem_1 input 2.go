package main

import "fmt"

func main() {

	if cekPrima(1500450271) {
		fmt.Println("merupakan bilangan prima")
	} else {
		fmt.Println("bukan merupakan bilangan prima")
	}

}

func cekPrima(input int) bool {
	var result bool

	if input == 2 || input == 3 || input == 5 || input == 7 {
		result = true
	} else if input%2 == 0 || input%3 == 0 || input%5 == 0 || input%7 == 0 {
		result = false
	} else {
		result = true
	}
	return result
}
