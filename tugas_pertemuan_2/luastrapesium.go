package main

import "fmt"

func main() {
	var a, b, t float64
	fmt.Print("Masukkan panjang sisi sejajar atas (a) trapesium: ")
	fmt.Scanln(&a)
	fmt.Print("Masukkan panjang sisi sejajar bawah (b) trapesium: ")
	fmt.Scanln(&b)
	fmt.Print("Masukkan tinggi (t) trapesium: ")
	fmt.Scanln(&t)

	luas := ((a + b) * t) / 2
	fmt.Printf("Luas trapesium dengan sisi sejajar atas %v dan sisi sejajar bawah %v serta tinggi %v adalah %v\n", a, b, t, luas)
}
