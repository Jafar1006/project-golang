package main

import "fmt"

func adaBilanganM(n, m int) string {
	var last int
	var hasil string
	for n >= 10 {
		last = n % 10
		if last == m {
			hasil = "YA"
		} else if last != m {
			hasil = "TIDAK"
		}
		n = n / 10
	}

	return hasil
}

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	if n <= 10 || n >= 999999 {
		fmt.Println("Angka masukan untuk n hanya bisa untuk diantara 10-999999")
	}

	if m <= 0 || m >= 9 {
		fmt.Print("Angka masukan untuk m hanya bisa untuk diantara 0-9")
	}
	fmt.Print(adaBilanganM(n, m))
}
