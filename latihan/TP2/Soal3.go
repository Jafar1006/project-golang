package main

import "fmt"

func hitungJumlah(x, y int) {
	hasil := x + y
	fmt.Println("Hasil penjumlahan:", hasil)
}
func hitungKali(x, y int) {
	hasil := x * y
	fmt.Println("Hasil perkalian:", hasil)
}
func hitungBagi(x, y int) {
	var hasil = float64(x) / float64(y)
	fmt.Println("Hasil pembagian:", hasil)
}
func menu() {
	fmt.Println("---------------------------")
	fmt.Printf("%17s\n", " M E N U")
	fmt.Println("---------------------------")
	fmt.Println("1. Hitung Penjumlahan")
	fmt.Println("2. Hitung Perkalian")
	fmt.Println("3. Hitung Pembagian")
	fmt.Println("4. Exit")
	fmt.Println("---------------------------")
	fmt.Print("Pilih (1/2/3/4)?")
}
func main() {
	var n, x, y int
	menu()
	fmt.Scanln(&n)
	for n != 4 {
		if n == 1 {
			fmt.Println("Masukan dua bilangan yang akan dijumlahkan: ")
			fmt.Scanln(&x, &y)
			hitungJumlah(x, y)
		} else if n == 2 {
			fmt.Println("Masukan dua bilangan yang akan dikalikan: ")
			fmt.Scanln(&x, &y)
			hitungKali(x, y)
		} else if n == 3 {
			fmt.Println("Masukan dua bilangan yang akan dibagikan: ")
			fmt.Scanln(&x, &y)
			hitungBagi(x, y)
		} else {
			fmt.Println("Tolong pilih sesuai dengan instruksi yang diberikan!")
		}
		menu()
		fmt.Scanln(&n)
	}
}
