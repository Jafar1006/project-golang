package main

import "fmt"

func Fahrenheit(C float32) float32 {
	var hasil float32
	hasil = 9*C/5 + 32
	return hasil
}

func Reamur(C float32) float32 {
	var hasil float32
	hasil = 4 * C / 5
	return hasil
}
func Kelvin(C float32) float32 {
	var hasil float32
	hasil = C + 273
	return hasil
}
func main() {
	var CelciusAwal, CelciusAkhir, step float32

	fmt.Scan(&CelciusAwal)
	fmt.Scan(&CelciusAkhir)
	fmt.Scan(&step)
	fmt.Println("Celcius ", "Reamur ", "Fahrenheit ", "Kelvin")
	for i := CelciusAwal; i <= CelciusAkhir; i = i + step {
		fmt.Println(CelciusAwal, "       ", Reamur(CelciusAwal), "       ", Fahrenheit(CelciusAwal), "       ", Kelvin(CelciusAwal))
		CelciusAwal = CelciusAwal + step
	}
}
