package main

import "fmt"

func main() {
	if isPrime(10) {
		fmt.Println("Bilangan Prima")
	} else {
		fmt.Println("Bukan Bilangan Prima")
	}
}

func isPrime(number int) bool {
	pembagi := 0

	for i := 1; i <= number; i++ {
		if number%i == 0 {
			pembagi++ // ini adalah operasi primitif
		}
	}

	if pembagi == 2 {
		return true
	}

	return false
}
