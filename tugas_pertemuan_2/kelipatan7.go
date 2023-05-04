package main

import "fmt"

func main() {
	var n int
	fmt.Print("Masukkan angka: ")
	fmt.Scanln(&n)

	if n%7 == 0 {
		fmt.Printf("%d adalah kelipatan 7\n", n)
	} else {
		fmt.Printf("%d bukan kelipatan 7\n", n)
	}
}
