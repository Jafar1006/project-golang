package main

import "fmt"

const NMAX int = 100

type Subscription struct {
	namaLangganan string
	biaya         int
}
type finansialManage [NMAX]Subscription

func main() {
	var arr finansialManage
	var n int
	var perintah string
	cetakUI()
	fmt.Scan(&perintah)
	for {
		switch perintah {
		case "1":
			baca(&arr, &n)
			fmt.Println(n)
			cetakUI()
			fmt.Scan(&perintah)
			// tambah subscription apa saja
		case "2":
			cetakarr(arr, n)
			hapusarray(&arr, &n)
			fmt.Println(n)
			cetakarr(arr, n)
			cetakUI()
			fmt.Scan(&perintah)
			//hapus langganan
		case "3":
			fmt.Printf("Banyaknya langganan %d", n)
			cetakUI()
			fmt.Scan(&perintah)
			//detail semua langganan
		case "4":
			// detail semua pengeluaran
		case "5":
			return
		default:
			fmt.Println("✖ Pilihan tidak valid.")
		}
	}
}
func cetakUI() {
	fmt.Println("\n===== Aplikasi Manajemen Subskripsi & Keuangan =====")
	fmt.Println("1. Tambah Langganan")
	fmt.Println("2. Hapus Langganan")
	fmt.Println("3. Lihat Semua Langganan")
	fmt.Println("4. Lihat Total Pengeluaran Bulanan")
	fmt.Println("5. Keluar")
	fmt.Print("Pilih menu (1-5): ")
}
func baca(A *finansialManage, n *int) {
	var nLangganan string
	var cost int
	fmt.Scan(&nLangganan)
	for nLangganan != "none" {
		fmt.Scan(&cost)
		A[*n].namaLangganan = nLangganan
		A[*n].biaya = cost
		*n++
		fmt.Scan(&nLangganan)
	}
}
func hapusarray(A *finansialManage, n *int){
	var i , j int
	fmt.Print("Pilih nomor yang ingin dihapus")
	fmt.Scan(&j)
	for i = 0 ; i<*n ; i++{
		A[j-1]=A[j]
		j++
	}	
	*n = i-1
}
func cetakarr(A finansialManage, n int){
	var i int
	for i = 0; i<n ; i++{
		fmt.Printf("%d. %s %d\n",i+1 , A[i].namaLangganan, A[i].biaya )
	}
}
