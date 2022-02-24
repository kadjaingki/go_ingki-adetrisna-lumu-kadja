package main

import "fmt"

func main() {
	var bilangan int

	fmt.Print("Masukkan angka : ")
	fmt.Scanln(&bilangan)

	for a := 1; a <= bilangan; a++ {
		if bilangan%a == 0 {
			fmt.Println(a)
		}
	}
}
