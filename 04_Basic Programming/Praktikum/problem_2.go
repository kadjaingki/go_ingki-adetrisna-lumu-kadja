package main

import "fmt"

func main() {
	var nilai int

	fmt.Print("Input Nilai : ")
	fmt.Scanln(&nilai)

	if nilai >= 80 && nilai <= 100 {
		fmt.Println("Nilai A")
	} else {
		fmt.Println("Nilai Invalid")
	}
}
