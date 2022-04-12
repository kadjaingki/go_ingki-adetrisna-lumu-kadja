package main

import "fmt"

func SimpleEquationsBrutoForce(a, b, c int) {
	found := false
	for i := 1; i <= 10; i++ {
		for j := 1; j <= 10; j++ {
			for k := 1; k <= 10; k++ {
				if i+j+k == a && i*j*k == b && (i*i)+(j*j)+(k*k) == c {
					fmt.Println(i, j, k)
					found = true
					break
				}
			}
			if found {
				break
			}
		}
		if found {
			break
		}
	}

	if !found {
		fmt.Println("no solution")
	}
}

func SimpleEquations(a, b, c int) {

	for x := 1; x <= a; x++ {
		for y := x; y <= a; y++ {
			if b%(x*y) != 0 {
				continue
			}
			z := a - x - y
			if z < 1 || z >= a {
				continue
			}
			if x*x+y*y+z*z == c {
				fmt.Println(x, y, z)
				return
			}
		}
	}
	fmt.Println("no solution")
}

func main() {
	SimpleEquations(1, 2, 3)
	SimpleEquations(6, 6, 14)
}
