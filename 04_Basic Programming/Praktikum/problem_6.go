package main

import "fmt"

func main() {
	// fmt.Println("hallo")
	fmt.Println(Exponent(2, 3))
	fmt.Println(Exponent(5, 3))
	fmt.Println(Exponent(10, 2))
	fmt.Println(Exponent(2, 5))
	fmt.Println(Exponent(7, 3))
}

func Exponent(bilangan, pangkat int) (hasilExponent int) {
	if pangkat == 0 {
		hasilExponent = 1
		return hasilExponent
	}
	hasilExponent = bilangan
	for i := 1; i < pangkat; i++ {
		hasilExponent = hasilExponent * bilangan
	}
	return hasilExponent
}