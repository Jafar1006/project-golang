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
	hasil := x / y
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
	fmt.Scan(&n)
	for n != 4 {
		if n == 1 {
			fmt.Print("Masukan dua bilangan yang akan dijumlahkan: ")
			fmt.Scan(&x, &y)
			hitungJumlah(x, y)
		} else if n == 2 {
			fmt.Print("Masukan dua bilangan yang akan dikalikan: ")
			fmt.Scan(&x, &y)
			hitungKali(x, y)
		} else if n == 3 {
			fmt.Print("Masukan dua bilangan yang akan dibagikan: ")
			fmt.Scan(&x, &y)
			hitungBagi(x, y)
		}
		menu()
		fmt.Scan(&n)
	}
}
