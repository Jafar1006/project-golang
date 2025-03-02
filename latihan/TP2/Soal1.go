package main

import (
	"fmt"
)

func hitungLuasKelilingLingkaran(r float64, l, k *float64) {
	var pi float64 = 3.14
	*l = pi * r * r
	*k = 2 * pi * r
}

func hitungLuasKelilingPersegi(s float64, l, k *float64) {
	*l = s * s
	*k = 4 * s
}

func hitungTotal(lL, lP, kL, kP float64, totLuas, totKel *float64) {
	*totLuas = lL + lP
	*totKel = kL + kP
}
func main() {
	var r, s, LL, KL, LP, KP, TL, TP float64
	fmt.Scan(&r, &s)
	hitungLuasKelilingLingkaran(r, &LL, &KL)
	hitungLuasKelilingPersegi(s, &LP, &KP)
	hitungTotal(LL, LP, KL, KP, &TL, &TP)

	fmt.Printf("%-10s   %-10s   %-10s   %-10s   %-10s   %-10s   %-10s   %-10s\n", "R", "S", "LL", "LP", "KL", "KP", "TL", "TP")
	fmt.Printf("%-10.2f | %-10.2f | %-10.2f | %-10.2f | %-10.2f | %-10.2f | %-10.2f | %-10.2f\n", r, s, LL, LP, KL, KP, TL, TP)
}
