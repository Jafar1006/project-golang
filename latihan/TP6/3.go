package main

import "fmt"

type pegawai struct {
	nama                   string
	gaji_pokok, masa_kerja int
	besar_bonus            float64
}

func main() {
	var pekerja pegawai

	//baca data pegawai
	fmt.Scan(&pekerja.nama, &pekerja.gaji_pokok, &pekerja.masa_kerja)

	//hitung bonus dengan memanggil prosedur hitung_bonus
	hitung_bonus(&pekerja)
	//cetak besar bonus
	fmt.Printf("Pegawai dengan nama %s mendapatkan bonus sebesar Rp %.f\n", pekerja.nama, pekerja.besar_bonus)

}
func hitung_bonus(p *pegawai) {
	if p.masa_kerja >= 10 {
		p.besar_bonus = float64(p.gaji_pokok) * 1.5
	} else if p.masa_kerja < 10 && p.masa_kerja >= 5 {
		p.besar_bonus = float64(p.gaji_pokok) * 0.75
	} else {
		p.besar_bonus = float64(p.gaji_pokok) * 0.5
	}
}
