package main

import (
	"fmt"
	"math"
)

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}

	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}

func main() {
	var num int
	fmt.Println("Masukkan Angka : ")
	fmt.Scanln(&num)
	if isPrime(num) {
		fmt.Println(num, "adalah bilangan prima")
	} else {
		fmt.Println(num, "bukan bilangan prima")
	}
}
