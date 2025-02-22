package main

import "fmt"

func adaBilanganM(n, m int) string {
	var last, count int
	var hasil string
	count = 0
	for n >= 1 {
		last = n % 10
		if last == m {
			count += 1
		} else if last != m {
			count += 0
		}
		n = n / 10
	}
	if count >= 1 {
		hasil = "YA"
	} else if count == 0 {
		hasil = "TIDAK"
	}

	return hasil
}

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	if n < 10 || n > 999999 {
		fmt.Println("Angka masukan untuk n hanya bisa untuk diantara 10-999999")
	}

	if m < 0 || m > 9 {
		fmt.Print("Angka masukan untuk m hanya bisa untuk diantara 0-9")
	}
	fmt.Print(adaBilanganM(n, m))
}
