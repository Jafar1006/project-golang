package main

import "fmt"

func hitungMenang(jm *int) {
	*jm = *jm + 1

}

func hitungDraw(jd *int) {
	*jd = *jd + 1
}

func hitungKalah(jk *int) {
	*jk = *jk + 1
}

func hitungjumGolKegolanSelisih(g, k int, jg, jkegolan, jsg *int) {
	*jg += g
	*jkegolan += k
	*jsg = *jg - *jkegolan
}
func hitungJumPoint(jm, jd int, jp *int) {
	*jp = 3*jm + jd
}
func main() {
	var n, g, k, jg, jk, jkegolan, jsg, jd, jm, jp int
	fmt.Print("masukan jumlah pertandingan: ")
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&g, &k)
		if g > k {
			hitungMenang(&jm)
		} else if g < k {
			hitungKalah(&jk)
		} else if g == k {
			hitungDraw(&jd)
		}

		hitungjumGolKegolanSelisih(g, k, &jg, &jkegolan, &jsg)
		hitungJumPoint(jm, jd, &jp)
	}
	fmt.Println(n, jm, jd, jk, jg, jkegolan, jsg, jp)
}
