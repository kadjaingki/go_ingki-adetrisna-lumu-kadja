package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(maximumBuyProduct(50000, []int{25000, 25000, 10000, 14000}))
	fmt.Println(maximumBuyProduct(30000, []int{15000, 10000, 12000, 5000, 3000}))
}

func maximumBuyProduct(money int, products []int) int {
	sort.Ints(products)
	var i = 0
	var countProduct = 0
	for money >= 0 && i < len(products) {
		money = money - products[i]
		countProduct++
		i++
	}
	return countProduct - 1
}
