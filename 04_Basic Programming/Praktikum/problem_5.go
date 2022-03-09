package main

import "fmt"

func main() {
	fmt.Println(palindrome("katak"))
	fmt.Println(palindrome("mistar"))
	fmt.Println(palindrome("civic"))
	fmt.Println(palindrome("kasur rusak"))
	fmt.Println(palindrome("lion"))
	
}

func palindrome(input string) (isPalindrome bool) {
	for i := 0; i < len(input); i++ {
		if input[i] != input[len(input)-(i+1)] {
			return false
		}
	}
	return true
}