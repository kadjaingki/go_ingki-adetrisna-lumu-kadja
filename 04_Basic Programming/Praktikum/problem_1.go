package main

import "fmt"

func main() {
	//input
	var t float64 = 20
	const r = 4
	const pi = 3.14

	fmt.Printf("Masukan luas : ")
	fmt.Scanf("%f", &t)
	luas := 2 * pi * r * (r + t)
	fmt.Print("Total adalah : ", luas)

}
