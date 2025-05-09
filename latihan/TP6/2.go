package main

import "fmt"

type titik struct {
	x, y  float64
	warna string
}

func main() {
	var x1, y1, x2, y2 float64
	var w1, w2 string
	var t1, t2 titik

	//baca data x1, y1, w2, x2, y2,w2
	fmt.Scan(&x1, &y1, &w1)
	fmt.Scan(&x2, &y2, &w2)

	//pembuatan titik t1 dan t2
	t1 = pembuatan_titik_baru(x1, y1, w1)
	t2 = pembuatan_titik_baru(x2, y2, w2)

	//pencetakan titik t1 dan t2
	fmt.Printf("Data titik 1: Koordinat (%v,%v), warna %s\n", t1.x, t1.y, t1.warna)
	fmt.Printf("Data titik 2: Koordinat (%v,%v), warna %s\n", t2.x, t2.y, t2.warna)
}

func pembuatan_titik_baru(x, y float64, w string) titik {
	//mengembalikan sebuah titik dengan koordinat x dan y, serta warna
	var b titik
	b.x = x
	b.y = y
	b.warna = w
	return b
}
